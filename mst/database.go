package mst

import (
	"fmt"
	//"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
)

type kv struct {
	Key []byte
	Value []byte
}

type nodeData struct {
	Hash string
	ParentHash string
	ChildHash  string
	Key        []string
	Value      []uint
	IsLeaf     bool
	IsExtend   bool
}

type MSTIiterator struct {
	db *leveldb.DB
	currentHash []byte
}

func put()  {
	db, err := leveldb.OpenFile("path/to/db", nil)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	kv1 := kv{Key: []byte("jac"), Value:[]byte("ca")}
	e := db.Put(kv1.Key, kv1.Value,nil)
	fmt.Println(e)
}

func get() {
	db, err := leveldb.OpenFile("path/to/db", nil)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	data, _ := db.Get([]byte("jac"), nil)
	fmt.Println(data)
	fmt.Printf("%c\n",data)
}

//func putMst(t *MST) {
//	//node1 := nodekv{Hash:"ha",ParentHash:"hh",ChildHash:"h1",IsExtend:true}
//	db, err := leveldb.OpenFile("path/to/db", nil)
//	defer db.Close()
//	if err != nil {
//		fmt.Println(err)
//	}
//	arrdata, err := rlp.EncodeToBytes(node1)
//	err = db.Put([]byte(node1.Hash),arrdata,nil)
//	data, _ := db.Get([]byte(node1.Hash), nil)
//	data1 := nodekv{}
//	rlp.DecodeBytes(data,&data1)
//	fmt.Println(data1)
//}

//func (i *MSTIiterator) getMST(db *leveldb.DB) *MST{
//
//}




