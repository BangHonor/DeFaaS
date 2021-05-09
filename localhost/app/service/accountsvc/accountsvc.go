package accountsvc

import (
	"fmt"
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

	svc.items = []*model.AccountItem{}

	keyStoreDirPath := path.Join(accountsDirPath, "keystore")
	svc.ks = keystore.NewKeyStore(keyStoreDirPath, keystore.StandardScryptN, keystore.StandardScryptP)

	if err := svc.load(); err != nil {
		return nil, err
	}

	jsonDirPath := path.Join(accountsDirPath, "json")
	_ = jsonDirPath

	return svc, nil
}

func (svc *AccountSvc) load() error {

	accounts := svc.ks.Accounts()

	_ = accounts

	for _, a := range accounts {
		fmt.Println(a.Address.Hex())
	}

	return nil
}

func (svc *AccountSvc) storeOne(item *model.AccountItem) error {
	return nil
}

func (svc *AccountSvc) store() error {

	for _, item := range svc.items {
		if err := svc.storeOne(item); err != nil {
			return err
		}
	}

	return nil
}

func (svc *AccountSvc) Create(password string) (*model.AccountItem, error) {

	account, err := svc.ks.NewAccount(password)
	if err != nil {
		return nil, err
	}

	item := NewAccountItem(account.Address.Hex(), password)
	svc.items = append(svc.items, item)

	g.Log().Printf("create account %v\n", item)

	return item, nil
}

func (svc *AccountSvc) List() ([]*model.AccountItem, error) {
	return svc.items, nil
}
