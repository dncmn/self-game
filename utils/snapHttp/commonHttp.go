package snapHttp

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// HttpJson sends a http request and unmarshal json response to v to return.
func HttpJson(method, url string, contentType string, body io.Reader, v interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, v)
		if err != nil {
			return fmt.Errorf("%s %s", string(data), err.Error())
		}
	}

	return nil
}

// HttpXml sends a http request and unmarshal xml response to v to return.
func HttpXml(method, url string, contentType string, body io.Reader, v interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = xml.Unmarshal(data, v)
		if err != nil {
			return fmt.Errorf("%s %s", string(data), err.Error())
		}
	}

	return nil
}
