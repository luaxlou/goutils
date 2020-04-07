<center>[<img src="goland.png" width="100" height="100">](https://www.jetbrains.com/?from=goutils)</center> 
<center>感谢Jetbrain 为开源事业提供的支持。</center>
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


### envutils
APP环境变量管理工具,如envutils.IsProd()，用于判断当前是否属于生产环境。

具体看代码 [tools/envutils/envutils.go](tools/envutils/envutils.go)



## wechat 微信相关API的开发
### corpwechat 
企业微信API封装， 对 access_token 做了lazy load 和过期刷新。

- [x] GetAccessToken 单独获得access_token，以便做扩展开发用。
- [x] GetUserInfo 获取 UserId。
- [x] SendTextMessage 获取 发送消息。

### wechatapp
小程序API封装

- [x] Login 登陆
- [x] 手机号解密 摘自 https://github.com/medivhzhan/weapp

### payment

微信支付封装,对 github.com/liyoung1992/wechatpay 项目进行了再一次包装

- [x] Pay 支付。
- [x] VerifyNotify 对支付结果进行验证。
- [x] GetPaySign 小程序二次签名。



### redisdb
简化redis 初始化

### mongo
对官方驱动Mongo 官方驱动的封装，对常规CRUD进行了简化。
具体看代码 [mongo/crud.go](mongo/crud.go)


### mysqldb
对mysql实例化过程进行了封装，默认使用gorm包装。
 
 
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


### ip2location 
自动下载ip数据库，内存中识别ip对应地址。
使用方法 

```
ip2location.GetLocation(ip)
```

依赖项目：[https://github.com/lionsoul2014/ip2region/](https://github.com/lionsoul2014/ip2region/)

具体看测试用例： [ip2location/ip2location_test.go](ip2location/ip2location_test.go)
