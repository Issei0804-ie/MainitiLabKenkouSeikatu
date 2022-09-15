package attendance

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func newUpdateCmd() *cobra.Command {
	subCmd := &cobra.Command{
		Use:   "update ",
		Short: "Get your attendance. If you not given arguments, it passes all week's attendance.",
	}
	subCmd.AddCommand(newCreateWeekCmd())
	return subCmd
}

func newUpdateWeekCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "week [user] [week id] [value]",
		Short:   "update data value",
		Example: "attendance update week sampleUser 0 10",
		RunE:    execUpdateWeekCmd,
	}
}

func execUpdateWeekCmd(_ *cobra.Command, args []string) error {
	if len(args) != 3 {
		return errors.New("invalid argument")
	}
	user := args[0]
	rawWeekID := args[1]
	rawValue := args[2]

	weekID, err := strconv.Atoi(rawWeekID)
	if err != nil {
		log.Println(err)
		return err
	}

	value, err := strconv.Atoi(rawValue)
	if err != nil {
		log.Println(err)
		return err
	}

	data, err := read(DataFile)
	if err != nil {
		log.Println(err)
		return err
	}
	data.Models = data.byUser(user)
	data.Models = setID(data.Models)
	timeWorked, err := newTimeWorked(value)
	if err != nil {
		log.Println(err)
		return err
	}
	models, err := update(timeWorked, weekID, data.Models)
	data.Models = models
	return nil
}
