package leveldbwrap

func (slf *Wrap) StringOfKey(key string) string {
	return slf.stringOfKey(key)
}

func (slf *Wrap) DataOfKey(key string) []byte {
	return slf.dataOfKey(key)
}

func (slf *Wrap) Int64OfKey(key string) int64 {
	return slf.int64OfKey(key)
}

func (slf *Wrap) BoolOfKey(key string) bool {
	return slf.boolOfKey(key)
}

func (slf *Wrap) Uint64OfKey(key string) uint64 {
	return slf.uint64OfKey(key)
}

func (slf *Wrap) Uint32OfKey(key string) uint32 {
	return slf.uint32OfKey(key)
}

func (slf *Wrap) StringArrayOfKey(key string) []string {
	return slf.stringArrayOfKey(key)
}
