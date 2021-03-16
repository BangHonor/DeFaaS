package market

import "math/big"

// FaaSLevel is ...
type FaaSLevel struct {
	ID   int
	Core *big.Int
	Mem  *big.Int
}

// GetFaaSLevels ...
func GetFaaSLevels() ([]FaaSLevel, error) {

	levels := []FaaSLevel{
		{0, big.NewInt(1), big.NewInt(512)},
		{1, big.NewInt(1), big.NewInt(1024)},
		{2, big.NewInt(2), big.NewInt(2048)},
		{3, big.NewInt(4), big.NewInt(2048)},
		{4, big.NewInt(4), big.NewInt(4096)},
	}

	return levels, nil
}
