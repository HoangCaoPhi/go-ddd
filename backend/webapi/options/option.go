package options

type Config struct {
	SqlServerOptions SqlServerOptions `mapstructure:"sqlserver"`
}

type SqlServerOptions struct {
	Host     string `mapstructure:"host"`
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
}
