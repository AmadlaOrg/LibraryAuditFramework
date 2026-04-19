package command

import (
	"encoding/json"
	"fmt"

	"github.com/AmadlaOrg/LibraryJudgeFramework/internal/command/judge"
	"github.com/spf13/cobra"
)

// PluginMeta holds metadata for the info subcommand.
type PluginMeta struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Supports    []string `json:"supports"`
	Description string   `json:"description"`
}

// New wires up subcommands onto the root cobra command.
func New(
	cmd *cobra.Command,
	meta PluginMeta,
	runJudge judge.RunJudge) {

	judgeService := judge.New(runJudge)
	cmd.AddCommand(judgeService.JudgeCmd())

	// info subcommand — outputs plugin metadata as JSON per UNIX plugin protocol
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "Display plugin metadata",
		Run: func(c *cobra.Command, args []string) {
			data, _ := json.Marshal(meta)
			fmt.Fprintln(c.OutOrStdout(), string(data))
		},
	}
	cmd.AddCommand(infoCmd)
}
