package accountsvc

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gogf/gf/frame/g"
)

var (
	_service     *UserSvc
	_serviceOnce sync.Once
)

func Service() *UserSvc {

	_serviceOnce.Do(func() {

		// g.Cfg().Set("local.keystore", "/home/dds/kitchen/defaas/tmp/keystore")
		// keyStoreDirPath := g.Cfg().GetString("local.keystore")
		keyStoreDirPath := "/home/dds/kitchen/defaas/tmp/keystore"
		_service = NewUserSvc(keyStoreDirPath)

	})

	return _service
}

type AccountItem struct {
	Address          string
	Password         string
	KeyStoreFilePath string
	BalanceOf        string
	IsWitness        bool
	IsProvider       bool
	WitnessState     string
}

func String(item AccountItem) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("address [%v]\n", item.Address))

	return b.String()
}

func NewAccountItem(address string, password string) *AccountItem {

	item := &AccountItem{}

	item.Address = address
	item.Password = password
	item.KeyStoreFilePath = ""
	item.BalanceOf = "0"
	item.IsWitness = false
	item.IsProvider = false
	item.WitnessState = "Offline"

	return item
}

type UserSvc struct {
	users []*AccountItem
	ks    *keystore.KeyStore
}

func NewUserSvc(keyStoreDirPath string) *UserSvc {

	svc := &UserSvc{}

	svc.users = []*AccountItem{}

	svc.ks = keystore.NewKeyStore(keyStoreDirPath, keystore.StandardScryptN, keystore.StandardScryptP)

	return svc
}

func (svc *UserSvc) LoadUsers() error {

	accounts := svc.ks.Accounts()

	_ = accounts

	// TODO

	return nil
}

func (svc *UserSvc) CreateAccount(password string) (*AccountItem, error) {

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
