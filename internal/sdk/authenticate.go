package sdk

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "net/http"
    "strconv"
    "strings"
    "time"
)

/**

Authenticate() is heavily based on fabioberger/coinbase-go implementation
From: https://github.com/fabioberger/coinbase-go/blob/8328539b18ab1c8b492ddf19321acf0b8c26f7a3/api_key_authentication.go

License as of 02/21/2021
https://github.com/fabioberger/coinbase-go/blob/master/LICENSE

(The MIT License)
Copyright (C) 2015, Alex Browne
Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in the
Software without restriction, including without limitation the rights to use, copy,
modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
and to permit persons to whom the Software is furnished to do so, subject to the
following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

**/
// API Key + Secret authentication requires a request header of the HMAC SHA-256
// signature of the "message" as well as an incrementing nonce and the API key
func Authenticate(req *http.Request, key, secret, method, path string, body []byte) error {
    // unix is time in seconds
    // path probably requires query params
    // body can be nil
    // method must be upper case
    unix := strconv.FormatInt(time.Now().Unix(), 10)
	message := unix + strings.ToUpper(method) + path
    if body != nil {
        message = message + string(body)
    }

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))

	signature := hex.EncodeToString(h.Sum(nil))

    req.Header.Set("CB-ACCESS-KEY",       key)
	req.Header.Set("CB-ACCESS-SIGN",      signature)
	req.Header.Set("CB-ACCESS-TIMESTAMP", unix)
    req.Header.Set("CB-VERSION", "2021-02-21")
    req.Header.Set("Content-Type", "application/json")

	return nil
}
