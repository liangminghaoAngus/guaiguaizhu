package data

import (
	"liangminghaoangus/guaiguaizhu/log"
	"time"
)

// just do kill enemy task now

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
	IsFinish   int       `json:"is_finish"` // 0 no 1 doing 2 done
	CreateTime time.Time `json:"create_time"`
	FinishTime time.Time `json:"finish_time"`
}

func (p *PlayerTask) TableName() string {
	return "player_task"
}

type PlayTaskItem struct {
	ID     int
	Name   string
	Intro  string
	Target string

	IsFinish int // 0 no 1 doing 2 done
}

// task only main,so it's just one by one

func GetPlayerTask(gameID int) []PlayTaskItem {
	res := make([]PlayTaskItem, 0)
	err := getDb().Table("player_task").
		Joins("task on task.id = player_task.task_id").
		Where("game_id = ?", gameID).
		Select("task.id as id,task.name as name,task.intro as intro,task.target as target,player_task.is_finish as is_finish").
		Find(&res).Error
	if err != nil {
		log.Error("%s", err.Error())
	}
	return res
}

func CreateTask(gameID int, taskID int) bool {
	err := getDb().Create(PlayerTask{
		GameID:     gameID,
		TaskID:     taskID,
		IsFinish:   1,
		CreateTime: time.Now(),
	}).Error
	if err != nil {
		log.Error("%s", err.Error())
	}
	return err == nil
}

func FinishTask(gameID int, taskID int) bool {
	err := getDb().Model(PlayerTask{}).
		Where("task_id = ? and gameID = ?", taskID, gameID).
		Updates(map[string]interface{}{
			"is_finish":   2,
			"finish_time": time.Now(),
		}).Error
	if err != nil {
		log.Error("%s", err.Error())
	}
	return err == nil
}
