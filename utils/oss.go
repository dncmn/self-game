package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"path"
	"strings"
)

// oss constants
const (
	OssEndPoint        = "xxxxxxx"
	OssAccessKeyId     = "yyyyyy"
	OssAccessKeySecret = "zzzzzz"
	OssBucketName      = "qa-game"
)

var (
	osClient, err = getOss()
)

func getOss() (client *oss.Client, err error) {
	client, err = oss.New(OssEndPoint,
		OssAccessKeyId,
		OssAccessKeySecret)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func GetOssBucket(bucketName string) (bkt *oss.Bucket, err error) {
	bkt, err = osClient.Bucket(bucketName)
	return
}

func PutObject(filePath, key string) (ossURL string, err error) {
	bucket, err := osClient.Bucket(OssBucketName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = bucket.PutObjectFromFile(key, filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	ossURL = fmt.Sprint("http://video.snaplingo.net/", key)
	return
}

func PutFilesToOSS(dir_path string) (err error) {
	bkt, err := GetOssBucket(OssBucketName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的文件
	fs, err := GetAllFiles(dir_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range fs {

		// 找到oss_key
		idx := strings.Index(f, path.Base(dir_path))
		oss_key := string(([]byte(f)[idx:]))
		err = bkt.PutObjectFromFile(oss_key, f)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return
}
