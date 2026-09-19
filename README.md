<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# LeetCodeSolutions

Three independent Go modules, one per solved LeetCode problem (21, 27, 2129), each with its own go.mod and a small main that exercises the solution.

**English** · [简体中文](README.zh-CN.md)

[![CI](https://github.com/anyingiit/LeetCodeSolutions/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/LeetCodeSolutions/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/LeetCodeSolutions)](LICENSE)

[Report a bug](https://github.com/anyingiit/LeetCodeSolutions/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/LeetCodeSolutions/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

This repository collects personal solutions to individual LeetCode problems, each isolated in its own directory as a self-contained Go module: `problem21/go/problem21.go` merges two sorted linked lists, `problem27/go/problem27.go` removes a value from a slice in place, and `problem2129/go/problem2129.go` capitalizes the words of a title. Each solution ships with its own `go.mod`, a `main.go` entry point, and comments walking through the approach.

See the [open issues](https://github.com/anyingiit/LeetCodeSolutions/issues) for planned features and known issues.

## Getting Started

### Prerequisites

- Go 1.22 or newer, the floor every module's `go.mod` declares (for example `problem21/go/go.mod`)

### Installation

There is no root module: clone the repository, then build whichever problem's module you want from inside its own directory.

```sh
git clone https://github.com/anyingiit/LeetCodeSolutions.git
cd LeetCodeSolutions
(cd problem21/go && go build ./...)
(cd problem27/go && go build ./...)
(cd problem2129/go && go build ./...)
```

## Usage

Run a solution's `main.go` directly from its module directory:

```sh
cd problem21/go && go run .
cd ../../problem27/go && go run .
cd ../../problem2129/go && go run .
```

Each module also carries its own automated test, for example:

```sh
cd problem21/go && go test ./...
```

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/LeetCodeSolutions](https://github.com/anyingiit/LeetCodeSolutions)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
