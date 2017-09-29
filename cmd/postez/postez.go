package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gentlemanautomaton/stathat"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	ezkey     = kingpin.Arg("ezkey", "StatHat EZKey").Required().String()
	name      = kingpin.Arg("name", "stat name").Required().String()
	value     = kingpin.Arg("value", "stat value").Required().Float64()
	count     = kingpin.Flag("count", "act as a counter instead of value").Short('c').Bool()
	timestamp = kingpin.Flag("timestamp", "timestamp").Short('t').Int64()
)

func main() {
	kingpin.Parse()

	kind := stathat.KindValue
	if *count {
		kind = stathat.KindCounter
	}

	var t time.Time
	if *timestamp != 0 {
		t = time.Unix(*timestamp, 0)
	}

	s := stathat.New().EZKey(*ezkey)
	err := s.PostEZ(*name, kind, *value, &t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
