package interfaces_test

import (
	"github.com/bruteforce1414/study-go/interfaces/interfaces"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)


func TestManager_AddFileError(t *testing.T) {
	a := assert.New(t)
	pm := interfaces.GetPluginManager()
	err := pm.AddFile("/______")
	a.Error(err)
}

func TestManager_AddDir(t *testing.T) {
	a := assert.New(t)
	pm := interfaces.GetPluginManager()
	path := filepath.Join("../testdata/plugins")
	err := pm.AddDir(path)
	a.NoError(err)
	a.Len(pm.List(), 4)
}

func TestManager_AddFileNoError(t *testing.T) {
	a := assert.New(t)
	pm2 := interfaces.GetPluginManager()
	path := filepath.Join("../testdata/pluginSingle", "singleplugin.txt")
	err := pm2.AddFile(path)
	a.NoError(err)
	a.Len(pm2.List(), 5)

}



func TestManager_AddDirError(t *testing.T) {
	a := assert.New(t)
	pm := interfaces.GetPluginManager()
	path := filepath.Join("/_____")
	err := pm.AddDir(path)
	a.Error(err)
	a.Len(pm.List(), 5)
}
