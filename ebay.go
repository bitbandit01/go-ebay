package ebay

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func Execute(request *Request) ([]byte, error) {
	return execute(request)
}

func execute(request *Request) ([]byte, error) {
	config := request.Config
	request.Command.SetToken(config.AuthToken)

	xmlBody := []byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>")
	xmlBody = append(xmlBody, request.Command.GetRequestBody()...)
	body := bytes.NewReader(xmlBody)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/ws/api.dll", config.BaseUrl),
		body,
	)

	req.Header.Add("X-EBAY-API-DEV-NAME", config.DevId)
	req.Header.Add("X-EBAY-API-APP-NAME", config.AppId)
	req.Header.Add("X-EBAY-API-CERT-NAME", config.CertId)
	req.Header.Add("X-EBAY-API-CALL-NAME", request.Command.CallName())
	req.Header.Add("X-EBAY-API-SITEID", strconv.Itoa(config.SiteId))
	req.Header.Add("X-EBAY-API-COMPATIBILITY-LEVEL", strconv.Itoa(1055))
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if urlErr, ok := err.(*url.Error); ok {
		log.Printf("Error while trying to send req: %v", urlErr)

		return []byte{}, urlErr
	} else if resp.StatusCode != 200 {
		rspBody, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(rspBody))

		return []byte{}, fmt.Errorf("<%d> - %s", resp.StatusCode, rspBody)
	}

	rspBody, _ := ioutil.ReadAll(resp.Body)

	return rspBody, nil
}
