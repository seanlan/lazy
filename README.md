# lazy
code generator

代码生成器

支持数据库 mysql postgresql

新建工程
```
mkdir project
cd project
mkdir gen
cd gen
lazy save --path ./templates
lazy save --path ./ --t sh
cd ../
sh gen/gen_dao.sh
sh gen/gen_api.sh login UserLogin
```

常用目录结构
```
project
├── cmd  # cobra 命令行工具生成
├── gen
│   ├── templates
│   ├── gen_api.sh
│   └── gen_dao.sh
├── init
│   └── init.go
├── pkg # 自定义包
├── app
│   ├── api/v1 请求处理(自动生成)
│   ├── dao # 数据库操作(自动生成)
│   ├── e    # 错误码定义
│   ├── handler 
│   ├── middleware # 自定义中间件
│   ├── model  # 请求/响应 数据模型(自动生成)
│   ├── router # 路由定义
│   ├── service # 逻辑(自动生成),逻辑部分需要自己实现
│   └── utils   # 工具方法
└── main.go  # cobra 命令行工具生成
```
