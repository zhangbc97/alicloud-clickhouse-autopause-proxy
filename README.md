# alicloud-clickhouse-autopause-proxy

阿里云Clickhouse企业版代理服务（配合autopause服务使用）

## 简介

阿里云目前已推出Clickhouse企业版，企业版采用存算分离架构，在不使用时可通过暂停实例来节省计算资源的费用。  
目前已经提供了一个gRPC服务提供自动启停功能，但是为了简化开发工作，实现一个代理服务来避免对业务代码产生侵入。

### 关联项目

- [alicloud-clickhouse-autopause](https://github.com/zhangbc97/alicloud-clickhouse-autopause)
- [Envoy](https://github.com/envoyproxy/envoy)

## 实现原理

本项目实现了两个Envoy的Filter，分别用于实现Clickhouse的TCP代理和HTTP代理。

### TCP代理

当新连接建立时，将会向alicloud-clickhouse-autopause服务发起KeepAlive请求，如果alicloud-clickhouse-autopause服务返回了需要启动实例的响应，则会阻塞当前连接，直到实例启动成功或者启动失败。
数据传输过程中，将会定时像alicloud-clickhouse-autopause服务发起KeepAlive请求，避免实例被暂停。同时为了降低对alicloud-clickhouse-autopause服务的压力，存在最小时间间隔限制，即在最小时间间隔内只会向alicloud-clickhouse-autopause服务发起一次KeepAlive请求。

### HTTP代理

当存在请求时，将会向alicloud-clickhouse-autopause服务发起KeepAlive请求，同时为了降低对alicloud-clickhouse-autopause服务的压力，存在最小时间间隔限制，即在最小时间间隔内只会向alicloud-clickhouse-autopause服务发起一次KeepAlive请求。

## 使用方式

- 编写配置文件

```yaml


```

## 使用限制

- 该项目需要配合alicloud-clickhouse-autopause服务使用，需先部署alicloud-clickhouse-autopause服务