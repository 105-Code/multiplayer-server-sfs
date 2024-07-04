package config

const (
	TickRateDefault     int  = 60
	PortDefault         int  = 4000
	TotalPlayersDefault int  = 25
	DebugDefault        bool = true
)

var (
	AppConfig ApplicationConfig
)

type ApplicationConfig struct {
	Debug        bool `mapstructure:"debug"`
	TotalPlayers int  `mapstructure:"total_player_default"`
	Port         int  `mapstructure:"port"`
}

func SetDefaults() {
	AppConfig.Debug = DebugDefault
	AppConfig.Port = PortDefault
	AppConfig.TotalPlayers = TotalPlayersDefault
}

func init() {
	SetDefaults()
}
