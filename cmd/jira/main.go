package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jimmyjames85/jira"
	"github.com/jimmyjames85/jira/endpoints"
	"github.com/spf13/viper"
)

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadSettings(cfgFile string) *jira.Config {

	fi, err := os.Stat(cfgFile)
	die(err)
	if fi.IsDir() {
		log.Fatalf("%q is a directory: I was expecting a config file", cfgFile)
	}

	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("properties")
	err = viper.ReadInConfig()
	die(err)

	ret := jira.Config{
		BasicAuth: viper.GetString("JIRA_BASIC_AUTH"),
		BaseURL:   viper.GetString("JIRA_BASE_URL"),
		UserName:  viper.GetString("JIRA_USERNAME"),
	}
	return &ret
}

func main() {
	cfg := loadSettings(filepath.Join(os.Getenv("HOME"), ".jira"))
	jql := fmt.Sprintf("assignee=%s AND createdDate>-10d", cfg.UserName)

	s, err := endpoints.Search(cfg, jql)
	die(err)
	fmt.Printf("%s\n", s)
}
