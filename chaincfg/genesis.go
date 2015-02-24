// Copyright (c) 2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"fmt"
	"time"
	"github.com/melange-app/nmcd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  wire.ShaHash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0x7f, 0x00, 0x1c, 0x01, 0x4c, 0x4b, /* |........| */
				0x2e, 0x2e, 0x2e, 0x20, 0x63, 0x68, 0x6f, 0x6f, /* |... choo| */
				0x73, 0x65, 0x20, 0x77, 0x68, 0x61, 0x74, 0x20, /* |se what | */
				0x63, 0x6f, 0x6d, 0x65, 0x73, 0x20, 0x6e, 0x65, /* |comes ne| */
				0x78, 0x74, 0x2e, 0x20, 0x20, 0x4c, 0x69, 0x76, /* |xt.  Liv| */
				0x65, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x79, 0x6f, /* |es of yo| */
				0x75, 0x72, 0x20, 0x6f, 0x77, 0x6e, 0x2c, 0x20, /* |ur own, | */
				0x6f, 0x72, 0x20, 0x61, 0x20, 0x72, 0x65, 0x74, /* |or a ret| */
				0x75, 0x72, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x63, /* |urn to c| */
				0x68, 0x61, 0x69, 0x6e, 0x73, 0x2e, 0x20, 0x2d, /* |hains. -| */
				0x2d, 0x20, 0x56, 0x0a, /* |- V|      */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0x6a, 0x77, 0xfa, 0x46, 0x49, 0x3d,
				0x61, 0x98, 0x5c, 0x11, 0x57, 0xa6, 0xe3, 0xe4,
				0x98, 0xb3, 0xb9, 0x7c, 0x87, 0x8c, 0x9c, 0x23,
				0xe5, 0xb4, 0x72, 0x9d, 0x35, 0x4b, 0x57, 0x4e,
				0xb3, 0x3a, 0x20, 0xc0, 0x48, 0x35, 0x51, 0x30,
				0x8e, 0x2b, 0xd0, 0x82, 0x95, 0xce, 0x23, 0x8e,
				0x8a, 0xd0, 0x9a, 0x7a, 0x24, 0x77, 0x73, 0x2e,
				0xb2, 0xe9, 0x95, 0xa3, 0xe2, 0x04, 0x55, 0xe9,
				0xd1, 0x37, 0xac,
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x70, 0xc7, 0xa9, 0xf0, 0xa2, 0xfb, 0x3d, 0x48,
	0xe6, 0x35, 0xa7, 0x0d, 0x5b, 0x15, 0x7c, 0x80,
	0x7e, 0x58, 0xc8, 0xfb, 0x45, 0xeb, 0x2c, 0x5e,
	0x2c, 0xb7, 0x62, 0x00, 0x00, 0x00, 0x00, 0x00,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x0d, 0xcb, 0xd3, 0xe6, 0xf0, 0x61, 0x21, 0x5b,
	0xf3, 0xb3, 0x38, 0x3c, 0x8c, 0xe2, 0xec, 0x20,
	0x1b, 0xc6, 0x5a, 0xcd, 0xe3, 0x25, 0x95, 0x44,
	0x9a, 0xc8, 0x68, 0x90, 0xbd, 0x2d, 0xc6, 0x41,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // 41c62dbd9068c89a449525e3cd5ac61b20ece28c3c38b3f35b2161f0e6d3cb0d
		Timestamp:  time.Unix(1303000001, 0), // 2011-04-17 00:26:41 +0000 UTC
		Bits:       0x1c007fff,               // 469794815 [TODO]
		Nonce:      0xa21ea192,               // 2719916434

	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

func init() {
	sha, _ := genesisBlock.BlockSha()
	if !genesisHash.IsEqual(&sha) {
		panic(fmt.Sprintf("Genesis Block Hash (%s) doesn't match the expected hash (%s).", sha, genesisHash))
	}
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x06, 0x22, 0x6e, 0x46, 0x11, 0x1a, 0x0b, 0x59,
	0xca, 0xaf, 0x12, 0x60, 0x43, 0xeb, 0x5b, 0xbf,
	0x28, 0xc3, 0x4f, 0x3a, 0x5e, 0x33, 0x2a, 0x1f,
	0xc7, 0xb2, 0xb7, 0x3c, 0xf1, 0x88, 0x91, 0x0f,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// test network (version 3).
var testNet3GenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x43, 0x49, 0x7f, 0xd7, 0xf8, 0x26, 0x95, 0x71,
	0x08, 0xf4, 0xa3, 0x0f, 0xd9, 0xce, 0xc3, 0xae,
	0xba, 0x79, 0x97, 0x20, 0x84, 0xe9, 0x0e, 0xad,
	0x01, 0xea, 0x33, 0x09, 0x00, 0x00, 0x00, 0x00,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNet3GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},            // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet3GenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0),  // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,                // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x18aea41a,                // 414098458
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0xf6, 0x7a, 0xd7, 0x69, 0x5d, 0x9b, 0x66, 0x2a,
	0x72, 0xff, 0x3d, 0x8e, 0xdb, 0xbb, 0x2d, 0xe0,
	0xbf, 0xa6, 0x7b, 0x13, 0x97, 0x4b, 0xb9, 0x91,
	0x0d, 0x11, 0x6d, 0x5c, 0xbd, 0x86, 0x3e, 0x68,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
