package command

import (
	"github.com/AmadlaOrg/LibraryAuditFramework/internal/command/amadla"
	"github.com/AmadlaOrg/LibraryAuditFramework/internal/command/audit"
	"github.com/spf13/cobra"
)

// NewCommandService
func NewCommandService(
	cmd *cobra.Command,
	supportedApplications, supportedEntities map[string]string,
	runAudit audit.RunAudit) {
	amadlaService := amadla.NewAmadlaService(supportedApplications, supportedEntities)
	auditService := audit.NewAuditService(runAudit)

	cmd.AddCommand(amadlaService.AmadlaCmd())
	cmd.AddCommand(auditService.AuditCmd())
}
