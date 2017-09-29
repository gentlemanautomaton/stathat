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
	for _, stat := range list {
		fmt.Printf("[%s] %s\n", stat.ID, stat.Name)
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
