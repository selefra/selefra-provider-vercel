package provider

import (
	"context"
	"os"

	"github.com/selefra/selefra-provider-vercel/constants"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"

	"github.com/selefra/selefra-provider-vercel/vercel_client"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Vercel,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var vercelConfig vercel_client.Config
				err := config.Unmarshal(&vercelConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}

				if vercelConfig.APIToken == "" {
					vercelConfig.APIToken = os.Getenv("VERCEL_API_TOKEN")
				}

				if vercelConfig.APIToken == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing APIToken in configuration")
				}

				if vercelConfig.TeamId == "" {
					vercelConfig.TeamId = os.Getenv("VERCEL_TEAM_ID")
				}

				if vercelConfig.TeamId == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing Team id in configuration")
				}

				clients, err := vercel_client.NewClients(vercelConfig)

				if err != nil {
					clientMeta.ErrorF(constants.Newclientserrs, err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Accountinformationnotfound)
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# api_token: <Your Vercel Api Token>
# team_id: <Your Vercel Team id>`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var vercelConfig vercel_client.Config
				err := config.Unmarshal(&vercelConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_0,
				constants.NA,
				constants.Notsupported,
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
