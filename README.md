# sql2struct

## 软件说明
基于cobra 的命令行工具,直接编译使用
```
go build -o sql2struct -v
```

### 目前支持功能
1. 数据表转换为结构体
使用方式
```
go run main.go sql struct --username root --password root --db blog --table users
```
### 后续功能
尝试直接使用packege的方式使用
## 版本 v0.01 