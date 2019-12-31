package mst

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
	"time"
)


func TestFileToKey (t *testing.T) {
	testCount := 10000
	start := time.Now()
	var list = []string{"a"}
	for i := testCount; i > 0; i-- {
		fileToKey("name", list)
		//fmt.Println(key_file_list1)
	}
	fmt.Println(time.Now().Sub(start))
}

func TestInsert(t *testing.T) {
	db, _ := leveldb.OpenFile("path/to/db_invert", nil)
	defer db.Close()
	index := Inverted_list{Db: db}
	start := time.Now()
	key_file1 := key_file{"tom", []string{"c"}}
	index.insert(key_file1)
	index.insert(key_file{"jack", []string{"a", "b"}})
	index.insert(key_file{"chris", []string{"x", "y", "z", "o"}})
	index.insert(key_file{"peter", []string{"b", "c", "d"}})
	index.insert(key_file{"alice", []string{"g", "u", "e", "o","s","p"}})
	index.insert(key_file{"tony", []string{"c", "f", "g"}})
	index.insert(key_file{"jack", []string{"a"}})
	index.insert(key_file{"chris", []string{"x", "y", "z", "o"}})
	index.k = len(index.list)
	fmt.Println(index)
	index.list_sort()
	fmt.Println(index)
	fmt.Println(time.Now().Sub(start))
}

func TestSearch(t *testing.T) {
	db, _ := leveldb.OpenFile("path/to/db_invert", nil)
	defer db.Close()
	index := Inverted_list{Db: db}
	key_file1 := index.searchKey("alice")
	fmt.Println(key_file1)

	index.RenewList()
	fmt.Println(index)
}
