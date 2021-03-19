package data

import "math/big"

// FaaSLevel is ...
type FaaSLevel struct {
	ID   int
	Core *big.Int
	Mem  *big.Int
}
