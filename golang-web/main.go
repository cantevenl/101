package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func testGet() {
	// https://www.juhe.cn/box/index/id/73
	url := "http://apis.juhe.cn/simpleWeather/query?key=087d7d10f700d20e27bb753cd806e40b&city=成都"
	//url := "https://www.baidu.com"
	//Request
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("b: %v\n", string(b))
}

func testGet2() {
	params := url.Values{}
	Url, err := url.Parse("http://apis.juhe.cn/simpleWeather/query")
	if err != nil {
		return
	}
	params.Set("key", "087d7d10f700d20e27bb753cd806e40b")
	params.Set("city", "北京")
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)

	resp, err := http.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func testParseJson() {
	type result struct {
		Args    string            `json:"args"`
		Headers map[string]string `json:"headers"`
		Origin  string            `json:"origin"`
		Url     string            `json:"url"`
	}

	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%#v", res)
}

//添加请求头参数
func testAddHeader() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
	req.Header.Add("name", "猪")
	req.Header.Add("age", "1")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}

//post请求
func testPost() {
	path := "http://apis.juhe.cn/simpleWeather/query"
	urlValues := url.Values{}
	urlValues.Add("key", "087d7d10f700d20e27bb753cd806e40b")
	urlValues.Add("city", "三亚")
	r, err := http.PostForm(path, urlValues)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("b: %v\n", string(b))
}

func testPost2() {
	urlValues := url.Values{
		"name": {"pig"},
		"age":  {"2"},
	}
	reqBody := urlValues.Encode()
	resp, _ := http.Post("http://httpbin.org/post", "text/html", strings.NewReader(reqBody))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func testPostJson() {
	data := make(map[string]interface{})
	data["site"] = "www.rvakva.com"
	data["name"] = "小咖"
	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func testClient() {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	url := "http://apis.juhe.cn/simpleWeather/query?key=087d7d10f700d20e27bb753cd806e40b&city=北京"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("referer", "http://apis.juhe.cn/")
	res, err2 := client.Do(req)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("b: %v\n", string(b))
}

func testHttpServer() {
	// 请求处理函数
	f := func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, "hello world")
	}
	// 响应路径,注意前面要有斜杠 /
	http.HandleFunc("/hello", f)
	// 设置监听端口，并监听，注意前面要有冒号:
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func testHttpServer2() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	testHttpServer2()
}
