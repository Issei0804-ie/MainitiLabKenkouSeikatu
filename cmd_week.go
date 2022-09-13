package attendance

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
	"time"
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
	user := args[0]
	rawWeek := args[1]

	layout := "20060102"
	splinted := strings.Split(rawWeek, "-")
	beginWeek, err := time.Parse(layout, splinted[0])
	if err != nil {
		return err
	}
	endWeek, err := time.Parse(layout, splinted[1])
	if err != nil {
		return err
	}

	model := Model{
		ID:         0,
		User:       user,
		BeginWeek:  beginWeek,
		EndWeek:    endWeek,
		TimeWorked: TimeWorked{},
	}
	return errors.New("hoge")
}
