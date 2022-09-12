package attendance

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strconv"
)

func newGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [id]",
		Short: "Get your attendance. If you not given arguments, it passes all week's attendance.",
		RunE:  execGetCmd,
	}
}

func execGetCmd(_ *cobra.Command, args []string) error {
	id := 0
	if len(args) != 0 {
		rawID, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		id = rawID
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	workDir := filepath.Join(homedir, ".config", "com", "issei0804")
	err = os.MkdirAll(workDir, 0755)
	if err != nil {
		return err
	}
	path := filepath.Join(workDir, "attendance_data.json")
	lstat, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create()
		}
		return err
	}

	read()
	return nil
}
