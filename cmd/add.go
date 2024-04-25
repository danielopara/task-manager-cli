package cmd

import (
	"fmt"
	"strings"

	"github.com/danielopara/task-manager/cmd/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := strings.Join(args, " ")
		_, err := db.CreateTask(tasks)
		if err != nil {
			fmt.Printf("Something went wrong")
		}
		fmt.Printf("Added \"%s\" to your task list", tasks)
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}