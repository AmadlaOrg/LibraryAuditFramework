package display

import "github.com/spf13/cobra"

func NewDisplayService(cmd *cobra.Command, tableHeaders []string, data map[string]any) IDisplay {
	display := &SDisplay{
		cmd:          cmd,
		tableHeaders: tableHeaders,
		data:         data,
	}

	display.attachFlags()

	return display
}
