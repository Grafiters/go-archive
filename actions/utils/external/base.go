package external

import (
	"archive/actions/utils/interfaces"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiUrl = "https://9gag.com"

type Response struct {
	Meta interfaces.Meta `json:"meta"`
	Data interfaces.Data `json:"data"`
}

func GetRequest(path string) (*Response, error) {
	fmt.Println("====== REQUEST ======")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	URL := fmt.Sprintf("%s/%s", apiUrl, path)
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", URL, nil)
	headers := map[string]string{
		"Host": "9gag.com",
		"User-Agent": "	Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/119.0",
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var apiResponse Response
	err = json.Unmarshal([]byte(body), &apiResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}
	fmt.Println("====== RESPONSE ======")

	return &apiResponse, nil
}
