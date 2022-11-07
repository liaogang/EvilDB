package leveldbwrap

import (
	"github.com/syndtr/goleveldb/leveldb"
)

const DBWriteKVDebug = true
const DBReadKVDebug = true

/*
leveldb 的包装
*/

//goland:noinspection ALL
const G_Level_Db_Index_Key = "_g_level_db_index_key_"

//goland:noinspection GoSnakeCaseUsage
type Wrap struct {
	inner *leveldb.DB

	//db folder in disk
	folder_in_disk string

	//visual key path of db
	path string

	updates *leveldb.Batch

	arrCache map[string]*Array
}

func NewWrap(path string) (*Wrap, error) {
	return newWrap(path)
}

func (slf *Wrap) SubWrap(folder string) *Wrap {
	return slf.subWrap(folder)
}

func (slf *Wrap) DeleteItem(key string) {
	slf.deleteItem(key)
}

func (slf *Wrap) BeginUpdate() {
	slf.beginUpdate()
}

func (slf *Wrap) ApplyUpdate() error {
	return slf.applyUpdate()
}

func (slf *Wrap) HasDir(dir string) bool {
	return slf.hasDir(dir)
}
