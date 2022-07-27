# blog-service（博客后端）

blog-service 博客后端 是《Go 语言编程之旅：一起用 Go 做项目》中的项目，是第二章 [HTTP 应用：写一个完整的博客后端] 的附属源码。

# 项目结构

```powershell
blog-service
├── README.md
├── configs # 配置文件
├── docs # 各种文档集合
├── global # 全局变量
├── go.mod
├── go.sum
├── internal # 内部模块
│   ├── dao # 数据访问层
│   ├── middleware # HTTP中间件
│   ├── model # 模型层，用于存放model对象
│   ├── routers # 路由相关的逻辑
│   └── service # 项目业务逻辑
├── main.go
├── pkg # 项目相关模块包
│   ├── errcode # 错误相关
│   ├── logger # 日志相关
│   └── setting # 配置相关
├── scripts # 各类构建、安装、分析等操作的脚本
├── storage # 项目零时文件
└── third_party # 第三方的资源功能，例如Swagger UI
```

