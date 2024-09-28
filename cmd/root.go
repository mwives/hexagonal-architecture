package cmd

import (
	"database/sql"
	"os"

	dbInfra "github.com/mwives/hexagonal-architecture/adapters/db"
	"github.com/mwives/hexagonal-architecture/app"
	"github.com/spf13/cobra"
)

var db, _ = sql.Open("sqlite3", "sqlite.db")
var productDB = dbInfra.NewProductDB(db)
var productService = app.ProductService{Persistence: productDB}

var rootCmd = &cobra.Command{
	Use:   "hexagonal-architecture",
	Short: "Hexagonal Architecture Example",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
