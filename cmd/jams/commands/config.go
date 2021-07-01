package commands

import (
	"io"
	"log"
	"os"
	"path/filepath"

	gap "github.com/muesli/go-app-paths"
)

func init() {
	scope := gap.NewScope(gap.User, "jams-go")
	path, err := scope.ConfigPath("jams.conf")
	if err != nil {
		log.Fatal("Error getting Config Path: ", err)
	}
	loadConfig(path)
}

func loadConfig(path string) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(dir)
	}
	if _, err := os.Stat(dir); err != nil {
		log.Fatal("Error Checking Config Directory: ", err)
	}
	if _, err := os.Stat(path); err != nil {
		f, _ := os.Create(path)
		io.WriteString(f, defaultConfig)
		f.Close()
	}

}

type AppConfig struct {
	Host     string
	Scheme   string
	BasePath string
	Username string
}

const defaultConfig = `Host = "1.2.3.4"
Scheme = "http"
BasePath = "/JAMS/"
Username = "user"
`
