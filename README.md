# E-commerce-system-based-on-Golang
Record learning of building an e-commerce system 

##V1.1 
这是一个基于 Golang 的电商系统后端项目，采用了 redis 作为缓存与存储组件，主要实现了基础的用户、商品、购物车、订单模块。可以作为 Golang + Redis 开发实战项目学习和实践的参考。
(redis正在被加入项目中，预计未来会设计一个前端)

## 主要功能

- 用户注册和登录（示例）
- 商品信息管理
- 购物车增删查改
- 订单生成与查询
- Redis 高效缓存购物车与订单
- RESTful API 设计

## 技术栈

- Go（1.24.3）
- Redis（推荐 6.0+，建议使用 Docker 快速部署）
- go-redis 客户端
- Gin Web 框架
- GORM

## 项目各模块说明

本项目采用分层架构，各模块职责清晰，便于维护和扩展。以下是主要目录和文件的作用说明：

| 文件/目录                      | 作用描述                                                                                      |
|:-------------------------------|:---------------------------------------------------------------------------------------------|
| `handlers/product_handler.go`   | 商品相关接口的处理器，负责接收并处理商品的 HTTP 请求，如查询商品列表、新增商品、商品详情等。    |
| `middlewares/auth.go`           | 鉴权中间件，用于校验用户的登录身份或权限（常见如校验 JWT Token），保护需要权限的接口。         |
| `models/product.go`             | 商品数据模型，定义商品结构体以及商品相关的数据库/存储操作方法。                              |
| `models/users.go`               | 用户数据模型，定义用户结构体以及用户注册、登录等相关的数据访问逻辑。                          |
| `utils/jwt.go`                  | JWT 生成与解析工具，封装了生成、验证、刷新 JSON Web Token 的方法。                           |
| `utils/ParseToken.go`           | Token 解析辅助工具，负责从请求中读取 JWT 并解析获取用户身份信息，通常被鉴权中间件调用。        |

---



#### config/
- `config.go`：加载和管理项目配置（如 Redis 连接参数、端口等）
- `.env`、`config.yaml`：环境变量或配置信息

---


### 测试

使用 curl/Postman 等工具访问已开放的 API，或根据代码进行功能测试。



## 贡献 & 问题反馈

如有建议、bug 反馈或想要贡献代码，请提交 issue 或 PR，欢迎交流！

---
