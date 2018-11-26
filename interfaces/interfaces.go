package main

import "sync"

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

}

func (m *manager) AddFile(path string) {

}

func (m *manager) List() []string {

}