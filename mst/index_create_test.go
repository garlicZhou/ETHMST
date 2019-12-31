package mst

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

func TestIndexCreate(T *testing.T) {
	//db_invert, _ := leveldb.OpenFile("path/to/db_invert", nil)
	//defer db_invert.Close()
	//db_mst, _ := leveldb.OpenFile("path/to/db_mst", nil)
	//defer db_mst.Close()
	//index := Inverted_list{Db: db_invert}
	////index.RenewList()
	//fmt.Println(index)
	//mst := MST{RootHash: [32]byte{}, Db:db_mst}
	//mst.ReNewMst()
	//mst.printMst()
	//CreateIndex(File{Name: "tom",Keys: []string{"block","chain"}},88,&index,&mst)
	//mst.printMst()
	//fmt.Println(mst.search([]string{"篮球"}))
	db_invert, _ := leveldb.OpenFile("path/to/db_invert", nil)
	defer db_invert.Close()
	db_mst, _ := leveldb.OpenFile("path/to/db_mst", nil)
	defer db_mst.Close()
	index := Inverted_list{Db:db_invert}
	index.RenewList()
	preMst := MST{RootHash:[common.HashLength]byte{114, 95, 45, 150, 211, 241, 92, 46, 82, 253, 99, 89, 17, 239, 89, 26, 8, 38, 22, 88, 77, 89, 187, 235, 43, 123, 216, 38, 88, 2, 150, 53},Db:db_mst}
	preMst.ReNewMst()
	CreateIndex(File{Name: "h",Keys: []string{"a","g","m"}},7,&index,&preMst)
	CreateIndex(File{Name: "h",Keys: []string{"bb","g","m"}},7,&index,&preMst)
	CreateIndex(File{Name: "h",Keys: []string{"af","c"}},7,&index,&preMst)
	CreateIndex(File{Name: "h",Keys: []string{"a","g","f","w"}},7,&index,&preMst)
	preMst.printMst()
	preMst.PutRootHash()
	fmt.Println(preMst.RootHash)
}