package EvilDB

import (
	"github.com/syndtr/goleveldb/leveldb"
)

const DBWriteKVDebug = false
const DBReadKVDebug = false

/*
leveldb 的包装
*/

//goland:noinspection GoSnakeCaseUsage
type DB struct {
	inner *leveldb.DB

	//db folder in disk
	folder_in_disk string

	//visual key path of db
	path string

	updates *leveldb.Batch

	arrCache map[string]*Array
}

func NewDB(path string) (*DB, error) {
	return newWrap(path)
}

func (slf *DB) GetInner() *leveldb.DB {
	return slf.inner
}

func (slf *DB) SubDB(folder string) *DB {
	return slf.subWrap(folder)
}

func (slf *DB) DeleteItem(key string) {
	slf.deleteItem(key)
}

func (slf *DB) BeginUpdate() {
	slf.beginUpdate()
}

func (slf *DB) ApplyUpdate() error {
	return slf.applyUpdate()
}

func (slf *DB) HasDir(dir string) bool {
	return slf.hasDir(dir)
}
