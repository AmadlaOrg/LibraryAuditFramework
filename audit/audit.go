package audit

import (
	"fmt"
	"github.com/AmadlaOrg/LibraryAuditFramework/internal/command"
	"github.com/spf13/cobra"
	"os"
)

// TODO: Put all this code in LibraryFramework

// Audit
func Audit(name, title, version string) {
	appName := fmt.Sprintf("auditor-%s", name)
	appTitleName := fmt.Sprintf("Auditor %s", title)

	var (
		rootCmd = &cobra.Command{
			Use:     appName,
			Short:   appTitleName + " CLI application",
			Version: version,
		}
		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Print the version number of " + appName,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println(appName + " version " + version)
			},
		}
	)

	commandService := command.NewCommandService()

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(commandService.AmadlaCmd())
	rootCmd.AddCommand(commandService.AuditCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
