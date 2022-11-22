package EvilDB

import "testing"

func Test_levelDB_all_key_val2(t *testing.T) {

	//var s = "/Users/apple/GoProject/ww_lite/level_db"
	//g_leveldb_root_dir = &s
	//
	//core.SharedUserMgr().Setup()
	//
	//var wrap, _ = core.SharedUserMgr().dbForUser("w4H4aWRus7wqtGTDG3wrKc")
	//
	//var db = wrap.db
	//
	//AllKeyVal(db.inner)
}

func Test_levelDB_all_key_val(t *testing.T) {

	//var s = "/Users/apple/GoProject/ww_lite/level_db"
	//g_leveldb_root_dir = &s
	//
	//core.SharedUserMgr().Setup()
	//
	//var wrap, _ = core.SharedUserMgr().dbForUser("w4H4aWRus7wqtGTDG3wrKc")
	//
	//var db = wrap.db
	//
	//AllKeyVal(db.inner)
}

func Test_EmptyVal(t *testing.T) {

	var wrap, err = NewDB("./.db")
	if err != nil {

	} else {
		wrap.writeData("test_empty", nil)

		//var f = wrap.haveDir("test_empty")
		//println(f)
	}

}
