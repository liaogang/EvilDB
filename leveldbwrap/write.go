package leveldbwrap

func (slf *Wrap) WriteString(key string, val string) {
	slf.writeString(key, val)
}

func (slf *Wrap) WriteData(key string, val []byte) {
	slf.writeData(key, val)
}

func (slf *Wrap) WriteAny(key string, val any) {
	slf.writeAny(key, val)
}

type WrapMap struct {
	w *Wrap
}

func (slf *WrapMap) WriteKeyVal(key string, val any) {

}

func (slf *WrapMap) DeleteItem(key string, val any) {

}

func (slf *WrapMap) ValOfKey(key string) {

}
