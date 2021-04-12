package rate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ryanuber/columnize"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mgutz/ansi"
)

var url string

var Cmd = &cobra.Command{
	Use:   "rate_limit",
	Short: "CGC rate-limit information endpoint",
	Long: `Support for rate limit information API endpoint
Ref: https://docs.cancergenomicscloud.org/docs/rate-limit`,
	RunE: func(cmd *cobra.Command, args []string) error {
		url = viper.Get("rooturl").(string) + "rate_limit"
		return GetRateStatus()
	},
}

type Response struct {
	Rate struct {
		Limit     int       `json:"limit"`
		Remaining int       `json:"remaining"`
		ResetRaw  int64     `json:"reset"`
		Reset     time.Time `json:"resetTime"`
	} `json:"rate"`
	InstanceLimit struct {
		Limit     int `json:"limit"`
		Remaining int `json:"remaining"`
	} `json:"instance_limit"`
}

func GetRateStatus() error {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	devKey, ok := viper.Get("token").(string)
	if !ok {
		return errors.New("Failed to find devlopers key.")
	}

	req.Header.Set("X-SBG-Auth-Token", devKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return err
	}

	ret := Response{}
	json.Unmarshal(body, &ret)
	ret.Rate.Reset = time.Unix(ret.Rate.ResetRaw, 0)

	ret.Print()

	return nil
}

func (r Response) Print() {
	ul := ansi.ColorCode("white+u")
	reset := ansi.ColorCode("reset")

	var output []string
	output = append(output, ul+"Rate Limits"+reset)
	output = append(output, fmt.Sprintf("Limit|%d", r.Rate.Limit))
	output = append(output, fmt.Sprintf("Remaining|%d", r.Rate.Remaining))
	output = append(output, fmt.Sprintf("Reset|%s", r.Rate.Reset.String()))

	output = append(output, ul+"Instance Limits"+reset)
	output = append(output, fmt.Sprintf("Limit|%d", r.InstanceLimit.Limit))
	output = append(output, fmt.Sprintf("Remaining|%d", r.InstanceLimit.Remaining))

	result := columnize.SimpleFormat(output)
	fmt.Println(result)
}
