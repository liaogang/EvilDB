package EvilDB

import (
	"bytes"
	"encoding/gob"
)

func (slf *DB) StringOfKey(key string) string {
	return slf.stringOfKey(key)
}

func (slf *DB) DataOfKey(key string) []byte {
	return slf.dataOfKey(key)
}

func (slf *DB) Int64OfKey(key string) int64 {
	return slf.int64OfKey(key)
}

func (slf *DB) BoolOfKey(key string) bool {
	return slf.boolOfKey(key)
}

func (slf *DB) Uint64OfKey(key string) uint64 {
	return slf.uint64OfKey(key)
}

func (slf *DB) Uint32OfKey(key string) uint32 {
	return slf.uint32OfKey(key)
}

func (slf *DB) StringArrayOfKey(key string) []string {
	return slf.stringArrayOfKey(key)
}

func (slf *DB) AnyOfKey(key string) (ret any, err error) {
	var w = slf
	var fullKey = w.fullKeyPath(key)

	var val []byte
	val, err = w.inner.Get(fullKey, nil)
	if err != nil {
		return nil, err
	}

	var buf = bytes.NewBuffer(val)
	var decoder = gob.NewDecoder(buf)

	err = decoder.Decode(ret)
	if err != nil {
		return nil, err
	} else {
		return ret, nil
	}

}
