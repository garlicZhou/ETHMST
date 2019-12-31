# RPC测试

geth开启RPC测试：

```
./geth --datadir data0 --networkid 1108 --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcapi admin,miner,db,eth,net,web3,personal 2>geth.log
```

关键词查询：

```
{"jsonrpc":"2.0","method":"eth_search","params":[["b"]],"id":11}
```

发送交易：

```
{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{
  "from": "0xa9cf6de0905c2bd58f48e0b13a73b89eef95e0e7",
  "to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
  "value": "0x988090",
  "filename":"haha",
  "key":["a","b"]
}],"id":1}
```

