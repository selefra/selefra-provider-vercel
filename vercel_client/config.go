package vercel_client

type Config struct {
	APIToken string `yaml:"api_token"  mapstructure:"api_token"`
	TeamId   string `yaml:"team_id"  mapstructure:"team_id"`
}
