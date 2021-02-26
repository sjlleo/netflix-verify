# netflix-verify

流媒体NetFlix解锁检测脚本，使用Go语言编写

## 鸣谢

1. 感谢 [@CoiaPrant](https://github.com/CoiaPrant) 指出对于地域检测更简便的方法
2. 感谢 [@XmJwit](https://github.com/XmJwit) 解决了IPV6 Only VPS无法下载脚本的问题

## 功能实现

- [X] 解锁情况判断
- [X] 地域信息显示
- [X] 双栈检测

## 相关名词解释

1. **不提供服务** - 所在的地区NF没开通，连自制剧也看不了
2. **宽松版权** - 有些NF拍摄的影片不是特别注重版权，所以限制放的很开
3. **解锁自制剧** - 代表可以看由NF自己拍摄的影片
4. **解锁非自制剧** - 代表可以看NF买下的第三方版权影片
5. **地域解锁** - NF在不同的地区可以看的片源都是不同的，有些影片只能在特定区观看

一般通俗意义上来说，需要能看非自制剧才算真正意义上的NF解锁

## 使用方法
#### 1、部署 `golang` 环境，执行 `go run nf.go` 运行本小应用

#### 2、懒人一键运行包（使用编译好的二进制文件执行本小程序）

##### 主下载链接:

`wget -O nf https://github.com/sjlleo/netflix-verify/releases/download/2.01/nf_2.01_linux_amd64 && chmod +x nf && clear && ./nf`

##### 备用下载链接(支持IPV6):

`wget -O nf https://cdn.jsdelivr.net/gh/sjlleo/netflix-verify/CDNRelease/nf_2.01_linux_amd64 && chmod +x nf && clear && ./nf`

## 效果演示图

![07e78b00ec35bc32e9434be6937f585d.png](https://img.leo.moe/images/2021/02/26/07e78b00ec35bc32e9434be6937f585d.png)

## @CoiaPrant の Linux Shell脚本

这是一个由 @CoiaPrant 开发的shell脚本哦，支持双栈的检测！

https://github.com/CoiaPrant/Netflix_Unlock_Information/

