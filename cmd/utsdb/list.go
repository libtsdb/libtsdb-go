package main

import (
	"github.com/spf13/cobra"
	"github.com/libtsdb/libtsdb-go/libtsdb"
	"fmt"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list supported databases",
	Run: func(cmd *cobra.Command, args []string) {
		dbs := libtsdb.Databases()
		for _, db := range dbs {
			// TODO: print full meta
			fmt.Println(db)
		}
	},
}
