package gemini

import (
  	"net/http"
  	"bytes"
  	"encoding/json"
  	"fmt"
	"encoding/hex"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha512"
	"time"
)


type Client struct {
  BaseURL         string
  Passphrase      string
  Secret          string
  Key             string
}

func NewClient(secret, key, passphrase string) *Client {
	client := Client{
		BaseURL:    "http://api.sandbox.gemini.com/v1",
		Secret:     secret,
		Key:        key,
		Passphrase: passphrase,
	}

	return &client
}

func (c *Client) Request(method string, url string, params, result interface{}) (res *http.Response, err error) {
	var data []byte
	body := bytes.NewReader(make([]byte, 0))

	if params != nil {
		data, err = json.Marshal(params)
		if err != nil {
			return res, err
		}
		body = bytes.NewReader(data)
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, url)
	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return res, err
	}
	
	reqStr, _ := json.Marshal(&params)
	payload := base64.StdEncoding.EncodeToString([]byte(reqStr))

	signature := makeSig(c.Secret, payload)
	
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Content-Type", "text/plain")

	req.Header.Add("X-GEMINI-APIKEY", c.Key)
	req.Header.Add("X-GEMINI-PAYLOAD", payload)
	req.Header.Add("X-GEMINI-SIGNATURE", signature)	
	fmt.Println(req)

	client := http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		defer res.Body.Close()
		geminiError := Error{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&geminiError); err != nil {
			return res, err
		}
		return res, error(geminiError)
	}

	if result != nil {
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(result); err != nil {
			return res, err
		}
	}
	return res, nil
}

func Nonce() int64 {
	return time.Now().UnixNano()
}

func makeSig(secret, payload string) string {
	mac := hmac.New(sha512.New384, []byte(secret))
	mac.Write([]byte(payload))

	signature := hex.EncodeToString(mac.Sum(nil))

	return signature
}
