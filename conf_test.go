package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	os.Clearenv()
	os.Setenv("x", "y")

	type config struct {
		A string
	}

	c := config{
		A: "HELLO",
	}

	err := Load(c)

	assert.NoError(t, err)
}
