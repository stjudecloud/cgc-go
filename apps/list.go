package apps

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/mgutz/ansi"
	"github.com/ryanuber/columnize"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ListParams listParams

var appsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List apps",
	Long: `Lists the apps the the provided developers key has access to
Ref: https://docs.cancergenomicscloud.org/docs/apps`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url = viper.Get("rooturl").(string) + "apps"
		return ListApps()
	},
}

type listParams struct {
	fields       string
	project      string
	projectOwner string
	visability   string
	query        string
	id           string
}

type ListResponse struct {
	Href  string `json:"href"`
	Items []struct {
		Href string `json:"href"`
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
}

func (r ListResponse) Print() {
	ul := ansi.ColorCode("white+u")
	reset := ansi.ColorCode("reset")
	var output []string

	output = append(output, fmt.Sprint(ul+"Name"+reset+" | "+ul+"ID"+reset+" | "+ul+"Latest Rev"+reset))

	for _, i := range r.Items {
		params := getParams(`(?P<project>.*)/(?P<rev>\d+)$`, i.Id)
		output = append(output, fmt.Sprintf("%s | %s | %s", i.Name, params["project"], params["rev"]))
	}

	result := columnize.SimpleFormat(output)
	fmt.Println(result)
}

func ListApps() error {
	ret, err := getApps(ListParams)
	if err != nil {
		log.Error(err)
		return err
	}
	ret.Print()

	return nil
}

func getParams(regEx, in string) (paramsMap map[string]string) {
	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(in)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

func getApps(p listParams) (ListResponse, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return ListResponse{}, err
	}
	devKey, ok := viper.Get("token").(string)
	if !ok {
		return ListResponse{}, errors.New("Failed to find developers key.")
	}

	req.Header.Set("X-SBG-Auth-Token", devKey)
	q := req.URL.Query()

	if p.fields != "" {
		q.Add("fields", p.fields)
	}
	if p.project != "" {
		q.Add("project", p.project)
	}
	if p.projectOwner != "" {
		q.Add("project_owner", p.projectOwner)
	}
	if p.visability != "" {
		q.Add("visibility", p.visability)
	}
	if p.query != "" {
		q.Add("q", p.query)
	}
	if p.id != "" {
		q.Add("id", p.id)
	}
	req.URL.RawQuery = q.Encode()
	log.Debug("Raw Query URL: " + req.URL.RawQuery)

	resp, err := client.Do(req)
	if err != nil {
		return ListResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return ListResponse{}, err
	}
	ret := ListResponse{}
	json.Unmarshal(body, &ret)

	return ret, nil
}
