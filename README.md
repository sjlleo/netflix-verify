# netflix-verify
A script used to determine whether to watch native NetFlix movies / NetFlix可看地区检测脚本

Tips: 不知道为啥被人传到了梨园上，有点受宠若惊

亦推荐[@CoiaPrant @zerocloud的Shell版本](https://github.com/Netflixxp/NF)

## 使用方法
#### 1、部署 `golang` 环境，执行 `go run nf.go` 运行本小应用

![image.png](https://img.leo.moe/images/2021/02/23/image.png)

#### 2、使用编译好的二进制文件执行本小程序

> 这里演示IP实际地理位置与NF的IP库位置不相同的情况

![image1126583f0c05440f.png](https://img.leo.moe/images/2021/02/23/image1126583f0c05440f.png)

### 懒人一键运行包

`wget https://github.com/sjlleo/netflix-verify/releases/download/1.0/nf-1.0-linux-amd64 && chmod +x nf-1.0-linux-amd64 && mv nf-1.0-linux-amd64 nf && clear && ./nf`
