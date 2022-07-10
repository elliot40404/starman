package env

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io/ioutil"
	"os"
)

// fomat input as command
func FORMAT_CMD(name string, path string) string {
	return fmt.Sprintf("@REM APP %s\nSTART \"\" \"%s\"", name, path)
}

func READ_FILE() string {
	// read config file and return contents
	file, err := ioutil.ReadFile(CONFIG_FILE())
	if err != nil {
		panic(err)
	}
	return string(file)
}

func WRITE_FILE(data string) {
	// write data to config file
	err := ioutil.WriteFile(CONFIG_FILE(), []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

func TABLE(data []APP) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Path"})
	for _, row := range data {
		t.AppendRow(table.Row{row.Name, row.Path})
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}
