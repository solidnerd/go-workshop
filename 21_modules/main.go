package main

//https://github.com/uber-go/zap

import (
	"time"

	"github.com/uber-go/zap"
)

// TODO:
// 1. Init a Go Modules
// 2. Add a http client and call https://httpbin.org/status/404
// 3. Check your Error's correctly and log it correctly

// More Infos https://httpbin.org/#/Status_codes/get_status__codes_
func main() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	url := "https://httpbin.org/status/404"
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")
}
