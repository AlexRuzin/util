/*
 * Copyright (c) 2017 AlexRuzin (stan.ruzin@gmail.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Thanks to matishsiao for his code for getsysinfo.
 *  https://github.com/matishsiao/goInfo
 */

package util

import (
    "runtime"
    "os/user"
    "os"
    "github.com/matishsiao/goInfo"
    "strings"
    "time"
    "os/exec"
    "bytes"
    "fmt"
    "crypto/rand"
    "encoding/hex"
    "io/ioutil"
)

type SystemInfo struct {
    /*
     * Interface MAC/IP addresses, GeoIP Location in a juicy serialized structure,
     *  and global IP address as reported by "ipinfo.io"
     */
    GlobalIP                *GeoIP
    LocalIP                 string

    /* Operating system version/build */
    OSName                  string

    /* Local host information -- should be universal */
    Hostname                string
    Username                string

    /* GoInfoObject -- kernel and detailed system info */
    GoInfo                  *goInfo.GoInfoObject

    /* Windows specific output of "systeminfo" command */
    SystemInfoCommand       *string
}

func GetSystemInfo() (sysinfo *SystemInfo, err error) {
    hostname, username := func () (hostname string, username string) {
        hostname = "<unknown>"
        username = "<unknown>"

        u, err := user.Current()
        if err == nil {
            username = u.Username
        }

        h, err := os.Hostname()
        if err == nil {
            hostname = h
        }

        return
    } ()

    var report = &SystemInfo{
        LocalIP: "",
        GlobalIP: nil,
        OSName: runtime.GOOS,
        Hostname: hostname,
        Username: username,
        GoInfo: nil,
    }

    localip, err := GetLocalIP()
    if err == nil {
        report.LocalIP = *localip
    }

    globalip, err := GetGeoIP()
    if err == nil {
        report.GlobalIP = globalip
    }

    switch report.OSName {
    case "darwin":
        report.GoInfo = func () *goInfo.GoInfoObject {
            out := getInfoDarwin()
            for strings.Index(out,"broken pipe") != -1 {
                out = getInfoDarwin()
                time.Sleep(500 * time.Millisecond)
            }
            osStr := strings.Replace(out,"\n","",-1)
            osStr = strings.Replace(osStr,"\r\n","",-1)
            osInfo := strings.Split(osStr," ")
            gio := &goInfo.GoInfoObject{Kernel:osInfo[0],Core:osInfo[1],Platform:osInfo[2],OS:osInfo[0],GoOS:runtime.GOOS,CPUs:runtime.NumCPU()}
            gio.Hostname,_ = os.Hostname()
            return gio
        } ()
    case "windows":
        report.GoInfo = func () *goInfo.GoInfoObject {
            cmd := exec.Command("cmd","ver")
            cmd.Stdin = strings.NewReader("some input")
            var out bytes.Buffer
            var stderr bytes.Buffer
            cmd.Stdout = &out
            cmd.Stderr = &stderr
            err := cmd.Run()
            if err != nil {
                panic(err)
            }
            osStr := strings.Replace(out.String(),"\n","",-1)
            osStr = strings.Replace(osStr,"\r\n","",-1)
            tmp1 := strings.Index(osStr,"[Version")
            tmp2 := strings.Index(osStr,"]")
            var ver string
            if tmp1 == -1 || tmp2 == -1 {
                ver = "unknown"
            } else {
                ver = osStr[tmp1+9:tmp2]
            }
            gio := &goInfo.GoInfoObject{Kernel:"windows",Core:ver,Platform:"unknown",OS:"windows",GoOS:runtime.GOOS,CPUs:runtime.NumCPU()}
            gio.Hostname,_ = os.Hostname()
            return gio
        } ()

        report.SystemInfoCommand = func () *string {
            var output = "['systeminfo' output]:\n\n"

            r := make([]byte, 8)
            rand.Read(r)
            output_filename := os.Getenv("TEMP") + "\\_" + hex.EncodeToString(r) + ".txt"

            cmd := exec.Command("systeminfo", " > " + output_filename)

            err := cmd.Start()
            SleepSeconds(10)

            if err != nil {
                output += "<command failed>"
            } else {
                if _, err := os.Stat(output_filename); !os.IsNotExist(err) {
                    raw_file, err := ioutil.ReadFile(output_filename)
                    if err != nil {
                        output += "<error: Report file not found>"
                    }
                    output += string(raw_file)
                } else {
                    /* FIXME -- The command does not start, nor generate the file */
                    output += "<error: Report file not found>"
                }
            }

            return &output
        } ()
    case "linux":
        report.GoInfo = func () *goInfo.GoInfoObject {
            out := getInfoLinux()
            for strings.Index(out,"broken pipe") != -1 {
                out = getInfoLinux()
                time.Sleep(500 * time.Millisecond)
            }
            osStr := strings.Replace(out,"\n","",-1)
            osStr = strings.Replace(osStr,"\r\n","",-1)
            osInfo := strings.Split(osStr," ")
            gio := &goInfo.GoInfoObject{Kernel:osInfo[0],Core:osInfo[1],Platform:osInfo[2],OS:osInfo[3],GoOS:runtime.GOOS,CPUs:runtime.NumCPU()}
            gio.Hostname,_ = os.Hostname()
            return gio
        } ()
    case "freebsd":
        report.GoInfo = func () *goInfo.GoInfoObject {
            out := getInfoFreeBSD()
            for strings.Index(out,"broken pipe") != -1 {
                out = getInfoFreeBSD()
                time.Sleep(500 * time.Millisecond)
            }
            osStr := strings.Replace(out,"\n","",-1)
            osStr = strings.Replace(osStr,"\r\n","",-1)
            osInfo := strings.Split(osStr," ")
            gio := &goInfo.GoInfoObject{Kernel:osInfo[0],Core:osInfo[1],Platform:runtime.GOARCH,OS:osInfo[2],GoOS:runtime.GOOS,CPUs:runtime.NumCPU()}
            gio.Hostname,_ = os.Hostname()
            return gio
        } ()
    default:
    }

    return nil, nil
}

func getInfoLinux() string {
    cmd := exec.Command("uname","-srio")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("getInfo:",err)
    }
    return out.String()
}

func getInfoDarwin() string {
    cmd := exec.Command("uname","-srm")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("getInfo:",err)
    }
    return out.String()
}

func getInfoFreeBSD() string {
    cmd := exec.Command("uname","-sri")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("getInfo:",err)
    }
    return out.String()
}
