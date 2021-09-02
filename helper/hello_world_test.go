package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Vento")
	require.Equal(t, "Hello Vento", result, "Result must be 'Hello Vento'")
	fmt.Println("Test HalloWorldRequire Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Test")

	assert.Equal(t, "Hello Test", result, "Should same")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Vento")

	if result != "Hello Vento" {
		// unit test failed
		panic("Result is not Vento")
	}
}
