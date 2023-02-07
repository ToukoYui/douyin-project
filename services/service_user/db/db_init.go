package db

import (
	"context"
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
func getSSH() *ssh.Client {
	client, err := ssh.Dial("tcp", "8.134.208.93:22",
		&ssh.ClientConfig{
			User: "root",
			Auth: []ssh.AuthMethod{
				ssh.Password("Yukino123"),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})

	if err != nil {
		panic(any("ssh连接失败"))
	}
	return client
}

// InitDb 获取数据库连接
func InitDb() {
	mysql2.RegisterDialContext("mysql+tcp", (&ViaSSHDialer{getSSH()}).Dial)

	s1 := "root:Touko217@mysql+tcp(127.0.0.1:3306)/dy_user?charset=utf8mb4&parseTime=true&loc=Local"

	var err error
	Db, err = gorm.Open(mysql.Open(s1), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(any("数据库连接失败！！！"))
	}
}
