package attendance

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

func

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

	_, err := os.Lstat(DataFile)
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
