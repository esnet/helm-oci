package cmd

import (
	"fmt"

	"github.com/esnet/helm-oci/pkg/registry"
	"github.com/spf13/cobra"
)

func newVersionsCmd() *cobra.Command {
	return &cobra.Command{
		Use:               "versions <name>",
		Short:             "List available versions for a bookmarked OCI chart",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeBookmarkNames,
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := getStore().Get(args[0])
			if err != nil {
				return err
			}

			tags, err := registry.ListTags(b.URL, false)
			if err != nil {
				return fmt.Errorf("listing versions for %q: %w", b.Name, err)
			}

			if len(tags) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "No versions found for %q\n", b.Name)
				return nil
			}

			for _, tag := range tags {
				fmt.Fprintln(cmd.OutOrStdout(), tag)
			}
			return nil
		},
	}
}
