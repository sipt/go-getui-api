package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url, authToken string, reply interface{}) error {

	req, err := http.NewRequest("GET", url, nil)
	req.Header["authtoken"] = append(req.Header["authtoken"], authToken)
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)

	client := &http.Client{
		Timeout: DEFAULT_CONNECTION_TIMEOUT * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(result), reply)
}

func Post(url, authToken string, body interface{}, reply interface{}) error {
	var bodyByte []byte
	var err error
	if body != nil {
		bodyByte, err = json.Marshal(body)
		if err != nil {
			return err
		}
	} else {
		bodyByte = make([]byte, 0, 0)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))
	req.Header["authtoken"] = append(req.Header["authtoken"], authToken)
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)

	client := &http.Client{
		Timeout: DEFAULT_CONNECTION_TIMEOUT * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(result, reply)
}

func Delete(url string, authToken string, body interface{}, reply interface{}) error {
	var bodyByte []byte
	var err error
	if body != nil {
		bodyByte, err = json.Marshal(body)
		if err != nil {
			return err
		}
	} else {
		bodyByte = make([]byte, 0, 0)
	}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(bodyByte))
	req.Header["authtoken"] = append(req.Header["authtoken"], authToken)
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(result, reply)
}

func GetBody(Param interface{}) ([]byte, error) {

	body, err := json.Marshal(Param)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}
