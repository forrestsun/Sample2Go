package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	API_Key = "" //私有KEY
)

type ShortResult struct {
	Status_code int    `json:"status_code"` //状态编码
	Longurl     string `json:"longUrl"`     //长地址
	Status_txt  string `json:"status_txt"`  //状态信息
	Url         string `json:"url"`         //短地址
}

func Shorten(uri string) (string, int, error) {
	u := "http://126.am/api!shorten.action?key=" + API_Key + "&longUrl=" + url.QueryEscape(uri)
	response, err := http.Get(u)

	if err != nil {
		return "", 0, err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	str_body := string(body)
	str_body = strings.Replace(str_body, "\\", "", -1)

	var shortUri ShortResult

	err = json.Unmarshal([]byte(str_body), &shortUri)

	if err != nil {
		return "err address", 0, err
	}

	return shortUri.Url, shortUri.Status_code, nil
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "请输入原始地址！")
		return
	}

	uri, code, err := Shorten(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	//可以通过code进一步显示详细信息
	fmt.Println("反馈状态编码为：", code)
	fmt.Println("短地址：", uri)

	return
}
