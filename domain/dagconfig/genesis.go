// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"math/big"

	"github.com/coinsec/coinsecd/domain/consensus/model/externalapi"
	"github.com/coinsec/coinsecd/domain/consensus/utils/blockheader"
	"github.com/coinsec/coinsecd/domain/consensus/utils/subnetworks"
	"github.com/coinsec/coinsecd/domain/consensus/utils/transactionhelper"
	"github.com/kaspanet/go-muhash"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, //script version
	0x01,                                           // Varint
	0x00,                                           // OP-FALSE
	0x63, 0x6F, 0x69, 0x6E, 0x73, 0x65, 0x63, 0x2D, 0x6D, 0x61, 0x69, 0x6E, 0x6E, 0x65, 0x74, // coinsec mainnet
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x4b, 0xaf, 0xed, 0x65, 0x39, 0x28, 0x32, 0x06, 0xdb, 0x5e, 0xa4, 0x23, 0xbd, 0xf9, 0x20, 0x50, 0x87, 0x34, 0xab, 0x5a, 0xb0, 0x3d, 0xc7, 0xde, 0x2b, 0xa2, 0x95, 0xa8, 0xad, 0xc6, 0x7c, 0x89,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x9d, 0x59, 0xc2, 0xff, 0xe6, 0xf2, 0x2c, 0x16, 0x9d, 0x5c, 0x63, 0x1e, 0x6b, 0xb2, 0x0b, 0xf7, 0x4a, 0x8f, 0xb3, 0x99, 0x09, 0xe3, 0x2f, 0x80, 0xe1, 0x2d, 0x6d, 0x7c, 0x8e, 0x0b, 0xf0, 0x70,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1721340128267, 511705087, 83330,
		0, // Checkpoint DAA score
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x63, 0x6F, 0x69, 0x6E, 0x73, 0x65, 0x63, 0x2d, 0x64, 0x65, 0x76, 0x6e, 0x65, 0x74, // coinsec-devnet
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x39, 0x17, 0xbc, 0xe2, 0x15, 0x45, 0xad, 0x15, 
	0x8c, 0x84, 0x66, 0xb8, 0xd1, 0xa8, 0xc7, 0xa6, 
	0x01, 0x92, 0x30, 0xdf, 0x1c, 0x3e, 0x01, 0x4e, 
	0x30, 0xfe, 0xd2, 0x35, 0x51, 0x29, 0xc4, 0x02,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x10, 0xe2, 0x17, 0xad, 0x08, 0xa7, 0x09, 0x8a, 0x71, 0x36, 0xaf, 0x92, 0x96, 0x35, 0x91, 0x86, 0x4f, 0x81, 0x4d, 0x48, 0x28, 0x17, 0x82, 0xcf, 0x1b, 0x2f, 0xf6, 0x41, 0x42, 0x60, 0x90, 0xfd,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x190C7DE0E31,
		525264379,
		0x48e5e,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x63, 0x6F, 0x69, 0x6E, 0x73, 0x65, 0x63, 0x2d, 0x73, 0x69, 0x6d, 0x6e, 0x65, 0x74, // coinsec-simnet
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x25, 0x35, 0xdb, 0x24, 0x3e, 0xb4, 0xe1, 0xb5, 
	0x18, 0x36, 0x56, 0x7e, 0xc9, 0x34, 0x9a, 0x81, 
	0xac, 0x60, 0xd4, 0xe7, 0xfe, 0x9c, 0xfe, 0x3e, 
	0x4f, 0xe0, 0x4f, 0xff, 0xd7, 0xbe, 0x51, 0x3c,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x5e, 0xb1, 0x5a, 0xf4, 0x5a, 0xa3, 0xf2, 0x23, 0x54, 0xea, 0x6e, 0x6e, 0xab, 0x3a, 0xfc, 0xfc, 0xc4, 0x1c, 0x50, 0x2f, 0xad, 0xfb, 0xbb, 0x84, 0x52, 0x4c, 0x8c, 0x04, 0xea, 0x0c, 0xae, 0xf5,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x190C7DE8EDB,
		0x207fffff,
		0x2,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                         // Varint
	0x00,                                                                         // OP-FALSE
	0x63, 0x6F, 0x69, 0x6E, 0x73, 0x65, 0x63, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x65, 0x74, // coinsec-testnet
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xcf, 0xfe, 0x2a, 0xc6, 0x3e, 0xa1, 0xfd, 0xce, 
	0x8d, 0x2e, 0x8e, 0xc1, 0x41, 0x4f, 0x9f, 0x19, 
	0x24, 0xdf, 0xdc, 0xe0, 0x6e, 0xb0, 0x89, 0x51, 
	0xce, 0x87, 0x8b, 0xf8, 0x1c, 0xe2, 0xd5, 0xf0,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xa9, 0xfe, 0xd5, 0x6e, 0xd1, 0x98, 0x80, 0xc5, 0x2c, 0xc7, 0xb0, 0x49, 0x9a, 0x77, 0x33, 0x1e, 0x9c, 0xb0, 0x38, 0xe5, 0x38, 0xa0, 0x46, 0x25, 0xc5, 0xba, 0xca, 0xf2, 0xef, 0x68, 0x24, 0x77,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x190C7DF5D96,
		0x1e7fffff,
		0x14582,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
