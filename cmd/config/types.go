package config

var Config = &config{}

type config struct {
	Server   Server
	Database Database
	Token    Token
}

type Server struct {
	UserAddr    string `env:"SERVER_ADDR" envDefault:":8081"`
	ProjectAddr string `env:"SERVER_ADDR" envDefault:":8082"`
	MemberAddr  string `env:"SERVER_ADDR" envDefault:":8083"`
	ImageAddr   string `env:"SERVER_ADDR" envDefault:":8084"`
	TokenAddr   string `env:"SERVER_ADDR" envDefault:":8085"`
}

type Database struct {
	Host     string `env:"PSQL_HOST" envDefault:"localhost"`
	Port     int    `env:"PSQL_PORT" envDefault:"5432"`
	User     string `env:"PSQL_USER" envDefault:"postgres"`
	Password string `env:"PSQL_PASSWORD" envDefault:"postgres"`
	DBName   string `env:"PSQL_DBNAME" envDefault:"vol10"`
	SSLMode  string `env:"PSQL_SSLMODE" envDefault:"disable"`

	ConnectTimeout  int  `env:"PSQL_CONNECT_TIMEOUT" envDefault:"10"`
	ConnectWaitTime int  `env:"PSQL_CONNECT_WAIT_TIME" envDefault:"10"`
	ConnectAttempts int  `env:"PSQL_CONNECT_ATTEMPTS" envDefault:"3"`
	ConnectBlocks   bool `env:"PSQL_CONNECT_BLOCKS" envDefault:"false"`
	CloseTimeout    int  `env:"PSQL_CLOSE_TIMEOUT" envDefault:"10"`
}

type Token struct {
	SymmetricKey string `env:"TOKEN_SYMMETRIC_KEY" envDefault:"abcdefghijabcdefghijabcdefghijab"`
}
