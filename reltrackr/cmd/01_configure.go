package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/planctae/enhancements-tracking-ng/pkg/git"
	"github.com/planctae/enhancements-tracking-ng/pkg/settings"
)

// proposeCmd represents the propose command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure the reltrackr tool for use",
	Long: `configure walks through a series of interactive prompts to create
a valid configuration file for the reltrackr tool. Configure should only have to
be run one time.
`,
	Args: cobra.ExactArgs(0), // accept no arguments
	RunE: func(cmd *cobra.Command, args []string) error {
		promptReceiptsLocation := promptui.Prompt{
			Label:    "path to enhancements receipts repository",
			Validate: validatePossiblyExpandablePath,
		}

		repoLocation, err := promptReceiptsLocation.Run()
		if err != nil {
			return err
		}

		expandedLocation, err := homedir.Expand(repoLocation)
		if err != nil {
			return err
		}

		_, err = os.Stat(expandedLocation)
		switch {
		case os.IsNotExist(err):
			promptCloneReceipts := promptui.Prompt{
				Label:    "it doesn't seem like you have a copy of the repo. Clone?",
				Validate: validateYNChoice,
			}

			cloneRepo, err := promptCloneReceipts.Run()
			if err != nil {
				return err
			}

			if strings.HasPrefix(cloneRepo, "y") {
				// one day we'll properly fork this
				err = git.Clone("https://github.com/kubernetes/enhancements", expandedLocation, os.Stdout)
				if err != nil {
					return err
				}
			}

		case err != nil:
			return err
		}

		promptSaveSettings := promptui.Prompt{
			Label:    "save settings?",
			Validate: validateYNChoice,
		}

		saveSettings, err := promptSaveSettings.Run()
		if err != nil {
			return err
		}

		if strings.HasPrefix(saveSettings, "y") {
			runtime, err := settings.NewRuntime(expandedLocation)
			if err != nil {
				return err
			}

			promptSaveLocation := promptui.Prompt{
				Label:    "file with *.yaml extension to write",
				Validate: validatePossiblyExpandablePath,
			}

			saveLocation, err := promptSaveLocation.Run()
			if err != nil {
				return err
			}

			expandedLocation, err = homedir.Expand(saveLocation)
			if err != nil {
				return err
			}

			err = runtime.Persist(expandedLocation)
			if err != nil {
				return err
			}
		}

		fmt.Println("")
		fmt.Println("congratulations! preflight check complete")
		return nil
	},
}

var validateYNChoice = func(input string) error {
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
		return fmt.Errorf("sorry, didn't understand choice: %s. Try one of y|n|yes|no", input)
	}
}

var validatePossiblyExpandablePath = func(input string) error {
	_, err := homedir.Expand(input)
	return err
}
