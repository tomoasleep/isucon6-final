package job

import "github.com/catatsuy/isucon6-final/portal/score"

type Job struct {
	ID     int    `json:"id"`
	TeamID int    `json:"teamID"`
	IPAddr string `json:"ipAddress"`
}

type Result struct {
	Job    *Job
	Output *score.Output
}
