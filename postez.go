package stathat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Kind indicates the type of stat posting
type Kind int

// Kinds
const (
	KindValue Kind = iota + 1
	KindCounter
)

// ErrKindMissing means the `Kind` was missing.
var ErrKindMissing = errors.New("kind missing")

// PostEZ posts the value to the stat using the EZ API.
// If you supply a non-nil time, it'll use that as the time value, otherwise it will not send a time value.
// This is probably what you want to do if posting stats.
func (s StatHat) PostEZ(name string, kind Kind, v float64, t *time.Time) error {
	u, _ := url.Parse(s.ezPrefix())
	q := u.Query()

	if len(s.ezkey) == 0 {
		return ErrMissingEZKey
	}
	q.Add("ezkey", s.ezkey)
	q.Add("stat", name)

	if t != nil && !t.IsZero() {
		q.Add("t", strconv.FormatInt(t.Unix(), 10))
	}

	if kind == KindValue {
		q.Add("value", strconv.FormatFloat(v, 'g', -1, 64))
	} else if kind == KindCounter {
		q.Add("count", strconv.FormatFloat(v, 'g', -1, 64))
	} else {
		return ErrKindMissing
	}

	u.RawQuery = q.Encode()
	fmt.Println(u)
	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return err
	}
	resp, err := httpDo(req)
	if err != nil {
		return err
	}
	var respJSON struct {
		// {"msg":"stat deleted."}
		Msg string
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &respJSON)
	if respJSON.Msg != "ok" {
		err = errors.New(respJSON.Msg)
	}
	return err
}
