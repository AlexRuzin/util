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
    "fmt"
    _"strconv"
    "errors"
)

func Check(err error) {
    if err != nil {
        panic(err)
    }
}

func CheckR(err error) error {
    if err != nil {
        return err
    }
    return nil
}

func CheckN(err error, d string) {
    if err != nil {
        panic(d)
    }
}

/*
 * Throw a panic
 */
func ThrowN(d string) {
    panic(d)
}

/*
 * Standard debug method
 */
func DebugOut(debug string) {
    fmt.Println(debug)
}

/*
 * Prints hex output
 */
func DebugOutHex(debug []byte) {
    fmt.Printf("%v\r\n", debug)
}

/*
 * Returns a new error object with the specified prefix
 */
func RetErrStr(text string) (err error) {
    err = errors.New("RetErrStr(): " + text)
    return
}

/*
 * Displays the method name of the caller method
 */
