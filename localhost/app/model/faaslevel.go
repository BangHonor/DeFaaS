package model

import (
	"encoding/json"
	"log"
)

type FaaslevelItem struct {
	ID  string `json:"id"`
	CPU string `json:"cpu"`
	Mem string `json:"mem"`
}

func (item FaaslevelItem) String() string {

	j, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return string(j)
}
