package stathat

import (
	"errors"
	"net/http"
)

// UserAgent string to be sent to StatHat.
// Override this by assigning a string right to it.
// This is encouraged.
var UserAgent = "github.com/dustywilson/stathat"

// ErrNotFound means the StatHat service reports that the requested object doesn't exist.
var ErrNotFound = errors.New("not found")

func httpDo(req *http.Request) (*http.Response, error) {
	c := &http.Client{}
	req.Header.Add("User-Agent", UserAgent)
	resp, err := c.Do(req)
	if err == nil && resp.StatusCode == 404 {
		err = ErrNotFound
	}
	return resp, err
}
