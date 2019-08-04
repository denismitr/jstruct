package jstruct

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSONString(t *testing.T) {
	tcs := []struct {
		name string
	}{
		{"f1"},
		{"f2"},
		{"f3"},
		{"f4"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			in := fmt.Sprintf("./fixtures/%s.json", tc.name)
			out := fmt.Sprintf("./results/%s.go", tc.name)
			err := Convert(in, out, "fixtures", tc.name)

			expected, err := ioutil.ReadFile(fmt.Sprintf("./fixtures/%s.go", tc.name))
			assert.Nil(t, err)

			// Format the output file
			cmd := exec.Command("gofmt", "-w", out)
			cmd.Run()

			actual, err := ioutil.ReadFile(out)

			assert.Nil(t, err)
			assert.Equal(t, string(expected), string(actual))
		})
	}
}
