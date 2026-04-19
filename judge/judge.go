package judge

import (
	"github.com/AmadlaOrg/LibraryJudgeFramework/internal/command"
	internaljudge "github.com/AmadlaOrg/LibraryJudgeFramework/internal/command/judge"
	"github.com/AmadlaOrg/LibraryFramework/cli"
	"github.com/spf13/cobra"
)

// RunJudge is the callback type that judge plugin implementors must satisfy.
type RunJudge = internaljudge.RunJudge

// New sets up the judge CLI application with UNIX plugin protocol support.
//
// The generated CLI includes:
//   - info subcommand: outputs plugin metadata as JSON
//   - judge subcommand: runs the judge callback on entity input
//   - version subcommand: prints the version
//
// Parameters:
//   - name: plugin binary name (e.g., "judge-application")
//   - title: display name (e.g., "Judge Application")
//   - version: semantic version (e.g., "1.0.0")
//   - supports: list of entity type URIs this plugin handles (e.g., ["amadla.org/entity/application@^v1.0.0"])
//   - description: short description of what the plugin does
//   - runJudge: callback that performs the validation/audit
func New(
	name, title, version string,
	supports []string,
	description string,
	runJudge RunJudge) {

	meta := command.PluginMeta{
		Name:        name,
		Version:     version,
		Supports:    supports,
		Description: description,
	}

	cli.New(name, title, version, func(c *cobra.Command) {
		command.New(c, meta, runJudge)
	})
}
