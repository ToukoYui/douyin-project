package utils

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var bucket = "https://douyin-1313537069.cos.ap-guangzhou.myqcloud.com"

var secretId = "AKID1yiHdpoY6QXQ6JRPDqGWdcQ2VpxtcAVi"
var secretKey = "4ygFd7VtFTVKyPhomX5RabYM1Xt2ygDt"

// NewOssClient 生成Client
func NewOssClient() *cos.Client {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息

	u, _ := url.Parse(bucket)
	// 用于 Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(secretId),  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: os.Getenv(secretKey), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return client
}

func UploadVideo(FileHeader *multipart.FileHeader) (string, string, error) {
	file, err := FileHeader.Open()
	if err != nil {
		panic(any("打开文件失败!!!"))
	}
	c := NewOssClient()
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	//name := "video/" + string(time.Now().Unix())

	// 拼接文件名-> video/yyyy/MM/dd/filename.xxx
	now := time.Now()
	fileName := FileHeader.Filename
	pathSlice := []string{"video", strconv.Itoa(now.Year()), strconv.Itoa(int(now.Month())), strconv.Itoa(now.Day()), fileName}
	filePath := strings.Join(pathSlice, "/")
	fmt.Println("上传路径--->", filePath)

	// 通过本地文件上传对象
	//_, err = c.Object.PutFromFile(context.Background(), name, "../test", nil)
	//if err != nil {
	//	panic(err)
	//}

	//通过文件流上传对象
	//fd, err := os.Open("./test")
	//if err != nil {
	//	panic(any(err))
	//}
	//defer fd.Close()
	//reader, _ := io.Reader.Read(data)

	_, err1 := c.Object.Put(context.Background(), filePath, file, nil)

	if err1 != nil {
		return "", "", err1
	}
	fmt.Println("上传文件成功！！！")

	// 生成play_url
	playUrl := strings.Join([]string{bucket, filePath}, "/")
	// 生成cover_rul
	coverUrl := "https://douyin-1313537069.cos.ap-guangzhou.myqcloud.com/picture/cover.jpg"
	//coverUrl := strings.Join([]string{bucket,coverPath},"/")

	return playUrl, coverUrl, nil
}
