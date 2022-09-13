package attendance

import "github.com/spf13/cobra"

func newCreateCmd() *cobra.Command {
	subCmd := &cobra.Command{
		Use:   "create",
		Short: "Get your attendance. If you not given arguments, it passes all week's attendance.",
	}
	subCmd.AddCommand(newCreateWeekCmd())
	return subCmd
}
