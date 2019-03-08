package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new enhancements tracking directory",
	Long: `init creates a new release directory for tracking
enhancements
`,
	Args: cobra.ExactArgs(0), // accept no arguments
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime, err := settings.Open(cfgFile)
		if err != nil {
			return nil
		}

		err = directory.New(runtime.ReceiptsLocation(), releaseName)
		if err != nil {
			return nil
		}

		println("")
		println("Please make sure everything looks ok (`git status`) and submit these changes to the enhancements tracking reposistory. Happy releasing!")

		return nil
	},
}
