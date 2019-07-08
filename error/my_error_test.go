package error

import (
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	for _, request := range []string{"", "hello world"} {
		fmt.Printf("request: %s\n", request)
		response, err := echo(request)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}

		fmt.Printf("response: %s\n", response)
	}
}
