/*
 * Copyright (c) 2019. Michael Zhao <michaelxzhao@icloud.com>
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiaofeizhao/xnote/note"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [note(s) name]",
	Short: "Create note(s)",
	Run: func(cmd *cobra.Command, args []string) {
		for _, fileName := range args {
			note.GetNote().Add(fileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
