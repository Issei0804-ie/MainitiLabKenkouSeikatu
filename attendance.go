package attendance

import (
	"github.com/spf13/cobra"
)

const cmdName = "attendance"

func NewAttendanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  cmdName,
		Long: "attendance tool.",
	}

	cmd.AddCommand(
		newGetCmd(),
		newCreateCmd(),
	)

	return cmd
}
