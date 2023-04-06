package platform_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-platform/platform_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// SysClientDao - User DAO Repository
type SysClientDao interface {
	InitializeDao(client utils.Map)
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	// Count(filter string, sort string, skip int64, limit int64) (int64, int64, error)
	// Update - Update Collection
	Update(clientid string, indata utils.Map) (utils.Map, error)
	// Find - Find by code
	Authenticate(clientid string, clientsecret string) (utils.Map, error)
	// Insert - Insert Collection
	Create(indata utils.Map) (string, error)
	// Find - Find by code
	Find(filter string) (utils.Map, error)

	// Delete - Delete Collection
	Delete(clientid string) (int64, error)

	Get(clientid string) (utils.Map, error)
}

// type appClientBaseDao struct {
// 	app_db_repository.SysClientDBDao
// 	instancename string
// }

// NewSysClientDao - Contruct User Dao
func NewSysClientDao(client utils.Map) SysClientDao {
	var daoSysClient SysClientDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoSysClient = &mongodb_repository.SysClientMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
		daoSysClient = nil
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
		daoSysClient = nil
	}

	if daoSysClient != nil {
		// Initialize the Dao
		daoSysClient.InitializeDao(client)
	}

	return daoSysClient
}
