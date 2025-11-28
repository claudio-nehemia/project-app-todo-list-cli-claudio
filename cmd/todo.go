package cmd

import (
	"fmt"
	"os"
	"mini_project2/dto"
	"mini_project2/service"

	"github.com/spf13/cobra"
)

var (
	taskName   string
	priority   string
	updateId   int
	newStatus  string
	deleteId   int
	searchName string
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Simple CLI Todo App",
	Long: `Todo CLI App untuk manajemen task.
Command yang tersedia:
  add     : Menambahkan task baru
  list    : Menampilkan list task
  update  : Mengubah status task
  delete  : Menghapus task
  search  : Mencari task berdasarkan nama`,
}

// ====================== ADD ======================
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add new todo",
	Long:    "Menambahkan todo baru ke list. Task name harus unik.",
	Example: `add --task "Belajar Go" --priority high`,
	Run: func(cmd *cobra.Command, args []string) {
		s := service.NewTodoService()
		req := dto.CreateTodoRequest{
			Task_name: taskName,
			Priority:  priority,
		}

		result, err := s.CreateTodo(req)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		fmt.Printf("Added todo: %s (priority: %s)\n", result.Task_name, result.Priority)
	},
}

// ====================== LIST ======================
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show todo list",
	Long:  "Menampilkan seluruh todo dalam bentuk tabel di terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		s := service.NewTodoService()
		items, err := s.ListTodo()
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		fmt.Println("+----+----------------------+---------+---------+")
		fmt.Println("| ID | TASK                 | PRIORITY| STATUS  |")
		fmt.Println("+----+----------------------+---------+---------+")

		for _, t := range *items {
			status := t.Status
			switch t.Status {
			case "progress":
				status = "\033[33m" + t.Status + "\033[0m"
			case "done":
				status = "\033[32m" + t.Status + "\033[0m"
			}

			fmt.Printf("| %-2d | %-20s | %-7s | %-7s |\n",
				t.Id, t.Task_name, t.Priority, status)
			fmt.Println("+----+----------------------+---------+---------+")
		}
	},
}

// ====================== UPDATE ======================
var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update todo status (pending/progress/done)",
	Long:    "Mengubah status todo berdasarkan ID.",
	Example: `update --id 1 --status done`,
	Run: func(cmd *cobra.Command, args []string) {
		s := service.NewTodoService()
		req := dto.UpdateTodoRequest{
			Id:        updateId,
			NewStatus: newStatus,
		}

		result, err := s.UpdateStatus(req)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		fmt.Printf("Updated ID %d → status: %s\n", result.Id, result.Status)
	},
}

// ====================== DELETE ======================
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete todo by ID",
	Long:    "Menghapus todo berdasarkan ID",
	Example: `delete --id 2`,
	Run: func(cmd *cobra.Command, args []string) {
		s := service.NewTodoService()
		err := s.DeleteTodo(deleteId)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}
		fmt.Printf("Deleted todo ID %d\n", deleteId)
	},
}

// ====================== SEARCH ======================
var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search todo by task name",
	Long:    "Mencari todo berdasarkan nama task (support partial match)",
	Example: `search --task belajar`,
	Run: func(cmd *cobra.Command, args []string) {
		s := service.NewTodoService()

		req := dto.SearchTodoRequest{
			Task_Name: searchName,
		}

		results, err := s.SearchTodo(req)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}

		fmt.Println("+----+----------------------+---------+---------+")
		fmt.Println("| ID | TASK                 | PRIORITY| STATUS  |")
		fmt.Println("+----+----------------------+---------+---------+")

		for _, r := range results {
			fmt.Printf("| %-2d | %-20s | %-7s | %-7s |\n",
				r.Id, r.Task_name, r.Priority, r.Status)
			fmt.Println("+----+----------------------+---------+---------+")
		}
	},
}

// ====================== INIT ======================
func init() {
	addCmd.Flags().StringVarP(&taskName, "task", "t", "", "Task name")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "", "Task priority")

	updateCmd.Flags().IntVarP(&updateId, "id", "i", 0, "Todo ID")
	updateCmd.Flags().StringVarP(&newStatus, "status", "s", "", "New status")

	deleteCmd.Flags().IntVarP(&deleteId, "id", "i", 0, "Todo ID")

	searchCmd.Flags().StringVarP(&searchName, "task", "t", "", "Keyword to search")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(searchCmd)
}

// ====================== EXECUTE ======================
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
