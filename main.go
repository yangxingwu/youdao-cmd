package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	keyFrom   = "wufeifei"
	key       = "716426270"
	docType   = "json"
	youdaoAPI = "http://fanyi.youdao.com/openapi.do?keyfrom=%s&key=%s&type=data&doctype=%s&version=1.1&q=%s"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <query string>\n", os.Args[0])
	}

	url := fmt.Sprintf(youdaoAPI, keyFrom, key, docType, os.Args[1])
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	result := struct {
		Query       string   `json:"query"`
		ErrorCode   int      `json:"errorCode"`
		Translation []string `json:"translation"`
		Basic       struct {
			US       string   `json:"us-phonetic"`
			UK       string   `json:"uk-phonetic"`
			Explains []string `json:"explains"`
		} `json:"basic"`
		Web []struct {
			Key string   `json:"key"`
			Val []string `json:"value"`
		} `json:"web"`
	}{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	if result.ErrorCode != 0 {
		log.Fatalf("Error: return code is NOT 0\n")
	}

	fmt.Printf("###################################\n")
	fmt.Printf("#  %s", result.Query)
	for _, s := range result.Translation {
		fmt.Printf(" %s", s)
	}
	fmt.Printf("\n")
	fmt.Printf("#  (U: %s E: %s)\n", result.Basic.US, result.Basic.UK)
	fmt.Printf("#\n")
	for _, s := range result.Basic.Explains {
		fmt.Printf("#  %s\n", s)
	}
	fmt.Printf("#\n")
	for _, keyVal := range result.Web {
		fmt.Printf("#  %s\n", keyVal.Key)
		for _, val := range keyVal.Val {
			fmt.Printf("\t\t%s\n", val)
		}
	}
	fmt.Printf("###################################\n")
}
