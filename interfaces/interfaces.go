package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

var instance *manager
var once sync.Once

type PluginManager interface {
	// method adds all files from given directory to list of plugins
	AddDir(path string)
	// method adds path to list of plugins, must check file exists or not
	AddFile(path string)
	// returns list of all plugins
	List() []string
}

type manager struct {
	// list of plugins, it stores absolute path to files, ex: /user/john/plugins/plugin.txt
	plugins []string
}

func GetPluginManager() PluginManager {
	once.Do(func() {
		instance = &manager{
			plugins: []string{},
		}
	})
	return instance
}

func (m *manager) AddDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		}
	for _, file := range files {
		m.plugins=append(m.plugins,path+file.Name())

	}
}


func (m *manager) AddFile(path string) {
	_, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		}else{
		m.plugins=append(m.plugins,path)
	}
}

func (m *manager) List() []string {
	if len(m.plugins)==0 {
		fmt.Println("\nПлагинов нет!!!")
	}else {
		fmt.Println("\nСписок загруженных плагинов:\n")
		for _, value := range m.plugins {
			fmt.Println(value)}
		}
	return m.plugins
}
