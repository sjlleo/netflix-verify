# netflix-verify

流媒体NetFlix解锁检测脚本，使用Go语言编写

## 鸣谢

1. 感谢 [@CoiaPrant](https://github.com/CoiaPrant) 指出对于地域检测更简便的方法

## 功能实现

- [X] 解锁情况判断
- [X] 地域信息显示
- [ ] 双栈检测 (暂不支持)

#### 脚本已知缺陷（不支持双栈会发生什么？）

如果这台VPS的网络同时支持IPV4/IPV6，脚本只会选择其一进行检测，如仅IPV6解锁而IPV4不解锁，但脚本是通过IPV4检测的，则该脚本会报告不解锁，反之同理。这样会导致明明可以解锁的(IPV4/IPV6)漏掉，这也是本脚本的缺陷，除非在VPS服务商仅提供单栈的情况下没有此问题。

对于这种需要检测双栈的情况，建议拉到最下方，使用CoiaPrant开源的Unix Shell版本的脚本进行检测。

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

`wget -O nf https://cdn.jsdelivr.net/gh/XmJwit/netflix-verify/nf_1.1_linux_amd64 && chmod +x nf && clear && ./nf`

## 效果演示图

![f015c50b55fa4f97a87eae0babcb3cc7.png](https://img.leo.moe/images/2021/02/25/f015c50b55fa4f97a87eae0babcb3cc7.png)

## @CoiaPrant の Linux Shell脚本

这是一个由 @CoiaPrant 开发的shell脚本哦，支持双栈的检测！

https://github.com/CoiaPrant/Netflix_Unlock_Information/

