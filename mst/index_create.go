package mst

import (
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
)

type File struct {
	Name string
	Keys []string
}

func CreateIndex(f1 File, blockNumber uint, in *Inverted_list, t *MST) *MST {
	key_file_list1 := fileToKey(f1.Name, f1.Keys)
	index := Inverted_list{Db: in.Db}
	for i := range key_file_list1.list {
		index.insert(key_file_list1.list[i])
	}
	index.list_sort()
	f1.keysSort(in.Db)
	t.root_insert(index_info{key: f1.Keys, pos: blockNumber})
	return t
}

func (file1 *File) keysSort(db ethdb.Database) {
	var key_file_pre key_file
	var key_file_next key_file
	for j := len(file1.Keys); j > 0; j-- {
		for i := 0;i < j - 1;i++ {
			data, _ := db.Get([]byte(file1.Keys[i]))
            key_file_pre.Key = file1.Keys[i]
			rlp.DecodeBytes(data,&key_file_pre.File_list)
			key_file_next.Key = file1.Keys[i]
			rlp.DecodeBytes(data,&key_file_next.File_list)
            if len(key_file_pre.Key) < len(key_file_next.Key) {
            	file1.swap(i, i + 1)
			}
		}
	}
}

func (file *File) swap(i, j int) {
	flag := file.Keys[i]
	file.Keys[j] = file.Keys[i]
	file.Keys[i] = flag
}
