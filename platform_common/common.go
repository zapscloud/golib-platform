package platform_common

import (
	"log"

	"github.com/zapscloud/golib-dbutils/db_common"
)

// ************************************
//
//  Database tables (collection) Names
//
// ************************************

// Sys module tables ========================================
const (
	DbPrefix = db_common.DB_COLLECTION_PREFIX
	// Platform Tables
	DbPlatformSettings = DbPrefix + "platform_settings"

	DbPlatformSysClients    = DbPrefix + "platform_sysclients"
	DbPlatformSysUsers      = DbPrefix + "platform_sysusers"
	DbPlatformSysRoles      = DbPrefix + "platform_sysroles"
	DbPlatformSysRoleCreds  = DbPrefix + "platform_sysrole_creds"
	DbPlatformSysRoleUsers  = DbPrefix + "platform_sysrole_users"
	DbPlatformSysUserAccess = DbPrefix + "platform_sysuser_business_access"

	DbPlatformAppClients   = DbPrefix + "platform_appclients"
	DbPlatformAppUsers     = DbPrefix + "platform_appusers"
	DbPlatformAppRoles     = DbPrefix + "platform_approles"
	DbPlatformAppRoleUsers = DbPrefix + "platform_approle_users"
	DbPlatformAppRoleCreds = DbPrefix + "platform_approle_creds"

	DbPlatformBusinessUser = DbPrefix + "platform_business_users"
	DbPlatformBusinesses   = DbPrefix + "platform_businesses"
	DbPlatformRegions      = DbPrefix + "platform_regions"
	DbPlatformCountries    = DbPrefix + "platform_countries"
	DbPlatformIndustries   = DbPrefix + "platform_industries"

	DbSites       = DbPrefix + "platform_sites"
	DbDepartments = DbPrefix + "platform_departments"
)

const (
	//
	// Sys Settings table fields
	FLD_SETTING_ID    = "setting_id"
	FLD_SETTING_VALUE = "setting_value"

	// Sys Access table fields
	FLD_SYS_ACCESS_ID            = "access_id"
	FLD_SYS_ACCESS_ROLE_ID       = "role_id"
	FLD_SYS_ACCESS_SITE_ID       = "site_id"
	FLD_SYS_ACCESS_DEPARTMENT_ID = "department_id"

	// Sys Role
	FLD_SYS_ROLE_ID      = "sys_role_id"
	FLD_SYS_ROLE_USER_ID = "sys_role_user_id"

	// Sys User
	FLD_SYS_USER_ID        = "user_id"
	FLD_SYS_USER_PASSWORD  = "password"
	FLD_SYS_USER_FIRSTNAME = "firstname"
	FLD_SYS_USER_LASTNAME  = "lastname"
	FLD_SYS_USER_EMAILID   = "email_id"
	FLD_SYS_USER_PHONE     = "phone"

	// App User
	FLD_APP_USER_ID       = "user_id"
	FLD_APP_USER_PASSWORD = "password"
	FLD_APP_USER_EMAILID  = "email_id"
	FLD_APP_USER_PHONE    = "phone"
	FLD_APP_USER_FNAME    = "first_name"
	FLD_APP_USER_LNAME    = "last_name"
	FLD_APP_USER_OTP      = "password_otp"

	// App Region table fields
	FLD_REGION_ID             = "region_id"
	FLD_REGION_NAME           = "region_name"
	FLD_REGION_DB_TYPE        = "database_type"
	FLD_REGION_MONGODB_SERVER = "mongodb_server"
	FLD_REGION_MONGODB_NAME   = "mongodb_name"
	FLD_REGION_MONGODB_USER   = "mongodb_user"
	FLD_REGION_MONGODB_SECRET = "mongodb_secret"
	FLD_REGION_ZAPSDB_APP     = "zapsdb_app"
	FLD_REGION_ZAPSDB_KEY     = "zapsdb_key"
	FLD_REGION_ZAPSDB_SECRET  = "zapsdb_secret"

	// App Business table fields
	FLD_BUSINESS_ID            = "business_id"
	FLD_BUSINESS_NAME          = "business_name"
	FLD_BUSINESS_EMAILID       = "business_email_id"
	FLD_BUSINESS_REGION_ID     = "region_id"
	FLD_BUSINESS_IS_TENANT_DB  = "tenant_db"
	FLD_BUSINESS_APPROVAL_CODE = "approval_code"

	// App Business User table fields
	FLD_BUSINESS_USER_ID = "business_user_id"

	// App Client Table
	FLD_CLIENT_ID     = "client_id"
	FLD_CLIENT_SECRET = "client_secret"
	FLD_CLIENT_TYPE   = "client_type"
	FLD_CLIENT_SCOPE  = "client_scope"

	// App Role
	FLD_APP_ROLE_ID      = "role_id"
	FLD_APP_ROLE_NAME    = "role_name"
	FLD_APP_ROLE_DESC    = "role_description"
	FLD_APP_ROLE_ISADMIN = "role_is_admin"

	FLD_APP_ROLE_USER_ID    = "user_id"
	FLD_APP_ROLE_CRED_ID    = "cred_id"
	FLD_APP_ROLE_CREDENTIAL = "credential"

	// Industry
	FLD_INDUSTRY_ID = "industry_id"

	// Country
	FLD_COUNTRY_ID = "country_id"
)

const (
	DEF_APP_ROLE_ADMIN_NAME = "admin"
	DEF_APP_ROLE_ADMIN_DESC = "Administrator"

	DEF_APP_ROLE_USER_NAME = "user"
	DEF_APP_ROLE_USER_DESC = "Normal User"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

}

func GetServiceModuleCode() string {
	return "S1"
}
