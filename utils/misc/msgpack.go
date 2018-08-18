package misc

import (
	"bytes"

	"github.com/ugorji/go/codec"
)

var msgHandle = new(codec.MsgpackHandle)

func Serialize(obj interface{}) ([]byte, error) {
	buf := make([]byte, 0, 2048)
	writer := bytes.NewBuffer(buf)
	var encoder = codec.NewEncoder(writer, msgHandle)
	err := encoder.Encode(obj)
	return writer.Bytes(), err

	//	return msgpack.Marshal(obj...)
	// if e != nil {
	//	return b, e
	// }
	// x := PackByGzip(b)
	// return x, e
}

func Deserialize(data []byte, v interface{}) error {
	reader := bytes.NewBuffer(data)
	decoder := codec.NewDecoder(reader, msgHandle)
	err := decoder.Decode(v)
	return err
	// e, x := UnpackByGzip(data)
	// if e != nil {
	//	return e
	// }
	//	return msgpack.Unmarshal(data, v...)
}

func init() {
	msgHandle.BasicHandle.TypeInfos = codec.NewTypeInfos([]string{"msgpack"})
}
