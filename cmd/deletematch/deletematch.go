package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/gentlemanautomaton/stathat"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide access token as first argument.  See also https://www.stathat.com/access")
	}
	if len(os.Args) < 3 {
		fmt.Println("Provide stat name regex pattern as second argument.")
	}
	if len(os.Args) < 4 {
		fmt.Println("Provide exactly this value as third argument: TRYMATCH")
		fmt.Println("Or if you want the deletion to actually happen: DELETEMYDATA")
		os.Exit(1)
	}
	token := os.Args[1]
	pattern := os.Args[2]
	mode := os.Args[3]
	if mode != "TRYMATCH" && mode != "DELETEMYDATA" {
		fmt.Println("Provide exactly this value as third argument: TRYMATCH")
		fmt.Println("Or if you want the deletion to actually happen: DELETEMYDATA")
		os.Exit(1)
	}
	match, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := stathat.New().Token(token)
	list, err := s.StatListAll()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, stat := range list {
		if match.MatchString(stat.Name) {
			fmt.Printf("[%s] %s\n", stat.ID, stat.Name)
			if mode == "DELETEMYDATA" {
				msg, e := stat.Delete()
				fmt.Println("\t(DELETE) " + msg)
				if err != nil {
					fmt.Println(e.Error())
				}
			} else {
				fmt.Println("\t(NOOP) Not deleting in TRYMATCH mode.")
			}
		}
	}
}
