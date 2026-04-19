package judge

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// Judge defines the interface for the judge subcommand.
type Judge interface {
	IsPassed() bool
	JudgeDetails() map[string]any
	JudgeResults() map[string]any
	JudgeCmd() *cobra.Command
}

// judgeImpl implements the Judge interface.
type judgeImpl struct {
	isPassed     bool
	judgeDetails map[string]any
	judgeResult  map[string]any
	judgeCmd     *cobra.Command
}

var (
	osOpen = os.Open
	osExit = os.Exit
)

// RunJudge is the callback type for judge plugin implementors.
//
// Params:
// - entity data reader
// Results:
// - isPass bool
// - details map[string]any
type RunJudge func(*io.Reader) (bool, map[string]any)

// attachFlags adds the entity input flags.
func (s *judgeImpl) attachFlags() {
	s.judgeCmd.Flags().StringP(
		"entity",
		"e",
		"",
		"Specify the entity file path (reads from stdin if omitted)",
	)

	// -o/--output flag for display format
	s.judgeCmd.Flags().StringP("output", "o", "table", "Output format: table, json, or yaml")
}

// inputEntity reads the entity from the --entity flag or stdin.
func (s *judgeImpl) inputEntity() (*io.Reader, error) {
	entityPath, err := s.judgeCmd.Flags().GetString("entity")
	if err != nil {
		return nil, err
	}

	var entity io.Reader

	if entityPath == "" || entityPath == "-" {
		// Read from stdin
		stat, err := os.Stdin.Stat()
		if err != nil {
			return nil, fmt.Errorf("failed to stat stdin: %v", err)
		}
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			entity = os.Stdin
		} else {
			return nil, fmt.Errorf("no entity file path or stdin input provided")
		}
	} else {
		entityFile, err := osOpen(entityPath)
		if err != nil {
			if os.IsNotExist(err) {
				s.judgeCmd.PrintErrln("Entity file does not exist")
			} else {
				s.judgeCmd.PrintErrf("Failed to open entity file: %v\n", err)
			}
			return nil, err
		}
		entity = entityFile
	}

	return &entity, nil
}

// judge runs the judge callback and sets exit code per UNIX protocol.
// Exit 0 = pass, Exit 1 = fail.
func (s *judgeImpl) judge(runJudge RunJudge) {
	inputEntity, err := s.inputEntity()
	if err != nil {
		s.judgeCmd.PrintErrln(err)
		osExit(2)
		return
	}

	pass, details := runJudge(inputEntity)

	s.isPassed = pass
	s.judgeDetails = details

	status := "failed"
	if pass {
		status = "pass"
	}

	s.judgeResult = map[string]any{
		"status":  status,
		"details": details,
	}

	if !pass {
		// Per UNIX plugin protocol: exit 1 on failure
		osExit(1)
	}
}

// IsPassed returns whether the judgment passed.
func (s *judgeImpl) IsPassed() bool {
	return s.isPassed
}

// JudgeDetails returns the judgment details.
func (s *judgeImpl) JudgeDetails() map[string]any {
	return s.judgeDetails
}

// JudgeResults returns the full judgment results including status.
func (s *judgeImpl) JudgeResults() map[string]any {
	return s.judgeResult
}

// JudgeCmd returns the cobra command.
func (s *judgeImpl) JudgeCmd() *cobra.Command {
	return s.judgeCmd
}
