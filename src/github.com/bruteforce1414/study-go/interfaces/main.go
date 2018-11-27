package main

import (
	"flag"
	"fmt"
	"github.com/bruteforce1414/study-go/interfaces/interfaces"
	"path/filepath"
)

// Программа может принять два флага @see https://golang.org/pkg/flag/
// флаг --plugin указывает путь к файлу, который нужно загрузить
// флаг --dir указывает путь к директории из которой нужно загрузить все файлы как плагины
func main() {
	m := interfaces.GetPluginManager()

	dirParameter2,_:=filepath.Abs("testdata/plugins")
	dir:=flag.String("path",dirParameter2,"a string")

	pluginParameter2,_:=filepath.Abs("testdata/pluginSingle/singleplugin.txt")
	plugin:=flag.String("pathFile",pluginParameter2,"a string")
	// тут нужно получить данные из флагов, если флаги указаны

	if dir!=nil{
		errDir:=m.AddDir(*dir)
		if errDir!=nil{
			fmt.Println("Возникла ошибка:",errDir)
		}
	}

	if plugin!=nil{
		errFile:=m.AddFile(*plugin)
		if errFile!=nil{
			fmt.Println("Возникла ошибка:", errFile)
		}

	}

	if len(m.List())==0 {
		fmt.Println("\nПлагинов нет!!!")
	}else {
		fmt.Println("\nСписок загруженных плагинов:\n")
		for _, value := range m.List() {
			fmt.Println(value)}
	}
	// выводим список всех плагинов, если список пуст, то выводим "плагинов нет"
}