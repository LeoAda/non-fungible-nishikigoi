<h3 align="center">non-fungible-nishikigoi</h3>
  <p align="center">
    <br />
    <a href="https://github.com/LeoAda/non-fungible-nishikigoi"><strong>Back to the project »</strong></a>
    <br />
    <br />
  </p>
</div>

## Documentation

### Structure

The project is divided into three package :
* [`blockchain`](pkg/blockchain/README.md) - the core of the project, containing the smart contract code and the blockchain code.
* [`security`](pkg/security/README.md) - the security package, containing the cryptographic code.
* [`storage`](pkg/storage/README.md) - the storage package, containing the storage code. Such as the interface and implementation with BoldDb.

All of this is used in the [`cli`](cmd/cli/README.md).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Naming

* All getters have name of the variable with a capital letter