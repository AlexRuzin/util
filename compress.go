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
    "bytes"
    "io/ioutil"
    "compress/zlib"
)

/* Use gzip compression to generate a compressed stream */
func CompressStream(p []byte) ([]byte, error) {
    var in bytes.Buffer

    zlib := zlib.NewWriter(&in)
    zlib.Write(p[:])
    defer zlib.Close()

    return nil, nil
}

/* Decompression subroutine */
func DecompressStream(p []byte) ([]byte, error) {
    b := bytes.NewReader(p[:])
    zlib, err := zlib.NewReader(b)
    if err != nil {
        return nil, err
    }

    defer zlib.Close()

    decompressed, err := ioutil.ReadAll(zlib)
    if err != nil {
        return nil, err
    }

    return decompressed, nil

}
