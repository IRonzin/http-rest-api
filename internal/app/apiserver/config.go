package apiserver

type Config struct {
	PprofBindAddr string `toml:"pprof_bind_addr"`
	BindAddr      string `toml:"bind_addr"`
	LogLevel      string `toml:"log_level"`
	DatabaseUrl   string `toml:"database_url"`
	SessionKey    string `toml:"session_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
