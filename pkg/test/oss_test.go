package test

import (
	"bytes"
	"fmt"
	"meta-mall/pkg"
	"os"
	"testing"
)

func TestSetFileOss(t *testing.T) {

	data, err := os.ReadFile("/Users/guodayang/Downloads/榴派产品图-公众号图-48.png")
	fmt.Println("err:", err)

	url, err := pkg.SetFileOss(bytes.NewReader(data), ".png", pkg.UserAvatar)

	fmt.Println(url, err)
}
