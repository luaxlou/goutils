[<img src="goland.png" width="100" height="100">](https://www.jetbrains.com/?from=goutils) 
 感谢Jetbrain 为开源事业提供的支持。

!!! 作者友情提示，请不要使用该项目。该项目有诸多地方欠考虑。

# Golang Utils

Golang 典型常用工具包，解决服务端开发的常用场景，提高每个项目的生产力。
 
定位：每个工具可单独被使用，尽可能保持简洁，并且普适性，大多数典型的服务端可能会用到其中80%的工具。  

## tools 工具库 
### logutils 

打印对象更漂亮
```
logutils.PrintObj(some Obj)
```

### dateutils 
日期工具类，后面根据需求添加。

具体看代码 [tools/dateutils/dateutils.go](tools/dateutils/dateutils.go)

### fileutils 
文件工具类， 后面根据需求添加。


- [x] Exists
- [x] IsDir  
- [x] MkDir 
- [x] MkDirIfNotExists 
- [x] Remove 


### iputils 
ip工具类


- [x] GetWanIp


### debugutils 
调试工具类，设置全局调试开关，可以开启和关闭调试

具体看代码

### envutils
APP环境变量管理工具,如envutils.IsProd()，用于判断当前是否属于生产环境。

具体看代码 [tools/envutils/envutils.go](tools/envutils/envutils.go)


### 其他

更多工具详见 [tools/](tools/)


## redisdb
简化redis客户端 初始化
 

## mysqldb
对mysql客户端实例化过程进行了封装，默认使用gorm包装。
 
 
### gindefault

为经典http server ，gin 提供的默认模板，避免每个项目一上来就设置一堆。比如时区默认设置为Asia/Shanghai，跨域什么的。

example:
```
gindefault.Run(":80", func(engine *gin.Engine) {
    //在此定义gin的其他
})
```

牺牲灵活性，对返回进行了标准化。

{status:1,msg:'success',data:some Object}

```
gindefault.returnSuccess()
gindefault.returnFail()
gindefault.returnError()
```

具体看代码 [gindefault/utils.go](gindefault/utils.go)

### safeexit  

优雅退出进程
### grpc server and grpc client  

对grpc客户端和服务端进行封装，省去自动连接等细节