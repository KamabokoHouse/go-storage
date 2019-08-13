package application

import (
	"github.com/spf13/cobra"
)

// RootCmd ...
var (
	RootCmd = &cobra.Command{
		Use:   "go-storage",
		Short: "GCS tools cli",
		Long:  "GCS tools include upload,download,list. maybe.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

// Run execute cmd
func Run() error {
	err := RootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	cobra.OnInitialize()
}
