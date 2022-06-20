package test

import (
	"github.com/d3code/binance-client-go/binance"
	"os"
)

var testClient = binance.Client(os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_SECRET")).PrintResponse(true)
