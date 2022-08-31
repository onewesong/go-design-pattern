package builder

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConnection(t *testing.T) {
	timeout := 10 * time.Second
	c, err := NewConnection("localhost:8080", WithTimeout(timeout), WithCaching(true))
	assert.NoError(t, err)
	assert.Equal(t, "localhost:8080", c.addr)
	assert.Equal(t, true, c.cache)
	assert.Equal(t, timeout, c.timeout)
}
