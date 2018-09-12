package snapHttp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	logs "log"
)

type SnapHttp struct {
	ReSend  int // 重发次数 default 3
	Header  map[string]string
	Timeout time.Duration
}

// PostJson 发送Json数据
func (p *SnapHttp) PostJson(url string, data interface{}, target interface{}) (err error) {

	var (
		jsonData []byte
		payload  *strings.Reader
	)

	logs.Infof("[PostJson Url ] %s", url)
	if jsonData, err = json.Marshal(data); err != nil {
		logs.Errorf("[postArs JsonEncode Error] %v", err)
		return
	}

	logs.Infof("[PostJson Body] %s", string(jsonData))

	payload = strings.NewReader(string(jsonData))

	// 发送消息
	if err = p.SendReq("POST", url, payload, target); err != nil {
		return
	}

	return
}

// GetJson 获取数据
func (p *SnapHttp) GetJson(url string, target interface{}) (err error) {

	logs.Infof("[Get Url] %s", url)

	// 发送消息
	err = p.SendReq("GET", url, nil, target)
	return
}

// 重发处理
func (p *SnapHttp) SendReq(method string, url string, body io.Reader, target interface{}) (err error) {

	var (
		res *http.Response
		req *http.Request
	)

	if p.ReSend == 0 {
		p.ReSend = 3
	}

	var timeout time.Duration
	if p.Timeout != 0 {
		timeout = p.Timeout
	} else {
		timeout = time.Duration(5 * time.Second)
	}

	// 重试处理
	for i := 0; i <= p.ReSend; i++ {

		client := http.Client{
			Timeout: timeout,
		}

		req, err = http.NewRequest(method, url, body)

		req.Header.Add("content-type", "application/json;charset=utf-8")
		for key, value := range p.Header {
			req.Header.Set(key, value)
		}

		logs.Debugf("[SnapHttp Headers] %v", req.Header)

		if res, err = client.Do(req); err == nil {

			// 获取返回数据的对象
			defer res.Body.Close()
			x, _ := ioutil.ReadAll(res.Body)
			logs.Infof("[Request Url Response] %s", string(x))
			err = json.Unmarshal(x, target)
			return

		}

		logs.Errorf("[Http Send Error Retry] %d %v", i, err)
		time.Sleep(2)
	}

	return
}
