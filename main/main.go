package main

import (
	ldb "github.com/liaogang/LevelDBWrap/leveldbwrap"
)

func main() {

	var wrap, err = ldb.NewWrap("./", "sharedDBForMgr")
	if err != nil {

	} else {
		wrap.WriteData("test_empty", nil)

		//var f = wrap.haveDir("test_empty")
		//println(f)
	}

}
