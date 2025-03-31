package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	JWT    JwtSetting    `mapstructure:"jwt"`
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbName"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type JwtSetting struct {
	TOKEN_HOUR_LIFESPAN uint   `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	JWT_EXPIRATION      string `mapstructure:"JWT_EXPIRATION"`
	API_SECRET_KEY      string `mapstructure:"API_SECRET_KEY"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Max_size      int    `mapstructure:"max_size"`
	Compress      bool   `mapstructure:"compress"`
}
