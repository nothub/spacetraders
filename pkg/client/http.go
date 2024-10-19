package st

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

func get(addr string, query map[string]string, model any) error {
	return do(http.MethodGet, addr, query, nil, model)
}

func post(addr string, query map[string]string, body any, model any) error {
	return do(http.MethodPost, addr, query, body, model)
}

func do(method string, addr string, query map[string]string, body any, model any) error {

	// TODO(#13): do proper url encoding

	// apply query params to url
	if query != nil && len(query) > 0 {
		q := url.Values{}
		for k, v := range query {
			q.Add(k, v)
		}
		addr = fmt.Sprintf("%s?%s", addr, q.Encode())
	}

	//log.Printf("%s %s %T %T\n", method, addr, body, model)

	var buf = bytes.NewBuffer(nil)
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, addr, buf)
	if err != nil {
		return err
	}

	if Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Token))
	}

	// we always want response data formatted as json
	req.Header.Set("Accept", "application/json")

	// if we send a body, it will always be json formatted
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
		// TODO(#14): what error codes have to be checked for an ErrorResponse body?
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
