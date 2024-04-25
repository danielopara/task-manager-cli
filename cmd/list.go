package cmd

import (
	"fmt"
	"os"

	"github.com/danielopara/task-manager/cmd/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Print("Something went wrong", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Print("No tasks available 😢")
			return
		}
		fmt.Println("Here are your following tasks")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
