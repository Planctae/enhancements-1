package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/planctae/enhancements-tracking-ng/pkg/git"
)

// proposeCmd represents the propose command
var proposeCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure the reltrackr tool for use",
	Long: `configure walks through a series of interactive prompts to create
a valid configuration file for the reltrackr tool. Configure should only have to
be run one time.
`,
	Args: cobra.ExactArgs(0), // accept no arguments
	RunE: func(cmd *cobra.Command, args []string) error {
		validateReceiptsLocation := func(input string) error {
			if isAbs := path.IsAbs(input); isAbs != true {
				return errors.New("enhancements receipts repository must be a full path")
			}

			return nil
		}

		promptReceiptsLocation := promptui.Prompt{
			Label:    "full path to enhancements receipts repository",
			Validate: validateReceiptsLocation,
		}

		repoLocation, err := promptReceiptsLocation.Run()
		if err != nil {
			return err
		}

		_, err = os.Stat(repoLocation)
		switch {
		case os.IsNotExist(err):
			validateCloneChoice := func(input string) error {
				switch strings.ToLower(input) {
				case "y":
					return nil
				case "n":
					return nil
				case "yes":
					return nil
				case "no":
					return nil
				default:
					return fmt.Errorf("sorry, don't understand choice: %s. Try one of y|n|yes|no", input)
				}
			}

			promptCloneReceipts := promptui.Prompt{
				Label:    "it doesn't seem like you have a copy of the repo. Clone?",
				Validate: validateCloneChoice,
			}

			cloneRepo, err := promptCloneReceipts.Run()
			if err != nil {
				return err
			}

			if strings.HasPrefix(cloneRepo, "y") {
				// one day we'll properly fork this
				err = git.Clone("https://github.com/planctae/enhancements-tracking-ng", repoLocation, os.Stdout)
				if err != nil {
					return err
				}
			}

		case err != nil:
			return err
		}

		return nil
	},
}
