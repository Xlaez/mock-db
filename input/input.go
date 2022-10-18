package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// take inputs for al struct string fields
func InputValidatorStrings() string {
	var i string
	reader := bufio.NewReader(os.Stdin)

	i, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	i = strings.TrimSpace(i)
	return i
}

// takes input for all struct int fields
func InputValidatorInt() int {
	var i int
	_, err := fmt.Scanln(&i)
	if err != nil {
		panic(err)
	}
	return i
}
