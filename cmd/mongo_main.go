package main

import (
	"os"

	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-platform/platform_common"
	"github.com/zapscloud/golib-platform/platform_services"
	"github.com/zapscloud/golib-utils/utils"
)

func GetMongoDBCreds() utils.Map {
	dbtype := db_common.DATABASE_TYPE_MONGODB
	dbuser := os.Getenv("READ_MONGODB_USER")
	dbsecret := os.Getenv("READ_MONGODB_SECRET")
	dbserver := os.Getenv("READ_MONGODB_SERVER")
	dbname := os.Getenv("READ_MONGODB_NAME")

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

	readDBCreds := GetMongoDBCreds()
	// Append BusinessId
	readDBCreds[platform_common.FLD_BUSINESS_ID] = "biz_ckht3d111ltrgeq49os0"

	usersrv, _ := platform_services.NewAppUserService(readDBCreds)
	bizsrv, _ := platform_services.NewBusinessService(GetMongoDBCreds())
	regionsrv, _ := platform_services.NewRegionService(GetMongoDBCreds())
	clientsrv, _ := platform_services.NewClientsService(GetMongoDBCreds())
	settingsrv, _ := platform_services.NewSysSettingService(GetMongoDBCreds())

	return usersrv, bizsrv, regionsrv, clientsrv, settingsrv
}
