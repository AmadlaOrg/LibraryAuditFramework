package judge

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestIsPassed(t *testing.T) {
	s := &judgeImpl{isPassed: true}
	assert.True(t, s.IsPassed())
}

func TestJudgeDetails(t *testing.T) {
	details := map[string]any{"key": "value"}
	s := &judgeImpl{judgeDetails: details}
	assert.Equal(t, details, s.JudgeDetails())
}

func TestJudgeResults(t *testing.T) {
	result := map[string]any{"status": "pass"}
	s := &judgeImpl{judgeResult: result}
	assert.Equal(t, result, s.JudgeResults())
}

func TestJudgeCmd(t *testing.T) {
	cmd := &cobra.Command{Use: "judge"}
	s := &judgeImpl{judgeCmd: cmd}
	assert.Equal(t, cmd, s.JudgeCmd())
}

func TestJudge_Pass(t *testing.T) {
	originalExit := osExit
	defer func() { osExit = originalExit }()
	var exitCode int
	osExit = func(code int) { exitCode = code }

	runJudge := func(r *io.Reader) (bool, map[string]any) {
		return true, map[string]any{"check": "ok"}
	}

	s := &judgeImpl{judgeCmd: &cobra.Command{Use: "judge"}}
	s.attachFlags()

	tmpFile, err := os.CreateTemp(t.TempDir(), "entity")
	assert.NoError(t, err)
	_, _ = tmpFile.WriteString("test entity content")
	tmpFile.Close()

	originalOpen := osOpen
	defer func() { osOpen = originalOpen }()
	osOpen = func(name string) (*os.File, error) {
		return os.Open(tmpFile.Name())
	}

	s.judgeCmd.Flags().Set("entity", tmpFile.Name())
	s.judge(runJudge)

	assert.True(t, s.isPassed)
	assert.Equal(t, "pass", s.judgeResult["status"])
	assert.Equal(t, map[string]any{"check": "ok"}, s.judgeResult["details"])
	assert.Equal(t, 0, exitCode)
}

func TestJudge_Fail(t *testing.T) {
	originalExit := osExit
	defer func() { osExit = originalExit }()
	var exitCode int
	osExit = func(code int) { exitCode = code }

	runJudge := func(r *io.Reader) (bool, map[string]any) {
		return false, map[string]any{"check": "failed"}
	}

	s := &judgeImpl{judgeCmd: &cobra.Command{Use: "judge"}}
	s.attachFlags()

	tmpFile, err := os.CreateTemp(t.TempDir(), "entity")
	assert.NoError(t, err)
	tmpFile.Close()

	originalOpen := osOpen
	defer func() { osOpen = originalOpen }()
	osOpen = func(name string) (*os.File, error) {
		return os.Open(tmpFile.Name())
	}

	s.judgeCmd.Flags().Set("entity", tmpFile.Name())
	s.judge(runJudge)

	assert.False(t, s.isPassed)
	assert.Equal(t, "failed", s.judgeResult["status"])
	assert.Equal(t, 1, exitCode)
}

func TestInputEntity_FileNotFound(t *testing.T) {
	originalOpen := osOpen
	defer func() { osOpen = originalOpen }()
	osOpen = func(name string) (*os.File, error) {
		return nil, os.ErrNotExist
	}

	s := &judgeImpl{judgeCmd: &cobra.Command{Use: "judge"}}
	s.attachFlags()
	s.judgeCmd.Flags().Set("entity", "/nonexistent/path")

	reader, err := s.inputEntity()
	assert.Error(t, err)
	assert.Nil(t, reader)
}

func TestInputEntity_FileSuccess(t *testing.T) {
	tmpFile, err := os.CreateTemp(t.TempDir(), "entity")
	assert.NoError(t, err)
	_, _ = tmpFile.WriteString("test content")
	tmpFile.Close()

	s := &judgeImpl{judgeCmd: &cobra.Command{Use: "judge"}}
	s.attachFlags()
	s.judgeCmd.Flags().Set("entity", tmpFile.Name())

	reader, err := s.inputEntity()
	assert.NoError(t, err)
	assert.NotNil(t, reader)

	data, err := io.ReadAll(*reader)
	assert.NoError(t, err)
	assert.Equal(t, "test content", string(data))

	if closer, ok := (*reader).(io.Closer); ok {
		closer.Close()
	}
}

func TestInputEntity_EmptyPath_NoStdin(t *testing.T) {
	s := &judgeImpl{judgeCmd: &cobra.Command{Use: "judge"}}
	s.attachFlags()

	reader, err := s.inputEntity()
	if err != nil {
		assert.Nil(t, reader)
	}
}

func TestInputEntity_Dash_NoStdin(t *testing.T) {
	s := &judgeImpl{judgeCmd: &cobra.Command{Use: "judge"}}
	s.attachFlags()
	s.judgeCmd.Flags().Set("entity", "-")

	reader, err := s.inputEntity()
	if err != nil {
		assert.Nil(t, reader)
	}
}

func TestAttachFlags(t *testing.T) {
	s := &judgeImpl{judgeCmd: &cobra.Command{Use: "judge"}}
	s.attachFlags()

	flag := s.judgeCmd.Flags().Lookup("entity")
	assert.NotNil(t, flag)
	assert.Equal(t, "e", flag.Shorthand)

	outputFlag := s.judgeCmd.Flags().Lookup("output")
	assert.NotNil(t, outputFlag)
	assert.Equal(t, "o", outputFlag.Shorthand)
	assert.Equal(t, "table", outputFlag.DefValue)
}

func TestInputEntity_ReaderNotClosed(t *testing.T) {
	tmpFile, err := os.CreateTemp(t.TempDir(), "entity")
	assert.NoError(t, err)
	content := "entity data for reading"
	_, _ = tmpFile.WriteString(content)
	tmpFile.Close()

	s := &judgeImpl{judgeCmd: &cobra.Command{Use: "judge"}}
	s.attachFlags()
	s.judgeCmd.Flags().Set("entity", tmpFile.Name())

	reader, err := s.inputEntity()
	assert.NoError(t, err)

	buf := new(strings.Builder)
	_, err = io.Copy(buf, *reader)
	assert.NoError(t, err)
	assert.Equal(t, content, buf.String())

	if closer, ok := (*reader).(io.Closer); ok {
		closer.Close()
	}
}
