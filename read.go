package EvilDB

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
