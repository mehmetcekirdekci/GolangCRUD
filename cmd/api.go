/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mehmetcekirdekci/GolangCRUD/cmd/controllerInstance"
	"github.com/mehmetcekirdekci/GolangCRUD/cmd/dbConnections"
	"github.com/mehmetcekirdekci/GolangCRUD/cmd/repositoryInstance"
	"github.com/mehmetcekirdekci/GolangCRUD/cmd/serviceInstance"
	echoextention "github.com/mehmetcekirdekci/GolangCRUD/pkg/echoExtention"
	echoSwagger "github.com/swaggo/echo-swagger"
	"time"

	"github.com/spf13/cobra"
)

type api struct {
	instance *echo.Echo
	command  *cobra.Command
	Port     string
}

// apiCmd represents the api command
var apiCmd = &api{
	command: &cobra.Command{
		Use:   "api",
		Short: "Service is up.",
		Long:  "Golang api project is up.",
	},
	Port: "5000",
}

func init() {
	rootCmd.AddCommand(apiCmd.command)
	apiCmd.apiCmdHandler()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (apiCmd *api) apiCmdHandler() {
	apiCmd.instance = echo.New()
	apiCmd.instance.Use(middleware.Logger())
	apiCmd.instance.Use(middleware.Recover())

	apiCmd.command.RunE = func(cmd *cobra.Command, args []string) error {

		postgreConnection := dbConnections.PostgreConnectionOpen()

		productRepository := repositoryInstance.GetProductRepositoryInstance(postgreConnection)

		mongoConnection := dbConnections.MongoConnectionOpen()

		userRepository := repositoryInstance.GetUserRepositoryInstance(mongoConnection)

		userBaseService := serviceInstance.GetUserBaseServiceInstance(userRepository)

		productBaseService := serviceInstance.GetProductBaseServiceInstance(productRepository)

		controllerInstance.ProductControllerInstance(productBaseService, apiCmd.instance)

		controllerInstance.UserControllerInstance(userBaseService, apiCmd.instance)

		apiCmd.instance.GET("/swagger/*", echoSwagger.WrapHandler)
		go func() {
			apiCmd.instance.Start(fmt.Sprintf(":%s", apiCmd.Port))
		}()
		//apiCmd.instance.Logger.Fatal(apiCmd.instance.Start(":5000"))
		echoextention.Shutdown(apiCmd.instance, 2*time.Second)

		return nil
	}
}
