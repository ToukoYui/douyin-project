package utils_test

import (
	"fmt"
	"github.com/sony/sonyflake"
	"testing"
)

func Test(t *testing.T) {
	snow := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := snow.NextID()
	fmt.Println(id)
}
