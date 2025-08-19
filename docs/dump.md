# dumping ideas

build the chain indexer.

## What is the type of Tf Chain

- evm/non-evm is the protocol/backbone of most of blockchains.

evm

- uses eth virtual machine, specific environment to run contracts in solidity,
- follow eth consensus
- eth is not just a blockchain, its evm protocol become popular in a way makes
  alot of chains trying to be compitable with it.
- you write a smart contract in soladity> compiled to bytecode with evm> run on eth blockchain.
- other chain support evm could run the same smart contracts

non-evm

- not specifc environtment, can run any language
- can be different consensus, Pos
- scallable
- chians not support the eth environtment, it usally be in different probramming launage for
  the smart contract, and runtime.
- it is more flexable in the decision and have a better performance.
- it allow diversity in the writen programming lang
- is is called substrate-based or polkdadot chain.

the chain is a substrate-based.

## How to indexer the chain?

- evm chains has the graph
- for the substrate based chains there is
  - [hydra](https://github.com/Joystream/hydra): inspired by the graph, former to subsquid
  - subsquid: built on top of the archive
  - [substrate-archive](https://github.com/paritytech/substrate-archive): top performar, from partiy, not mintaned
  - substrate-graph: looks not stabel
  - subquery
  - the graphq also was trieng to make indexer for this type
- [tools](https://wiki.polkadot.network/docs/build-data)

## which indexer will use, why? and how?
