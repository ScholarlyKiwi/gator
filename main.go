package main

import (
	"fmt"

	"github.com/ScholarlyKiwi/gator/internal/config"
)

func main() {
	gatorConfig, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	gatorConfig.SetUser("david")

	gatorConfig2, err := config.Read()
	fmt.Println(gatorConfig2)
}
