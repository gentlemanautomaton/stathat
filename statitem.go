package stathat

import (
	"errors"
	"time"
)

// StatItem is a StatHat stat item
type StatItem struct {
	stathat        StatHat
	ID             string    `json:",omitempty"`
	Name           string    `json:",omitempty"`
	Public         bool      `json:",omitempty"`
	Counter        bool      `json:",omitempty"`
	Unit           string    `json:",omitempty"`
	ClassicKey     string    `json:",omitempty"`
	EmbedKey       string    `json:",omitempty"`
	DataReceivedAt time.Time `json:",omitempty"`
	CreatedAt      time.Time `json:",omitempty"`
}

// {"id":"ABCD"
//  "name":"Some Stat Name Here"
//  "public":false
//  "counter":false
//  "unit":"smoots"
//  "classic_key":"xxxxx"
//  "embed_key":"xxxxx"
//  "data_received_at":1487412183
//  "created_at":1414232922}

// ErrStatItemInvalid means that a private value is unset, which will cause helper methods to fail.
// This means that you can not create your own `StatItem`, but instead get them from `StatHat.Stat`, `StatHat.StatList`, or `StatHat.StatListAll` if you wish to use the helper methods.
var ErrStatItemInvalid = errors.New("StatItem is invalid")
