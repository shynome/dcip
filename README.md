# 简介

通过 ssh 转发容器未暴露的端口到本地进行操作. 比如: 数据库端口转发到本地进行管理之类的.

# 使用

## 获取容器的主机可访问 ip

获取本机上的容器 ip

```sh
dcip of main_pg
```

获取远程 ssh 主机上的容器 ip

```sh
dcip of debian@example.host main_pg
```

## 端口转发到本地

基于命令 `ssh -NT -L 0.0.0.0:5432:172.17.0.5:5432 debian@example.host`

```sh
# main_pg 可以是容器名或服务名, 但只会选中最新的容器
dcip export debian@example.host main_pg:5432 0.0.0.0:5432
# 可省略本地端口以及监听地址, 监听地址默认是 0.0.0.0, 本地端口默认为容器端口
dcip export debian@example.host main_pg:5432
```

## 进阶使用

现在可以在 host 上传递 ssh options 了, 示例: `dcip of ' -J cn-host usa-host' pg`, 用法就是加上单引号并且使用空格开头
