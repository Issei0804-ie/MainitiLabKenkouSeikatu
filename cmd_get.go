package attendance

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

func newGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [user]",
		Short: "Get your attendance. If you not given arguments, it passes all week's attendance.",
		RunE:  execGetCmd,
	}
}

func execGetCmd(_ *cobra.Command, args []string) error {
	user := ""
	if len(args) != 0 {
		user = args[0]
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

	_, err = os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = touch(path)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}

	data, err := read(path)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(user)
	log.Println(data)
	return nil
}
