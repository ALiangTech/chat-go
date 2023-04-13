# chat-go

## 项目结构

├─ cmd/
│  └─ main.go
├─ internal/
│  ├─ config/
│  │  └─ config.go
│  ├─ handlers/
│  │  ├─ health.go
│  │  └─ todo.go
│  ├─ middleware/
│  │  └─ auth.go
│  ├─ models/
│  │  └─ todo.go
│  ├─ repositories/
│  │  └─ todo.go
│  └─ server/
│     └─ server.go
├─ migrations/
├─ pkg/
├─ scripts/
├─ static/
├─ templates/
├─ tests/
└─ go.mod

### 文件夹说明

cmd 目录用于存放命令行应用程序，例如主函数和命令行工具。
internal 目录用于存放您的应用程序的内部代码，例如配置、处理程序、中间件、模型、存储库和服务器等。
migrations 目录用于存放数据库迁移脚本。
pkg 目录用于存放可重用的库代码。
scripts 目录用于存放一些有用的脚本。
static 目录用于存放静态文件。
templates 目录用于存放 HTML 模板文件。
tests 目录用于存放测试代码。
go.mod 文件用于记录您的项目的依赖关系

config：存放应用程序的配置文件，包括数据库连接、日志等。
router：存放路由定义和处理函数。可以根据业务逻辑组织子路由，例如router/user、router/article等。
middleware：存放中间件，例如身份验证、跨域、日志记录等。
model：存放数据模型和数据库操作相关的代码。可以根据业务逻辑组织子模型，例如model/user、model/article等。
service：存放业务逻辑相关的代码。可以根据业务逻辑组织子服务，例如service/user、service/article等。
controller：存放请求处理函数和响应相关的代码。可以根据业务逻辑组织子控制器，例如controller/user、controller/article等。
utils：存放通用的工具函数和结构体。
