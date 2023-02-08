# 运行
- 执行 go mod tidy
- 建数据库运行sql文件
- 运行main.go
- 抖声ip改成运行主机的内网ip:8080

# 工具类使用
- 雪花算法生成唯一id：go get github.com/sony/sonyflake
- jwt生成Token : go get github.com/golang-jwt/jwt/v4
- viper读取配置类文件xxx.yml : go get github.com/spf13/viper
- 对象存储 : go get -u github.com/tencentyun/cos-go-sdk-v5

#补充
- 测试用例：https://www.apifox.cn/apidoc/shared-581228eb-4ef0-4e15-ae1e-35e08986483d
- 目前实现了基本接口，还没加Redis缓存
- app上无法发布视频，还没找到原因，不过使用接口测试是能成功的
- 如何让每个用户的token进行线程隔离的问题还没解决，所以目前没有进行token检验
- 尝试使用protoc生成model，但因为玩不明白后面的model就改为手敲了，会有点乱见谅
- test包的测试用例没有实际用处
- 为了方便查看密码目前还没加md5加密
- 本人没啥项目经验就更别说微服务项目了，项目结构都是按自己理解自行创建，如果有不对的地方欢迎指正
