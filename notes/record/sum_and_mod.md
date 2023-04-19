# Go项目中的sum，mod文件

## mod文件

### 背景

* 老版本的go 使用GOPATH模式，将自己的包或者别人的包，都放在GOPATH/src下面进行管理
* 这样会遇到一个问题，无法做到不同项目依赖同一个库的不同版本，你无法在你的项目中使用指定版本的包

#### vendor

* 在每个项目下都创建一个vendor目录，

## sum文件

* go.sum 文件的作用是记录项目依赖的hash值，防止被人修改

点开文件，可以看到每个依赖包的哈希值，在构建时，如果本地的依赖包hash值与sum文件的hash值不一致，则会拒绝构建

* go.sum文件的意义：希望别人在构建此项目的时候，使用的依赖包和go.sum文件的保持一致

go.sum文件中每行记录有module名、版本和哈希组成，并由空格分开

`<module><version>` [/go.mod] `<hash>`

依赖包版本中任何一个文件改动，都会改变其整体哈希值，每条记录中的哈希值前均有一个表示哈希算法的 `h1:`，表示后面的哈希值是由算法SHA-256计算出来的。

### 生成记录

* 当我们在GOMODULE模式中引入一个新依赖时，通常会用go get命令获取该依赖

#### go get

* go get首先会将该依赖包下载到本地缓存目录，下载完成后会对该包做哈希运算，并将结果存放到后缀为.ziphash的文件中
* 如果是在项目中运行go get命令，go get会同步更新go.mod文件和go.sum文件
* go.mod中记录的是依赖名及其版本

```
require (
   github.com/go-playground/validator/v10 v10.11.0 // indirect
)
```

go.sum文件则会记录依赖包的哈希值

```
github.com/go-playground/validator/v10 v10.11.0 h1:0W+xRM511GY47Yy3bZUbJVitCNg2BOGlCyvTqsp/xIw=
github.com/go-playground/validator/v10 v10.11.0/go.mod h1:i+3WkQ1FvaUjjxh1kSvIA4dMGDBiPU55YFDl0WbKdWU=
```

### 校验

当我们拿某项目在本地构建的时候，go命令会从本地缓存中查找所有go.mod中记录的依赖，并计算本地依赖包的哈希值，然后与go.sum中的记录进行对比；

如果校验失败，说明本地缓存目录中依赖包版本的哈希值和项目中go.sum中记录的哈希值不一致，go命令将拒绝构建。就需要确认到底是本地缓存错了，还是go.sum记录错了。更新依赖后再重新构建。
