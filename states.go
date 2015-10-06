package mesos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) GetStateFromCluster() (*State, error) {
	resp, err := c.GetHTTPResponseFromCluster(c.GetURLForStateFile)
	return c.parseStateResponse(resp, err)
}

func (c *Client) GetStateFromLeader() (*State, error) {
	resp, err := c.GetHTTPResponseFromLeader(c.GetURLForStateFilePid)
	return c.parseStateResponse(resp, err)
}

func (c *Client) GetStateFromPid(pid *Pid) (*State, error) {
	u := c.GetURLForStateFilePid(*pid)
	resp, err := c.GetHTTPResponse(&u)
	return c.parseStateResponse(resp, err)
}

func (c *Client) parseStateResponse(resp *http.Response, err error) (*State, error) {
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var state State
	err = json.Unmarshal(body, &state)
	if err != nil {
		return nil, err
	}

	c.State = &state

	return c.State, nil
}

func (c *Client) GetURLForStateFile(instance url.URL) url.URL {
	instance.Path = "master/state.json"
	return instance
}

func (c *Client) GetURLForStateFilePid(pid Pid) url.URL {
	host := pid.Host + ":" + strconv.Itoa(pid.Port)

	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   pid.Role + "/state.json",
	}

	return u
}
