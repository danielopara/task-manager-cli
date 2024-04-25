package main

import (
	"path/filepath"

	"github.com/danielopara/task-manager/cmd"
	"github.com/danielopara/task-manager/cmd/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		print(err)
	}
	cmd.RootCmd.Execute()
}