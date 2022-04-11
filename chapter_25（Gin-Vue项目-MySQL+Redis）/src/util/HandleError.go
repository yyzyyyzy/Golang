package util

import (
	"fmt"
	"os"
)

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
