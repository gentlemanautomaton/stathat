package stathat

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Stat returns a single stat by name.
func (s StatHat) Stat(name string) (StatItem, error) {
	item := StatItem{}
	req, err := http.NewRequest(http.MethodGet, s.urlPrefix()+`/stat?name=`+url.PathEscape(name), nil)
	if err != nil {
		return item, err
	}
	resp, err := httpDo(req)
	if err != nil {
		return item, err
	}
	defer resp.Body.Close()
	j := json.NewDecoder(resp.Body)
	err = j.Decode(&item)
	item.stathat = s
	return item, err
}
