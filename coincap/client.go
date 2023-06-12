package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c Client) GetAssests() ([]Asset, error) {
	resp, err := c.client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r assetsResponse

	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.Assets, nil

}

func (c Client) GetAssest(name string) (*Asset, error) {
	url := fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r assetResponse

	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r.Asset, nil

}
