wire
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/coinsec/coinsecd/wire)
=======

Package wire implements the coinsec wire protocol.

## Coinsec Message Overview

The coinsec protocol consists of exchanging messages between peers. Each message
is preceded by a header which identifies information about it such as which
coinsec network it is a part of, its type, how big it is, and a checksum to
verify validity. All encoding and decoding of message headers is handled by this
package.

To accomplish this, there is a generic interface for coinsec messages named
`Message` which allows messages of any type to be read, written, or passed
around through channels, functions, etc. In addition, concrete implementations
of most all coinsec messages are provided. All of the details of marshalling and 
unmarshalling to and from the wire using coinsec encoding are handled so the 
caller doesn't have to concern themselves with the specifics.

## Reading Messages Example

In order to unmarshal coinsec messages from the wire, use the `ReadMessage`
function. It accepts any `io.Reader`, but typically this will be a `net.Conn`
to a remote node running a coinsec peer. Example syntax is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main coinsec network.
	pver := wire.ProtocolVersion
	coinsecnet := wire.Mainnet

	// Reads and validates the next coinsec message from conn using the
	// protocol version pver and the coinsec network coinsecnet. The returns
	// are a appmessage.Message, a []byte which contains the unmarshalled
	// raw payload, and a possible error.
	msg, rawPayload, err := wire.ReadMessage(conn, pver, coinsecnet)
	if err != nil {
		// Log and handle the error
	}
```

See the package documentation for details on determining the message type.

## Writing Messages Example

In order to marshal coinsec messages to the wire, use the `WriteMessage`
function. It accepts any `io.Writer`, but typically this will be a `net.Conn`
to a remote node running a coinsec peer. Example syntax to request addresses
from a remote peer is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main bitcoin network.
	pver := wire.ProtocolVersion
	coinsecnet := wire.Mainnet

	// Create a new getaddr coinsec message.
	msg := wire.NewMsgGetAddr()

	// Writes a coinsec message msg to conn using the protocol version
	// pver, and the coinsec network coinsecnet. The return is a possible
	// error.
	err := wire.WriteMessage(conn, msg, pver, coinsecnet)
	if err != nil {
		// Log and handle the error
	}
```
