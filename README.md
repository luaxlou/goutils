# goutils

golang 典型常用工具包
 
对于该工具包的定位：每个工具可单独被使用,尽可能保持简介，并且普适性，80%的项目都会用到。


## tools 工具库 
* logutils 打印对象更漂亮

## db 数据库
* mongodb 对官方驱动Mongo 官方驱动的封装，对常规CRUD进行了简化，具体看代码 [db/mongodb/utils.go](db/mongodb/utils.go)

 
 
## net 网络相关工具
* gindefault 为经典http server ，gin 提供的默认模板，避免每个项目一上来就设置一堆。比如时区默认设置为Asia/Shanghai，跨域什么的。

    example:
```
    gindefault.Run(":80", func(engine *gin.Engine) {
        //在此定定义gin的其他
    })
```