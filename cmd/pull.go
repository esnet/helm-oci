package cmd

import (
	"github.com/spf13/cobra"
)

func newPullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "pull <name> [flags]",
		Short:              "Pull a bookmarked OCI chart",
		Args:               cobra.MinimumNArgs(1),
		DisableFlagParsing: true,
		ValidArgsFunction:  completeBookmarkNames,
		RunE: func(cmd *cobra.Command, args []string) error {
			name, extra := args[0], args[1:]
			b, err := getStore().Get(name)
			if err != nil {
				return err
			}

			helmArgs := append([]string{"pull", b.URL}, extra...)
			return runner.Run(helmArgs, cmd.OutOrStdout(), cmd.ErrOrStderr())
		},
	}
	return cmd
}
