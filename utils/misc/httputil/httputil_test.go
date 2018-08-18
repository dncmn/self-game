package httputil

import (
	"sync"
	"testing"
	//	"fmt"
	"fmt"
	"sync/atomic"
)

func doTest() int32 {
	wg := &sync.WaitGroup{}
	size := 1000
	wg.Add(size)
	c := int32(0)
	for i := size; i > 0; i-- {
		go func() {
			resp, err := Post("http://localhost/api/login", P{"uid": "yy", "keysum": "1", "nick": "yy"})
			if err == nil {
				//				fmt.Println(resp)
				atomic.AddInt32(&c, int32(len(resp)))
			} else {
				fmt.Println(c / 1024)
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return c
}

func TestHttp(t *testing.T) {
	c1 := doTest()
	if c1 == 0 {
		t.Fail()
	}
	c2 := doTest()
	if c2 == 0 {
		t.Fail()
	}
	c3 := doTest()
	if c3 == 0 {
		t.Fail()
	}
	c4 := doTest()
	if c4 == 0 {
		t.Fail()
	}
	fmt.Println((c1 + c2 + c3 + c4) / 1024)
}
