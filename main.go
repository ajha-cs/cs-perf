package main

import (
	"fmt"

	"cloudsmith.io/orgcs/gorepo/usinggorightway/pkg"
)

func main() {
	message := pkg.Greet("World")
	fmt.Println(message)
}
