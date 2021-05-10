package accountsvc

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gogf/gf/frame/g"

	"defaas/localhost/app/model"
)

var (
	service     *AccountSvc
	serviceOnce sync.Once
)

func Service() *AccountSvc {

	serviceOnce.Do(func() {

		// g.Cfg().Set("local.keystore", "/home/dds/kitchen/defaas/tmp/keystore")
		// keyStoreDirPath := g.Cfg().GetString("local.keystore")
		accountsDirPath := "/home/dds/kitchen/defaas/tmp/accounts"
		service, _ = NewAccountSvc(accountsDirPath)

	})

	return service
}

func NewAccountItem(address string, password string) *model.AccountItem {

	item := &model.AccountItem{}

	item.Address = address
	item.Password = password
	item.BalanceOf = "0"
	item.ETH = "0"
	item.IsWitness = false
	item.WitnessState = "offline"
	item.WitnessReward = "0"
	item.IsProvider = false
	item.ProviderState = "offline"

	return item
}

type AccountSvc struct {
	dirPath string
	items   []*model.AccountItem
	ks      *keystore.KeyStore
}

func NewAccountSvc(accountsDirPath string) (*AccountSvc, error) {

	svc := &AccountSvc{}

	svc.dirPath = accountsDirPath
	svc.items = []*model.AccountItem{}
	svc.ks = keystore.NewKeyStore(
		path.Join(accountsDirPath, "keystore"), keystore.StandardScryptN, keystore.StandardScryptP)

	svc.loadJSONAll()

	return svc, nil
}

func (svc *AccountSvc) loadJSONAll() {

	accounts := svc.ks.Accounts()

	for _, account := range accounts {

		var (
			item model.AccountItem
		)

		jsonFilePath := path.Join(svc.dirPath, "json", account.Address.Hex()+".json")

		f, err := os.OpenFile(jsonFilePath, os.O_RDONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		dec := json.NewDecoder(f)
		if err := dec.Decode(&item); err != nil {
			log.Fatal(err)
		}

		svc.items = append(svc.items, &item)

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}

	g.Log().Println("[accountsvc] load loadJSONAll done")
}

func storeAccountInJSON(jsonFilePath string, item *model.AccountItem) error {

	f, err := os.OpenFile(jsonFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")

	if err := enc.Encode(item); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (svc *AccountSvc) storeJSON(item *model.AccountItem) error {

	jsonFilePath := path.Join(svc.dirPath, "json", item.Address+".json")

	if err := storeAccountInJSON(jsonFilePath, item); err != nil {
		return err
	}

	return nil
}

func (svc *AccountSvc) Create(password string) (*model.AccountItem, error) {

	account, err := svc.ks.NewAccount(password)
	if err != nil {
		return nil, err
	}

	item := NewAccountItem(account.Address.Hex(), password)

	if err := svc.storeJSON(item); err != nil {
		svc.ks.Delete(account, password)
		return nil, err
	}

	svc.items = append(svc.items, item)
	g.Log().Printf("[accountsvc] create account %v\n", item.Address)

	return item, nil
}

func (svc *AccountSvc) List() ([]*model.AccountItem, error) {
	return svc.items, nil
}
