# HackBot
HackBot is an AI driven security scanning tool that combines OpenAI and Projectdiscovery.

HackBot是一款人工智能驱动的安全扫描工具，结合了OpenAI和Projectdiscovery。


make:
```go
GOOS=windows GOARCH=amd64 go build   -o HackBot.exe .
GOOS=linux GOARCH=amd64 go build  -o HackBot .
GOOS=drawin GOARCH=amd64 go build  -o HackBot .
```

use:
```
./HackBot -p http://proxy -t sk-xxxxx


Usage of ./HackBot:
  -p string
    	Proxy
  -t string
    	Openai Token
```

原文：
![image](https://github.com/UltimateSec/ultimaste-nuclei-templates/blob/main/qrcode.jpg)

# 免责声明：
本文章或工具仅供安全研究使用，请勿利用从事非法测试，由于传播、利用此文所提供的信息而造成的任何直接或者间接的后果及损失，均由使用者本人负责，极致攻防实验室及作者不为此承担任何责任。

