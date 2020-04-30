package cmd

import (
	"bytes"
	"testing"

	"github.com/jeandias/monitoring/cmd"
	"github.com/stretchr/testify/assert"
)

func TestGetOption(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("2\n"))

	result, err := cmd.GetOption(&stdin)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)
}
