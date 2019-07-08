package error

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty content")
		return
	}
	response = fmt.Sprintf("echo: %s\n", request)
	return
}