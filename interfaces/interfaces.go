package main

import (
	"io/ioutil"
	"path/filepath"
	"sync"
)

var instance *manager
var once sync.Once

type PluginManager interface {
	// method adds all files from given directory to list of plugins
	AddDir(path string) error
	// method adds path to list of plugins, must check file exists or not
	AddFile(path string) error
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

func (m *manager) AddDir(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
		}
	for _, file := range files {
		m.plugins=append(m.plugins,filepath.Join(path,file.Name()))
	}
	return nil
}


func (m *manager) AddFile(path string) error {
	_, err := ioutil.ReadFile(path)
	if err != nil {
		return err
		}else{
		m.plugins=append(m.plugins,filepath.Join(path))
	}
	return nil
}

func (m *manager) List() []string {
	return m.plugins
}
