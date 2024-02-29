# grapi

## What is this?
- by: axiner
- grapi
- This is a gin restful api.

## 脚手架使用
- 1）mod及proxy设置(若已执行或不需要请忽略)
  - `go env -w GO111MODULE=on`
  - `go env -w GOPROXY=https://goproxy.io,direct`
- 2）下载安装
  - `go install github.com/atpuxiner/gtools/gtcli@latest`
- 3）创建项目
  - `gtcli grapi new -p <项目名称> -m <模块名称> -d <目录(不指定则默认.)>`
- 4）添加api
  - `cd到上面创建的项目根目录`
  - `gtcli grapi add -a <api名称> -v <版本号(不指定则默认v1)>`

## 项目运行
- 1）cd到项目根目录
- 2）初始化相关
  - 第三方模块:
    - `go get -u`
    - `go mod tidy`
    - swag cmd:
      - `go install github.com/swaggo/swag/cmd/swag@latest`
  - 代码格式化:
    - `go fmt ./...`
  - swagger:
    - `swag init`
- 3）编译运行
  - win:
    - `go build -o grapi.exe main.go`
    - `./grapi.exe runserver`
  - linux:
    - `go build -o grapi main.go`
    - `./grapi runserver`

## 项目结构
- ABD: ABD模式
  - A   api
  - B   business
  - D   datatype
- 调用过程: main.go(initializer) - router(middleware) - api - business - (datatype)
- 结构如下: (命名经过多次修改敲定，简洁易懂，ABD目录贴合避免杂乱无章)
  ```
  └── grapi
      ├── app                         (应用)
      │   ├── api                     ├── (api)
      │   │   └── v1                  │   └── (v1)
      │   ├── business                ├── (业务)
      │   ├── datatype                ├── (数据类型)
      │   │   ├── entity              │   ├── (实体)
      │   │   └── model               │   └── (模型)
      │   ├── initializer             ├── (初始化)
      │   │   ├── conf                │   ├── (配置)
      │   │   ├── db                  │   ├── (数据库)
      │   │   ├── logger              │   ├── (日志)
      │   │   └── redis               │   └── (redis)
      │   ├── middleware              ├── (中间件)
      │   ├── router                  ├── (路由)
      │   └── utils                   └── (utils)
      ├── cmd                         (命令目录)
      ├── config                      (配置目录)
      ├── deploy                      (部署目录)
      ├── docs                        (文档目录)
      ├── log                         (日志目录)
      ├── .gitignore
      ├── go.mod
      ├── LICENSE
      ├── main.go
      └── README.md
  ```

## LICENSE
This project is released under the MIT License (MIT). See [LICENSE](LICENSE)
