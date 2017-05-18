package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type jiraConfig struct {
	BasicAuth string
	BaseURL   string
	UserName  string
}

func loadSettings(cfgFile string) jiraConfig {

	fi, err := os.Stat(cfgFile)
	die(err)
	if fi.IsDir() {
		log.Fatalf("%q is a directory: I was expecting a config file", cfgFile)
	}

	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("properties")
	err = viper.ReadInConfig()
	die(err)

	ret := jiraConfig{
		BasicAuth: viper.GetString("JIRA_BASIC_AUTH"),
		BaseURL:   viper.GetString("JIRA_BASE_URL"),
		UserName:  viper.GetString("JIRA_USERNAME"),
	}
	return ret
}

func main() {
	cfg := loadSettings(filepath.Join(os.Getenv("HOME"), ".jira"))
	client := &http.Client{}
	searchUrl := fmt.Sprintf("%s%s%s%s", cfg.BaseURL, "/rest/api/latest/search?jql=assignee=", cfg.UserName, " AND createdDate>-10d")

	req, err := http.NewRequest("GET", searchUrl, nil)
	die(err)


	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", cfg.BasicAuth))

	res, err := client.Do(req)
	die(err)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	fmt.Printf("%s\n", buf.String())
}