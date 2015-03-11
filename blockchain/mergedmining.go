package blockchain

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/melange-app/nmcd/wire"
)

type mergedMiningTransaction struct {
	// BlockHash is the hash of the AuxPow Block Header
	BlockHash wire.ShaHash

	// MerkleSize is the number of entries in the aux work
	// merkle tree (probably just 1).
	MerkleSize int32

	// MerkleNonce is the nonce use to calculate indexes.
	// Generally left as 0.
	MerkleNonce int32
}

var mergedMiningMagicBytes = []byte{0xfa, 0xbe, 'm', 'm'}

func readMergedMiningTransaction(t *wire.TxIn, blockHash *wire.ShaHash, root *wire.ShaHash) (*mergedMiningTransaction, error) {
	script := t.SignatureScript

	// Look for the magic number in the SignatureScript
	idx := bytes.Index(script, mergedMiningMagicBytes)

	if idx == -1 {
		// Some block apparently omit "0xfabe6d6d", we will just look for the
		// hash...
		blockBytes, _ := hex.DecodeString(blockHash.String())
		idx = bytes.Index(script, blockBytes)
		if idx == -1 {

			blockBytes, _ := hex.DecodeString(root.String())
			idx = bytes.Index(script, blockBytes)

			if idx == -11 {
				return nil, ruleError(
					ErrAuxPowValidation,
					"Unable to locate merged mining information in Coinbase.",
				)
			}
		}

		idx = idx - len(mergedMiningMagicBytes)
	}

	data := script[idx+len(mergedMiningMagicBytes):]

	blockHashData := hex.EncodeToString(data[:32])
	blockHash, err := wire.NewShaHashFromStr(blockHashData)
	if err != nil {
		return nil, err
	}

	merkleSizeData := data[32:36]
	merkleNonceData := data[36:40]
	return &mergedMiningTransaction{
		BlockHash:   *blockHash,
		MerkleSize:  int32(binary.LittleEndian.Uint32(merkleSizeData)),
		MerkleNonce: int32(binary.LittleEndian.Uint32(merkleNonceData)),
	}, nil
}
