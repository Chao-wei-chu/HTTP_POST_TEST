package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"rsc.io/quote"
)

func loopfunc() {
	/*---------------錯誤回收與安全退出-----------------*/
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("[ERROR] ", err, ", Program Shutdown.") //注意! 如果沒有這句，Server仍會運行，但DB掛了
		}
	}()

	for {
		time.Sleep(time.Duration(500) * time.Millisecond)
		reqBody := `
		"PK01": "IOT2",
		"PK02": "",
		"SC": "no",
		"Parm01": "` + strconv.Itoa(rand.Intn(100)) + `",
		"Parm02": "` + strconv.Itoa(rand.Intn(100)) + `",
		"Parm03": "` + strconv.Itoa(rand.Intn(100)) + `",
		`

		sEnc := base64.StdEncoding.EncodeToString([]byte(reqBody))

		reqBody2 := `
		{` + reqBody + `"EvidenceInfo":"` + sEnc + `"
		}
		`

		//fmt.Println(string(reqBody2))

		//creating the proxyURL
		proxyURL, err := url.Parse("http://10.160.3.88:8080")
		if err != nil {
			panic(err)
		}

		//adding the proxy settings to the Transport object
		transport := &http.Transport{
			Proxy:           http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
		}

		client := http.Client{
			Transport: transport,
			Timeout:   time.Duration(30) * time.Second,
		}

		req, err := http.NewRequest("POST", "https://iot.cht.com.tw/apis/CHTIoT/blockchain/v2/evidence", strings.NewReader(reqBody2))
		if err != nil {
			panic(err)
		}
		req.Header.Set("accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", "21e298b1-0c86-4bed-9ab4-c9c8b6436bef")

		response, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		//處理回傳值
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println("[ERROR] (response.Body)")
			panic(err)
		}
		fmt.Println(string(data))
	}
}

func main() {
	//afwgwaofgalwfglawiafwfwfawfawfwfwfawfawgwawfwfwf
	fmt.Println("HTTP_POST_TEST start!!")
	rand.Seed(time.Now().UnixNano())
	//go loopfunc()
	//go loopfunc()
	fmt.Println(quote.Hello())
	/*
		for {

		}*/
}
