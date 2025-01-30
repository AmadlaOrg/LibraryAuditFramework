package amadla

import "github.com/spf13/cobra"

type IAmadla interface {
	Supported() map[string]any
	Cmd() *cobra.Command
}

type SAmadla struct {
	supportedApplications map[string]string
	supportedEntities     map[string]string
	amadlaCmd             *cobra.Command
}

// Cmd
func (s *SAmadla) Cmd() *cobra.Command {
	return s.amadlaCmd
}

// Supported
func (s *SAmadla) Supported() map[string]any {
	return map[string]any{
		"applications": s.processSupportedApplications(),
		"entities":     s.supportedEntities,
	}
}

// TODO:
// processSupportedApplications
func (s *SAmadla) processSupportedApplications() map[string]string {
	// TODO: Add default
	// TODO: Accept input
	return map[string]string{
		"hery":  "^0", // TODO:
		"judge": "^0", // TODO:
	}
}
