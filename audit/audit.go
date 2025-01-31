package audit

import (
	"github.com/AmadlaOrg/LibraryAuditFramework/internal/command"
	"github.com/AmadlaOrg/LibraryAuditFramework/internal/command/audit"
	"github.com/AmadlaOrg/LibraryFramework/cli"
	"github.com/spf13/cobra"
)

// Audit
func Audit(
	name, title, version string,
	supportedApplications, supportedEntities map[string]string,
	runAudit audit.RunAudit) {
	cli.New(name, title, version, func(c *cobra.Command) {
		command.NewCommandService(c, supportedApplications, supportedEntities, runAudit)
	})
}
