package witness

import (
	"defaas/contracts/go/witnesspool"
	basic "defaas/core/client/basic"
	"defaas/core/config"
	"defaas/core/helper"
	"io/ioutil"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type WitnessClient struct {
	basic.BasicClient

	stopRunningTrigger chan struct{}

	IsLogin      bool
	WitnessState string // "online", "offline", "busy"
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

	client := &WitnessClient{}

	_basicClient, err := basic.NewBasicClient(dfc, key)
	if err != nil {
		return nil, err
	}
	client.BasicClient = *_basicClient

	client.stopRunningTrigger = make(chan struct{})

	return client, nil
}

// Login registers the account to Wintess Pool, then the account becomes a witness.
// The default state of just registered witnesses is "Offline".
func (client *WitnessClient) Login() error {

	depoist, err := client.WitnessPool.StdWitnessDepoist()
	if err != nil {
		return err
	}

	log.Printf("[witness] approval for login...\n")

	transferTx, err := client.FaaSToken.Approve(client.Suite.WitnessPoolAddress, depoist)
	if err != nil {
		return err
	}
	if err := client.ConfirmTxByPolling(transferTx.Hash(), helper.NumBlockToWaitRecommended); err != nil {
		return err
	}

	log.Printf("[witness] approval for login done\n")

	log.Printf("[witness] login ...\n")

	loginTx, err := client.WitnessPool.WitnessLogin()
	if err != nil {
		return err
	}
	if err := client.ConfirmTxByPolling(loginTx.Hash(), helper.NumBlockToWaitRecommended); err != nil {
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
	if err := client.ConfirmTxByPolling(logoutTx.Hash(), helper.NumBlockToWaitRecommended); err != nil {
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
	if err := client.ConfirmTxByPolling(turnOnTx.Hash(), helper.NumBlockToWaitRecommended); err != nil {
		return err
	}

	client.StartWatching()

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
	if err := client.ConfirmTxByPolling(turnOffTx.Hash(), helper.NumBlockToWaitRecommended); err != nil {
		return err
	}

	client.StoptWatching()

	log.Printf("[witness] turn off done\n")

	return nil
}

// Report reports the monitoring result for specified SLA.
func (client *WitnessClient) Report(slaID *big.Int, isProviderViolated bool) error {

	reportTx, err := client.WitnessPool.ReportViolation(slaID)
	if err != nil {
		return err
	}
	if err := client.ConfirmTxByPolling(reportTx.Hash(), helper.NumBlockToWaitRecommended); err != nil {
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
	if err := client.ConfirmTxByPolling(drawRewardTx.Hash(), helper.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}

func (client *WitnessClient) StartWatching() {

	var (
		filter              = client.WitnessPool.Contract.WitnessPoolFilterer
		sinkWitnessSelected = make(chan *witnesspool.WitnessPoolWitnessSelectedEvent)
	)

	// subWitnessSelected, err := filter.WatchWitnessSelectedEvent(nil, sinkWitnessSelected, []common.Address{client.Key.Address})
	subWitnessSelected, err := filter.WatchWitnessSelectedEvent(nil, sinkWitnessSelected, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[witness] start watching\n")

	go func() {
		for {
			select {

			case <-client.stopRunningTrigger:
				log.Printf("[witness] stop watching\n")
				return

			case err := <-subWitnessSelected.Err():
				log.Fatal(err)

			case event := <-sinkWitnessSelected:

				log.Printf("[witness] [%v] is selected for SLA [%v]\n", event.Witness, event.SlaID)

				// TODO: monitoring...
				// TODO: random time monitor
				// TODO: random time report

				log.Println(event)
				time.Sleep(5 * time.Second)

				client.Report(event.SlaID, false)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()
}

func (client *WitnessClient) StoptWatching() {
	client.stopRunningTrigger <- struct{}{}
}
