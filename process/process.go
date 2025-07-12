package process

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Process struct {
	ID      string `json:"id"`
	Script  string `json:"script"`
	PID     int    `json:"pid"`
	Status  string `json:"status"`
	Created string `json:"created"`
}

func Create(script string, id string) Process {
	outFile, _ := os.OpenFile(fmt.Sprintf("./logs/%s.out.log", id), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	errFile, _ := os.OpenFile(fmt.Sprintf("./logs/%s.err.log", id), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	cmd := exec.Command("node", script)
	cmd.Stdout = outFile
	cmd.Stderr = errFile
	cmd.Start()

	return Process{
		ID:      id,
		Script:  script,
		PID:     cmd.Process.Pid,
		Status:  "running",
		Created: time.Now().Format(time.RFC3339),
	}
}
