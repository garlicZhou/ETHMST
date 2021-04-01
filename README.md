# MSTDB

Mstdb is a semantic blockchain system based on go-ethereum(v1.8).

## Requirements

On Linux ( > Ubuntu 16.04)

Go 1.12 or later

a C compiler

## Building the source

Before the source code is compiled, rename the ETHMST to github.com/ethereum/go-ethereum

```
mv ETHMST/ github.com/ethereum/go-ethereum
```

Once the dependencies are installed, run

```
cd github.com/ethereum/go-ethereum
make geth
cd build/bin
```

## Running `geth` 

Create a file named genesis.json and write the following content

```
{
   "config": {
        "chainId": 15,
        "homesteadBlock": 0,
        "eip150Block": 0,
        "eip155Block": 0
    },
    "coinbase" : "0x0000000000000000000000000000000000000000",
    "difficulty" : "0x00400",
    "extraData" : "",
    "gasLimit" : "0xffffffff",
    "nonce" : "0x0000000000000042",
    "mixhash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
    "parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
    "timestamp" : "0x00",
    "alloc": { }
}
```

Initialize the genesis block

```
./geth --datadir data0 init genesis.json
```

Start a private chain and enter the geth console

```
./geth --datadir data0 --networkid 1108 console 
```

or write the log information into a file

```
./geth --datadir data0 --networkid 1108 console 2>geth.log
```

Create a new account:

```
personal.newAccount('123') 
//123 is the password of the account
```

Start mining

```
miner.start()

//if you want to mine only one block
miner.start(1);admin.sleepBlocks(1);miner.stop() 
```

Before send the transaction, the account must be unlocked

```
personal.unlockAccount(eth.accounts[0])
```

Send a transaction, while the filename and key[] are the meta-data from off-chain data

```
eth.sendTransaction({from:eth.accounts[0],to:"0xc9194a8ea28d76389af4c7e9c81222386a6ab47a",value:1000000,filename:"math",key:["a","b","c"]})
```

Search a block by key words

```
eth.search(["d","c"])
```

Other console functions please search in https://web3js.readthedocs.io/en/v1.2.0/

## Running `geth` with rpc model

Start a private chain and add rpc parameter

```
./geth --datadir data0 --networkid 1108 --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcapi admin,miner,db,eth,net,web3,personal 2>geth.log
```

If you want to use the console at the same time, add the console parameter

```
./geth --datadir data0 --networkid 1108 --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcapi admin,miner,db,eth,net,web3,personal console 2>geth.log
```

If the console stops you sending a transaction through http, add the unlock parameter

```
 ./geth --datadir data0 --networkid 1108 --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcapi admin,miner,db,eth,net,web3,personal console --allow-insecure-unlock 2>geth.log
```

Send a transaction

```
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{
  "from": "0xa9cf6de0905c2bd58f48e0b13a73b89eef95e0e7",
  "to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
  "value": "0x988090",
  "filename":"haha",
  "key":["a","b"]
}],"id":1}'
```

Search a block by key words

```
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_search","params":[["b"]],"id":11}'
```

Other json-rpc functions please search in https://eth.wiki/json-rpc/API