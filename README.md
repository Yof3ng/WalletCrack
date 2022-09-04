# WalletCrack
Generate eth wallets with a large number of mnemonics and get the balance, golang implementation.

`metamask`
`tokenPocket`
`Trust wallet`

## Usage
```bash
Usage of ./getit:
  -api string
    	API string, [https://mainnet.infura.io/v3/key] or [Local Node].
  -st string
    	Wallet address, query balance of target account.
  -t int
    	Thread number, default 10. (default 10)
```

## ScreenShot
Query account balance.

```bash
➜  metatest ./getit -h
Usage of ./getit:
  -api string
    	API string, [https://mainnet.infura.io/v3/key] or [Local Node].
  -st string
    	Wallet address, query balance of target account.
  -t int
    	Thread number, default 10. (default 10)
➜  metatest ./getit -api http://127.0.0.1:8545 -st 0x4b229002a83923ee10ba9228389426efa15b7e09
[0x4b229002a83923ee10ba9228389426efa15b7e09] => Balance:0.4573155683
```

Generate wallet and query balance of it.
```bash
➜  metatest ./getit -api http://127.0.0.1:8545 -t 3
start...
mnemonic: athlete course street brother result sort ritual become record marble junior elevator
mnemonic: crunch reason tomato dragon vague pilot current stage dry similar version mixed
mnemonic: tiny you danger patrol balance divide setup nominee crisp monkey divert enlist
[0x82fDb7071a68158b3F9258059A2641702bD0E948] => Balance:0
[0x6C322B60A75c86A478CA20119D8579B051dc09A7] => Balance:0
[0x6b7e7A87A64E58D90a702748200E6b3Ab7099460] => Balance:0
[0x813d3EDD5AbBa6195889dd1D9CB3BA8e5adae5c8] => Balance:0
[0x2bDB394498fb7a476f483626991438e01099B9A5] => Balance:0
[0x84DC10AC0F5475158b303733a8104DA2af31a512] => Balance:0
```


