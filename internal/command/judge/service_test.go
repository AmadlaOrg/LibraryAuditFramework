package judge

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	runJudge := func(r *io.Reader) (bool, map[string]any) {
		return true, nil
	}

	svc := New(runJudge)

	assert.NotNil(t, svc)
	assert.NotNil(t, svc.JudgeCmd())
	assert.Equal(t, "judge", svc.JudgeCmd().Use)
}
