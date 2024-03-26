# WOL
一个简单的 WOL (Wake-On-Line) 工具

用以唤醒网络内的某个设备

无依赖，纯 golang 实现

编译成 arm 二进制可以直接塞到路由器里面去用

解决一些路由器无自带 wol 工具的问题(例如小米路由器)

## 使用
```shell
ericwyn@Desktop-Godzilla:/dev/go/wol$ wol -h
Usage of ./wol:
  -b string
        Broadcast address, default is 255.255.255.255
  -m string
        MAC address
ericwyn@Desktop-Godzilla:/dev/go/wol$ wol -m "1a:2b:3c:4e:5f:66" # 唤醒 mac 为 1a:2b:3c:4e:5f:66 的设备
```

## 编译
```shell
# 编译 arm 32 位
GOOS=linux GOARCH=arm go build -o wol_arm wol.go

# 编译 amd 64 位
GOOS=linux GOARCH=amd64 go build -o wol_amd64 wol.go
```