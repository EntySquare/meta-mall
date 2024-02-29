package pkg

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/savsgio/gotils/uuid"
	"io"
	"strconv"
)

const (
	imgUrl          = "https://xuanwu-nft.oss-cn-beijing.aliyuncs.com"
	endpoint        = "http://oss-cn-beijing.aliyuncs.com"
	accessKeyID     = "LTAI5tLvK7ZSuCXHDKQAJZWp"
	accessKeySecret = "XFX63hmUufFXKLUqydk3vJaT9p4oqc"
	bucketName      = "xuanwu-nft"
	//fileFolder      = "img"
)

type FileType int

const (
	UserAvatar FileType = iota
	UserBackground
	NFT
)

func (f FileType) GetString() string {
	switch f {
	case UserAvatar:
		return "UserAvatar"
	case UserBackground:
		return "UserBackground"
	case NFT:
		return "NFT"
	default: //d其他文件方Img
		return "Img"
	}
}

// oss添加文件
func SetFileOss(reader io.Reader, format string, fileFolder FileType) (string, error) {
	var (
		fileName = uuid.V4() + "-" + strconv.Itoa(TimeNow().Nanosecond())
		client   *oss.Client
		bucket   *oss.Bucket
		err      error
	)
	// 创建OSSClient实例。
	if client, err = oss.New(endpoint, accessKeyID, accessKeySecret); err != nil {
		return "", err
	}
	// 获取存储空间。
	if bucket, err = client.Bucket(bucketName); err != nil {
		return "", err
	}

	urlRoute := fileFolder.GetString() + "/" + fileName + format

	if err = bucket.PutObject(urlRoute, reader); err != nil {
		return "", err
	}

	fmt.Println(imgUrl + "/" + urlRoute)

	return imgUrl + "/" + urlRoute, nil
}
