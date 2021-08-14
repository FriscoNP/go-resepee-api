package main

import (
	"os"

	"go-resepee-api/app/boot"
	appMiddleware "go-resepee-api/app/middleware"
	dbDriver "go-resepee-api/db/driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config/config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
}

func main() {
	configDB := dbDriver.ConfigDB{
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Database: viper.GetString("database.db"),
	}
	db := configDB.InitialDB()

	configJWT := appMiddleware.ConfigJWT{
		SecretJWT:       viper.GetString("jwt.secret"),
		ExpiresDuration: viper.GetInt("jwt.expired"),
	}

	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())

	bootApp := boot.BootApp{
		DB:        db,
		JwtAuth:   &configJWT,
		JwtConfig: configJWT.Init(),
		Echo:      e,
	}

	bootApp.RegisterRoutes()

	log.Fatal(e.Start(viper.GetString("server.address")))
}
