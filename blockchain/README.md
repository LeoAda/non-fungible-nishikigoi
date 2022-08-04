<a name="readme-top"></a>
<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->


[![Go][Go-version]][Golang-url]

[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <!--<a href="https://github.com/LeoAda/non-fungible-nishikigoi">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>-->

<h3 align="center">non-fungible-nishikigoi</h3>

  <p align="center">
    <!--project_description-->
    <br />
    <a href="https://github.com/LeoAda/non-fungible-nishikigoi/tree/master/blockchain/docs"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/LeoAda/non-fungible-nishikigoi">View Demo</a>
    ·
    <a href="https://github.com/LeoAda/non-fungible-nishikigoi/issues">Report Bug</a>
    ·
    <a href="https://github.com/LeoAda/non-fungible-nishikigoi/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!--[![Product Name Screen Shot][product-screenshot]](https://example.com)-->
This project is a non-fungible token system and an image generator.
This part is the blockchain side.
It is based written in Go and is inspired by several other blockchain such as ethereum.


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Golang][Golang]][Golang-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

You need to have a Go compiler and Go runtime installed. [https://go.dev/](https://go.dev/)
If you are on windows, you may need to install "make for windows" to properly use makefile command.

### Installation

1. Install go [https://go.dev/](https://go.dev/), git, and make (if on windows).
2. Clone the repo
   ```sh
   git clone https://github.com/LeoAda/non-fungible-nishikigoi.git
   ```
3. Install Go packages (make), in /blockchain/ directory
   ```sh
   make setup
   ```
   or

   ```sh
   go get ./...
   ```
4. Run the client
   ```sh
    make run
   ```
   or

    ```sh
    go run ./cmd/cli/
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [ ] Handle error properly
- [ ] Make CLI
- [ ] Add POW


See the [open issues](https://github.com/LeoAda/non-fungible-nishikigoi/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>






<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [https://jeiwan.net/](https://jeiwan.net/)
* [https://refactoring.guru/](https://refactoring.guru/)
* []()

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/LeoAda/non-fungible-nishikigoi.svg?style=for-the-badge
[contributors-url]: https://github.com/LeoAda/non-fungible-nishikigoi/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/LeoAda/non-fungible-nishikigoi.svg?style=for-the-badge
[forks-url]: https://github.com/LeoAda/non-fungible-nishikigoi/network/members
[stars-shield]: https://img.shields.io/github/stars/LeoAda/non-fungible-nishikigoi.svg?style=for-the-badge
[stars-url]: https://github.com/LeoAda/non-fungible-nishikigoi/stargazers
[issues-shield]: https://img.shields.io/github/issues/LeoAda/non-fungible-nishikigoi.svg?style=for-the-badge
[issues-url]: https://github.com/LeoAda/non-fungible-nishikigoi/issues
[license-shield]: https://img.shields.io/github/license/LeoAda/non-fungible-nishikigoi.svg?style=for-the-badge
[license-url]: https://github.com/LeoAda/non-fungible-nishikigoi/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/leo-ada
[product-screenshot]: images/screenshot.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[Go-version]: https://img.shields.io/github/go-mod/go-version/LeoAda/non-fungible-nishikigoi?filename=blockchain%2Fgo.mod
[Golang-url]: https://golang.org/
[Golang]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Python]: https://img.shields.io/badge/Python-14354C?style=for-the-badge&logo=python&logoColor=white