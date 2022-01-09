package app

import "fmt"

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}
}
