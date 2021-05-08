package model

import (
	"encoding/json"
	"log"
)

type AccountItem struct {
	Address       string `json:"address"`
	Password      string `json:"password"`
	BalanceOf     string `json:"balanceOf"`
	ETH           string `json:"eth"`
	IsWitness     bool   `json:"isWitness"`
	WitnessState  string `json:"witnessState"`
	WitnessReward string `json:"witnessReward"`
	IsProvider    bool   `json:"isProvider"`
	ProviderState string `json:"providerState"`
}

func String(item AccountItem) string {

	j, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return string(j)
}
