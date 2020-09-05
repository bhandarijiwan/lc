//https://leetcode.com/problems/defanging-an-ip-address/
package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {
	fmt.Println(defangIPaddr("255.100.50.0"))
}
