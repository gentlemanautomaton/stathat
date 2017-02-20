package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dustywilson/stathat"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	token    = kingpin.Arg("token", "StatHat access token (https://www.stathat.com/access)").Required().String()
	stats    = kingpin.Arg("stats", "stat IDs").Required().Strings()
	start    = kingpin.Flag("start", "data start").Short('s').Int64()
	interval = kingpin.Flag("interval", "data interval").Short('i').String()
	period   = kingpin.Flag("period", "data period").Short('p').String()
	timezone = kingpin.Flag("timezone", "timezone, cosmetic").Short('z').String()
)

func main() {
	kingpin.Parse()

	var startTime time.Time
	if *start > 0 {
		startTime = time.Unix(*start, 0)
	}
	s := stathat.New().Token(*token)
	ds, err := s.Get(stathat.GetOptions{
		Start:    &startTime,
		Period:   *period,
		Interval: *interval,
	}, *stats...)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	loc := time.Now().Location()
	if timezone != nil {
		var err error
		loc, err = time.LoadLocation(*timezone)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	for _, stat := range ds {
		fmt.Printf("[%s] [%s]\n", stat.Name, stat.Timeframe)
		for _, point := range stat.Points {
			fmt.Printf("\t%s\t%f\n", point.Time.In(loc), point.Value)
		}
	}
}
