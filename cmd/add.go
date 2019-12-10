/*
 * Copyright (c) 2019. Michael Zhao <michaelxzhao@icloud.com>
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xiaofeizhao/xnote/note"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [note name]",
	Short: "Add a note",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			filename := viper.GetString("filepath") + args[0]
			note.Add(filename)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
