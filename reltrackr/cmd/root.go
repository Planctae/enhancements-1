package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "reltrackr",
	Short: "release enhancements tracking tool",
	Long: `The enhancement tracking tool supports members of the release
team as well as contributors requesting that an enhancement be tracked during
the course of the release. The following workflows are supported:

-> A release team member opening the release:

   0. ensure you understand the enhancements tracking proceedure:
        https://github.com/planctae/enhancements-tracking-ng/blob/master/guides/release-team.md 
   1. ensure your application config is valid (reltrackr configure)
   2. reltrackr init release-1.15
   3. submit changes via a pull request for final review

-> A Kubernetes contributor importing an existing enhancements tracking issue:

   0. ensure your existing issue is ready for import:
	https://github.com/planctae/enhancements-tracking-ng/blob/master/guides/prepare-existing-tracking-issue.md
   1. ensure your application config is valid (reltrackr configure)
   2. reltrackr import enhancements/42 --release release-1.15

-> A Kubernetes contributor adding a new enhancement tracking receipt:

   0. ensure your enhancement is ready for tracking:
        https://github.com/planctae/enhancements-tracking-ng/blob/master/guides/enhancement-author.md
   1. ensure your application config is valid (reltrackr configure)
   2. reltrackr add --release release-1.15
   3. submit changes as a pull request for further collaboration with the release team

-> A release team member tracking the status of the release:

   0. ensure you understand the enhancements tracking proceedure:
        https://github.com/planctae/enhancements-tracking-ng/blob/master/guides/release-team.md 
   1. ensure your application config is valid (reltrackr configure)
   2. reltrackr status --release release-1.15
   3. submit changes via a pull request for further collaboration between team members and contributors

`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var release string

	rootCmd.PersistentFlags().StringVar(&release, "release", "", "release to operate on")

	rootCmd.AddCommand(configureCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
