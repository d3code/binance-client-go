package binance

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *client) GetAccountTradeList(q url.Values) (*[]AccountTradeList, error) {
	requestURL := "/api/v3/myTrades"

	resBody, err := doGetRequest(http.MethodGet, requestURL, q, true, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse []AccountTradeList

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetCoinInfo(q url.Values) (*[]CoinInfo, error) {
	requestURL := "/sapi/v1/capital/config/getall"

	resBody, err := doGetRequest(http.MethodGet, requestURL, q, true, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse []CoinInfo

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetDustLog(q url.Values) (*DustLog, error) {
	requestURL := "/sapi/v1/asset/dribblet"

	resBody, err := doGetRequest(http.MethodGet, requestURL, q, true, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse DustLog

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}

func (c *client) GetOrder(q url.Values) (*[]Order, error) {
	requestURL := "/api/v3/allOrders"

	resBody, err := doGetRequest(http.MethodGet, requestURL, q, true, c)
	if err != nil {
		return nil, err
	}

	var cmcResponse []Order

	unmarshalError := json.Unmarshal(resBody, &cmcResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &cmcResponse, nil
}
