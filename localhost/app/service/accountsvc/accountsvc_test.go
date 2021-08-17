package accountsvc

import (
	"fmt"
	"path"
	"testing"
)

func TestLoad(t *testing.T) {
	Service()
}

func TestStoreJSON(t *testing.T) {

	item := NewAccountItem("0x123", "123")
	fpath := path.Join("/home/dds/kitchen/defaas/tmp/accounts", "json", item.Address+".json")

	if err := storeAccountInJSON(fpath, item); err != nil {
		t.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	if _, err := Service().Create("123"); err != nil {
		t.Fatal(err)
	}
}

func TestList(t *testing.T) {

	items, err := Service().List()
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}
}
