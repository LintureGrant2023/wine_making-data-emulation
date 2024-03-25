# 部署仿真程序

## 项目介绍

```
该项目是 基于大数据智能的白酒固态发酵生物反应器系统的子模块--数据仿真模块，按照以下要求即可部署、运行该模块
```



## 1.拉取并修改代码

```
//从github仓库拉取代码
git clone https://github.com/LintureGrant2023/wine_making-data-emulation.git

//修改代码的配置
1.修改mysql的配置
修改 utils/mysql.go 的 ConnectMysql() 函数，将dsn换成当前mysql的ip和port

2.修改influxdb的配置
修改 utils/influxdb.go 的 NewClient() 函数，将token和url换成当前influxdb的
修改 utils/influxdb.go 的 WriteIntoInfluxdb() 函数，将org和bucket换成当前influxdb的
修改 utils/influxdb.go 的 QueryData() 函数，将org和bucket换成当前influxdb的
建立mysql表，数据库的表字段如下图：

```

![image-20240325161755016](/img/img.png)

## 2.制作镜像

```
//前提：切换到下载项目的路径
//前提安装24.0.7以上版本的docker
//测试的docker版本如下
zgl@master:~/go/liquor_fermentation/EdgeCompute$ docker version
Client: Docker Engine - Community
 Version:           24.0.7
 API version:       1.43
 Go version:        go1.20.10
 Git commit:        afdd53b
 Built:             Thu Oct 26 09:08:01 2023
 OS/Arch:           linux/amd64
 Context:           default

//执行指令
make
make docker
make clean

//执行指令，查看本地的docker镜像，有data-emulation 则制作镜像成功
root@master:/home/zgl/go/liquor_fermentation/EdgeCompute# docker images
REPOSITORY                                                                    TAG       IMAGE ID       CREATED         SIZE
data-emulation                                                                v1        2325a69f3856   2 minutes ago   71.8MB

```

## 3.部署容器

```
#执行以下指令，运行docker容器
docker run --rm data-emulation:v1
```

## 4.备用方案

```
//前提：1.mysql能远程连接 2.mysql已经建表 3.influxdb已经安装
//在满足前提的情况下，制作镜像失败，可以选择下面这种方法：远程拉取镜像

//1.登录docker
docker login --username=aliyun5503283891 registry.cn-hangzhou.aliyuncs.com -pcao123123

//2.拉取镜像
docker pull registry.cn-hangzhou.aliyuncs.com/cscs77/data-emulation:v1.0

//3.运行并进入容器
docker run --it --rm data-emulation:v1

//4.按照【第一步：拉取并修改代码】中的要求，修改代码
//完成之后，程序正常运行
```

