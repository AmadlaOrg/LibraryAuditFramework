package command

import (
	"github.com/AmadlaOrg/LibraryAuditFramework/internal/command/amadla"
	"github.com/AmadlaOrg/LibraryAuditFramework/internal/command/audit"
)

// NewCommandService
func NewCommandService() {
	amadla.NewAmadlaService()
	audit.NewAuditService()
}
