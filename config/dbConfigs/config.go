package dbConfigs

// db configurations
type (
	PostgreSqlConfiguration struct {
		Host       string
		User       string
		Password   string
		DbName     string
		Port       string
		SslMode    string
		RetryCount int
	}
)
