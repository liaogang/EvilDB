package leveldbwrap

import (
	"strconv"
	"sync"
)

func (slf *Wrap) ArrayOfName(name string) *Array {
	var arr = new(Array)

	var folderName = ".array_" + name

	var db = slf.subDBWithSubFolder(folderName)

	arr.inner = db
	arr.name = name

	return arr
}

/*
.array_name/

.array_name/size -> size

.array_name/0 -> item0
.array_name/1 -> item1
.
.
.
.array_name/last -> itemE

last = size - 1
*/

type Array struct {
	inner *Wrap
	name  string

	lastIndex int

	lock *sync.Mutex
}

func (slf *Array) Append(item any) {

	slf.lock.Lock()

	var index = slf.lastIndex + 1
	var sIndex = strconv.Itoa(index)
	slf.inner.writeAny(sIndex, item)
	slf.lastIndex = index

	slf.lock.Unlock()
}

func (slf *Array) Size() int {
	return slf.lastIndex + 1
}

func (slf *Array) DeleteLastItem() {
}
