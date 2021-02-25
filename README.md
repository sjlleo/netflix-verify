# netflix-verify
A script used to determine whether to watch native NetFlix movies / NetFlix可看地区检测脚本

本脚本不支持双栈测试，如果您的服务商同时提供IPV4/IPV6，在使用本脚本测试时可能会更偏向于使用IPV6进行测试，这些受限于Go语言的HTTP接口底层的决策限制。如果您这方面的测试需求，可以考虑使用 [@CoiaPrant的Shell脚本](https://github.com/CoiaPrant/Netflix_Unlock_Information)。

## 相关名词解释

1. **不提供服务** - 所在的地区NF没开通，连自制剧也看不了
2. **宽松版权** - 有些NF拍摄的影片不是特别注重版权，所以限制放的很开
3. **解锁自制剧** - 代表可以看由NF自己拍摄的影片
4. **解锁非自制剧** - 代表可以看NF买下的第三方版权影片
5. **地域解锁** - NF在不同的地区可以看的片源都是不同的，有些影片只能在特定区观看

一般通俗意义上来说，需要能看非自制剧才算真正意义上的NF解锁

## 使用方法
#### 1、部署 `golang` 环境，执行 `go run nf.go` 运行本小应用

![image.png](https://img.leo.moe/images/2021/02/23/image.png)

#### 2、使用编译好的二进制文件执行本小程序

> 这里演示IP实际地理位置与NF的IP库位置不相同的情况

![image1126583f0c05440f.png](https://img.leo.moe/images/2021/02/23/image1126583f0c05440f.png)

### 懒人一键运行包

`wget -O nf https://github.com/sjlleo/netflix-verify/releases/download/1.1/nf_1.1_linux_amd64 && chmod +x nf && clear && ./nf`
