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
    "errors"

    /* Preference for the gunzip compressor, although using zlib follows the same I/O API */
    "compress/gzip"
    _"compress/zlib"
)

/* Use gzip compression to generate a compressed stream */
func CompressStream(p []byte) ([]byte, error) {
    var in bytes.Buffer

    gz := gzip.NewWriter(&in)
    defer gz.Close()

    if _, err := gz.Write(p); err != nil {
        return nil, err
    }

    gz.Flush()

    return in.Bytes(), nil
}

/* Decompression subroutine */
func DecompressStream(p []byte) ([]byte, error) {
    in := bytes.NewReader(p)

    gz, err := gzip.NewReader(in)
    if err != nil {
        return nil, err
    }

    defer gz.Close()

    decompressed, _ := ioutil.ReadAll(gz)
    if len(decompressed) == 0 {
        return nil, errors.New("decompressed stream is 0 bytes")
    }

    return decompressed, nil
}

func GetCompressedSize(p []byte) int {
    var (
        err         error = nil
        compressed  []byte
    )

    compressed, err = CompressStream(p)
    if err != nil {
        panic(err)
    }

    return len(compressed)
}
