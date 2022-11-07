package leveldbwrap

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func allKeysWithPrefix(w *Wrap, prefix string) []string {
	var list []string
	var iter = w.inner.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	for iter.Next() {
		var key = iter.Key()
		var s = string(key)
		list = append(list, s)
	}
	iter.Release()

	return list
}

func AllKeys(w *Wrap) []string {
	var list []string
	var iter = w.inner.NewIterator(nil, nil)
	for iter.Next() {
		var key = iter.Key()
		var s = string(key)
		list = append(list, s)
	}
	iter.Release()

	return list
}

func AllKeyVal(db *leveldb.DB) {
	var iter = db.NewIterator(nil, nil)
	for iter.Next() {
		var key = iter.Key()
		var s = string(key)

		var val = iter.Value()

		print("key: " + s + " val: ")
		println(hex.EncodeToString(val))

	}
	iter.Release()
}

func gobDecodeDataToTheType[T int64 | uint64 | uint32 | bool | []string](data []byte) T {
	var buf = bytes.NewBuffer(data)
	var decoder = gob.NewDecoder(buf)

	var ret T
	var err = decoder.Decode(&ret)
	if err != nil {
		print("decode data to type error", data, ret)
		panic(err)
	} else {
		return ret
	}
}

//goland:noinspection ALL
func levelDbWrap_objectOfKey[T int64 | uint64 | uint32 | bool | []string](w *Wrap, key string) T {

	var fullKey = w.fullKeyPath(key)
	var val, err = w.inner.Get(fullKey, nil)

	if err != nil {

		if DBReadKVDebug {
			println("levelDbWrap_objectOfKey not exist, key: "+string(fullKey), err)
		}

		var ret T
		return ret
	} else {
		return gobDecodeDataToTheType[T](val)
	}
}
