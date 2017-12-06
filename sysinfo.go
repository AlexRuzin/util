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

package util

import (
    "runtime"
)

type SystemInfo struct {
    /*
     * Interface MAC/IP addresses, GeoIP Location in a juicy serialized structure,
     *  and global IP address as reported by "ipinfo.io"
     */
    GlobalIP    *GeoIP
    LocalIP     string

    /* Operating system version/build */
    OSName      string

    /* Local host information -- should be universal */
    Hostname    string
    Username    string
}

func GetSystemInfo() (sysinfo *SystemInfo, err error) {
    var report = &SystemInfo{
        LocalIP: "",
        GlobalIP: nil,
        OSName: runtime.GOOS,
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

    case "windows":

    case "linux":

    default:
    }

    return nil, nil
}