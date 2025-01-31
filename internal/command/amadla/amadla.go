package amadla

import "github.com/spf13/cobra"

type IAmadla interface {
	Supported() map[string]any
	Cmd() *cobra.Command
	AmadlaCmd() *cobra.Command
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
	// Set default values if needed
	s.processSupportedApplications()

	return map[string]any{
		"applications": s.supportedApplications,
		"entities":     s.supportedEntities,
	}
}

// AmadlaCmd
func (s *SAmadla) AmadlaCmd() *cobra.Command {
	return s.amadlaCmd
}

// processSupportedApplications
func (s *SAmadla) processSupportedApplications() {
	var (
		heryPresent  bool
		judgePresent bool
	)

	for appName, appVersion := range s.supportedApplications {
		if s.supportedEntities[appName] == "hery" && appVersion != "" {
			heryPresent = true
		} else if s.supportedEntities[appName] == "judge" && appVersion != "" {
			judgePresent = true
		}
	}

	if !heryPresent {
		s.supportedApplications["hery"] = "^0"
	}
	if !judgePresent {
		s.supportedApplications["judge"] = "^0"
	}
}
