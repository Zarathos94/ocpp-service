package myservice

import (
	"encoding/xml"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// Type of string defining identification token, e.g. RFID or credit card number. To be treated as case insensitive.
type IdToken string

// Defines the authorization-status-value
type AuthorizationStatus string

const (
	AuthorizationStatusAccepted AuthorizationStatus = "Accepted"

	AuthorizationStatusBlocked AuthorizationStatus = "Blocked"

	AuthorizationStatusExpired AuthorizationStatus = "Expired"

	AuthorizationStatusInvalid AuthorizationStatus = "Invalid"

	AuthorizationStatusConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// Defines the unlock-status-value
type UnlockStatus string

const (
	UnlockStatusAccepted UnlockStatus = "Accepted"

	UnlockStatusRejected UnlockStatus = "Rejected"
)

// Defines the reset-type-value
type ResetType string

const (
	ResetTypeHard ResetType = "Hard"

	ResetTypeSoft ResetType = "Soft"
)

// Defines the reset-status-value
type ResetStatus string

const (
	ResetStatusAccepted ResetStatus = "Accepted"

	ResetStatusRejected ResetStatus = "Rejected"
)

// Defines the availability-type-value
type AvailabilityType string

const (
	AvailabilityTypeInoperative AvailabilityType = "Inoperative"

	AvailabilityTypeOperative AvailabilityType = "Operative"
)

// Defines the availability-status-value
type AvailabilityStatus string

const (
	AvailabilityStatusAccepted AvailabilityStatus = "Accepted"

	AvailabilityStatusRejected AvailabilityStatus = "Rejected"

	AvailabilityStatusScheduled AvailabilityStatus = "Scheduled"
)

// Defines the clear-cache-status-value
type ClearCacheStatus string

const (
	ClearCacheStatusAccepted ClearCacheStatus = "Accepted"

	ClearCacheStatusRejected ClearCacheStatus = "Rejected"
)

// Defines the configuration-status-value
type ConfigurationStatus string

const (
	ConfigurationStatusAccepted ConfigurationStatus = "Accepted"

	ConfigurationStatusRejected ConfigurationStatus = "Rejected"

	ConfigurationStatusNotSupported ConfigurationStatus = "NotSupported"
)

// Defines the remote-start-stop-status-value
type RemoteStartStopStatus string

const (
	RemoteStartStopStatusAccepted RemoteStartStopStatus = "Accepted"

	RemoteStartStopStatusRejected RemoteStartStopStatus = "Rejected"
)

type CancelReservationStatus string

const (
	CancelReservationStatusAccepted CancelReservationStatus = "Accepted"

	CancelReservationStatusRejected CancelReservationStatus = "Rejected"
)

// Defines the status returned in DataTransfer.conf
type DataTransferStatus string

const (
	DataTransferStatusAccepted DataTransferStatus = "Accepted"

	DataTransferStatusRejected DataTransferStatus = "Rejected"

	DataTransferStatusUnknownMessageId DataTransferStatus = "UnknownMessageId"

	DataTransferStatusUnknownVendorId DataTransferStatus = "UnknownVendorId"
)

type ReservationStatus string

const (
	ReservationStatusAccepted ReservationStatus = "Accepted"

	ReservationStatusFaulted ReservationStatus = "Faulted"

	ReservationStatusOccupied ReservationStatus = "Occupied"

	ReservationStatusRejected ReservationStatus = "Rejected"

	ReservationStatusUnavailable ReservationStatus = "Unavailable"
)

type UpdateType string

const (
	UpdateTypeDifferential UpdateType = "Differential"

	UpdateTypeFull UpdateType = "Full"
)

type UpdateStatus string

const (
	UpdateStatusAccepted UpdateStatus = "Accepted"

	UpdateStatusFailed UpdateStatus = "Failed"

	UpdateStatusHashError UpdateStatus = "HashError"

	UpdateStatusNotSupported UpdateStatus = "NotSupported"

	UpdateStatusVersionMismatch UpdateStatus = "VersionMismatch"
)

type IdTagInfo struct {
	Status AuthorizationStatus `xml:"status,omitempty"`

	ExpiryDate time.Time `xml:"expiryDate,omitempty"`

	ParentIdTag IdToken `xml:"parentIdTag,omitempty"`
}

type UnlockConnectorRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ unlockConnectorRequest"`

	ConnectorId int32 `xml:"connectorId,omitempty"`
}

type UnlockConnectorResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ unlockConnectorResponse"`

	Status *UnlockStatus `xml:"status,omitempty"`
}

type ResetRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ resetRequest"`

	Type_ *ResetType `xml:"type,omitempty"`
}

type ResetResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ resetResponse"`

	Status *ResetStatus `xml:"status,omitempty"`
}

type ChangeAvailabilityRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ changeAvailabilityRequest"`

	ConnectorId int32 `xml:"connectorId,omitempty"`

	Type_ *AvailabilityType `xml:"type,omitempty"`
}

type ChangeAvailabilityResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ changeAvailabilityResponse"`

	Status *AvailabilityStatus `xml:"status,omitempty"`
}

type GetDiagnosticsRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ getDiagnosticsRequest"`

	Location *AnyURI `xml:"location,omitempty"`

	StartTime time.Time `xml:"startTime,omitempty"`

	StopTime time.Time `xml:"stopTime,omitempty"`

	Retries int32 `xml:"retries,omitempty"`

	RetryInterval int32 `xml:"retryInterval,omitempty"`
}

type GetDiagnosticsResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ getDiagnosticsResponse"`

	FileName string `xml:"fileName,omitempty"`
}

type ClearCacheRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ clearCacheRequest"`
}

type ClearCacheResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ clearCacheResponse"`

	Status *ClearCacheStatus `xml:"status,omitempty"`
}

type UpdateFirmwareRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ updateFirmwareRequest"`

	RetrieveDate time.Time `xml:"retrieveDate,omitempty"`

	Location *AnyURI `xml:"location,omitempty"`

	Retries int32 `xml:"retries,omitempty"`

	RetryInterval int32 `xml:"retryInterval,omitempty"`
}

type UpdateFirmwareResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ updateFirmwareResponse"`
}

type ChangeConfigurationRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ changeConfigurationRequest"`

	Key string `xml:"key,omitempty"`

	Value string `xml:"value,omitempty"`
}

type ChangeConfigurationResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ changeConfigurationResponse"`

	Status *ConfigurationStatus `xml:"status,omitempty"`
}

type RemoteStartTransactionRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ remoteStartTransactionRequest"`

	IdTag IdToken `xml:"idTag,omitempty"`

	ConnectorId int32 `xml:"connectorId,omitempty"`
}

type RemoteStartTransactionResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ remoteStartTransactionResponse"`

	Status *RemoteStartStopStatus `xml:"status,omitempty"`
}

type RemoteStopTransactionRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ remoteStopTransactionRequest"`

	TransactionId int32 `xml:"transactionId,omitempty"`
}

type RemoteStopTransactionResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ remoteStopTransactionResponse"`

	Status *RemoteStartStopStatus `xml:"status,omitempty"`
}

type CancelReservationRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ cancelReservationRequest"`

	ReservationId int32 `xml:"reservationId,omitempty"`
}

type CancelReservationResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ cancelReservationResponse"`

	Status *CancelReservationStatus `xml:"status,omitempty"`
}

type DataTransferRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ dataTransferRequest"`

	VendorId string `xml:"vendorId,omitempty"`

	MessageId string `xml:"messageId,omitempty"`

	Data string `xml:"data,omitempty"`
}

type DataTransferResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ dataTransferResponse"`

	Status *DataTransferStatus `xml:"status,omitempty"`

	Data string `xml:"data,omitempty"`
}

type GetConfigurationRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ getConfigurationRequest"`

	Key []string `xml:"key,omitempty"`
}

type KeyValue struct {
	Key string `xml:"key,omitempty"`

	Readonly bool `xml:"readonly,omitempty"`

	Value string `xml:"value,omitempty"`
}

type GetConfigurationResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ getConfigurationResponse"`

	ConfigurationKey []*KeyValue `xml:"configurationKey,omitempty"`

	UnknownKey []string `xml:"unknownKey,omitempty"`
}

type GetLocalListVersionRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ getLocalListVersionRequest"`
}

type GetLocalListVersionResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ getLocalListVersionResponse"`

	ListVersion int32 `xml:"listVersion,omitempty"`
}

type ReserveNowRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ reserveNowRequest"`

	ConnectorId int32 `xml:"connectorId,omitempty"`

	ExpiryDate time.Time `xml:"expiryDate,omitempty"`

	IdTag IdToken `xml:"idTag,omitempty"`

	ParentIdTag IdToken `xml:"parentIdTag,omitempty"`

	ReservationId int32 `xml:"reservationId,omitempty"`
}

type ReserveNowResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ reserveNowResponse"`

	Status ReservationStatus `xml:"status,omitempty"`
}

type AuthorisationData struct {
	IdTag IdToken `xml:"idTag,omitempty"`

	IdTagInfo IdTagInfo `xml:"idTagInfo,omitempty"`
}

type SendLocalListRequest struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ sendLocalListRequest"`

	UpdateType UpdateType `xml:"updateType,omitempty"`

	ListVersion int32 `xml:"listVersion,omitempty"`

	LocalAuthorisationList []AuthorisationData `xml:"localAuthorisationList,omitempty"`

	Hash string `xml:"hash,omitempty"`
}

type SendLocalListResponse struct {
	XMLName xml.Name `xml:"urn://Ocpp/Cp/2012/06/ sendLocalListResponse"`

	Status UpdateStatus `xml:"status,omitempty"`

	Hash string `xml:"hash,omitempty"`
}

type ChargePointService interface {
	UnlockConnector(request *UnlockConnectorRequest) (*UnlockConnectorResponse, error)

	Reset(request *ResetRequest) (*ResetResponse, error)

	ChangeAvailability(request *ChangeAvailabilityRequest) (*ChangeAvailabilityResponse, error)

	GetDiagnostics(request *GetDiagnosticsRequest) (*GetDiagnosticsResponse, error)

	ClearCache(request *ClearCacheRequest) (*ClearCacheResponse, error)

	UpdateFirmware(request UpdateFirmwareRequest) (UpdateFirmwareResponse, error)

	ChangeConfiguration(request *ChangeConfigurationRequest) (*ChangeConfigurationResponse, error)

	RemoteStartTransaction(request *RemoteStartTransactionRequest) (*RemoteStartTransactionResponse, error)

	RemoteStopTransaction(request *RemoteStopTransactionRequest) (*RemoteStopTransactionResponse, error)

	CancelReservation(request *CancelReservationRequest) (*CancelReservationResponse, error)

	DataTransfer(request *DataTransferRequest) (*DataTransferResponse, error)

	GetConfiguration(request *GetConfigurationRequest) (*GetConfigurationResponse, error)

	GetLocalListVersion(request *GetLocalListVersionRequest) (*GetLocalListVersionResponse, error)

	ReserveNow(request *ReserveNowRequest) (*ReserveNowResponse, error)

	SendLocalList(request *SendLocalListRequest) (*SendLocalListResponse, error)
}

type chargePointService struct {
	client *soap.Client
}

func NewChargePointService(client *soap.Client) ChargePointService {
	return &chargePointService{
		client: client,
	}
}

func (service *chargePointService) UnlockConnector(request *UnlockConnectorRequest) (*UnlockConnectorResponse, error) {
	response := new(UnlockConnectorResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) Reset(request *ResetRequest) (*ResetResponse, error) {
	response := new(ResetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) ChangeAvailability(request *ChangeAvailabilityRequest) (*ChangeAvailabilityResponse, error) {
	response := new(ChangeAvailabilityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) GetDiagnostics(request *GetDiagnosticsRequest) (*GetDiagnosticsResponse, error) {
	response := new(GetDiagnosticsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) ClearCache(request *ClearCacheRequest) (*ClearCacheResponse, error) {
	response := new(ClearCacheResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) UpdateFirmware(request UpdateFirmwareRequest) (UpdateFirmwareResponse, error) {
	response := new(UpdateFirmwareResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) ChangeConfiguration(request *ChangeConfigurationRequest) (*ChangeConfigurationResponse, error) {
	response := new(ChangeConfigurationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) RemoteStartTransaction(request *RemoteStartTransactionRequest) (*RemoteStartTransactionResponse, error) {
	response := new(RemoteStartTransactionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) RemoteStopTransaction(request *RemoteStopTransactionRequest) (*RemoteStopTransactionResponse, error) {
	response := new(RemoteStopTransactionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) CancelReservation(request *CancelReservationRequest) (*CancelReservationResponse, error) {
	response := new(CancelReservationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) DataTransfer(request *DataTransferRequest) (*DataTransferResponse, error) {
	response := new(DataTransferResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) GetConfiguration(request *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	response := new(GetConfigurationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) GetLocalListVersion(request *GetLocalListVersionRequest) (*GetLocalListVersionResponse, error) {
	response := new(GetLocalListVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) ReserveNow(request *ReserveNowRequest) (*ReserveNowResponse, error) {
	response := new(ReserveNowResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *chargePointService) SendLocalList(request *SendLocalListRequest) (*SendLocalListResponse, error) {
	response := new(SendLocalListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
