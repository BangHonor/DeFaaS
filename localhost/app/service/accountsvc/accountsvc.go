package accountsvc

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gogf/gf/frame/g"

	"defaas/localhost/app/model"
)

var (
	service     *UserSvc
	serviceOnce sync.Once
)

func Service() *UserSvc {

	serviceOnce.Do(func() {

		// g.Cfg().Set("local.keystore", "/home/dds/kitchen/defaas/tmp/keystore")
		// keyStoreDirPath := g.Cfg().GetString("local.keystore")
		keyStoreDirPath := "/home/dds/kitchen/defaas/tmp/keystore"
		service = NewUserSvc(keyStoreDirPath)

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

type UserSvc struct {
	users []*model.AccountItem
	ks    *keystore.KeyStore
}

func NewUserSvc(keyStoreDirPath string) *UserSvc {

	svc := &UserSvc{}

	svc.users = []*model.AccountItem{}

	svc.ks = keystore.NewKeyStore(keyStoreDirPath, keystore.StandardScryptN, keystore.StandardScryptP)

	return svc
}

func (svc *UserSvc) LoadUsers() error {

	accounts := svc.ks.Accounts()

	_ = accounts

	// TODO

	return nil
}

func (svc *UserSvc) CreateAccount(password string) (*model.AccountItem, error) {

	account, err := svc.ks.NewAccount(password)
	if err != nil {
		return nil, err
	}

	item := NewAccountItem(account.Address.Hex(), password)
	svc.users = append(svc.users, item)

	g.Log().Printf("create account %v\n", item)

	return item, nil
}

func (svc *UserSvc) ListAccount() error {

	return nil
}
