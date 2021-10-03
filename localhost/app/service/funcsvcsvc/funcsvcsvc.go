package funcsvcsvc

import (
	"defaas/localhost/app/model"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"io/ioutil"
	"log"
	"os"
	"path"
	//"sync"
)

var (
	service *FuncsvcSvc
	//serviceOnce sync.Once
)

func Service() *FuncsvcSvc {

	// serviceOnce.Do(func() {
	// })

	funcsvcsDirPath := "/home/dds/kitchen/defaas/tmp/funcsvcs"
	service, _ = NewFuncsvcSvc(funcsvcsDirPath)

	return service
}

type FuncsvcSvc struct {
	dirPath  string
	funcsvcs []model.FuncsvcItem
}

func NewFuncsvcSvc(funcsvcsDirPath string) (*FuncsvcSvc, error) {

	svc := &FuncsvcSvc{}
	svc.dirPath = funcsvcsDirPath
	svc.loadJSONAll()
	return svc, nil
}

func (svc *FuncsvcSvc) loadJSONAll() {
	//遍历所有文件
	dir, err := ioutil.ReadDir("/home/dds/kitchen/defaas/tmp/funcsvcs/json")
	if err != nil {
		fmt.Println(err)
	}

	for _, fi := range dir {

		if fi.IsDir() { // 忽略目录
		} else {
			fmt.Println(fi.Name() + " ")
			var item model.FuncsvcItem

			jsonFilePath := path.Join(svc.dirPath, "json", fi.Name())

			f, err := os.OpenFile(jsonFilePath, os.O_RDONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}

			dec := json.NewDecoder(f)
			if err := dec.Decode(&item); err != nil {
				log.Fatal(err)
			}
			svc.funcsvcs = append(svc.funcsvcs, item)

		}
	}
	g.Log().Println("[funccsvcsvc] load loadJSONAll done")
}

func storeFuncsvcInJSON(jsonFilePath string, item *model.FuncsvcItem) error {

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

func (svc *FuncsvcSvc) storeJSON(item *model.FuncsvcItem) error {

	jsonFilePath := path.Join(svc.dirPath, "json", item.Name+".json")

	fmt.Println("FuncsvcName", item.Name)
	if err := storeFuncsvcInJSON(jsonFilePath, item); err != nil {
		return err
	}

	return nil
}
func (svc *FuncsvcSvc) List() ([]model.FuncsvcItem, error) {

	return svc.funcsvcs, nil
}

func (svc *FuncsvcSvc) Add(item model.FuncsvcItem) (model.FuncsvcItem, error) {

	// add name of funcsvc
	service.storeJSON(&item)

	svc.funcsvcs = append(svc.funcsvcs, item)

	g.Log().Printf("[funcsvcsvc] add funcsvc %v\n", item)

	return item, nil
}

func (svc *FuncsvcSvc) Delete(item model.FuncsvcItem) (model.FuncsvcItem, error) {

	for index, value := range svc.funcsvcs {
		if value.Name == item.Name {
			svc.funcsvcs = append(svc.funcsvcs[:index], svc.funcsvcs[index+1:]...)
			break
		}
	}
	return item, nil
}
