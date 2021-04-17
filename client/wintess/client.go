package wintess

import (
	basic "defaas/client/basic"
	"defaas/core/config"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type WitnessClient struct {
	basic.BasicClient
}

func NewWitnessClientWithFile(configFilePath, keyStoreFilePath string, password string) (*WitnessClient, error) {

	// parse defaas config
	dfc, err := config.ParseConfigFile(configFilePath)
	if err != nil {
		return nil, err
	}

	// decrypt keystore
	keyjson, err := ioutil.ReadFile(keyStoreFilePath)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		return nil, err
	}

	return NewWitnessClient(dfc, key)
}

func NewWitnessClient(dfc *config.DeFaaSConfig, key *keystore.Key) (*WitnessClient, error) {

	// init
	_basicClient, _ := basic.NewBasicClient(dfc, key)

	client := &WitnessClient{
		BasicClient: *_basicClient,
	}

	return client, nil
}

// Login registers the account to Wintess Pool, then the account becomes a witness.
// The default state of just registered witnesses is "Offline".
func (client *WitnessClient) Login() error {

	log.Printf("[witness] login ...\n")

	depoist, err := client.WitnessPool.StdWitnessDepoist()
	if err != nil {
		return err
	}

	log.Printf("[witness] approve for login...\n")
	transferTx, err := client.FaaSToken.Approve(client.DeFaaSConfig.WitnessPoolContractAddress, depoist)
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(transferTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}
	log.Printf("[witness] approve for login done\n")

	loginTx, err := client.WitnessPool.WitnessLogin()
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(loginTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	log.Printf("[witness] login done\n")

	return nil
}

// Logout logouts the account from Wintess Pool, then the account is not a witness anymore.
func (client *WitnessClient) Logout() error {

	log.Printf("[witness] logout ...\n")

	logoutTx, err := client.WitnessPool.WitenessLogout()
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(logoutTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	log.Printf("[witness] logout done\n")

	return nil
}

// TurnOn changes the state of wintess from "Offline" to "Online".
func (client *WitnessClient) TurnOn() error {

	log.Printf("[witness] turn on ...\n")

	turnOnTx, err := client.WitnessPool.WintessTurnOn()
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(turnOnTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	log.Printf("[witness] turn on done\n")

	return nil
}

// TurnOff changes the state of wintess from "Online" to "Offline".
func (client *WitnessClient) TurnOff() error {

	log.Printf("[witness] turn off ...\n")

	turnOffTx, err := client.WitnessPool.WitnessTurnOff()
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(turnOffTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	log.Printf("[witness] turn off done\n")

	return nil
}

// Report reports the monitoring result for specified SLA.
func (client *WitnessClient) Report(slaID *big.Int, isProviderViolated bool) error {

	reportTx, err := client.WitnessPool.ReportViolation(slaID)
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(reportTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}

// QueryReward queries the reward that the witness earns.
// Note that a reward is just a record in WitnessPool.
// If a wintess wants to transfer reward to his account, use "WitnessClient.DrawReward()".
func (client *WitnessClient) QueryReward() (*big.Int, error) {

	reward, err := client.WitnessPool.WitnessQueryReward()
	if err != nil {
		return nil, err
	}

	return reward, nil
}

// DrawReward draws all reward from WitnessPool contract to the account of witness.
func (client *WitnessClient) DrawReward() error {

	drawRewardTx, err := client.WitnessPool.WitnessDrawReward()
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(drawRewardTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}

func (client *WitnessClient) Run() error {

	// TODO
	// listen monitoring task from WitnessPool contract

	return nil
}
