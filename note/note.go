/*
 * Copyright (c) 2019. Michael Zhao <michaelxzhao@icloud.com>
 */

package note

import (
	"log"
	"os"
	"os/exec"
)

const (
	filePath = "/Users/michael.zhao/GoogleDrive/Sync/Note/"
)

func Add(filename string) {
	path := filePath + filename
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	cmd := exec.Command("vim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func Edit(filename string) {
	path := filePath + filename

	cmd := exec.Command("vim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
