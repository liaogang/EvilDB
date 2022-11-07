package leveldbwrap

import (
	"strconv"
	"sync"
)

func (slf *Wrap) ArrayOfName(name string) (*Array, error) {

	//check from cache
	if slf.arrCache == nil {
		slf.arrCache = make(map[string]*Array, 0)
	}

	if arr, ok := slf.arrCache[name]; ok {
		return arr, nil
	}

	//create
	var arr = new(Array)

	var folderName = ".array_" + name

	var db = slf.SubWrap(folderName)

	arr.inner = db
	arr.name = name

	slf.arrCache[name] = arr

	arr.LoadAllMember()

	return arr, nil
}

/*
.array_name/

.array_name/LOCK 标记正在使用

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

	cache []string

	lastIndex int

	lock *sync.Mutex
}

func (slf *Array) LoadAllMember() {
	//load all array items
	var db = slf.inner

	//first load size
	var size = int(db.Uint32OfKey("size"))
	slf.cache = make([]string, 0)
	for i := 0; i < size; i++ {
		var indexStr = strconv.Itoa(i)
		slf.cache = append(slf.cache, db.stringOfKey(indexStr))
	}
	slf.lastIndex = size - 1
}

func (slf *Array) Append(item any) {

	if slf.lock == nil {
		slf.lock = new(sync.Mutex)
	}

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

func (slf *Array) Item(index int) string {
	return slf.cache[index]
}

func (slf *Array) DeleteLastItem() {
}
