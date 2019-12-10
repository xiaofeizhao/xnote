/*
 * Copyright (c) 2019. Michael Zhao <michaelxzhao@icloud.com>
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiaofeizhao/xnote/note"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [note name]",
	Short: "Edit a note",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			note.Edit(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
