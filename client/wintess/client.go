package wintess

import (
	basic "defaas/client/basic"
	"defaas/core/config"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type WitnessClient struct {
	basic.BasicClient
}

func NewWitnessClientWithFile(configFilePath, keyStoreFilePath string, password string) (*WitnessClient, error) {

	// init
	_basicClient, _ := basic.NewBasicClientWithFile(configFilePath, keyStoreFilePath, password)

	client := &WitnessClient{
		BasicClient: *_basicClient,
	}

	return client, nil
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

	depoist, err := client.WitnessPool.StdWitnessDepoist()
	if err != nil {
		return err
	}

	transferTx, err := client.FaaSToken.Transfer(client.DeFaaSConfig.WitnessPoolContractAddress, depoist)
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(transferTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	loginTx, err := client.WitnessPool.WitnessLogin()
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(loginTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}

// Logout logouts the account from Wintess Pool, then the account is not a witness anymore.
func (client *WitnessClient) Logout() error {

	logoutTx, err := client.WitnessPool.WitenessLogout()
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(logoutTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}

// TurnOn changes the state of wintess from "Offline" to "Online".
func (client *WitnessClient) TurnOn() error {

	turnOnTx, err := client.WitnessPool.WintessTurnOn()
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(turnOnTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}

// TurnOff changes the state of wintess from "Online" to "Offline".
func (client *WitnessClient) TurnOff() error {

	turnOffTx, err := client.WitnessPool.WitnessTurnOff()
	if err != nil {
		return err
	}
	if err := client.ComfirmTxByPolling(turnOffTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

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
