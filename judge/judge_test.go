package judge

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"judge-test", "version"}

	assert.NotPanics(t, func() {
		New("judge-test", "Judge Test", "1.0.0",
			[]string{"amadla.org/entity/test@^v1.0.0"},
			"Test judge plugin",
			func(r *io.Reader) (bool, map[string]any) {
				return true, nil
			})
	})
}
