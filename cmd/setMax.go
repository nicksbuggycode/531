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

	_ "github.com/lib/pq"
	db "github.com/nicksbuggycode/531/db/sqlc"
	"github.com/spf13/cobra"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/531?sslmode=disable"
)

// setMaxCmd represents the setMax command
var setMaxCmd = &cobra.Command{
	Use:   "setMax",
	Short: "accepts lift, weight -- will add this and calculated training max to the db",
	Long:  `Adds a new max to the maxes table. Must include 2 arguments: lift, weight`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("setMax called")
		wgt, err := strconv.ParseInt(args[1], 10, 32)
		wgt32 := int32(wgt)
		tmcalc := 0.9 * float64(wgt32)
		tmint := int32(tmcalc)
		if err != nil {
			log.Fatal("invalid value: ", err)
		}
		s := db.SetMaxParams{
			Lift:        sql.NullString{String: args[0], Valid: true},
			Onerepmax:   sql.NullInt32{Int32: wgt32, Valid: true},
			Trainingmax: sql.NullInt32{Int32: tmint, Valid: true},
		}
		conn, err := sql.Open(dbDriver, dbSource)
		if err != nil {
			log.Fatal("Can't connect to DB:", err)
		}
		defer conn.Close()
		queries := db.New(conn)
		queries.SetMax(context.Background(), s)

	},
}

func init() {

	rootCmd.AddCommand(setMaxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setMaxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setMaxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
