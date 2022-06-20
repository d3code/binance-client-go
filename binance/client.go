package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	cmcBaseUrl      = "https://api.binance.com"
	cmcApiKeyHeader = "X-MBX-APIKEY"
)

type client struct {
	apiKey        string
	apiSecret     string
	printResponse bool
}

func Client(apiKey string, apiSecret string) *client {
	c := &client{
		apiKey:        apiKey,
		apiSecret:     apiSecret,
		printResponse: false,
	}

	return c
}

func (c *client) PrintResponse(printResponse bool) *client {
	c.printResponse = printResponse
	return c
}

func doGetRequest(httpMethod string, requestUrl string, values url.Values, sign bool, c *client) ([]byte, error) {
	baseUrl := cmcBaseUrl + requestUrl

	req, err := http.NewRequest(httpMethod, baseUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(cmcApiKeyHeader, c.apiKey)

	values.Add("timestamp", strconv.FormatInt(time.Now().UnixMilli(), 10))

	if sign {
		shaHmac := hmac.New(sha256.New, []byte(c.apiSecret))
		shaHmac.Write([]byte(values.Encode()))

		sha := hex.EncodeToString(shaHmac.Sum(nil))
		values.Add("signature", sha)
	}

	req.URL.RawQuery = values.Encode()

	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}

	res, requestError := httpClient.Do(req)

	if requestError != nil {
		return nil, requestError
	}

	resBody, responseError := ioutil.ReadAll(res.Body)
	if responseError != nil {
		return nil, responseError
	}

	if c.printResponse && resBody != nil {
		responseString := string(resBody)
		println("Response: ", responseString)
	}

	return resBody, nil
}
