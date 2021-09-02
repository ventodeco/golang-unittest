package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")

	m.Run()

	fmt.Println("Setelah Unit Test")
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Vento",
			request: "Vento",
		},
		{
			name:    "Deco",
			request: "Deco",
		},
		{
			name:    "Tes",
			request: "Tes",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Vento", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Vento")
		}
	})

	b.Run("Deco", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Deco")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		HelloWorld("vento")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Vento",
			request:  "Vento",
			expected: "Hello Vento",
		},
		{
			name:     "Deco",
			request:  "Deco",
			expected: "Hello Deco",
		},
		{
			name:     "Coba",
			request:  "Coba",
			expected: "Hello Coba",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)

			require.Equal(t, test.expected, result)
		})
	}
}

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
