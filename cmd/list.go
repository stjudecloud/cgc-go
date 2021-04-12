package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List availabe CGC endpoints",
	Long: `This list of endpoints is provided by quering the CGC API, it is possible not all functionality will be supported by cgc-go
	Ref: https://docs.cancergenomicscloud.org/docs/new-1`,
	Run: func(cmd *cobra.Command, args []string) {
		getCGCEndpoints()
	},
}

func getCGCEndpoints() error {
	resp, err := http.Get(viper.Get("rooturl").(string))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ret := make(map[string]string)

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return err
	}

	var output []string

	for k, v := range ret {
		k2 := strings.TrimSuffix(k, "_url")
		d := fmt.Sprintf("%s | %s\n", k2, v)
		output = append(output, d)
		sort.Strings(output)
	}

	result := columnize.SimpleFormat(output)
	fmt.Println(result)

	return nil
}
