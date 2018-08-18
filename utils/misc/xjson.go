package misc

import (
	"bytes"
	"net/http"
	//	"compress/gzip"
	"io/ioutil"

	"github.com/ugorji/go/codec"
)

var (
	_END_MSG_TOKEN []byte = []byte{13}
	SendMsg        func(http.ResponseWriter, interface{})
	SendMsgs       func(http.ResponseWriter, []interface{})
)

var jsonHandle = new(codec.JsonHandle)

func jsonMarshal(obj interface{}) ([]byte, error) {
	buf := make([]byte, 0, 2048)
	writer := bytes.NewBuffer(buf)
	var encoder = codec.NewEncoder(writer, jsonHandle)
	err := encoder.Encode(obj)
	return writer.Bytes(), err
}

func jsonUnmarshal(data []byte, v interface{}) error {
	reader := bytes.NewBuffer(data)
	decoder := codec.NewDecoder(reader, jsonHandle)
	err := decoder.Decode(v)
	return err
}

// 强制关闭远程链接
func SendThenClose(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Connection", "close")
	//	w.WriteHeader(http.StatusBadRequest)
	SendMsg(w, v)
}

// 返回错误信息
func SendError(w http.ResponseWriter, v interface{}) {
	//	w.WriteHeader(http.StatusInternalServerError)
	SendMsg(w, v)
}

func SendMsgPack(w http.ResponseWriter, objs []interface{}) {
	for i, _ := range objs {
		b, _ := Serialize(objs[i]) //  json.Marshal(v) 测试用json
		w.Write(b)
		w.Write(_END_MSG_TOKEN)
	}
}

func SendBytes(w http.ResponseWriter, objs []interface{}) {
	for i, _ := range objs {
		w.Write(objs[i].([]byte))
		w.Write(_END_MSG_TOKEN)
	}
}

func SendByte(w http.ResponseWriter, obj interface{}) {
	w.Write(obj.([]byte))
	w.Write(_END_MSG_TOKEN)
}

func SendJsons(w http.ResponseWriter, objs []interface{}) {
	for i, _ := range objs {
		b, _ := jsonMarshal(objs[i])
		w.Write(b)
		w.Write(_END_MSG_TOKEN)
	}

	// zWriter, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	// defer zWriter.Close()
	// if err != nil {
	// 	for i, _ := range objs {
	// 		b, _ := xjson.MarshalIndent(objs[i], "", "\t")
	// 		w.Write(b)
	// 		w.Write(_END_MSG_TOKEN)
	// 	}
	// 	return
	// }
	// w.Header().Set("Content-Encoding", "gzip")
	// w.Header().Set("content-type", "text/plain; charset=utf-8")
	// for i, _ := range objs {
	// 	b, _ := xjson.MarshalIndent(objs[i], "", "\t")
	// 	zWriter.Write(b)
	// 	zWriter.Write(_END_MSG_TOKEN)
	// }
	// zWriter.Flush()
}

func SendJson(w http.ResponseWriter, v interface{}) {
	b, _ := jsonMarshal(v)
	w.Write(b)
	w.Write(_END_MSG_TOKEN)

	// zWriter, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	// defer zWriter.Close()
	// if err != nil {
	// 	w.Write(b)
	// 	w.Write(_END_MSG_TOKEN)
	// 	return
	// }
	// w.Header().Set("Content-Encoding", "gzip")
	// w.Header().Set("content-type", "text/plain; charset=utf-8")
	// zWriter.Write(b)
	// zWriter.Write(_END_MSG_TOKEN)
	// zWriter.Flush()
}

func FromJson(data []byte, obj interface{}) error {
	return jsonUnmarshal(data, obj)
}

func ToJson(obj interface{}) ([]byte, error) {
	return jsonMarshal(obj)
}

func FromJsonFile(path string, obj interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = jsonUnmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}

func ToJsonFile(path string, obj interface{}) error {
	ret, err := jsonMarshal(obj)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, ret, 0644)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	jsonHandle.Indent = 2
	jsonHandle.BasicHandle.TypeInfos = codec.NewTypeInfos([]string{"json"})
}
