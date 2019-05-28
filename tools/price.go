package tools

import (
	"ELRA/structs"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// BitcoinPrice is the global Bitcoin Price. Currency is defined by Bitcoin Price API (config)
var BitcoinPrice = 0.0

// StartPriceDaemon starts a service, which checks every minute for current bitcoin price
func StartPriceDaemon(config structs.Configuration) {
	delay := 1 * 60 * 1000 // Every minute

	priceCheck := func() {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		request, err := http.NewRequest("GET", config.BitcoinPriceAPI, nil)
		CheckError(err)
		resp, err := http.DefaultClient.Do(request.WithContext(ctx))
		if err != nil {
			log.Print("There is something wrong with your Bitcoin Price API.")
			log.Print(err.Error())
		} else {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Print("There is something wrong with your Bitcoin Price API.")
				log.Print(err.Error())
			} else {
				var returnValue map[string]interface{}
				json.Unmarshal(b, &returnValue)
				BitcoinPrice = returnValue[config.BitcoinPriceAPIKeyword].(float64)
				log.Print("Current Bitcoin Price: " + strconv.FormatFloat(BitcoinPrice, 'f', 2, 64))
			}

		}
		defer cancel()
	}

	priceCheck()

	_ = SetInterval(priceCheck, delay, true)
}
