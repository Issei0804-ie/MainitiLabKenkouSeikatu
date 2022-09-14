package attendance

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"
)

var DataFile = ""

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	workDir := filepath.Join(homedir, ".config", "com", "issei0804")
	err = os.MkdirAll(workDir, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	path := filepath.Join(workDir, "attendance_data.json")
	DataFile = path
}

func newModel(user string, begin time.Time, end time.Time) Model {
	timeWorker, _ := newTimeWorked(0)
	return Model{
		User:       user,
		BeginWeek:  begin,
		EndWeek:    end,
		TimeWorked: timeWorker,
	}
}

type Data struct {
	Models []Model
}

func (d Data) byUser(user string) []Model {
	var filtered []Model
	for _, model := range d.Models {
		if model.User == user {
			filtered = append(filtered, model)
		}
	}
	return filtered
}

type Model struct {
	ID         int
	User       string
	BeginWeek  time.Time
	EndWeek    time.Time
	TimeWorked TimeWorked
}

func (d *Model) AddTimeWorked(timeWorked TimeWorked) {
	d.TimeWorked = d.TimeWorked.add(timeWorked)
}

type TimeWorked struct {
	Value int
}

func newTimeWorked(value int) (TimeWorked, error) {
	if value < 0 {
		return TimeWorked{}, errors.New("illegal value")
	}
	return TimeWorked{Value: value}, nil
}

func (d TimeWorked) add(other TimeWorked) TimeWorked {
	return TimeWorked{Value: d.Value + other.Value}
}

func setID(models []Model) []Model {
	var set []Model

	for i, model := range models {
		model.ID = i + 1
		set = append(set, model)
	}
	return set
}

func read(path string) (Data, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return Data{}, err
	}
	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		switch err.(type) {
		case *json.InvalidUnmarshalError:
			return Data{}, nil
		}
	}

	return data, nil
}

func write(path string, data Data) error {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, marshaled, 0644)
	if err != nil {
		return err
	}
	return nil
}

func touch(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = write(path, Data{})
	return err
}
