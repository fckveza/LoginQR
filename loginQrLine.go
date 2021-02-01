package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var hosts = "https://api.vhtear.com/"
var apikey = "Premiumkey"
var headersSu = `{"User-Agent": "Line/6.3.2","X-Line-Application": "DESKTOPWIN\t6.3.2\tVH-PC\t10.0.14;SECONDARY","x-lal": "en_id"}`

func NewLoginV2() string {
	requestBody := strings.NewReader(headersSu)
	res, err := http.Post(hosts+"loginqr="+apikey, "application/json; charset=UTF-8", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	bodyString := string(data)
	get := gjson.Get(bodyString, "result.url")
	if get.Exists() {
		return get.String()
	} else {
		fmt.Println(bodyString)
		return ""
	}
}

func LoginWithCertificate(cert string) string {
	requestBody := strings.NewReader(headersSu)
	res, err := http.Post(hosts+"loginqrwithcert="+apikey+"&"+cert, "application/json; charset=UTF-8", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	bodyString := string(data)
	get := gjson.Get(bodyString, "result.url")
	if get.Exists() {
		return get.String()
	} else {
		fmt.Println(bodyString)
		return ""
	}
}

func GetCode() {
	for {
		time.Sleep(8 * time.Second)
		fmt.Println("Wait code")
		res, err := http.Get(hosts + "getcode=" + apikey)
		if err != nil {
			log.Fatal(err)
		}
		data, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		//fmt.Println(string(data))
		get := gjson.Get(string(data), "result.code")
		if get.String() != "" {
			fmt.Println(get.String())
			break
		}
	}
}

func GetTokens() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("Wait token")
		res, err := http.Get(hosts + "gettoken=" + apikey)
		if err != nil {
			log.Fatal(err)
		}
		data, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		token := gjson.Get(string(data), "result.token")
		cert := gjson.Get(string(data), "result.certificate")
		if token.String() != "" {
			fmt.Println(token.String())
			fmt.Println(cert.String())
			break
		
	}
}

func main() {
	//fmt.Println(LoginWithCertificate("certMu"))
	//fmt.Println(NewLoginV2())
	//GetCode()
	//GetTokens()
}
