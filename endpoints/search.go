package endpoints

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/jimmyjames85/jira"
)

func Search(cfg *jira.Config, jql string) (string, error) {

	client := &http.Client{}
	searchUrl := fmt.Sprintf("%s%s%s", cfg.BaseURL, "/rest/api/latest/search?jql=", jql)

	req, err := http.NewRequest("GET", searchUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", cfg.BasicAuth))
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
