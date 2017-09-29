package main

import (
	"fmt"
	"os"

	"github.com/gentlemanautomaton/stathat"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide access token as first argument.  See also https://www.stathat.com/access")
		os.Exit(1)
	}
	s := stathat.New().Token(os.Args[1])
	list, err := s.StatListAll()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Total Stats: %d\n", len(list))
}
