package aliases

import "github.com/ethereum/go-ethereum/common"

func HashToAddress(hash common.Hash) common.Address {
	return common.BytesToAddress(hash.Bytes())
}

func Address(str string) common.Address {
	return common.HexToAddress(str)
}

func HexToHash(str string) common.Hash {
	return common.HexToHash(str)
}
