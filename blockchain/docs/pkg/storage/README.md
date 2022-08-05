<h3 align="center">non-fungible-nishikigoi</h3>
  <p align="center">
    <br />
    <a href="https://github.com/LeoAda/non-fungible-nishikigoi/tree/master/blockchain/docs"><strong>Back to the documentation »</strong></a>
    <br />
    <a href="https://github.com/LeoAda/non-fungible-nishikigoi"><strong>Back to the project »</strong></a>
    <br />
  </p>
</div>

## Storage documentation

### Structure
The package is divided into 3 files:
* `db`: contains the database interface.
* `boltdb`: contains the boltdb implementation.
* `utils`: contains some useful methods.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Database
The `db` file contains the database interface.

The db is an `interface` with the following fields:
* `GetBlock`: returns a block from the database.
* `GetLastBlock`: returns the last block of the blockchain in the database.
* `GetBlockchain`: returns the blockchain from the database.
* `AddBlock`: adds a block to the database.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Boltdb
The `boltdb` file contains the boltdb implementation.
Bolt db is a key-value database that stores data in binary format.
The db is in the blockchain.db file.
It got the same methods as database but with `getDb` that returns the boltdb instance.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Utils
The `utils` file contains some useful methods.

#### Methods
The utils file has the following methods:
* `SerializeBlock`: serializes a block.
* `DeserializeBlock`: deserializes a block.
  
<p align="right">(<a href="#readme-top">back to top</a>)</p>
