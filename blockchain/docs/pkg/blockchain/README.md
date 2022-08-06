<h3 align="center">non-fungible-nishikigoi</h3>
  <p align="center">
    <br />
    <a href="https://github.com/LeoAda/non-fungible-nishikigoi/tree/master/blockchain/docs"><strong>Back to the documentation »</strong></a>
    <br />
    <a href="https://github.com/LeoAda/non-fungible-nishikigoi"><strong>Back to the project »</strong></a>
    <br />
  </p>
</div>

## Blockchain documentation

### Structure
The package is divided into 4 files:
* `blockchain`: contains the blockchain object code.
* `block`: contains the block object code.
* `explorer`: contains methods to explore the blockchain.
* `iterator`: contains methods to iterate over the blockchain.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Blockchain
The `blockchain` file contains the blockchain object code.
The blockchain object is a `struct` with the following fields:
* `blocks`: a `slice` of blocks.
* `index`: a counter of the last blocks.

#### Methods
The blockchain object has the following methods:
* `AddBlock`: adds a block to the blockchain.
* `AddBlockObject`: adds a block to the blockchain.
* `NewGenesisBlock`: creates a new genesis block.
* `NewBlockchain`: create a new blockchain.
* `NewBlockchainList`: create a new blockchain from an existing block list.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Block
The `block` file contains the block object code.
The block object is a `struct` with the following fields:
* `blockNumber`: the index of the block.
* `timestamp`: the time of the block creation.
* `parentHash`: the hash of the parent block.
* `data`: the data of the block.
* `blockHash`: the hash of the block.

#### Methods
The block object has the following methods:
* `NewBlock`: creates a new block from string.
* `NewBlockAll`: creates a new block from all metadata.
* `setHash`: sets the hash of the block.
* `calculateBlockHash`: calculates the hash of the block.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Explorer 
The `explorer` file contains methods to explore the blockchain.

#### Methods
The explorer object has the following methods associate to the blockchain:
* `FindBlock`: finds a block by its hash.
* `LastBlock`: returns the last block of the blockchain.
* `BlockExist`: checks if a block exists.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Iterator
The `iterator` file contains methods to iterate over the blockchain.

#### Methods
The Iterator object has the following methods :
* `Iterator`: creates a new iterator.
* `resetIterator`: resets the iterator.
* `Next`: returns the next block of the blockchain.
* `hasNext`: checks if there is a next block.

<p align="right">(<a href="#readme-top">back to top</a>)</p>