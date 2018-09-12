# sllog

一个基于logrus 的json 格式日志库，支持gorm 日志接入，支持日志滚动、日志删除、打印文件名、行号等功能。

# 使用试例

1、下载包
```
go get github.com/snaplingo-org/sllog
```

2、导入包，引用
```
package main

import "github.com/snaplingo-org/sllog"

func main() {
	sllog.Init(sllog.FileLogConfig{
		Path:     "./loggers",
		Filename: "./loggers/apps.log",
		MaxLines: 10000000000000,
		Maxsize:  10000000000000,
		Daily:    true,
		MaxDays:  3,
		Rotate:   true,
		Level:    sllog.DebugLevel,
	})

	logs := sllog.GetLogger()
  
        // 添加json字段
	testMap := make(map[string]interface{})
	testMap["a"] = "a"
	testMap["b"] = "b"

	logs.Info("hello")
	logs.InfoWithF(&testMap, "hello")
	logs.Infof("%s", "hello world!")
	logs.InfofWithF(&testMap, "%s", "hello ryan!")
        …
}
```
