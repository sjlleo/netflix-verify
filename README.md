# netflix-verify
A script used to determine whether to watch native NetFlix movies / NetFlix可看地区检测脚本

如果您有双栈（IPV4+IPV6）测试需求，可以考虑使用 [@CoiaPrant的Shell版本](https://github.com/CoiaPrant/Netflix_Unlock_Information)

本脚本不支持双栈测试，如果您的服务商同时提供IPV4/IPV6，在使用本脚本测试时可能会更偏向于使用IPV6进行测试，这些受限于Golang语言的HTTP接口底层的决策限制

## 使用方法
#### 1、部署 `golang` 环境，执行 `go run nf.go` 运行本小应用

![image.png](https://img.leo.moe/images/2021/02/23/image.png)

#### 2、使用编译好的二进制文件执行本小程序

> 这里演示IP实际地理位置与NF的IP库位置不相同的情况

![image1126583f0c05440f.png](https://img.leo.moe/images/2021/02/23/image1126583f0c05440f.png)

### 懒人一键运行包

`wget https://github.com/sjlleo/netflix-verify/releases/download/1.1/nf_1.1_linux_amd64 && chmod +x nf_1.1_linux_amd64 && mv nf_1.1_linux_amd64 nf && clear && ./nf`
