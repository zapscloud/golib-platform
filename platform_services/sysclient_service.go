package platform_services

import (
	"fmt"
	"log"

	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-dbutils/db_utils"
	"github.com/zapscloud/golib-platform/platform_common"
	"github.com/zapscloud/golib-platform/platform_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// SysClientService - Users Service structure
type SysClientService interface {
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	Get(clientid string) (utils.Map, error)
	Find(filter string) (utils.Map, error)
	Create(indata utils.Map) (string, error)
	Update(clientid string, indata utils.Map) (utils.Map, error)
	Delete(clientid string) error
	Authenticate(useremail string, userpwd string) (utils.Map, error)

	BeginTransaction()
	CommitTransaction()
	RollbackTransaction()

	EndService()
}

type sysClientBaseService struct {
	db_utils.DatabaseService
	daoAppClient platform_repository.AppClientDao
	child        SysClientService
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func NewSysClientService(props utils.Map) (SysClientService, error) {
	p := sysClientBaseService{}

	err := p.OpenDatabaseService(props)
	if err != nil {
		return nil, err
	}

	log.Printf("appClientService ")
	p.daoAppClient = platform_repository.NewSysClientDao(p.GetClient())
	p.child = &p
	return &p, nil
}

func (p *sysClientBaseService) EndService() {
	p.CloseDatabaseService()
}

// List - List All records
func (p *sysClientBaseService) List(filter string, sort string, skip int64, limit int64) (utils.Map, error) {

	log.Println("SysClientService::FindAll - Begin")

	dataresponse, err := p.daoAppClient.List(filter, sort, skip, limit)
	if err != nil {
		return nil, err
	}
	log.Println("SysClientService::FindAll - End ")
	return dataresponse, nil
}

// GetDetails - Find By Code
func (p *sysClientBaseService) Get(clientid string) (utils.Map, error) {
	log.Printf("SysClientService::GetDetails::  Begin %v", clientid)

	data, err := p.daoAppClient.Get(clientid)

	log.Println("SysClientService::GetDetails:: End ", data, err)
	return data, err
}

func (p *sysClientBaseService) Find(filter string) (utils.Map, error) {
	fmt.Println("SysClientService::GetDetails::  Begin ", filter)

	data, err := p.daoAppClient.Find(filter)

	log.Println("SysClientService::GetDetails:: End ", data, err)
	return data, err
}

// Create - Create Service
func (p *sysClientBaseService) Create(indata utils.Map) (string, error) {

	log.Println("ClientService::Create - Begin")

	var clientId string

	dataval, dataok := indata[platform_common.FLD_CLIENT_ID]
	if dataok {
		clientId = dataval.(string)
	} else {
		err := &utils.AppError{ErrorCode: "S3040101", ErrorMsg: "Missing client_id", ErrorDetail: "Missing required field client_id !!"}
		return "", err
	}
	log.Println("Provided Profile ID:", clientId)

	_, err := p.daoAppClient.Get(clientId)
	if err == nil {
		err := &utils.AppError{ErrorCode: "S3040102", ErrorMsg: "Existing client_id", ErrorDetail: "Given client_id is already exist"}
		return dataval.(string), err
	}

	_, dataok = indata[platform_common.FLD_CLIENT_SECRET]
	if !dataok {
		err := &utils.AppError{ErrorCode: "S3040103", ErrorMsg: "Missing client_secret", ErrorDetail: "Missing required field client_secret !!"}
		return "", err
	}

	_, dataok = indata[platform_common.FLD_CLIENT_TYPE]
	if !dataok {
		err := &utils.AppError{ErrorCode: "S3040104", ErrorMsg: "Missing app_client_type", ErrorDetail: "Missing required field app_client_type !!"}
		return "", err
	}

	// Update converted clientId back to indata
	indata[platform_common.FLD_CLIENT_ID] = clientId

	createdId, err := p.daoAppClient.Create(indata)
	if err != nil {
		return "", err
	}
	log.Println("ClientService::Create - End ")
	return createdId, nil
}

// Update - Update Service
func (p *sysClientBaseService) Update(clientid string, indata utils.Map) (utils.Map, error) {

	log.Println("ClientService::Update - Begin")

	// Delete the Key fields
	delete(indata, platform_common.FLD_CLIENT_ID)

	data, err := p.daoAppClient.Update(clientid, indata)

	log.Println("ClientService::Update - End ")
	return data, err
}

// Delete - Delete Service
func (p *sysClientBaseService) Delete(clientid string) error {

	log.Println("ClientService::Delete - Begin", clientid)

	result, err := p.daoAppClient.Delete(clientid)
	if err != nil {
		return err
	}

	log.Printf("ClientService::Delete - End %v", result)
	return nil
}

// GetDetails - Find By Code
func (p *sysClientBaseService) Authenticate(clientid string, clientsecret string) (utils.Map, error) {
	log.Println("Authenticate::  Begin ", clientid, clientsecret)

	log.Println("User Password from API", clientsecret)
	dataUser, err := p.daoAppClient.Authenticate(clientid, clientsecret)

	log.Println("Length of dataUser :", dataUser)

	if err != nil {
		err := &utils.AppError{ErrorCode: "S30340101", ErrorMsg: "Wrong Credentials", ErrorDetail: "Authenticate credentials is wrong !!"}
		return utils.Map{}, err
	}

	dataval, dataok := dataUser[db_common.FLD_IS_DELETED]
	if dataok && !dataval.(bool) {
		err := &utils.AppError{ErrorCode: "S30340102", ErrorMsg: "User not in Active Mode. Contact Admin!", ErrorDetail: "User not in Active Mode. Contact Admin!"}
		return utils.Map{}, err
	}
	return dataUser, nil
}
