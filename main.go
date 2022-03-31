package main

import (
	"fmt"
	"html"
)

func main() {
	var aConTilde = html.UnescapeString("&aacute;")
	fmt.Println(aConTilde)
}
