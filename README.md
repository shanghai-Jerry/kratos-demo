# kratos-demo
a demo service created by kratos

## 目录说明
1. api: 对外接口定义
2. cmd: 服务的运行入口
3. configs: 配置文件
4. internal: 内部不对外暴露的业务逻辑
   1. biz: 具体业务逻辑
   2. conf: 配置文件解析
   3. data: 数据model处理
   4. server: 具体服务的初始化，包括http和grpc
   5. service: 服务的具体接口实现
5. third_party: 外部依赖

## server具体逻辑依赖

具体server初始化的实现依赖如下：

> server <- service <- biz.UserCase <- data.Repo 

1. Service实现api中对应的接口，同时将proto request转换为biz中定义的req，请求得到biz中定义的resp，然后返回proto中定义的resp（这里为什么需要额外转一下？ 可以简单理解一下， 就是biz中处理的是具体和data直接相关的数据， proto定义的req和resp可能在biz中不完全是需要的，或者可能是来自其他逻辑）。
2. biz中定义对应的req和resp，和data中用来获取数据的接口，用来处理来自service的请求和响应（和biz紧密相关的数据获取逻辑，其他的数据获取逻辑，途径包括通过其他外部接口获取等，在service中进行resp的聚合）
3. data中具体实现biz中约定的接口定义，处理biz中对应约定的req， 返回约定的resp