package audit

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

type IAudit interface {
	IsPassed() bool
	AuditDetails() map[string]any
	AuditResults() map[string]any
}

type SAudit struct {
	isPassed     bool
	auditDetails map[string]any
	auditResult  map[string]any
	auditCmd     *cobra.Command
}

var (
	osOpen = os.Open
)

// Process
//
// Params:
// - entity data
// Results:
// - isPass bool
// - details map[string]any
type Process func(*io.Reader) (bool, map[string]any)

// attachFlags
func (s *SAudit) attachFlags() {
	// 1. Setup of the `entity` flags
	s.auditCmd.Flags().StringP(
		"entity",
		"e",
		"",
		"Specify the entity file path (optional)",
	)

	// 2. The entity flag is required
	err := s.auditCmd.MarkFlagRequired("entity")
	if err != nil {
		s.auditCmd.Println(err)
		return
	}
}

// TODO: Breakdown the function (find a way to test performance and linking)
// TODO: Breakdown for easier testing and making the code more readable
// inputEntity
func (s *SAudit) inputEntity() (*io.Reader, error) {
	entityPath, err := s.auditCmd.Flags().GetString("entity")
	if err != nil {
		// TODO:
		return nil, err
	}

	var entity io.Reader

	// 3. Check if the entity path or stdin is available
	if entityPath == "" {
		// If no `--entity` flag is provided, use stdin
		stat, err := os.Stdin.Stat()
		if err != nil {
			return nil, fmt.Errorf("Failed to stat stdin: %v\n", err)
		}
		// Check if stdin has data (not a terminal input)
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			entity = os.Stdin
		} else {
			s.auditCmd.Println("No entity file path or stdin input provided")
			if err := s.auditCmd.Execute(); err != nil {
				s.auditCmd.Println(err)
			}

			// TODO:
			return nil, err
		}
	} else {
		// Open the entity file if the `--entity` flag is provided
		entityFile, err := osOpen(entityPath)
		if err != nil {
			if os.IsNotExist(err) {
				s.auditCmd.Println("Entity file does not exist")
			} else {
				s.auditCmd.Printf("Failed to open entity file: %v\n", err)
			}

			// TODO:
			return nil, err
		}
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				s.auditCmd.Printf("Failed to close entity file: %v\n", err)
			}
		}(entityFile)
		entity = entityFile
	}

	// TODO: Test linking
	return &entity, nil
}

// audit
func (s *SAudit) audit(process Process) {
	inputEntity, err := s.inputEntity()
	if err != nil {
		s.auditCmd.Println(err)
	}

	pass, details := process(inputEntity)

	s.isPassed = pass
	s.auditDetails = details

	status := "failed"
	if pass {
		status = "pass"
	}

	s.auditResult = map[string]any{
		"status":  status,
		"details": details,
	}
}

// IsPassed
func (s *SAudit) IsPassed() bool {
	return s.isPassed
}

// AuditDetails
func (s *SAudit) AuditDetails() map[string]any {
	return s.auditDetails
}

// AuditResults
func (s *SAudit) AuditResults() map[string]any {
	return s.auditResult
}
