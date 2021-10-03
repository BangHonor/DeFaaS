package faaslevelsvc

import (
	"defaas/localhost/app/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/frame/g"
	//"github.com/jmoiron/sqlx"
	"log"
	"path"
	"strconv"
	//"sync"
)

var (
	service *FaaslevelSvc
	//serviceOnce sync.Once
)

func Service() *FaaslevelSvc {
	// serviceOnce.Do(func() { //只执行一次

	// })

	faaslevelsDirPath := "/home/dds/kitchen/defaas/tmp/faaslevels"
	service, _ = NewFaaslevelSvc(faaslevelsDirPath)
	//开发的时候放在外面，方便调试，不然只响应一次，改变了无法察觉

	return service
}

type FaaslevelSvc struct {
	dirPath string
	levels  []model.FaaslevelItem
}

func NewFaaslevelSvc(faaslevelsDirPath string) (*FaaslevelSvc, error) {
	svc := &FaaslevelSvc{}

	svc.dirPath = faaslevelsDirPath

	svc.loadJSONAll()
	return svc, nil
}

func (svc *FaaslevelSvc) loadJSONAll() {
	//遍历所有文件
	dir, err := ioutil.ReadDir("/home/dds/kitchen/defaas/tmp/faaslevels/json")
	if err != nil {
		fmt.Println(err)
	}

	for _, fi := range dir {

		if fi.IsDir() { // 忽略目录
		} else {
			fmt.Println(fi.Name() + " ")
			var item model.FaaslevelItem

			jsonFilePath := path.Join(svc.dirPath, "json", fi.Name())

			f, err := os.OpenFile(jsonFilePath, os.O_RDONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}

			dec := json.NewDecoder(f)
			if err := dec.Decode(&item); err != nil {
				log.Fatal(err)
			}
			svc.levels = append(svc.levels, item)

		}
	}
	g.Log().Println("[funccodesvc] load loadJSONAll done")
}

func storeLevelInJSON(jsonFilePath string, item *model.FaaslevelItem) error {

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

func (svc *FaaslevelSvc) storeJSON(item *model.FaaslevelItem) error {

	jsonFilePath := path.Join(svc.dirPath, "json", item.ID+".json")

	if err := storeLevelInJSON(jsonFilePath, item); err != nil {
		fmt.Println("stroe faaslvel failed!")
		return err
	}

	return nil
}

func (svc *FaaslevelSvc) List() ([]model.FaaslevelItem, error) {

	return svc.levels, nil
}

func (svc *FaaslevelSvc) Add(item model.FaaslevelItem) (*model.FaaslevelItem, error) {

	// add ID of faas level
	item.ID = strconv.FormatInt(int64(len(svc.levels)), 10)

	if err := svc.storeJSON(&item); err != nil {
		return nil, err
	}

	svc.levels = append(svc.levels, item)

	g.Log().Printf("[faaslevelsvc] add faaslevel %v\n", item)

	return &item, nil
}
