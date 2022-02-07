package config

import "time"

type Smtp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"host"`
	Port     int    `json:"port"`
}

type Mail struct {
	Smtp       Smtp     `json:"smtp"`
	Recipients []string `json:"recipients"`
}

type Process struct {
	Name    string `json:"name"`
	PidFile string `json:"pid_file"`
}

type Instance struct {
	Duration time.Duration `json:"duration"`
	Process  Process       `json:"process"`
	Mail     Mail          `json:"mail"`
}

type Instances []Instance
