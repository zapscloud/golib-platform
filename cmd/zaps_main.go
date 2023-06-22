package main

import (
	"os"

	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-platform/platform_services"
	"github.com/zapscloud/golib-utils/utils"
)

func GetZapsDBCreds() utils.Map {
	dbtype := db_common.DATABASE_TYPE_ZAPSDB
	dbuser := os.Getenv("ZAPS_APP_KEY")
	dbsecret := os.Getenv("ZAPS_APP_SECRET")
	dbapp := os.Getenv("ZAPS_APP")

	dbCreds := utils.Map{
		db_common.DB_TYPE:   dbtype,
		db_common.DB_APP:    dbapp,
		db_common.DB_KEY:    dbuser,
		db_common.DB_SECRET: dbsecret}

	return dbCreds
}

func ZapsDBMain() (platform_services.AppUserService, platform_services.BusinessService, platform_services.RegionService,
	platform_services.ClientsService, platform_services.SysSettingService) {

	usersrv, _ := platform_services.NewAppUserService(GetZapsDBCreds())
	bizsrv, _ := platform_services.NewBusinessService(GetZapsDBCreds())
	regionsrv, _ := platform_services.NewRegionService(GetZapsDBCreds())
	clientsrv, _ := platform_services.NewClientsService(GetZapsDBCreds())
	settingsrv, _ := platform_services.NewSysSettingService(GetZapsDBCreds())

	return usersrv, bizsrv, regionsrv, clientsrv, settingsrv
}
