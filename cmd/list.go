/*
 * Copyright (c) 2019. Michael Zhao <michaelxzhao@icloud.com>
 */

package cmd

import (
	"github.com/spf13/viper"
	"github.com/xiaofeizhao/xnote/note"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all notes",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("filepath")
		note.List(path)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
