package wire

import "io"

const (
	blockVersionDefault    = (1 << 0)
	blockVersionAuxPow     = (1 << 8)
	blockVersionChainStart = (1 << 16)
	blockVersionChainEnd   = (1 << 30)
)

type AuxPow struct {
	// CoinbaseTxn is the transaction that is in the parent block.
	CoinbaseTx *MsgTx

	// BlockHash is the hash of the ParentBlock header
	BlockHash ShaHash

	CoinbaseBranch   MerkleBranch
	BlockchainBranch MerkleBranch

	ParentBlock BlockHeader
}

type MerkleBranch struct {
	// BranchHash contains all of the ShaHash objects required to verify
	// that the specified object is in the MerkleTree.
	BranchHash []ShaHash

	// BranchSideMask is a bitmask of which side of the Merkle hash function
	// the BranchHash object should go on. Zero indicates that it should go on
	// the right. One indicates the left.
	BranchSideMask int32
}

func readAuxPow(r io.Reader, pver uint32, ap *AuxPow) error {
	ap.CoinbaseTx = &MsgTx{}
	if err := ap.CoinbaseTx.BtcDecode(r, pver); err != nil {
		return err
	}

	if err := readElement(r, &ap.BlockHash); err != nil {
		return err
	}

	if err := readMerkleBranch(r, pver, &ap.CoinbaseBranch); err != nil {
		return err
	}

	if err := readMerkleBranch(r, pver, &ap.BlockchainBranch); err != nil {
		return err
	}

	if err := readBlockHeader(r, pver, &ap.ParentBlock); err != nil {
		return err
	}

	return nil
}

func readMerkleBranch(r io.Reader, pver uint32, mb *MerkleBranch) error {
	count, err := readVarInt(r, pver)
	if err != nil {
		return err
	}

	mb.BranchHash = make([]ShaHash, count)
	for i := uint64(0); i < count; i++ {
		if err := readElement(r, &mb.BranchHash[i]); err != nil {
			return err
		}
	}

	if err := readElement(r, &mb.BranchSideMask); err != nil {
		return err
	}

	return nil
}
