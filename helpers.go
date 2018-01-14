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
    "os"
    "strconv"
    "math/rand"
    "time"
    "sync"
    "bytes"
    "encoding/base64"
    "unicode"
    "bufio"
    "encoding/gob"
    "syscall"
    "unsafe"
)

func GetStdin() *string {
   reader := bufio.NewReader(os.Stdin)
   DebugOut("in> ")
   data, _ := reader.ReadString('\n')

   /* Strip '\n' character */
   new := make([]byte, len(data) - 1)
   copy(new, data)
   output := string(new)

   return &output
}

func IntToString(n int) string {
    var output string

    output = strconv.FormatInt(int64(n), 10)
    if output == "" {
        ThrowN("Invalid input to IntToString")
    }

    return output
}

func RandIntRange(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func WaitForever() {
    m := sync.Mutex{}
    m.Lock()
    m.Lock()
}


func RandomString(l int) string {
    var result bytes.Buffer
    var temp string
    for i := 0; i < l; {
        if string(RandInt(65, 90)) != temp {
            temp = string(RandInt(65, 90))
            result.WriteString(temp)
            i++
        }
    }
    return result.String()
}

func RandInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}

func Sleep(val time.Duration) {
    time.Sleep(val)
}

func SleepSeconds(val time.Duration) {
    time.Sleep(val * time.Second)
}

func SleepHours(val time.Duration) {
    time.Sleep(val * time.Hour)
}

func B64E(d []byte) string {
    return base64.StdEncoding.EncodeToString(d)
}

func B64D(d string) (data []byte, err error) {
    output, is_ok := base64.StdEncoding.DecodeString(d)
    if is_ok != nil {
        return nil, is_ok
    }

    return output, nil
}

func IsAsciiPrintable(s string) bool {
    for _, r := range s {
        if r > unicode.MaxASCII || !unicode.IsPrint(r) {
            return false
        }
    }
    return true
}

func Serialize(d interface{}) ([]byte, error) {
    b := new(bytes.Buffer)
    e := gob.NewEncoder(b)
    if err := e.Encode(d); err != nil {
        return nil, err
    }
    return b.Bytes(), nil
}

func DeSerialize(data []byte) (* interface {}, error) {
    var output interface{}

    b := bytes.Buffer{}
    b.Write(data)
    d := gob.NewDecoder(&b)

    if err := d.Decode(&output); err != nil {
        return nil, err
    }

    return &output, nil
}

/*
 * Create mutex function for win32
 * https://stackoverflow.com/questions/23162986/restricting-to-single-instance-of-executable-with-golang
 */
var (
    kernel32        = syscall.NewLazyDLL("kernel32.dll")
    procCreateMutex = kernel32.NewProc("CreateMutexW")
)
func CreateMutexGlobal(name string) (uintptr, error) {
    globalMutex := "Global\\\\" + name
    ret, _, err := procCreateMutex.Call(
        0,
        0,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(globalMutex))),
    )
    switch int(err.(syscall.Errno)) {
    case 0:
        return ret, nil
    default:
        return ret, err
    }
}

/*
 * Generate a SQL DATETIME compliant string
 * https://stackoverflow.com/questions/21648842/output-go-time-in-rfc3339-like-mysql-format
 */
func CreateSqlDatetime() string {
    const createFormat = "2006-01-02 15:04:05"
    return time.Unix(1391878657, 0).Format(createFormat)
}