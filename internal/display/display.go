package display

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

type IDisplay interface {
	Display()
	SetTableHeaders(tableHeaders []string) IDisplay
}

type SDisplay struct {
	cmd          *cobra.Command
	tableHeaders []string
	data         map[string]any
}

var (
	jsonMarshal = json.Marshal
	yamlMarshal = yaml.Marshal
)

// attachFlags
func (s *SDisplay) attachFlags() {
	s.cmd.Flags().BoolP("json", "j", false, "Display output in JSON format")
	s.cmd.Flags().BoolP("yaml", "y", false, "Display output in YAML format")
}

// Display
func (s *SDisplay) Display() {
	// Retrieve flags
	jsonFlag, _ := s.cmd.Flags().GetBool("json")
	yamlFlag, _ := s.cmd.Flags().GetBool("yaml")

	// Error if both --json and --yaml are set
	if jsonFlag && yamlFlag {
		s.cmd.Println("Error: Cannot use both --json and --yaml flags at the same time.")
		os.Exit(1)
	}

	if jsonFlag {
		s.json()
	} else if yamlFlag {
		s.yaml()
	} else {
		s.table()
	}
}

// SetTableHeaders
func (s *SDisplay) SetTableHeaders(tableHeaders []string) IDisplay {
	s.tableHeaders = tableHeaders
	return s
}

// getTableHeaders
func (s *SDisplay) getTableHeaders() []string {
	if len(s.tableHeaders) == 0 {
		return []string{"Category", "Key", "Value"}
	}
	return s.tableHeaders
}

// table is the default way the data will be displayed and that is in a table (easy to read)
func (s *SDisplay) table() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(s.getTableHeaders())

	for category, data := range s.data {
		switch v := data.(type) {
		case map[string]string:
			for key, value := range v {
				table.Append([]string{category, key, value})
			}
		case map[string]any:
			for key, subData := range v {
				if subMap, ok := subData.(map[string]string); ok {
					for subKey, subValue := range subMap {
						table.Append([]string{
							category,
							fmt.Sprintf("%s.%s", key, subKey),
							subValue,
						})
					}
				}
			}
		}
	}

	table.Render()
}

// json is display's the data in JSON format
func (s *SDisplay) json() {
	jsonData, err := jsonMarshal(s.data)
	if err != nil {
		s.cmd.Println("Error encoding JSON:", err)
		os.Exit(1)
	}
	s.cmd.Println(string(jsonData))
}

// yaml display's the data in a YAML format
func (s *SDisplay) yaml() {
	yamlData, err := yamlMarshal(s.data)
	if err != nil {
		s.cmd.Println("Error encoding YAML:", err)
		os.Exit(1)
	}
	s.cmd.Println(string(yamlData))
}
