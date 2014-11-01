package stop

import (
	"os"
	"testing"

	"github.com/deis/systemd/logger"
	"github.com/stretchr/testify/assert"
)

func init() {
}

func TestBasic(t *testing.T) {
	assert.Equal(t, "", "")
}
