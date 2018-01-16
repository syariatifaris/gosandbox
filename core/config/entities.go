package config

type ConfigurationData struct {
	Database Database
}

type Database struct {
	DBHost     string
	DBType     string
	DBName     string
	DBUser     string
	DBPassword string
	DBPort     int64
}
