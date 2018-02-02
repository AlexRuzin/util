// +build linux, +build ignore
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
    "errors"
    "os"
    "syscall"
    "sync"
)

const ROOT_MUTEX_DIRECTORY = "/tmp/"

/*
 * A n-length map containing the mutex name, and the handle object as the value
 */
var (
    syncMap = make(map[string]uintptr)
    syncObj sync.Mutex
)

/*
 * These methods are not stable as of yet. Do not use until completed
 */
func CreateMutexGlobal(name string) (uintptr, error) {
    syncObj.Lock()
    defer syncObj.Unlock()

    mutexName := genMutexName(name)

    var fd_lock = syscall.Open(mutexName, syscall.O_CREAT)
    if fd_lock == -1 {
        return 0, RetErrStr("Failed to obtain mutex lock: " + mutexName)
    }

    syscall.Flock(fd_lock, syscall.LOCK_EX)

    syncMap[name] = fd_lock

    return 0, nil
}

func CloseGlobalMutex(name string) error {
    syncObj.Lock()
    defer syncObj.Unlock()

    mutexName := genMutexName(name)
    
    syscall.Flock(fd_lock, syscall.LOCK_UN)
    syscall.Close(syncMap[name])

    delete(syncMap[name])
}

func genMutexName(name string) string {
    hostname, _ := os.Hostname()
    return ROOT_MUTEX_DIRECTORY + hostname + "_" + name + "_util.mutex"
}