# goutils

Golang 典型常用工具包，解决服务端开发的常用场景，提高每个项目的生产力。
 
定位：每个工具可单独被使用,尽可能保持简洁，并且普适性，大多数典型的服务端可能会用到其中80%的工具。


## tools 工具库 
* logutils 打印对象更漂亮
```
    logutils.PrintObj(some Obj)
```

## db 数据库
* mongodb 对官方驱动Mongo 官方驱动的封装，对常规CRUD进行了简化，具体看代码 [db/mongodb/crud.go](db/mongodb/curd.go)

 
 
## net 网络相关工具
* gindefault 为经典http server ，gin 提供的默认模板，避免每个项目一上来就设置一堆。比如时区默认设置为Asia/Shanghai，跨域什么的。

    example:
```
    gindefault.Run(":80", func(engine *gin.Engine) {
        //在此定定义gin的其他
    })
```
* ginutils 牺牲灵活性，对返回进行了标准化。
{status:1,msg:'success',data:some Object}
具体看代码 [net/ginutils/utils.go](net/ginutils/utils.go)
