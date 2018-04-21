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
    "sync"
)

type QueueObject struct {
    count               uint64
    elements            []interface{}
    syncObj             sync.Mutex
}

/*
 * References a certain element in the array
 */
func (f *QueueObject) Index(c int) interface{} {
    if uint64(c) >= f.count {
        panic(RetErrStr("Queue: Invalid index"))
    }

    return f.elements[c]
}

/*
 * Return a read-only array of the queue
 */
func (f *QueueObject) Array() []interface{} {
    return f.elements
}

/*
 * Get the length of the elements
 */
func (f *QueueObject) Len() int {
    f.syncObj.Lock()
    defer f.syncObj.Unlock()

    return int(f.count)
}

/*
 * Input: Object that needs to be pushed to the Queue
 * Output: Number of objects in the array
 */
func (f *QueueObject) Push(p interface{}) int {
    f.syncObj.Lock()
    defer f.syncObj.Unlock()

    f.elements = append(f.elements, p)
    f.count += 1

    if int(f.count) != len(f.elements) {
        panic(RetErrStr("Queue.Push(): Critical indexing error"))
    }

    return len(f.elements)
}

func (f *QueueObject) Pop() (interface{}) {
    f.syncObj.Lock()
    defer f.syncObj.Unlock()

    ref := &f.elements[0]
    f.elements[0] = nil
    f.elements = f.elements[1:]

    f.count -= 1
    return ref
}

func NewQueue(load ...interface{}) (Queue *QueueObject) {
    var output = &QueueObject{}

    for _, v := range load {
        output.Push(v)
    }

    return output
}

func (f *QueueObject) CloseQueue() {
    f.syncObj.Lock()
    defer f.syncObj.Unlock()

    for k, _ := range f.elements {
        f.elements[k] = nil
    }

    f.count = 0
    f.elements = nil

    return
}