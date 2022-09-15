package attendance

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
)

func newCreateWeekCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "week [user] [begin week - end week]",
		Short:   "Create week.",
		Example: "attendance create week sampleUser 20210701-20210707",
		RunE:    execCreateWeek,
	}
}

func execCreateWeek(_ *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("invalid argument")
	}
	user := args[0]
	rawWeek := args[1]
	begin, end, err := parseWeek(rawWeek)
	if err != nil {
		log.Println(err)
		return err
	}

	model := Model{
		ID:         0,
		User:       user,
		BeginWeek:  begin,
		EndWeek:    end,
		TimeWorked: TimeWorked{},
	}

	data, err := read(DataFile)
	if err != nil {
		log.Println(err)
		return err
	}

	data.Models = append(data.Models, model)
	err = write(DataFile, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
