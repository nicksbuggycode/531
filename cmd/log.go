/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	db "github.com/nicksbuggycode/531/db/sqlc"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log a lift",
	Long:  `Accepts lift, weight, reps and will also provide a calculated max`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("log called")
		wgt, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		rep, err := strconv.ParseInt(args[2], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		wgt32 := int32(wgt)
		rep32 := int32(rep)
		p := db.LogLiftParams{
			Lift:   sql.NullString{String: args[0], Valid: true},
			Weight: sql.NullInt32{Int32: wgt32, Valid: true},
			Reps:   sql.NullInt32{Int32: rep32, Valid: true},
		}
		conn, err := sql.Open(dbDriver, dbSource)
		if err != nil {
			log.Fatal("Can't connect to DB:", err)
		}
		queries := db.New(conn)
		queries.LogLift(context.Background(), p)
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
