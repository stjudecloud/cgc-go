package apps

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var app App
var gha bool

func init() {
	Cmd.AddCommand(appsListCmd)
	//Cmd.AddCommand(appDetails)

	//details.go
	Cmd.Flags().StringVarP(&AppParams.upload, "upload", "u", "", "Upload a new version of your app")
	Cmd.Flags().BoolVarP(&gha, "gha", "", false, "Github Actions friendly output")

	//list.go
	appsListCmd.Flags().StringVarP(&ListParams.fields, "fields", "f", "", "Selector specifying a subset of fields to include in the response.")
	appsListCmd.Flags().StringVarP(&ListParams.project, "project", "p", "", "Enter a project, in the form {project_owner}/{project_short_name} to restrict the results to apps from that project only.")
	appsListCmd.Flags().StringVarP(&ListParams.projectOwner, "project_owner", "o", "", "Enter a CGC username to restrict the results to apps from that user's projects only. Note that you can only see apps within projects that you are a member of.")
	appsListCmd.Flags().StringVarP(&ListParams.visability, "visibility", "s", "", "Set this to public to see all public apps on the CGC.")
	appsListCmd.Flags().StringVarP(&ListParams.query, "query", "q", "", "Enter one or more search terms to query apps using the q parameter. Learn more about querying in documentation.")
	appsListCmd.Flags().StringVarP(&ListParams.id, "id", "i", "", "Use this parameter to query apps based on their ID.")
}

var Cmd = &cobra.Command{
	Use:   "apps <appid>",
	Short: "CGC apps endpoint",
	Long: `Support for apps API endpoint
Ref: https://docs.cancergenomicscloud.org/docs/apps`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing appid")
		}
		app = NewApp(args[0])
		url = viper.Get("rooturl").(string) + "apps"

		if AppParams.upload != "" {
			if err := app.validateUploadFile(AppParams.upload); err != nil {
				return errors.New("invalid upload file provided")
			}
		}

		if err := app.validAppID(); err != nil {
			return errors.New("invalid appid provided")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		url = viper.Get("rooturl").(string) + "apps"
		return app.Process()
	},
}
