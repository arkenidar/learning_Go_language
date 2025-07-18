// https://www.facebook.com/profile/1595905519/search/?q=read-line.go

// https://www.facebook.com/dario.cangialosi.arkenidar/posts/pfbid04xW9wK3KxGjoU4BF1eKzfRVPKv9DreZTAQLExD4At363fyUQJCrfGE7Fcfb8mbYRl?__cft__[0]=AZUBS2Kmh-caedw_FyC41nHIdR0UFUVtZ0pjGfqvn39dLIXHc4z-mzlQ_bdD5DZdFP2qfvgY6RJEomul4ZCFHBEYk3__lgR115LAtPdVqjlWiusOMF0n-LHhJ37mMgScUmyUtVgZq0kZr1Wth2keuVl7HAJvTg_H860qMlb5CP7sOw&__tn__=%2CO%2CP-R

// https://onlinegdb.com/cH-Pmfwe4
// --
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSuffix(input, "\n")
	return input, nil
}
func main() {
	fmt.Println("Enter text lines:")
	for {
		fmt.Print(">>> ")
		input, err := ReadLine()
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		if input == "" {
			break
		}
		fmt.Println("You entered:", strconv.Quote(input))
	}
	fmt.Println("==================================")
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}

// cd /home/arkenidar/temporary/go-lang && rlwrap go run read-line.go
