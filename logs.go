package mesos

import (
	"net/url"
	"strconv"
)

func (c *Client) GetStdOutOfTask(pid *Pid, directory string) ([]byte, error) {
	u := c.getBaseURLForLogs("stdout", pid, directory)
	body, err := c.GetBodyOfHTTPResponse(&u)
	return body, err
}

func (c *Client) GetStdErrOfTask(pid *Pid, directory string) ([]byte, error) {
	u := c.getBaseURLForLogs("stderr", pid, directory)
	body, err := c.GetBodyOfHTTPResponse(&u)
	return body, err
}

func (c *Client) getBaseURLForLogs(mode string, pid *Pid, directory string) url.URL {
	u := url.URL{
		Scheme:   "http",
		Host:     pid.Host + ":" + strconv.Itoa(pid.Port),
		Path:     "files/download.json",
		RawQuery: "path=" + directory + "/" + mode,
	}

	return u
}
