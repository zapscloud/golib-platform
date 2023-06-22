package main

import (
	"os"

	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-platform/platform_services"
	"github.com/zapscloud/golib-utils/utils"
)

func GetMongoDBCreds() utils.Map {
	dbtype := db_common.DATABASE_TYPE_MONGODB
	dbuser := os.Getenv("MONGO_DB_USER")
	dbsecret := os.Getenv("MONGO_DB_SECRET")
	dbserver := os.Getenv("MONGO_DB_SERVER")
	dbname := os.Getenv("MONGO_DB_NAME")

	dbCreds := utils.Map{
		db_common.DB_TYPE:   dbtype,
		db_common.DB_SERVER: dbserver,
		db_common.DB_NAME:   dbname,
		db_common.DB_USER:   dbuser,
		db_common.DB_SECRET: dbsecret}

	return dbCreds
}

func MongoDBMain() (
	platform_services.AppUserService,
	platform_services.BusinessService,
	platform_services.RegionService,
	platform_services.ClientsService,
	platform_services.SysSettingService) {

	usersrv, _ := platform_services.NewAppUserService(GetMongoDBCreds())
	bizsrv, _ := platform_services.NewBusinessService(GetMongoDBCreds())
	regionsrv, _ := platform_services.NewRegionService(GetMongoDBCreds())
	clientsrv, _ := platform_services.NewClientsService(GetMongoDBCreds())
	settingsrv, _ := platform_services.NewSysSettingService(GetMongoDBCreds())

	return usersrv, bizsrv, regionsrv, clientsrv, settingsrv
}
