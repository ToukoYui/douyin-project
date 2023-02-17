# 运行
- 执行 go mod tidy
- 运行main.go
- 抖声ip改成运行主机的内网ip:8080

# 工具类使用
- 雪花算法生成唯一id：go get github.com/sony/sonyflake
- jwt生成Token : go get github.com/golang-jwt/jwt/v4
- viper读取配置类文件xxx.yml : go get github.com/spf13/viper
- 对象存储 : go get -u github.com/tencentyun/cos-go-sdk-v5
- redis：go get github.com/go-redis/redis/v8
- 接口限流: "golang.org/x/time/rate"

# 说明
- 测试用例：https://www.apifox.cn/apidoc/shared-581228eb-4ef0-4e15-ae1e-35e08986483d
- 目前实现了基本接口，还没加Redis缓存
- 用的数据库在云服务器，也可以自己在本机上运行sql添加数据
- app在鸿蒙系统运行时视频发布的功能会失效，请求发不出来

- 生成token的负载中带了用户id和用户名，并按 [键（token）:用户id用户名字（值）]的方式存入redis中
- 如何让每个用户的token进行线程隔离的问题还没解决，所以目前的token检验方法是把redis的数据与解析前端token后获得的数据进行比对
- 尝试使用protoc生成model，但因为玩不明白后面的model就改为手敲了，会有点乱见谅
- test包的测试用例没有实际用处
- 为了方便查看密码目前还没加md5加密
- 如果要运行腾讯云密钥可以自行准备，也可以联系我
- 本人没啥项目经验就更别说微服务项目了，项目结构都是按自己理解自行创建，如果有不对的地方欢迎指正
