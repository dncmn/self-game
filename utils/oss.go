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
	ossURL = fmt.Sprint("http://video.happySelf.cn/", key)
	return
}

/*
	上传指定目录到oss
	dir_path:要上传的本地目录名,例如:"../Handler"
*/
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
		oss_key := string(([]rune(f)[idx:]))

		if strings.HasSuffix(f, ".DS_Store") {
			continue
		}
		// 一个Level,一个Level的资源进行更新的:为了测试更新level1的图片,特地注掉的
		//if strings.Contains(f, "L1") || strings.Contains(f, "L2") ||
		//	strings.Contains(f, "L3") || strings.Contains(f, "L4") ||
		//	strings.Contains(f, "L5") || strings.Contains(f, "L6") {
		//	continue
		//}

		// 只更新指定level的图片资源,不更新音频
		if strings.Contains(f, ".mp3") {
			continue
		}

		//// 只更新指定level的图片资源
		if !strings.Contains(f, "L3") {
			continue
		}

		b := fmt.Sprint("wxtools/", oss_key)
		// 更改oss_key-将osskey都指定成数字形式的
		if strings.Contains(b, "LessonExercise") {
			tmp := strings.Split(b, "/")
			pre := strings.Join(tmp[:len(tmp)-1], "/")

			cls_arr := strings.Split(tmp[4], "_")
			typ_arr := strings.Split(tmp[5], "_")
			title_arr := strings.Split(tmp[len(tmp)-1], "_")

			b = fmt.Sprint(pre, "/", cls_arr[1], "_", typ_arr[0], "_", title_arr[0], path.Ext(b))
			fmt.Println("oss_key=", b, "  f=", f)
		}
		err = bkt.PutObjectFromFile(b, f)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return
}
