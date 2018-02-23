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
    "compress/gzip"
    "io"
)

/* Decompression subroutine */
func DecompressStream(p []byte) ([]byte, error) {
    zippedRaw := bytes.NewReader(p)
    zipped, _ := gzip.NewReader(zippedRaw)
    defer zipped.Close()

    var decompressed []byte
    for {
        tmp := make([]byte, 100)
        read, err := zipped.Read(tmp)
        if err != io.EOF {
            return nil, err
        }
        if err == io.EOF {

        }

        decompressed = append(decompressed, tmp)
    }

    return nil, nil
}

/* Use gzip compression to generate a compressed stream */
func CompressStream(p []byte) ([]byte, error) {
    var compressedBuffer = bytes.NewBuffer(nil)
    gzipWriter := gzip.NewWriter(compressedBuffer)
    gzipWriter.Write(p)

    zipped := bytes.Buffer{}
    zipped.ReadFrom(compressedBuffer)

    rawBytes := make([]byte, zipped.Len())
    read, err := zipped.Read(rawBytes)
    if err != nil {
        return nil, err
    }

    if read == 0 {
        return nil, RetErrStr("Error in the compression of input stream")
    }

    return rawBytes, nil
}