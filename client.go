// Copyright (c) 2019 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package iex

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const apiURL = "https://cloud.iexapis.com/beta"

// Client models a client to consume the IEX Cloud API.
type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// NewClient creates a client with the given authorization toke.
func NewClient(token string, baseURL string) *Client {
	if baseURL == "" {
		baseURL = apiURL
	}
	return &Client{
		baseURL: baseURL,
		token:   token,
	}
}

// GetJSON gets the JSON data from the given endpoint.
func (c *Client) GetJSON(endpoint string, v interface{}) error {
	address := c.baseURL + endpoint + "?token=" + c.token
	resp, err := http.Get(address)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(b)
	log.Printf("Body as string = %s", bodyString)
	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}
