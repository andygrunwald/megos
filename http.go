package mesos

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Client) GetHTTPResponseFromCluster(f func(url.URL) url.URL) (*http.Response, error) {
	for _, instance := range c.Master {
		u := f(*instance)
		resp, err := c.GetHTTPResponse(&u)

		// If there is no error, we hit an instance / master that is online
		if err == nil {
			return resp, nil
		}
	}

	return nil, errors.New("No master online.")
}

func (c *Client) GetHTTPResponseFromLeader(f func(Pid) url.URL) (*http.Response, error) {
	u := f(*c.Leader)
	return c.GetHTTPResponse(&u)
}

func (c *Client) GetHTTPResponse(u *url.URL) (*http.Response, error) {
	resp, err := http.Get(u.String())

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetBodyOfHTTPResponse(u *url.URL) ([]byte, error) {
	resp, err := c.GetHTTPResponse(u)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
