package judge

import "github.com/spf13/cobra"

// New creates and returns a new Judge.
func New(runJudge RunJudge) Judge {
	j := &judgeImpl{
		judgeCmd: &cobra.Command{
			Use:   "judge",
			Short: "Run validation/audit on an entity",
		},
	}

	j.attachFlags()

	j.judgeCmd.Run = func(cmd *cobra.Command, args []string) {
		j.judge(runJudge)
	}

	return j
}
