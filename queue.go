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

)

type QueueObject struct {
    count               uint64
    elements            []interface{}
}

/*
 * Input: Object that needs to be pushed to the Queue
 * Output: Number of objects in the array
 */
func (f *QueueObject) Push(p interface{}) int {


    return len(f.elements)
}

func (f *QueueObject) Pop() interface{} {
    var topObject = f.elements[0]
    f.elements = f.elements[1:]

    return topObject
}

func BuildQueue(load ...interface{}) (Queue *QueueObject) {
    var output = &QueueObject{}

    for _, v := range load {
        output.Push(v)
    }

    return output
}

func DestroyQueue(Queue *QueueObject) {
    for k, _ := range Queue.elements {
        Queue.elements[k] = nil
    }

    Queue.count = 0
    Queue.elements = nil

    return
}