package attendance

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

func newModel(user string, week time.Time) Model {
	timeWorker, _ := newTimeWorked(0)
	return Model{
		User:       user,
		Week:       week,
		TimeWorked: timeWorker,
	}
}

type Model struct {
	User       string
	Week       time.Time
	TimeWorked MinWorked
}

func (d *Model) AddTimeWorked(timeWorked MinWorked) {
	d.TimeWorked = d.TimeWorked.add(timeWorked)
}

type MinWorked struct {
	Value int
}

func newTimeWorked(value int) (MinWorked, error) {
	if value < 0 {
		return MinWorked{}, errors.New("illegal value")
	}
	return MinWorked{Value: value}, nil
}

func (d MinWorked) add(other MinWorked) MinWorked {
	return MinWorked{Value: d.Value + other.Value}
}

func read(path string) ([]Model, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var models []Model
	err = json.Unmarshal(body, &models)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func write(path string, models Model) (bool, error) {
	marshaled, err := json.Marshal(models)
	if err != nil {
		return false, err
	}
	err = os.WriteFile(path, marshaled, 0644)
	if err != nil {
		return false, err
	}
	return true, nil
}
