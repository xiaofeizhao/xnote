/*
 * Copyright (c) 2019. Michael Zhao <michaelxzhao@icloud.com>
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiaofeizhao/xnote/note"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [note(s) name]",
	Short: "Delete note(s)",
	Run: func(cmd *cobra.Command, args []string) {
		for _, fileName := range args {
			note.GetNote().Delete(fileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
