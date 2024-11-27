# ARPostcard_server

## 项目结构
```text
- main.go 项目入口
- conf 配置文件文件夹
- biz 业务逻辑文件夹
    - conf 配置，里面存储用于配置初始化的函数
    - consts 常量
    - dao 数据访问对象
    - model 数据模型
    - service 服务
    - handler 处理器，用于存储各个路由实际调用的函数
    - infra 基础设施，用于存储各种基础设施，包括数据库等配置
    - mw 中间件
    - utils 工具
- run_log 运行日志
```

## 项目启动
1. 安装项目所需要的包，命令行中使用 `go mod tidy`
2. 修改 biz/conf/enter.go 中的常量，让 ConfigFile 指向正确的配置文件
3. 运行项目，命令行中使用 `go run main.go`，或者使用IDE中的运行按钮