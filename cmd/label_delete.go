package cmd

import (
	"log"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"github.com/zaquestion/lab/internal/action"
	lab "github.com/zaquestion/lab/internal/gitlab"
)

var labelDeleteCmd = &cobra.Command{
	Use:              "delete [remote] <name>",
	Aliases:          []string{"remove"},
	Short:            "Deletes an existing label",
	Long:             ``,
	PersistentPreRun: LabPersistentPreRun,
	Args:             cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rn, name, err := parseArgsRemoteAndProject(args)
		if err != nil {
			log.Fatal(err)
		}

		labels, err := MapLabels(rn, []string{name})
		if err != nil {
			log.Fatal(err)
		}

		err = lab.LabelDelete(rn, labels[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	labelCmd.AddCommand(labelDeleteCmd)
	carapace.Gen(labelCmd).PositionalCompletion(
		action.Remotes(),
	)
}
