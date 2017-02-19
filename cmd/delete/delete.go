package main

import (
	"fmt"
	"os"

	"github.com/dustywilson/stathat"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide access token as first argument.  See also https://www.stathat.com/access")
	}
	if len(os.Args) < 3 {
		fmt.Println("Provide stat ID as second argument.")
		os.Exit(1)
	}
	s := stathat.New(os.Args[1])
	msg, err := s.DeleteStat(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(msg)
}
