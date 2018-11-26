package main

import (
	"flag"
	"fmt"
)

// Программа может принять два флага @see https://golang.org/pkg/flag/
// флаг --plugin указывает путь к файлу, который нужно загрузить
// флаг --dir указывает путь к директории из которой нужно загрузить все файлы как плагины
func main() {
	m := GetPluginManager()
    dir:=flag.String("path","D:\\development\\GoProjects\\study-go\\interfaces\\testdata\\plugins","a string")
	plugin:=flag.String("pathFile","D:\\development\\GoProjects\\study-go\\interfaces\\testdata\\pluginSingle\\singleplugin.txt","a string")


	// тут нужно получить данные из флагов, если флаги указаны

	if dir!=nil{
		m.AddDir(*dir)
	}else {
		fmt.Println("Флаг --dir не установлен")
	}


	if plugin!=nil{
		m.AddFile(*plugin)
	}else {
		fmt.Println("Флаг --plugin не установлен")
	}

     m.List()

	// выводим список всех плагинов, если список пуст, то выводим "плагинов нет"
}
