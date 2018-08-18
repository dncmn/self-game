package httputil

import (
	"time"

	"github.com/valyala/fasthttp"
)

type P map[string]string

var (
	defaultHeaders = map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Content-Type":    {"application/x-www-form-urlencoded"}}
	// 设置Json发送，请用SetContent设置，不要直接使用Header，否则无效
	// defaultJsonHeaders = map[string][]string{
	//	"Content-Type": {"application/json"}}
)

const (
	DefaultTimeout = time.Second * 10
)

func buildPostData(p P) *fasthttp.Args {
	args := &fasthttp.Args{}
	for k, v := range p {
		args.Set(k, v)
	}
	return args
}

func Do(method string, uri string, args map[string]string) (body []byte, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI(uri)
	req.Header.SetMethod(method)
	var queryArgs *fasthttp.Args
	switch method {
	case "GET":
		queryArgs = req.URI().QueryArgs()
	case "POST":
		queryArgs = req.PostArgs()
		req.Header.SetContentType("application/x-www-form-urlencoded")
	default:
		queryArgs = req.URI().QueryArgs()
	}

	for k, v := range args {
		//queryArgs.Set(k, url.QueryEscape(v))
		queryArgs.Set(k, v)
	}
	if method == "POST" {
		req.SetBodyString(queryArgs.String())
	}
	err = fasthttp.DoTimeout(req, resp, DefaultTimeout)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func Get(url string) (string, error) {
	var (
		state int
		resp  []byte
		err   error
	)
	state, resp, err = fasthttp.GetTimeout(resp, url, DefaultTimeout)
	if state == 200 {
		return string(resp), nil
	}
	return string(resp), err
}

func GetBytes(url string) ([]byte, error) {
	var (
		state int
		resp  []byte
		err   error
	)
	state, resp, err = fasthttp.GetTimeout(resp, url, DefaultTimeout)
	if state == 200 {
		return resp, nil
	}
	return resp, err
}

func GetWithHeader(uri string, args map[string]string, header map[string]string) (body []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI(uri)
	req.Header.SetMethod("GET")
	req.Header.SetConnectionClose()
	var queryArgs *fasthttp.Args
	queryArgs = req.URI().QueryArgs()
	for k, v := range args {
		queryArgs.Set(k, v)
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	err = fasthttp.DoTimeout(req, resp, DefaultTimeout)
	return resp.Body(), err
}

func PostWithHeader(uri string, data []byte, header map[string]string) (body []byte, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI(uri)
	req.Header.SetMethod("POST")
	req.SetBody(data)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	err = fasthttp.DoTimeout(req, resp, DefaultTimeout)
	return resp.Body(), err
}

func postForm(uri string, args P) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI(uri)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded")
	postArgs := req.PostArgs()
	for k, v := range args {
		postArgs.Set(k, v)
	}
	req.SetBodyString(postArgs.String())
	err := fasthttp.DoTimeout(req, resp, DefaultTimeout)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func Post(url string, p P) (string, error) {
	resp, err := postForm(url, p)
	return string(resp), err
}

func PostWithRetry(url string, p P, t int) (string, error) {
	var (
		resp []byte
		err  error
	)
	if t <= 0 {
		panic("times must be more than 1")
	}
	for i := 0; i < t; i++ {
		resp, err = postForm(url, p)
		if err == nil {
			return string(resp), nil
		}
	}
	return string(resp), err
}

func PostData(url string, data []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.SetBody(data)
	err := fasthttp.DoTimeout(req, resp, DefaultTimeout)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func PostDataWithRetry(url string, data []byte, t int) ([]byte, error) {
	var (
		resp []byte
		err  error
	)
	if t <= 1 {
		panic("times must be more than 1")
	}
	for i := 0; i < t; i++ {
		resp, err = PostData(url, data)
		if err == nil {
			return resp, nil
		}
	}
	return resp, err
}
