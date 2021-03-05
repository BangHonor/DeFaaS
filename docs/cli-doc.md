# cli-doc

## cobra 库

https://github.com/spf13/cobra

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
cd cli/defaas && go build && cd -
```


----------------

## init 子命令

在`mydir`目录下执行：

```bash
defaas init
```

将产生一个使用 defaas 的模板。

模板包括以下文件：

| 名称               | 类型     | 描述                                                                              |
| ------------------ | -------- | --------------------------------------------------------------------------------- |
| client-config.toml | 配置文件 | 客户端链接到区块链的配置文件                                                       |
| defaas-config.toml | 配置文件 | defaas 的配置文件，包括三个合约的地址：token 合约，Market 合约 和 WitnessPool 合约 |
| functions/         | 函数目录 | 函数实体的目录，该目录下有多个子目录，一个子目录存放一个函数                       |
| accounts/          | 账户目录 | 地址账户的目录，该目录下有多个私钥文件（`.pem` 格式或 `.p12` 格式）                |
| tools/             | 工具目录 | 辅助工具的目录，该目录是不建议修改的                                               |

- `defaas init` 命令会默认生成一个账户 `accounts/xxxx.pem`，一个 hello 函数模板 `functions/hello/`。

初始化顺序：
1. 到服务器下载压缩包 `tools.zip`
2. 解压出目录 `tools/`
3. 创建目录 `accounts/`  使用 `tools/` 中的工具创建一个账户
4. 创建目录 `funcitons/` 使用 `tools/` 中的工具创建一个函数模板
5. 复制 `tools/` 中的配置模板到 `mydir/` 目录下
6. 尝试进行一系列检查：连接...


## function 子命令

### init

在 `mydir`目录下执行：  

```bash
defaas function init -name=hello -lang=go 
```

将在 `mydir/functions/` 目录下产生一个新的目录 `mydir/functions/hello/`。

`hello/` 目录包含一个生成的编写 FaaS 函数的模板。





## account 子命令

### generate

在 `mydir`目录下执行：  

```bash
defaas account generate
```

将在`mydir/accounts/` 目录下产生一个新的私钥文件，即产生一个新的地址账户。



### switch

切换当前使用的账户，要求目标账户的私钥文件在`mydir/accounts/` 目录下。

```
defaas account switch 0xYYYYYZZZZZZ
```





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

provider-config.toml  履行配置，包括对 接受部署服务器的配置，部署订单的过滤规则，竞价策略，资源负载的限制





## witness

### login



### logout



### turnon



### turnoff

