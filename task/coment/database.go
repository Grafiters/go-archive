package coment

import (
	database "archive/db"

	"fmt"

	"github.com/spf13/cobra"
)

type CommandDefinition struct {
	Use   string
	Short string
	Run   func(cmd *cobra.Command, args []string)
}

var Commands = []CommandDefinition{
	{
		Use:   "db:create",
		Short: "Run database create",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running Create database...")
			database.Create()
			fmt.Println("Done creating database")
		},
	},
	{
		Use:   "db:drop",
		Short: "Run database drop",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running database migrations...")
			database.Drop()
			fmt.Println("Done Droping database")
		},
	},
	{
		Use:   "drop:migrate",
		Short: "Run database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running database migrations...")

		},
	},
	{
		Use:   "db:seed",
		Short: "Seed the database with initial data",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Seeding the database...")
		},
	},
}
