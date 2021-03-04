# cli-doc

## cobra

https://github.com/spf13/cobra


## cli 设计



目录设计
account/
config.toml    - blockchain config file 连接到区块链的配置文件路径

functions/
    hello/
    www/

初始化模板
defaas init

账户管理
- 切换账户
defaas account switch xxx.pem
- 生成一个账户
defaas account generate

token 管理
- 查询余额
defaas token balance


函数辅助
- 生成初始化函数模板
defaas function init -lang=go

租户接口
- 部署在 functions 目录中的函数
defaas customer deploy hello

供应商接口
- 注册
defaas provider login   
- 注销
defaas provider logout  注销
- 加入服务
defaas provider serve fulfll-config.toml
    - fulfill config  履行配置，包括对 接受部署服务器的配置，部署订单的过滤规则，竞价策略，资源负载的限制

证人接口
- 注册
defaas witness login     
- 注销
defaas witness logout   
- 上线
defaas witness turnon
- 下线
defaas witness turnoff  