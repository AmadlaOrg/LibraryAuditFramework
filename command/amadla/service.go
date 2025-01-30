package amadla

import (
	"github.com/AmadlaOrg/LibraryAuditFramework/display"
	"github.com/spf13/cobra"
)

// NewAmadlaService to set up the support service
func NewAmadlaService(supportedApplications, supportedEntities map[string]string) IAmadla {
	amadla := &SAmadla{
		supportedApplications: supportedApplications,
		supportedEntities:     supportedEntities,
		amadlaCmd: &cobra.Command{
			Use:   "amadla",
			Short: "Amadla commands",
		},
	}

	amadla.amadlaCmd.Run = func(c *cobra.Command, args []string) {
		display.NewDisplayService(c, []string{"Category", "Supported", "Version Supported"}, amadla.Supported())
	}

	return amadla
}
