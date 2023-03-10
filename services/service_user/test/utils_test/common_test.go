package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type User struct {
	name string
}

func TestName(t *testing.T) {
	now := time.Now()
	fileName := "abc.mp4"
	pathSlice := []string{"video", strconv.Itoa(now.Year()), strconv.Itoa(int(now.Month())), strconv.Itoa(now.Day()), fileName}
	filePath := strings.Join(pathSlice, "/")
	fmt.Println(filePath)
	fmt.Println(strconv.Itoa(int(now.Month())))

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Unix(16757759911, 0))
	fmt.Println(time.Unix(1675775991, 0).Format("2006-01-02 15:04:05"))

	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))
}
