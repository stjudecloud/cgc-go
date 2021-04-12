package apps

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mgutz/ansi"
	"github.com/ryanuber/columnize"
	log "github.com/sirupsen/logrus"
	yaml2 "gopkg.in/yaml.v2"

	"github.com/spf13/viper"
)

var url string
var AppParams appParams

type appParams struct {
	upload string
}

type App struct {
	ShortID    string
	Href       string                 `json:"href"`
	ID         string                 `json:"id"`
	Project    string                 `json:"project"`
	Name       string                 `json:"name"`
	Revision   int                    `json:"revision"`
	Raw        map[string]interface{} `json:"raw"`
	Upload     []byte
	UploadType string
}

func NewApp(inID string) App {
	a := App{
		ShortID: inID,
	}
	return a
}

func (a App) Print() {
	ul := ansi.ColorCode("white+u")
	reset := ansi.ColorCode("reset")

	var output []string
	output = append(output, fmt.Sprintf("%s%s%s", ul, a.Name, reset))
	output = append(output, fmt.Sprintf("Project|%s", a.Project))
	output = append(output, fmt.Sprintf("ID|%s", a.ID))
	output = append(output, fmt.Sprintf("Revision|%d", a.Revision))

	result := columnize.SimpleFormat(output)
	log.Debugf("GHA -> %b", gha)
	if gha {
		fmt.Printf("::set-output name=status::%s\n", a.ID)
		return
	}
	fmt.Println(result)
}

func (a App) Process() error {
	a.getAppDetails()
	if a.Upload != nil {
		a.uploadNewApp()
		a.getAppDetails()
	}
	a.Print()
	return nil
}

func (a *App) getAppDetails() error {
	body, err := a.queryAPI("DETAILS", "GET", "")
	if err != nil {
		return err
	}
	json.Unmarshal(body, &a)
	return nil
}

func (a *App) getRaw() error {
	body, err := a.queryAPI("RAW", "GET", "raw")
	if err != nil {
		return err
	}
	json.Unmarshal(body, &a)
	return nil
}

func (a App) uploadNewApp() error {
	body, err := a.queryAPI("UPLOAD", "POST", fmt.Sprintf("%d/raw", a.Revision+1))
	if err != nil {
		return err
	}
	var b map[string]interface{}
	json.Unmarshal(body, &b)
	if b["message"] != nil {
		log.Error(b["message"])
		return errors.New(b["message"].(string))
	}
	return nil
}

//helper functions
func (a App) validAppID() error {
	p := listParams{id: a.ShortID}
	ret, err := getApps(p)

	if err != nil {
		log.Error(err)
		return err
	}
	if len(ret.Items) == 0 {
		log.Error("no apps returned")
		return errors.New("no apps returned with that appid")
	}
	return nil
}

func (a *App) validateUploadFile(inFile string) error {
	data, err := ioutil.ReadFile("./" + inFile)
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	err = json.Unmarshal(data, &obj)
	if err == nil {
		log.Debug("Parsed valid JSON")
		a.UploadType = "json"
		a.Upload = data
		return nil
	}

	err = yaml2.Unmarshal(data, &obj)
	if err == nil {
		log.Debug("Parsed valid YAML")
		a.Upload = data
		a.UploadType = "yaml"
		return nil
	}

	return err
}

func (a App) queryAPI(inFunc string, inMethod string, inURL string) ([]byte, error) {
	var req *http.Request
	var err error

	appURL := url + fmt.Sprintf("/%s/%s", a.ShortID, inURL)

	log.Debugf("[%s] URL -> %s", inFunc, appURL)
	client := http.Client{}

	switch inMethod {
	case "GET":
		{
			req, err = http.NewRequest("GET", appURL, nil)
			if err != nil {
				log.Error(err)
				return nil, err
			}
		}
	case "POST":
		{
			data := bytes.NewBuffer(a.Upload)
			req, err = http.NewRequest("POST", appURL, data)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			switch a.UploadType {
			case "json":
				req.Header.Set("Content-Type", "application/json")
			case "yaml":
				req.Header.Set("Content-Type", "application/yaml")
			}
		}
	}

	devKey, ok := viper.Get("token").(string)
	if !ok {
		return nil, errors.New("Failed to find developers key.")
	}
	req.Header.Set("X-SBG-Auth-Token", devKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return body, nil
}
