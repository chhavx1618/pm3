package manager

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"pm3/process"
)

const dbFile = "./data/process.json"

func ensureDirs() {
	os.MkdirAll("./data", os.ModePerm)
	os.MkdirAll("./logs", os.ModePerm)
}

func loadDB() []process.Process {
	ensureDirs()
	data, err := os.ReadFile(dbFile)
	if err != nil {
		return []process.Process{}
	}
	var procs []process.Process
	json.Unmarshal(data, &procs)
	return procs
}

func saveDB(procs []process.Process) {
	data, _ := json.MarshalIndent(procs, "", "  ")
	os.WriteFile(dbFile, data, 0644)
}

func Start(script string) {
	procs := loadDB()
	id := fmt.Sprintf("proc_%d", time.Now().UnixNano())
	p := process.Create(script, id)
	procs = append(procs, p)
	saveDB(procs)
	fmt.Printf("Started %s as %s (pid %d)\n", script, id, p.PID)
}

func Stop(id string) {
	procs := loadDB()
	for i, p := range procs {
		if p.ID == id {
			exec.Command("kill", strconv.Itoa(p.PID)).Run()
			procs[i].Status = "stopped"
			saveDB(procs)
			fmt.Println("Stopped", id)
			return
		}
	}
	fmt.Println("No process with id:", id)
}

func Restart(id string) {
	procs := loadDB()
	for i, p := range procs {
		if p.ID == id {
			exec.Command("kill", strconv.Itoa(p.PID)).Run()
			newP := process.Create(p.Script, id)
			procs[i] = newP
			saveDB(procs)
			fmt.Println("Restarted", id)
			return
		}
	}
	fmt.Println("No process with id:", id)
}

func Delete(id string) {
	procs := loadDB()
	for i, p := range procs {
		if p.ID == id {
			exec.Command("kill", strconv.Itoa(p.PID)).Run()
			procs = append(procs[:i], procs[i+1:]...)
			saveDB(procs)
			fmt.Println("Deleted", id)
			return
		}
	}
	fmt.Println("No process with id:", id)
}

func List() {
	procs := loadDB()
	fmt.Printf("%-20s %-15s %-8s %-10s %s\n", "ID", "Script", "PID", "Status", "Started")
	for _, p := range procs {
		fmt.Printf("%-20s %-15s %-8d %-10s %s\n", p.ID, filepath.Base(p.Script), p.PID, p.Status, p.Created)
	}
}
