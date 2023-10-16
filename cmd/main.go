package main

import (
	"fmt"

	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-platform/platform_common"
	"github.com/zapscloud/golib-platform/platform_services"
	"github.com/zapscloud/golib-utils/utils"

	"github.com/kr/pretty"
)

func main() {
	fmt.Println("test")
	usersrv, bizsrv, regionsrv, clientsrv, settingsrv := MongoDBMain()
	// usersrv, bizsrv, regionsrv, clientsrv, settingsrv := ZapsDBMain()

	if usersrv == nil {
		return
	}

	EmptyAppUser(usersrv)
	// UpdateAppUser(srv)
	// changePassword(srv)
	// AuthAppUser(srv)
	CreateAppUser(usersrv)
	//ListAppUsers(usersrv)
	//GetAppUser(usersrv)
	// FindAppUser(srv)
	// ListAppUserBusinesses(usersrv)

	EmptySysBusiness(bizsrv)
	// CreateSysBusiness(bizsrv)
	// UpdateSysBusiness(bizsrv)
	// DeleteSysBusiness(bizsrv)
	// FindSysBusiness(bizsrv)
	// GetSysBusiness(bizsrv)
	//ListSysBusiness(bizsrv)
	// AddUser(bizsrv)
	// GetAccessDetails(bizsrv)
	// ListSysBusinessUsers(bizsrv)

	EmptyAppRegion(regionsrv)

	// CreateAppRegion(regionsrv)
	// UpdateAppZapsRegion(regionsrv)
	// UpdateAppMongoRegion(regionsrv)
	// ListAppRegions(regionsrv)
	// GetAppRegion(regionsrv)

	EmptySysClient(clientsrv)
	// DeleteSysClient(clientsrv)
	// CreateSysClient(clientsrv)
	// UpdateSysClient(clientsrv)
	// ListSysClients(clientsrv)
	// FindSysClient(clientsrv)
	// GetSysClient(clientsrv)

	EmptySysSetting(settingsrv)
	// DeleteSysSetting(settingsrv)
	// CreateSysSetting(settingsrv)
	// UpdateSysSetting(settingsrv)
	// ListSysSettings(settingsrv)
	// FindSysSetting(settingsrv)
	// GetSysSetting(settingsrv)

}

func EmptyAppUser(srv platform_services.AppUserService) {
	fmt.Println("App User Service ")
}

func EmptyAppRegion(srv platform_services.RegionService) {
	fmt.Println("App Region Service ")
}

func EmptySysBusiness(srv platform_services.BusinessService) {
	fmt.Println("App Business Service ")
}

func EmptySysClient(srv platform_services.ClientsService) {
	fmt.Println("App Client Service ")
}

func EmptySysSetting(srv platform_services.SysSettingService) {
	fmt.Println("App Setting Service ")
}

func CreateAppUser(srv platform_services.AppUserService) {

	indata := utils.Map{
		platform_common.FLD_APP_USER_ID:       "user003",
		platform_common.FLD_APP_USER_PASSWORD: "secret003",
		platform_common.FLD_APP_USER_EMAILID:  "user003@zaps.com",
		platform_common.FLD_APP_USER_FNAME:    "Demo user 003",
	}

	res, err := srv.Create(indata)
	fmt.Println("Create AppUser", err)
	pretty.Println(res)

}

func GetAppUser(srv platform_services.AppUserService) {
	res, err := srv.Get("user003")
	fmt.Println("Get AppUser", err)
	pretty.Println(res)

}

func FindAppUser(srv platform_services.AppUserService) {

	filter := fmt.Sprintf(`{"%s":"%s"}`, platform_common.FLD_SYS_USER_EMAILID, "user003@zaps.com")
	res, err := srv.Find(filter)
	fmt.Println("Get AppUser", err)
	pretty.Println(res)

}

func UpdateAppUser(srv platform_services.AppUserService) {

	indata := utils.Map{
		platform_common.FLD_SYS_USER_ID:        "user003",
		platform_common.FLD_SYS_USER_FIRSTNAME: "Demo user 003 Updated",
		db_common.FLD_IS_DELETED:               false,
	}

	res, err := srv.Update("user003", indata)
	fmt.Println("Update AppUser", err)
	pretty.Println(res)

}

func AuthAppUser(srv platform_services.AppUserService) {

	res, err := srv.Authenticate(platform_common.FLD_SYS_USER_EMAILID, "user003@zaps.com", "secret003")
	fmt.Println("Auth AppUser Error", err)
	pretty.Println(res)

}

func ChangePassword(srv platform_services.AppUserService) {
	res, err := srv.ChangePassword("user003", "secret00300")
	fmt.Println("Auth AppUser Error", err)
	pretty.Println(res)

}

func ListAppUsers(srv platform_services.AppUserService) {
	res, err := srv.List("", "", 0, 0)
	fmt.Println("List User success ", err)
	fmt.Println("List User summary ", res)
	pretty.Print(res)
}

func ListAppUserBusinesses(srv platform_services.AppUserService) {
	filter := fmt.Sprintf(`{"%s":"%s"}`, "business_role", "staff")
	res, err := srv.BusinessList("user003", filter, "", 0, 0)
	fmt.Println("List User Business success ", err)
	fmt.Println("List User Business summary ", res)
	pretty.Print(res)
}

func CreateSysBusiness(srv platform_services.BusinessService) {

	indata := utils.Map{
		platform_common.FLD_BUSINESS_ID:        "business004",
		platform_common.FLD_BUSINESS_NAME:      "Business 004",
		platform_common.FLD_BUSINESS_REGION_ID: platform_common.DEF_REGION_ID,
	}

	res, err := srv.Create(indata)
	fmt.Println("Create SysBusiness", err)
	pretty.Println(res)

}

// func ListSysBusiness(srv platform_services.BusinessService) {
// 	res, err := srv.List("", "", 0, 0)
// 	fmt.Println("List Business success ", err)
// 	fmt.Println("List Business summary ", res)
// 	pretty.Print(res)
// }

func UpdateSysBusiness(srv platform_services.BusinessService) {

	indata := utils.Map{
		platform_common.FLD_BUSINESS_NAME:      "Business 002 Region Update",
		platform_common.FLD_BUSINESS_REGION_ID: platform_common.DEF_REGION_ID,
		db_common.FLD_IS_DELETED:               false,
	}

	res, err := srv.Update("business001", indata)
	fmt.Println("Update AppUser", err)
	pretty.Println(res)
}

func DeleteSysBusiness(srv platform_services.BusinessService) {

	err := srv.Delete("cadk0luv9mc1jtpu0ckg", false)
	fmt.Println("DeleteSysBusiness", err)
}

func FindSysBusiness(srv platform_services.BusinessService) {

	filter := fmt.Sprintf(`{"%s":"%s"}`, platform_common.FLD_BUSINESS_NAME, "Business 002")
	res, err := srv.Find(filter)
	fmt.Println("Find FindSysBusiness", err)
	pretty.Println(res)

}

func GetSysBusiness(srv platform_services.BusinessService) {
	res, err := srv.Get("business002")
	fmt.Println("Get AppUser", err)
	pretty.Println(res)

}

func AddUser(srv platform_services.BusinessService) {
	res, err := srv.AddUser("business004", "user003")
	fmt.Println("Create AddUser", err)
	pretty.Println(res)
}

func GetAccessDetails(srv platform_services.BusinessService) {
	res, err := srv.GetUserDetails("business004", "user003")
	fmt.Println("GetAccessDetails success ", err)
	fmt.Println("GetAccessDetails summary ", res)
	pretty.Print(res)
}

func ListSysBusinessUsers(srv platform_services.BusinessService) {
	filter := fmt.Sprintf(`{"%s":"%s"}`, "business_role", "staff")
	res, err := srv.GetUsers("business002", filter, "", 0, 0)
	fmt.Println("List User Business success ", err)
	fmt.Println("List User Business summary ", res)
	pretty.Print(res)
}

func ListAppRegions(srv platform_services.RegionService) {
	res, err := srv.List("", "", 0, 0)
	fmt.Println("List Region success ", err)
	fmt.Println("List Region summary ", res)
	pretty.Print(res)
}

func CreateAppRegion(srv platform_services.RegionService) {

	indata := utils.Map{
		platform_common.FLD_REGION_ID:   platform_common.DEF_REGION_ID,
		platform_common.FLD_REGION_NAME: platform_common.DEF_REGION_NAME,
	}

	res, err := srv.Create(indata)
	fmt.Println("Create CreateAppRegion", err)
	pretty.Println(res)

}

func UpdateAppZapsRegion(srv platform_services.RegionService) {

	indata := utils.Map{
		platform_common.FLD_REGION_ID:            platform_common.DEF_REGION_ID,
		platform_common.FLD_REGION_DB_TYPE:       db_common.DATABASE_TYPE_ZAPSDB,
		platform_common.FLD_REGION_ZAPSDB_APP:    "medishopdevapp",
		platform_common.FLD_REGION_ZAPSDB_KEY:    "medishopdevapp",
		platform_common.FLD_REGION_ZAPSDB_SECRET: "9a5659539f72f22db824f4a6b6a3ba7cf997f6a373f242e16d514b0e0d2f3bf26dae16a933a159b52bcc062fb8b4a6a13177e4f44fe2d1066fc84c946b63f4be",
	}

	res, err := srv.Update(platform_common.DEF_REGION_ID, indata)
	fmt.Println("Update UpdateAppZapsRegion", err)
	pretty.Println(res)

}

func UpdateAppMongoRegion(srv platform_services.RegionService) {

	indata := utils.Map{
		platform_common.FLD_REGION_ID:             platform_common.DEF_REGION_ID,
		platform_common.FLD_REGION_DB_TYPE:        db_common.DATABASE_TYPE_MONGODB,
		platform_common.FLD_REGION_MONGODB_SERVER: "clusterdev.v0ayj.mongodb.net",
		platform_common.FLD_REGION_MONGODB_NAME:   "zerveedev",
		platform_common.FLD_REGION_MONGODB_USER:   "zerdevdbuser",
		platform_common.FLD_REGION_MONGODB_SECRET: "O1sraRW5ePs3bala",
	}

	res, err := srv.Update(platform_common.DEF_REGION_ID, indata)
	fmt.Println("Update UpdateAppZapsRegion", err)
	pretty.Println(res)

}

func GetAppRegion(srv platform_services.RegionService) {
	res, err := srv.Get(platform_common.DEF_REGION_ID)
	fmt.Println("Get GetAppRegion", err)
	pretty.Println(res)

}

func CreateSysClient(srv platform_services.ClientsService) {

	indata := utils.Map{
		platform_common.FLD_CLIENT_ID:     "client001",
		platform_common.FLD_CLIENT_SECRET: "secret001",
		platform_common.FLD_CLIENT_TYPE:   "business", // app, business, user
		platform_common.FLD_CLIENT_SCOPE:  "business004",
	}

	res, err := srv.Create(indata)
	fmt.Println("Create SysClient", err)
	pretty.Println(res)

}

func ListSysClients(srv platform_services.ClientsService) {

	filter := "" //fmt.Sprintf(`{"%s":"%s"}`, platform_common.FLD_CLIENT_ID, "client001")
	res, err := srv.List(filter, "", 0, 0)
	fmt.Println("List clients success ", err)
	fmt.Println("List clients summary ", res)
	pretty.Print(res)
}

func FindSysClient(srv platform_services.ClientsService) {

	filter := fmt.Sprintf(`{platform_common.FLD_CLIENT_ID:"%s",platform_common.FLD_CLIENT_SECRET:"%s" }`, "client001", "secret001")
	res, err := srv.Find(filter)
	fmt.Println("Find FindSysBusiness", err)
	pretty.Println(res)

}

func GetSysClient(srv platform_services.ClientsService) {
	res, err := srv.Get("client001")
	fmt.Println("Get AppUser", err)
	pretty.Println(res)

}

func UpdateSysClient(srv platform_services.ClientsService) {

	sys_client_id := "client001"

	indata := utils.Map{
		platform_common.FLD_CLIENT_ID: sys_client_id,
		db_common.FLD_IS_DELETED:      false,
	}

	res, err := srv.Update(sys_client_id, indata)
	fmt.Println("Update AppUser", err)
	pretty.Println(res)
}

func DeleteSysClient(srv platform_services.ClientsService) {

	sys_client_id := "client001"
	err := srv.Delete(sys_client_id)
	fmt.Println("Delete SysClient", err)

}

func CreateSysSetting(srv platform_services.SysSettingService) {

	indata := utils.Map{
		platform_common.FLD_SETTING_ID:    "appname",
		platform_common.FLD_SETTING_VALUE: "Demo Application",
	}

	res, err := srv.Create(indata)
	fmt.Println("Create SysSetting", err)
	pretty.Println(res)

}

func ListSysSettings(srv platform_services.SysSettingService) {

	filter := "" //fmt.Sprintf(`{"%s":"%s"}`, platform_common.FLD_CLIENT_ID, "client001")
	res, err := srv.List(filter, "", 0, 0)
	fmt.Println("List clients success ", err)
	fmt.Println("List clients summary ", res)
	pretty.Print(res)
}

func FindSysSetting(srv platform_services.SysSettingService) {

	filter := fmt.Sprintf(`{platform_common.FLD_SETTING_ID:"%s" }`, "appname")
	res, err := srv.Find(filter)
	fmt.Println("Find FindSysBusiness", err)
	pretty.Println(res)

}

func GetSysSetting(srv platform_services.SysSettingService) {
	res, err := srv.Get("appname")
	fmt.Println("Get AppUser", err)
	pretty.Println(res)

}

func UpdateSysSetting(srv platform_services.SysSettingService) {

	sys_client_id := "appname"

	indata := utils.Map{
		platform_common.FLD_SETTING_ID: sys_client_id,
		db_common.FLD_IS_DELETED:       false,
	}

	res, err := srv.Update(sys_client_id, indata)
	fmt.Println("Update AppUser", err)
	pretty.Println(res)
}

func DeleteSysSetting(srv platform_services.SysSettingService) {

	app_setting_id := "appname"
	err := srv.Delete(app_setting_id)
	fmt.Println("Delete SysSetting", err)

}
