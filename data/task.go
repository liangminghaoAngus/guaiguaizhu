package data

import "time"

type Task struct {
	ID        int    `json:"id"`
	PreTaskID int    `json:"pre_task_id"`
	Name      string `json:"name"`
	Intro     string `json:"intro"`
	Target    string `json:"target"`
}

func (t *Task) TableName() string {
	return "task"
}

type PlayerTask struct {
	GameID     int       `json:"game_id"`
	TaskID     int       `json:"task_id"`
	IsFinish   int       `json:"is_finish"` // 0 no 1 yes
	CreateTime time.Time `json:"create_time"`
	FinishTime time.Time `json:"finish_time"`
}

func (p *PlayerTask) TableName() string {
	return "player_task"
}

// task only main,so it's just one by one
