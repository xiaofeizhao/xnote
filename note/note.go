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
	"path"
	"sync"

	"github.com/spf13/viper"
)

var (
	note NoteModel
	once sync.Once
)

type NoteModel interface {
	Add(fileName string)
	Edit(fileName string)
	Delete(fileName string)
	List(folder string)
}

func GetNote() NoteModel {
	once.Do(func() {
		note = &noteModelImpl{
			home:   viper.GetString("note_home"),
			editor: viper.GetString("editor"),
		}
	})
	return note
}

type noteModelImpl struct {
	home   string
	editor string
}

func (n *noteModelImpl) Add(fileName string) {
	fullName := path.Join(n.home, fileName)
	if _, err := os.Stat(fullName); os.IsNotExist(err) {
		file, err := os.Create(fullName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func (n *noteModelImpl) Edit(fileName string) {
	fullName := path.Join(n.home, fileName)

	if _, err := os.Stat(fullName); os.IsNotExist(err) {
		log.Fatal(err)
	}

	cmd := exec.Command(n.editor, fullName)
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

func (n *noteModelImpl) Delete(fileName string) {
	fullName := path.Join(n.home, fileName)
	if err := os.Remove(fullName); err != nil {
		log.Fatal(err)
	}
}

func (n *noteModelImpl) List(folder string) {
	fullName := path.Join(n.home, folder)
	files, err := ioutil.ReadDir(fullName)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
