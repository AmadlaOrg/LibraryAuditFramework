package command

import (
	"bytes"
	"io"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	meta := PluginMeta{
		Name:        "judge-test",
		Version:     "1.0.0",
		Supports:    []string{"amadla.org/entity/test@^v1.0.0"},
		Description: "Test plugin",
	}
	runJudge := func(r *io.Reader) (bool, map[string]any) {
		return true, nil
	}

	New(cmd, meta, runJudge)

	var names []string
	for _, sub := range cmd.Commands() {
		names = append(names, sub.Use)
	}
	assert.Contains(t, names, "judge")
	assert.Contains(t, names, "info")
}

func TestInfoSubcommand(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	meta := PluginMeta{
		Name:        "judge-test",
		Version:     "1.0.0",
		Supports:    []string{"amadla.org/entity/test@^v1.0.0"},
		Description: "Test plugin",
	}
	runJudge := func(r *io.Reader) (bool, map[string]any) {
		return true, nil
	}

	New(cmd, meta, runJudge)

	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetArgs([]string{"info"})
	err := cmd.Execute()

	assert.NoError(t, err)
	assert.Contains(t, buf.String(), `"name":"judge-test"`)
	assert.Contains(t, buf.String(), `"version":"1.0.0"`)
	assert.Contains(t, buf.String(), `"supports"`)
}
