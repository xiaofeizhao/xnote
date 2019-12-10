/*
 * Copyright (c) 2019. Michael Zhao <michaelxzhao@icloud.com>
 */

package note

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Add(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	cmd := exec.Command("vim", filename)
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
	cmd := exec.Command("vim", filename)
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

func List(root string) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
