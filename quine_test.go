package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuine(t *testing.T) {
	contents, err := os.ReadFile("main.go")
	require.NoError(t, err)

	output, err := exec.Command("go", "run", "main.go").CombinedOutput()
	require.NoError(t, err)

	assert.Equal(t, string(contents), string(output))
}
