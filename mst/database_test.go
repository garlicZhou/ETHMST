package mst

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"testing"
)

func TestDatabase(t *testing.T) {
	//db, err := leveldb.OpenFile("path/to/db", nil)
	//defer db.Close()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//i := db.Delete([]byte("jack"), nil)
	//fmt.Println(i)//<nil>
	//data, _ := db.Get([]byte("jack"), nil)
	//fmt.Println(data)
    put()
	//get()
}

func TestRLPEncode(t *testing.T) {
	//MST1 := MST{}
	//putMst(&MST1)

}

func TestRLPDecode(t *testing.T) {
	arrdata, err := rlp.EncodeToBytes(nodekv{Key: []string{"tom","jack","hah"}})
	fmt.Printf("unuse err:%v\n", err)
	fmt.Println(arrdata)
}

