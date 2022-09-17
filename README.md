# 项目功能
QQ频道机器人

# 项目运行
```
make bot
```
# 项目结构

```
project
│   
└───cmd //服务启动入口
│   └───bot //QQ频道机器人
│       │   main.go //启动入口
└───config
│   │   config.yaml //QQ频道机器人配置文件
└───dto //数据格式
└───err //错误类型
│   │   config.yaml //QQ频道机器人配置文件
└───internale //服务内部实现
│   │
│   └───bot //QQ频道机器人
│       │   botsdk //QQ机器人 API接口
│       │   config //配置和全局变量
│       │   event //ws 事件
│       │   token
│       │   websocket //websocket 客户端
│       │   common.go //通用处理
│       │   server.go //服务处理模块

```