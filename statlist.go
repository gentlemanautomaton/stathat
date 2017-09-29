package stathat

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// StatList returns a list of your stats.
// There is a maximum of 10000 results returned per the StatHat API.
// To get more, run it again with an offset of 10000 (or in multiples of 10000).
func (s StatHat) StatList(offset int) ([]StatItem, error) {
	req, err := http.NewRequest(http.MethodGet, s.apiPrefix()+`/statlist?offset=`+strconv.Itoa(offset), nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpDo(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	list := make([]StatItem, 0, 10000)
	j := json.NewDecoder(resp.Body)
	err = j.Decode(&list)
	for i := range list {
		list[i].stathat = s
	}
	return list, err
}

// StatListAll automatically runs StatList until all stats are collected.
// This is simply a helper.  If it doesn't work the way you'd like, you'll want to implement your own.
// If there is an error, it'll return all of what it has collected so far, but it won't necessarily be the full list; be sure to check err.
func (s StatHat) StatListAll() ([]StatItem, error) {
	offset := 0
	var all []StatItem
	for {
		list, err := s.StatList(offset)
		offset += len(list)
		all = append(all, list...)
		if err != nil {
			return all, err
		}
		if len(list) == 0 {
			break
		}
	}
	return all, nil
}
