package misc

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func PackByGzip(src []byte) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(src)
	w.Close()
	return b.Bytes()
}

func UnpackByGzip(src []byte) (error, []byte) {
	b := bytes.NewBuffer(src)
	r, err := gzip.NewReader(b)
	if err != nil {
		return err, nil
	}
	defer r.Close()
	s, err2 := ioutil.ReadAll(r)
	if err2 != nil {
		return err2, nil
	}
	return nil, s
}
