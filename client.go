package gemini

import (
  "net/http"
  "bytes"
  "encoding/json"
  "fmt"
)


type Client struct {
  BaseURL         string
  Passphrase      string
  Secret          string
  Key             string
}

func NewClient(secret, key, passphrase string) *Client {
	client := Client{
		BaseURL:    "https://api.gemini.com/v1",
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
