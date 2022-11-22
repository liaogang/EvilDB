package EvilDB

func (slf *DB) WriteString(key string, val string) {
	slf.writeString(key, val)
}

func (slf *DB) WriteData(key string, val []byte) {
	slf.writeData(key, val)
}

func (slf *DB) WriteAny(key string, val any) {
	slf.writeAny(key, val)
}
