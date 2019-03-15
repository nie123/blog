package models

import "time"

type User struct {
	id          uint
	name        string
	pwd         string
	email       string
	last_time   time.Time
	last_ip     string
	state       int8
	create_time time.Time
	update_time time.Time
}

func (user *User) TableName() string {
	return "test"
	//return TableName("user")
}
