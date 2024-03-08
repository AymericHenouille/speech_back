package settings

import (
	"flag"
	"fmt"
)

type AppSettings struct {
	Port       int
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

func New(port int, dbUser, dbPassword, dbHost, dbName string) *AppSettings {
	return &AppSettings{
		Port:       port,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBName:     dbName,
	}
}

func CreateSettingsFromFlags() *AppSettings {
	port := flag.Int("port", 8080, "Port to run the server on")
	dbUser := flag.String("dbuser", "postgres", "Database user")
	dbPassword := flag.String("dbpassword", "secret", "Database password")
	dbHost := flag.String("dbhost", "localhost:5432", "Database host")
	dbName := flag.String("dbname", "postgres", "Database name")
	flag.Parse()
	return New(*port, *dbUser, *dbPassword, *dbHost, *dbName)
}

func (a *AppSettings) AppPort() string {
	return fmt.Sprintf(":%d", a.Port)
}
