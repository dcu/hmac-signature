package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	keyFlag    = flag.String("key", "", "Key to sign the request.")
	nonceFlag  = flag.String("nonce", "", "Optional nonce sign the request.")
	methodFlag = flag.String("http_method", "", "Request http method.")
	urlFlag    = flag.String("url", "", "Request URL.")
	paramsFlag = flag.String("params", "", `Request Params. Example: "a=1 b=2"`)
)

func paramsToString() string {
	params := url.Values{}
	for _, param := range strings.Split(*paramsFlag, " ") {
		keyAndValue := strings.SplitN(param, "=", 2)

		if len(keyAndValue) == 2 {
			params[keyAndValue[0]] = []string{keyAndValue[1]}
		}
	}

	return params.Encode()
}

func main() {
	flag.Parse()

	if (*keyFlag) == "" || (*urlFlag) == "" {
		flag.Usage()
	}

	//  Create a string variable using the url without params
	println("URL:", *urlFlag)

	// Create a string variable with the HTTP method in upper case (GET, POST)
	println("Method:", *methodFlag)

	// Sort the list of parameters in case-sensitive order and convert them to URL format (both key and value should be URL-encoded)
	sortedParams := paramsToString()
	println("Params:", sortedParams)

	// Generate a unique nonce (number once):
	if len(*nonceFlag) == 0 {
		*nonceFlag = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	println("Nonce:", *nonceFlag)

	// Join http_method, nonce, url and params_in_url_format together with the | character:
	data := strings.Join([]string{*nonceFlag, strings.ToUpper(*methodFlag), *urlFlag, sortedParams}, "|")
	println("Data:", data)

	// Hash the resulting data using HMAC-SHA256, using your app_signing_key as the key
	mac := hmac.New(sha256.New, []byte(*keyFlag))
	mac.Write([]byte(data))
	digest := mac.Sum(nil)

	// Encode in base64 the digest:
	digestInBase64 := base64.StdEncoding.EncodeToString(digest)

	println("")
	println("Signature:", digestInBase64)
	println("Nonce:", *nonceFlag)
}

func init() {
	flag.Usage = func() {
		cmd := os.Args[0]
		if len(cmd) > 50 {
			cmd = "go run hmac-signature.go"
		}

		fmt.Fprintf(os.Stderr, "Usage of `%s`:\n", cmd)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample: %s -key=\"key\" -http_method=\"get\" -url=\"example.com/path\" -nonce=\"123\" -params=\"c=3 b=2 a=1\"\n", cmd)

		os.Exit(0)
	}
}
