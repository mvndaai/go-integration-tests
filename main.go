//go:generate go mod tidy
//go:generate go test ./...
//go:generate go test ./integration/...

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mvndaai/go-integration-tests/echo"
)

func main() {
	out := echo.Echo(strings.Join(os.Args[1:], " "))
	fmt.Println(out)
}
