package model

import (
	"encoding/json"
	"log"
)

type FunccodeItem struct {
  Name string `json:"name"`
  Tag string `json:"tag"`
  Files struct{
	Filename string `json:"filename"`
	Language string `json:"language"`
	Code string `json:"code"`
  }
}

func (item FunccodeItem) String() string {

	j, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return string(j)
}
