package display

import "github.com/spf13/cobra"

func NewDisplayService(cmd *cobra.Command, data map[string]any) IDisplay {
	display := &SDisplay{
		cmd:  cmd,
		data: data,
	}

	display.attachFlags()

	return display
}
