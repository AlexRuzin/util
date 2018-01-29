// +build windows
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
    "syscall"
    "unsafe"
    "sync"
    "strconv"
)

/*
 * Create mutex function for win32
 * https://stackoverflow.com/questions/23162986/restricting-to-single-instance-of-executable-with-golang
 */
var (
    kernel32        = syscall.NewLazyDLL("kernel32.dll")
    procCreateMutex = kernel32.NewProc("CreateMutexW")
    procCloseHandle = kernel32.NewProc("CloseHandle")
)

/*
 * A n-length map containing the mutex name, and the handle object as the value
 */
var (
    syncMap = make(map[string]uintptr)
    syncObj sync.Mutex
)

func CreateMutexGlobal(name string) (uintptr, error) {
    syncObj.Lock()
    defer syncObj.Unlock()

    globalMutex := "Global\\\\" + name
    ret, _, err := procCreateMutex.Call(0, 0,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(globalMutex))),
    )
    switch int(err.(syscall.Errno)) {
    case 0:
        syncMap[name] = ret
        return ret, nil
    default:
        return ret, err
    }
}

func CloseGlobalMutex(name string) error {
    syncObj.Lock()
    defer syncObj.Unlock()

    var mutexHandle uintptr = syncMap[name]
    if mutexHandle == 0 {
        return RetErrStr("no such object (" + name + ") exists")
    }

    /* Clean handle from local namespace */
    defer func () {
        delete(syncMap, name)
    } ()

    ret, _, err := procCloseHandle.Call(0, 0,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
    )

    switch int(err.(syscall.Errno)) {
    case 0:
        return nil
    default:
        return RetErrStr("CloseHandle() returned value: " + strconv.Itoa(int(ret)))
    }
}