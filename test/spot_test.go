package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetAccountTradeList(t *testing.T) {
	values := url.Values{}
	values.Add("symbol", "ETHAUD")
	values.Add("limit", "1000")
	values.Add("recvWindow", "60000")

	res, err := testClient.GetAccountTradeList(values)
	if err != nil || res == nil {
		t.Error(err)
	}

	tradeList := *res

	for i, s := range tradeList {
		fmt.Println(i, s)
	}
}

func TestGetCoinInfo(t *testing.T) {
	values := url.Values{}
	//values.Add("symbol", "ETHAUD")
	//values.Add("limit", "1000")
	//values.Add("recvWindow", "60000")

	res, err := testClient.GetCoinInfo(values)
	if err != nil || res == nil {
		t.Error(err)
	}

	tradeList := *res

	for i, s := range tradeList {
		fmt.Println(i, s)
	}
}

func TestDustLog(t *testing.T) {
	res, err := testClient.GetDustLog(url.Values{})
	if err != nil || res == nil {
		t.Error(err)
	}
}

func TestOrder(t *testing.T) {
	values := url.Values{}
	values.Add("symbol", "ETHAUD")

	res, err := testClient.GetOrder(values)
	if err != nil || res == nil {
		t.Error(err)
	}

	tradeList := *res

	for i, s := range tradeList {
		fmt.Println(i, s)
	}
}
