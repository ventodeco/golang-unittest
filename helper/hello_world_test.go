package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("TestHello", func(t *testing.T) {
		result := HelloWorld("Vento")
		require.Equal(t, "Hello Vento", result)
	})

	t.Run("TestHelloWithSpace", func(t *testing.T) {
		result := HelloWorld("Deco")
		require.Equal(t, "Hello Deco", result)
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")

	m.Run()

	fmt.Println("Setelah Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Hanya Jalan di Windows")
	}

	result := HelloWorld("tes")
	require.Equal(t, "Hello tes", result)
}

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
