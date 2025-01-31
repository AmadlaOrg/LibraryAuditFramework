package audit

import "github.com/spf13/cobra"

// NewAuditService to set up the audit service
func NewAuditService(runAudit RunAudit) IAudit {
	audit := &SAudit{
		auditCmd: &cobra.Command{
			Use:   "audit",
			Short: "Audit commands",
			Long:  ``,
		},
	}

	audit.attachFlags()

	audit.auditCmd.Run = func(cmd *cobra.Command, args []string) {
		audit.audit(runAudit)
	}

	return audit
}
