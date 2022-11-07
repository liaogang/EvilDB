package leveldbwrap

import (
	"bytes"
	"encoding/gob"
	"github.com/syndtr/goleveldb/leveldb"
	"path/filepath"
	"strings"
)

func newWrap(path string) (*Wrap, error) {

	var wrap = new(Wrap)

	var db, err = leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	} else {
		wrap.inner = db
		wrap.folder_in_disk = path
		wrap.path = "/"
		return wrap, nil
	}

}

func (slf *Wrap) subWrap(sub string) *Wrap {

	if strings.Contains(sub, "/") {
		panic("folder name should not contains `/`")
	}

	var wrap = new(Wrap)

	wrap.inner = slf.inner
	wrap.path = filepath.Join(slf.path, sub)

	//create dir
	slf.inner.Put([]byte(wrap.path), nil, nil)
	//create dir

	return wrap
}

func (slf *Wrap) hasDir(dir string) bool {
	var path = filepath.Join(slf.path, dir)

	var _, err = slf.inner.Get([]byte(path), nil)
	if err != nil {
		return false
	} else {
		return true
	}

}

func (slf *Wrap) writeAny(key string, val any) {

	if val == nil {
		panic("value of key is nil")
		return
	}

	var buf = new(bytes.Buffer)
	var encoder = gob.NewEncoder(buf)
	encoder.Encode(val)

	slf.WriteData(key, buf.Bytes())
}

func (slf *Wrap) writeString(key string, val string) {
	slf.WriteData(key, []byte(val))
}

func (slf *Wrap) writeData(key string, val []byte) {

	if DBWriteKVDebug {
		println("levelDbWrap write data, key", key, " val: ", val)
	}

	if slf.updates != nil {
		slf.updates.Put(slf.fullKeyPath(key), val)
	} else {
		slf.inner.Put(slf.fullKeyPath(key), val, nil)
	}

}

func (slf *Wrap) stringOfKey(key string) string {
	var fullKey = slf.fullKeyPath(key)
	var val, err = slf.inner.Get(fullKey, nil)
	if err != nil {
		println(err)
		return ""
	} else {
		return string(val)
	}
}

func (slf *Wrap) dataOfKey(key string) []byte {

	var fullKey = slf.fullKeyPath(key)
	var val, err = slf.inner.Get(fullKey, nil)
	if err != nil {
		println(err)
		return nil
	} else {
		return val
	}
}

func (slf *Wrap) int64OfKey(key string) int64 {
	return levelDbWrap_objectOfKey[int64](slf, key)
}

func (slf *Wrap) boolOfKey(key string) bool {
	return levelDbWrap_objectOfKey[bool](slf, key)
}

func (slf *Wrap) uint64OfKey(key string) uint64 {
	return levelDbWrap_objectOfKey[uint64](slf, key)
}

func (slf *Wrap) uint32OfKey(key string) uint32 {
	return levelDbWrap_objectOfKey[uint32](slf, key)
}

func (slf *Wrap) stringArrayOfKey(key string) []string {
	return levelDbWrap_objectOfKey[[]string](slf, key)
}

func (slf *Wrap) deleteItem(key string) {
	if slf.updates != nil {
		slf.updates.Delete(slf.fullKeyPath(key))
	} else {
		slf.inner.Delete(slf.fullKeyPath(key), nil)
	}
}

func (slf *Wrap) beginUpdate() {
	if slf.updates != nil {
		panic("last update not completed!")
	} else {
		slf.updates = new(leveldb.Batch)
	}
}

func (slf *Wrap) applyUpdate() error {
	var err = slf.inner.Write(slf.updates, nil)
	slf.updates = nil
	return err
}

func (slf *Wrap) fullKeyPath(leaf string) []byte {
	return []byte(filepath.Join(slf.path, leaf))
}
