# cli-doc

## cobra 库

https://github.com/spf13/cobra
https://pkg.go.dev/github.com/spf13/cobra


```bash
cobra add account  -a kitchen -l mit
cobra add init     -a kitchen -l mit
cobra add customer -a kitchen -l mit
cobra add provider -a kitchen -l mit
cobra add witness  -a kitchen -l mit
cobra add token    -a kitchen -l mit
cobra add function -a kitchen -l mit
```


```bash
cd cli/app && go build -o defaas && cd -
```


----------------


## function 子命令

### init

在 `mydir`目录下执行：  

```bash
defaas function init -name=hello -lang=go 
```

将在 `mydir/functions/` 目录下产生一个新的目录 `mydir/functions/hello/`。

`hello/` 目录包含一个生成的编写 FaaS 函数的模板。





## account 子命令

### new

一个新的私钥文件，即产生一个新的地址账户。



## customer


### deploy





## provider

### login



### logout



### serve

加入服务

```
defaas provider serve provider-config.toml
```

provider-config.toml  履行配置，包括 接受部署服务器的配置，部署订单的过滤规则，竞价策略，资源负载的限制





## witness

### login



### logout



### turnon



### turnoff

