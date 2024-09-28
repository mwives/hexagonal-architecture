package cmd

import (
	"fmt"

	"github.com/mwives/hexagonal-architecture/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productID string
var productName string
var productPrice float64

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Perform actions on products",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, action, productID, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable or disable a product")
	cliCmd.Flags().StringVarP(&productID, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "name", "n", "", "Product Name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product Price")
}
