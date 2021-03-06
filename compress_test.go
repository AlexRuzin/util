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
    "testing"

    "crypto/md5"
    "bytes"
    "errors"
)

func TestCompression(t *testing.T) {

    /* Generate a random length raw stream, take a md5sum */
    decompressedString := RandomString(RandInt(512, 4096))
    rawSum := md5.Sum([]byte(decompressedString))
    compressedStream, err := CompressStream([]byte(decompressedString))
    if err != nil {
        panic(err)
    }

    /* Decompress and check sum */
    newDecompressed, err := DecompressStream(compressedStream)
    if err != nil {
        panic(err)
    }
    newSum := md5.Sum(newDecompressed)

    if bytes.Compare(newSum[:], rawSum[:]) != 0 {
        panic(errors.New("Checksum failure in gzip decompression"))
    }

    return
}