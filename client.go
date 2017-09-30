package gemini

import (
  	"net/http"
  	"encoding/json"
  	"fmt"
	"encoding/hex"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha512"
	"time"
	"io/ioutil"
)

//Client is a Basic client which has information to send to the Gemini API
type Client struct {
  BaseURL         string
  Passphrase      string
  Secret          string
  Key             string
}

//NewClient creates a new client with the given API Secret, API Key and Passphrase (although I will probably remove passphrase.
func NewClient(secret, key, passphrase string) *Client {
	client := Client{
		BaseURL:    "https://api.gemini.com",
		Secret:     secret,
		Key:        key,
		Passphrase: passphrase,
	}

	return &client
}

//Request takes a client, an HTTP method, a URL and a set of params and creates an HTTP request which the Gemini Endpoints can respond to.
func (c *Client) Request(method string, url string, params Request, result interface{}) (res *http.Response, err error) {

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, params.GetRoute())
	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return res, err
	}

	
	payload := base64.StdEncoding.EncodeToString(params.GetPayload())

	h := hmac.New(sha512.New384, []byte(c.Secret))
	h.Write([]byte(payload))
	sig := h.Sum(nil)



	
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Content-Type", "text/plain")

	req.Header.Add("X-GEMINI-APIKEY", c.Key)
	req.Header.Add("X-GEMINI-PAYLOAD", payload)
	req.Header.Add("X-GEMINI-SIGNATURE", hex.EncodeToString(sig))	

	client := http.Client{}
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	fmt.Println(req)

	if res.StatusCode != 200 {
		defer res.Body.Close()
		geminiError := Error{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&geminiError); err != nil {
			return res, err
		}
		return res, error(geminiError)
	}
	
	var body []byte

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	
	json.Unmarshal(body, &result)
	
	return res, nil
}

//Nonce returns a nonce using a time, so it is always increasing
func Nonce() int64 {
	return time.Now().UnixNano()
}

