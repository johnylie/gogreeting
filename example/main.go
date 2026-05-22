package main

import (
	"fmt"

	"github.com/johnylie/gogreeting"
)

func main() {
	fmt.Println(
		gogreeting.Greeting(
			"NAME",
			gogreeting.ID,
			nil,
		),
	)

	fmt.Println(
		gogreeting.Greeting(
			"NAME",
			gogreeting.EN,
			nil,
		),
	)

	custom := &gogreeting.Translation{
		Morning: "Halo pagi",
	}

	fmt.Println(
		gogreeting.Greeting(
			"NAME",
			gogreeting.ID,
			custom,
		),
	)
}
