package lib

import (
	"cake-syrup-pools-lover/util"
	"fmt"
	"github.com/dgraph-io/badger"
	"strconv"
)

//type DB struct {
//	Database *badger.DB
//}

// @title 数据库初始化
//func InitDb() {
//	AllConfig := GetConfig()
//
//
//}

//func init() {
//	var Db *badger.DB = GetDbIns()
//	fmt.Println(Db)
//}

func GetDbIns() *badger.DB {
	AllConfig := GetConfig()

	// 先检查是否有数据库存在，如果存在将会给出警告并退出
	//if util.FileExists(AllConfig.BcFile) {
	//	fmt.Println("BcFile already exists")
	//	runtime.Goexit()
	//}

	// opts就是启动Badger的配置
	// 全部使用默认配置，并将地址指定为constcoe.BCPath
	opts := badger.DefaultOptions(AllConfig.BcPath)
	// 使数据库的操作信息不输出到标准输出中
	opts.Logger = nil
	// 按照我们的配置启动一个数据库（如果没有现成的数据库就会初始化一个）,将会返回该数据库的指针
	db, err := badger.Open(opts)
	util.Handle(err)
	return db
}

// @title 存储TgUser的Id
func SaveChatId(userId int) {
	var db = GetDbIns()
	defer db.Close()
	userIdStr := strconv.Itoa(userId)
	// 对Badger数据库进行更新操作
	err := db.Update(func(txn *badger.Txn) error {
		fmt.Printf("Save User Id: %d \n", userId)
		// Key值:userIdStr，Value:userIdStr
		err := txn.Set([]byte(userIdStr), []byte(userIdStr))
		return err
	})
	util.Handle(err)
}

// @title 存储SyrupPool信息
// @param key string "键值pools/pools_shortly"
// @param poolMsg string "通知信息"
func SaveSyrupPoolStr(key string, poolMsg string) {
	fmt.Println("[*] Try to update New SyrupPool Msg")
	var db = GetDbIns()
	defer db.Close()
	// 对Badger数据库进行更新操作
	err := db.Update(func(txn *badger.Txn) error {
		// Key值:固定字符串 pools，Value:poolMsg
		err := txn.Set([]byte(key), []byte(poolMsg))
		return err
	})
	util.Handle(err)
}

// @title 查询SyrupPool信息
// @param key string "键值pools/pools_shortly"
// @return poolMsg string "池子信息"
func QuerySyrupPoolStr(key string) string {
	var db = GetDbIns()
	defer db.Close()
	var poolMsg string
	// get
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		util.Handle(err)
		err = item.Value(func(val []byte) error {
			poolMsg = string(val)
			fmt.Printf("The pools is: %s \n", poolMsg)
			return nil
		})
		util.Handle(err)
		return err
	})
	util.Handle(err)
	return poolMsg
}

// @title 查询TgUser的Id
// @param userId int "键值"
// @return uid int "用户id"
func QueryChatId(userId int) int {
	var db = GetDbIns()
	defer db.Close()

	userIdStr := strconv.Itoa(userId)
	var userIdbyte []byte
	var uid int
	// get
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(userIdStr))
		util.Handle(err)
		err = item.Value(func(val []byte) error {
			userIdbyte = val
			uid, _ = strconv.Atoi(string(userIdbyte))
			fmt.Printf("The userIdbyte is: %d\n", uid)
			return nil
		})
		util.Handle(err)
		return err
	})
	util.Handle(err)
	return uid
}

// @title 查询全部TgUser的Id
func QueryAllChatId() []int {
	var db = GetDbIns()
	defer db.Close()

	var uids []int
	// get
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			uid, _ := strconv.Atoi(string(k))
			uids = append(uids, uid)
			fmt.Printf("key= %s\n", k)
		}
		return nil
	})
	util.Handle(err)
	return uids
}
