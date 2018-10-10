package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jswank/nodeping-go/nodeping"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Nodeping API Token: ")
	api_token, _ := r.ReadString('\n')

	tp := nodeping.BasicAuthTransport{
		Token: strings.TrimSpace(api_token),
	}

	client := nodeping.NewClient(tp.Client())

	checks, _, err := client.Checks.List()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving checks. %v", err)
		os.Exit(1)
	}

	for _, c := range checks {
		if c.Enable != "active" {
			fmt.Printf("This check is inactive. ")
		}
		fmt.Printf("id: %s, state: %d, label: %s\n", c.Id, c.State, c.Label)
	}

}
