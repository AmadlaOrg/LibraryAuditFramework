package amadla

import (
	"github.com/AmadlaOrg/LibraryAuditFramework/internal/display"
	"github.com/spf13/cobra"
)

// NewAmadlaService to set up the amadla service
func NewAmadlaService(supportedApplications, supportedEntities map[string]string) IAmadla {
	amadla := &SAmadla{
		supportedApplications: supportedApplications,
		supportedEntities:     supportedEntities,
		amadlaCmd: &cobra.Command{
			Use:   "amadla",
			Short: "Amadla supported applications and entities",
			Long:  `Displays in JSON (--json|-j), YAML (--yaml|-y) or table (default) format the supported applications and entities.`,
		},
	}

	amadla.amadlaCmd.Run = func(cmd *cobra.Command, args []string) {
		display.NewDisplayService(cmd, amadla.Supported()).
			SetTableHeaders([]string{"Category", "Supported", "Version Supported"}).
			Display()
	}

	return amadla
}
