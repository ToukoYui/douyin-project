package db

import (
	"context"
	"douyin-template/utils"
	"fmt"
	mysql2 "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
)

// Db 全局变量Db
var Db *gorm.DB

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(context context.Context, addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}

// getSSH 进行ssh连接
func getSSH(config utils.SSH) *ssh.Client {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", config.Addr, config.Port),
		&ssh.ClientConfig{
			User: config.Usr,
			Auth: []ssh.AuthMethod{
				ssh.Password(config.Secret),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})

	if err != nil {
		panic(any("ssh连接失败"))
	}
	return client
}

// InitDb 获取数据库连接
func InitDb(sshConfig utils.SSH, mysqlConfig utils.Mysql) {
	mysql2.RegisterDialContext("mysql+tcp", (&ViaSSHDialer{getSSH(sshConfig)}).Dial)
	dsn := fmt.Sprintf("%s:%s@mysql+tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		mysqlConfig.Usr, mysqlConfig.Pwd, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Db)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(any("数据库连接失败！！！"))
	}
}
