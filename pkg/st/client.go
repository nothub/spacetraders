package st

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var HttpClient = http.DefaultClient
var Token string

const BaseUrl = "https://api.spacetraders.io/v2"

var ErrRateLimit = errors.New("rate limit")

// Meta details for pagination.
type Meta struct {

	// Shows the total amount of items of this kind that exist.
	//
	//   >= 0
	//
	// (required)
	Total int `json:"total"`

	// A page denotes an amount of items, offset from the first item. Each page
	// holds an amount of items equal to the limit.
	//
	//   >= 1
	//
	// (required)
	Page int `json:"page"`

	// The amount of items in each page. Limits how many items can be fetched
	// at once.
	//
	//   >= 1 && <= 20
	Limit int `json:"limit"`
}

func get(url string, model any) error {
	return do(http.MethodGet, url, nil, model)
}

func post(url string, body any, model any) error {
	return do(http.MethodPost, url, body, model)
}

func do(method string, url string, body any, model any) error {

	var buf = bytes.NewBuffer(nil)
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}

	if Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
	}

	// we always want response data formatted as json
	req.Header.Set("Accept", "application/json")

	// if we have a body, it is always in json format
	// reminder: empty json '{}' is >0 len too
	if buf.Len() > 0 {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := HttpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode/100 != 2 {
		// handle all non 2xx status codes
		switch res.StatusCode {
		case http.StatusTooManyRequests:
			// TODO: handle directly here in the lib? OR: return rate limit error with relevant infos
			return ErrRateLimit
		default:
			return fmt.Errorf("http response: %s", res.Status)
		}
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(model)
	if err != nil {
		return err
	}

	return nil
}
