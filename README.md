
Coinsecd
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/wombatlabs/coinsecd)

Coinsecd is the reference full node Coinsec implementation written in Go (golang).

## What is coinsec

Coinsec is an attempt at a proof-of-work cryptocurrency with instant confirmations and sub-second block times. It is based on [the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a generalization of Nakamoto consensus.

## Requirements

Go 1.20 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install coinsecd including all dependencies:

```bash
$ git clone https://github.com/wombatlabs/coinsecd
$ cd coinsecd
$ go install . ./cmd/...
```

- Coinsecd (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.


## Getting Started

Coinsecd has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ coinsecd
```

## Discord
Join our discord server using the following link: https://discord.gg/DWzrk4ZwbA

## License

Coinsecd is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
