package dbConnections

import (
	"fmt"
	"github.com/mehmetcekirdekci/GolangCRUD/config/dbConfigs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgreConnectionOpen() *gorm.DB {
	return postgreConnectionOpen()
}

func postgreConnectionOpen() *gorm.DB {
	var connection *gorm.DB
	var err error
	//dsn := setDSN()
	//retryCount := getRetryCount()
	dsn := "host=localhost user=postgres password=5492 dbname=GolangCRUDPostgreSql port=5432 sslmode=disable"
	retryCount := 3

	for i := 0; i < retryCount; i++ {
		connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		isConnectionAlive := err == nil && connection != nil
		if isConnectionAlive {
			break
		}
	}
	if err != nil {
		return nil
	}
	return connection
}

func setDSN() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfigs.PostgreSqlConfiguration{}.Host,
		dbConfigs.PostgreSqlConfiguration{}.User,
		dbConfigs.PostgreSqlConfiguration{}.Password,
		dbConfigs.PostgreSqlConfiguration{}.DbName,
		dbConfigs.PostgreSqlConfiguration{}.Port,
		dbConfigs.PostgreSqlConfiguration{}.SslMode,
	)
	return dsn
}

func getRetryCount() int {
	return dbConfigs.PostgreSqlConfiguration{}.RetryCount
}
