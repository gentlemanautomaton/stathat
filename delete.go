package stathat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeleteStat will send a delete request to StatHat for this stat by ID.
// This is a destructive call.  Do not take this lightly.
// Deleted stats are able to be undeleted via https://www.stathat.com/v/stats/trash for 48 hours.
func (s StatHat) DeleteStat(stat string) (string, error) {
	if s.noop {
		return "", nil
	}

	// return "", errors.New("Would be deleting this: " + s.urlPrefix() + `/stats/` + stat)
	req, err := http.NewRequest(http.MethodDelete, s.apiPrefix()+`/stats/`+stat, nil)
	if err != nil {
		return "", err
	}
	resp, err := httpDo(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.Status != "" {
			return "", fmt.Errorf("unexpected http status code %d: %s", resp.StatusCode, resp.Status)
		}
		return "", fmt.Errorf("unexpected http status code %d", resp.StatusCode)
	}
	var respJSON struct {
		// {"msg":"stat deleted."}
		Msg string
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &respJSON)
	return respJSON.Msg, err
}

// Delete will send a delete request to StatHat for this StatItem.
// This is a destructive call.  Do not take this lightly.
// This is a helper that simply calls `StatHat.DeleteStat(StatItem.ID)` on your behalf.
func (i StatItem) Delete() (string, error) {
	if len(i.stathat.token) == 0 {
		return "", ErrStatItemInvalid
	}
	return i.stathat.DeleteStat(i.ID)
}
