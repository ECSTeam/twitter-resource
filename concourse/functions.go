package concourse

// adapted from https://github.com/concourse/s3-resource/blob/master/utils.go

import (
  "fmt"
  "encoding/json"
  "os"

	"github.com/mitchellh/colorstring"
)

func Fatal(doing string, err error) {
	Sayf(colorstring.Color("[red]error %s: %s\n"), doing, err)
	os.Exit(1)
}

func Sayf(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
}

func ReadRequest(request *interface{}) {
  if err := json.NewEncoder(os.Stdin).Encode(request); err != nil {
    Fatal("Error reading request: %v\n", err)
  }
}

func WriteResponse(response interface{}) {
  if err := json.NewDecoder(os.Stdout).Decode(response); err != nil {
    Fatal("Error writing response: %v\n", err)
  }
}
