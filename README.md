# NETFLIX-VERIFY

Latest Version: `v3.0-Alpha`

流媒体NetFlix解锁检测脚本，使用Go语言编写。

在VPS网络正常的情况下，哪怕是双栈网络也可在几秒内快速完成IPv4/IPv6的NF解锁情况判断。

## 其他常见流媒体脚本链接

DisneyPlus 解锁检测： https://github.com/sjlleo/VerifyDisneyPlus

Youtube 缓存节点、地域信息检测：https://github.com/sjlleo/TubeCheck

## 新特性

**2022/05/20**

由于实在看不下去自己刚学Go时写的垃圾代码...

所以趁着最近有空重构了所有的模块，引入`goroutine`并发机制，提升运行效率

## 鸣谢

1. 感谢 [@CoiaPrant](https://github.com/CoiaPrant) 指出对于地域检测更简便的方法
2. 感谢 [@XmJwit](https://github.com/XmJwit) 解决了IPV6 Only VPS无法下载脚本的问题
3. 感谢 [@samleong123](https://github.com/samleong123) 指出了文档的解释缺陷
4. 感谢 [@ymcoming](https://github.com/ymcoming) 指出了IPv6的检测Bug

## 功能实现

- [X] 解锁情况判断
- [X] 地域信息显示
- [X] 双栈网络测试
- [X] 代理检测 (Experiment)

~~半解锁检测（Deprecated）~~

## 如何食用

使用前，如果您不知道您使用的是什么架构的CPU，请先使用`uname -m`查看

**如果提示`Exec format error`是因为您下载了与您系统架构不对应的二进制文件**

对于`amd64`（`x86_64`），请使用如下命令下载运行
```bash
wget -O nf https://github.com/sjlleo/netflix-verify/releases/download/v3.0-Alpha.3/nf_linux_amd64 && chmod +x nf && ./nf
```

对于`arm64`，请使用如下命令下载运行
```bash
wget -O nf https://github.com/sjlleo/netflix-verify/releases/download/v3.0-Alpha.3/nf_linux_arm64 && chmod +x nf && ./nf
```

对于部分路由器，其SoC使用了`mips`架构，请使用如下命令下载运行
```bash
wget -O nf https://github.com/sjlleo/netflix-verify/releases/download/v3.0-Alpha.3/nf_linux_mips && chmod +x nf && ./nf
```

## 相关名词解释

1. **不提供服务** - 所在的地区NF没开通，连自制剧也看不了
2. **宽松版权** - 有些NF拍摄的影片不是特别注重版权，所以限制放的很开
3. **解锁自制剧** - 代表可以看由NF自己拍摄的影片
4. **解锁非自制剧** - 代表可以看NF买下的第三方版权影片
5. **地域解锁** - NF在不同的地区可以看的片源都是不同的，有些影片只能在特定区观看

一般来说，需要能看非自制剧才算真正意义上的NF解锁
## Stargazers over time

[![Stargazers over time](https://starchart.cc/sjlleo/netflix-verify.svg)](https://starchart.cc/sjlleo/netflix-verify)

