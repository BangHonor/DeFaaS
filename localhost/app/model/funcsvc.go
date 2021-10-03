package model

import (
	"encoding/json"
	"log"
)

type FuncsvcItem struct {
	Name                 string `json:"name"` //要写进json的，补充在json中的key
	AccountAddress       string `json:"accountAddress"`
	FunccodeName         string `json:"funccodeName"` //知道funncodename直接获取相关信息
	FaaslevelID          string `json:"faaslevelID"`
	BidDuration          string `json:"bidDuration"`
	ServiceDuration      string `json:"serviceDuration"`
	HighestUnitPrice     string `json:"highestUnitPrice"`
	UnitPrice            string `json:"unitPrice"`
	ServiceFee           string `json:"serviceFee"`
	DeploymentOrderID    string `json:"deploymentOrderID"`
	DeploymentOrderState string `json:"deploymentOrderState"`
}

func (item FuncsvcItem) String() string {

	j, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return string(j)
}
