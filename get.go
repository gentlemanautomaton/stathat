package stathat

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// GetOptions are passed into Get to provide optional values
type GetOptions struct {
	Start    *time.Time
	Period   string
	Interval string
}

// Dataset is a dataset
type Dataset struct {
	Name      string
	Timeframe string
	Points    []Datapoint
}

// Datapoint is a datapoint
type Datapoint struct {
	Timestamp int64 `json:"time"`
	Value     float64
	Time      time.Time `json:"-"`
}

// Get returns a data for a stat.
// If `start` is nil, StatHat will use their default of `start = now - period`.
// `period` is the span of time from which to return data, starting at `start` and ending at `start+period`.  This is in the same format as `interval`.
// `interval` is in the format of `1h` meaning "one hour".  Other units available are in the set `[mhdwMy]`.
// See https://www.stathat.com/manual/export for accepted time unit format for `interval` and `period`.
// `stats` can be range from one to five stats.  Stats listed beyond five are ignored.
func (s StatHat) Get(opts GetOptions, stats ...string) ([]Dataset, error) {
	rawurl := s.apiPrefix() + `/data/`

	if len(stats) == 0 {
		return nil, ErrNotFound // FIXME: maybe not the best error to return
	}

	// (StatHat only cares about the first five stats requested and ignores any extra)
	for i := 0; i < 5 && i < len(stats); i++ {
		rawurl += url.PathEscape(stats[i]) + "/"
	}

	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	q := u.Query()

	if opts.Start != nil && !opts.Start.IsZero() {
		q.Add("start", strconv.FormatInt(opts.Start.Unix(), 10))
	}

	if len(opts.Period) > 0 || len(opts.Interval) > 0 {
		q.Add("t", opts.Period+opts.Interval)
	}

	u.RawQuery = q.Encode()
	rawurl = u.String()

	req, err := http.NewRequest(http.MethodGet, rawurl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpDo(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ds []Dataset
	j := json.NewDecoder(resp.Body)
	err = j.Decode(&ds)

	for dsI := range ds {
		for point := range ds[dsI].Points {
			ds[dsI].Points[point].Time = time.Unix(ds[dsI].Points[point].Timestamp, 0)
		}
	}

	return ds, err
}
