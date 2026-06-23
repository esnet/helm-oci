package cmd

import (
	"github.com/spf13/cobra"
)

func newTemplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "template <name> <release> [flags]",
		Short:              "Render templates for a bookmarked OCI chart",
		Args:               cobra.MinimumNArgs(2),
		DisableFlagParsing: true,
		ValidArgsFunction:  completeBookmarkNames,
		RunE: func(cmd *cobra.Command, args []string) error {
			name, release, extra := args[0], args[1], args[2:]
			b, err := getStore().Get(name)
			if err != nil {
				return err
			}

			helmArgs := append([]string{"template", release, b.URL}, extra...)
			return runner.Run(helmArgs, cmd.OutOrStdout(), cmd.ErrOrStderr())
		},
	}
	return cmd
}
