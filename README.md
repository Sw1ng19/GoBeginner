# GoBeginner
一起由浅入深学习 Go 语言，从 Go 语言语法到分布式应用，持续更新中，敬请期待... 

## Go Structure
Go 语言官方推荐的工程结构：
```
golang #环境变量 GOPATH 所在路径
├── bin #编译生成的可执行文件目录
├── pkg #编译生成第三方库
├── src #golang 工程源代码
│   ├── github.com #来自 github 的第三方库
│   ├── golang.org #来自 golang.org 的第三方库
│   ├── tank #clone 下来的 tank 根目录
│   │   ├── build #用来辅助打包的文件夹
│   │   │   ├── conf #默认的配置文件
│   │   │   ├── doc #文档
│   │   │   ├── html #前端静态资源，从项目 tank-front 编译获得
│   │   │   ├── pack #打包的脚本
│   │   │   ├── service #将 tank 当作服务启动的脚本
│   │   ├── dist #运行打包脚本后获得的安装包目录
│   │   ├── rest #golang 源代码目录
│   │   │   ├── ... #golang 源代码 不同文件用前缀区分
│   │   ├── .gitignore #gitignore 文件
│   │   ├── CHNAGELOG #版本变化日志
│   │   ├── DOCKERFILE #构建 Docker 的文件
│   │   ├── LICENSE #证书说明文件
│   │   ├── main.go #程序入口文件
│   │   ├── README.md #README 文件
```
