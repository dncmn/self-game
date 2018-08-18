package misc

import (
	"testing"
)

func TestPackByGzip(t *testing.T) {
	packed := PackByGzip([]byte(src))
	err, dstb := UnpackByGzip(packed)
	if err != nil {
		t.Error("Gzip unpack error", err)
	}
	dst := string(dstb)
	if src != dst {
		t.Log(src)
		t.Log(dst)
		t.Error("Gzip compact error")
	}
}

var src = "Hello World"

func init() {
	for i := 0; i < 20; i++ {
		src += src
	}
}
func BenchmarkPackByGzip(b *testing.B) {
	packed := PackByGzip([]byte(src))
	UnpackByGzip(packed)
}
