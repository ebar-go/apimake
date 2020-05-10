# apimake
基于golang的CLI库以及本地存储，实现通过命令行进行的简单的管理api文档

## 编译
在build之前建议配置`GOPROXY=https://goproxy.cn`
```
git clone https://github.com/ebar-go/apimake
cd apimake
go build
```
## 功能
通过help查看使用方式
```
apimake --help
```

### 列表显示接口文档
通过`list`或者别名`ls`展示所有的接口文档
```
apimake ls
```
如图:
![图](https://github.com/ebar-go/apimake/blob/master/list.png)

### 创建接口文档
通过`create`命令创建接口文档
```
// 根据炫酷的终端交互，使用tab切换输入框,输入完成后切换到保存按钮再Enter
apimake create
```
如图:
![图](https://github.com/ebar-go/apimake/blob/master/create.png)


### 查看接口文档
通过`show`命令查看接口文档
```
apimake show --id=1
```
如图:
![图](https://github.com/ebar-go/apimake/blob/master/detail.png)


### 更新接口文档
通过`update`或者别名`edit`更新接口文档。如果选择已存在的参数，为更新，也可以新增参数
```
// 查看帮助
apimake update --help

// 更新基本信息,id为列表里的编号，每个接口的编号都是唯一的
apimake update --id=1

// 更新头部信息
apimake update --id=1 --type=header
// 更新请求参数
apimake update --id=1 --type=request
// 更新响应参数
apimake update --id=1 --type=response
```

### 删除接口文档
暂无

## TODO
- 通过`--language`实现中英文切换