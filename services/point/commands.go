package point

import (
	"errors"
	"fmt"
	"time"

	"github.com/Zarathos94/ocpp-service/services/ocpp"
)

// GetRemoteConfiguration -
func (cp *CPointInterface) GetRemoteConfiguration() (map[string]string, error) {
	cp.preSetActionHeaders("/GetConfiguration")
	conf, err := cp.CPService.GetConfiguration(&ocpp.GetConfigurationRequest{})
	if err != nil {
		return nil, err
	}
	resp := make(map[string]string)
	for _, v := range conf.ConfigurationKey {
		resp[v.Key] = v.Value
	}
	return resp, nil
}

// ChangeRemoteConfiguration -
func (cp *CPointInterface) ChangeRemoteConfiguration(key string, value string) (bool, error) {

	cp.preSetActionHeaders("/ChangeConfiguration")

	resp, err := cp.CPService.ChangeConfiguration(&ocpp.ChangeConfigurationRequest{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.ConfigurationStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}

	return true, nil
}

// ClearCache -
func (cp *CPointInterface) ClearCache() (bool, error) {

	cp.preSetActionHeaders("/ClearCache")
	resp, err := cp.CPService.ClearCache(&ocpp.ClearCacheRequest{})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.ClearCacheStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}

	return true, nil
}

// Reset -
func (cp *CPointInterface) Reset(resetType ocpp.ResetType) (bool, error) {

	cp.preSetActionHeaders("/Reset")
	resp, err := cp.CPService.Reset(&ocpp.ResetRequest{
		Type_: resetType,
	})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.ResetStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}

	return true, nil
}

// RemoteStartTransaction -
func (cp *CPointInterface) RemoteStartTransaction(tag string, connectorID int32) (bool, error) {

	cp.preSetActionHeaders("/RemoteStartTransaction")
	resp, err := cp.CPService.RemoteStartTransaction(&ocpp.RemoteStartTransactionRequest{
		IdTag:       ocpp.IdToken(tag),
		ConnectorId: connectorID,
	})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.RemoteStartStopStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}

	return true, nil
}

// RemoteStopTransaction -
func (cp *CPointInterface) RemoteStopTransaction(transactionID int32) (bool, error) {

	cp.preSetActionHeaders("/RemoteStopTransaction")
	resp, err := cp.CPService.RemoteStopTransaction(&ocpp.RemoteStopTransactionRequest{
		TransactionId: transactionID,
	})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.RemoteStartStopStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}

	return true, nil
}

// ChangeAvailability -
func (cp *CPointInterface) ChangeAvailability(connectorID int32, availabilityType ocpp.AvailabilityType) (bool, error) {

	cp.preSetActionHeaders("/ChangeAvailability")
	resp, err := cp.CPService.ChangeAvailability(&ocpp.ChangeAvailabilityRequest{
		Type_:       availabilityType,
		ConnectorId: connectorID,
	})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.AvailabilityStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}

	return true, nil
}

// ReserveNow -
func (cp *CPointInterface) ReserveNow(connectorID int32, expiryDate time.Time, idTag string, parentIDTag string, reservationID int32) (bool, error) {

	if expiryDate.Unix() < time.Now().Unix() {
		return false, errors.New("Invalid expiration time")
	}

	cp.preSetActionHeaders("/ReserveNow")
	resp, err := cp.CPService.ReserveNow(&ocpp.ReserveNowRequest{
		ConnectorId:   connectorID,
		ExpiryDate:    expiryDate, //ocpp.Time(expiryDate.Format("2006-01-02 15:04:05")),
		IdTag:         ocpp.IdToken(idTag),
		ParentIdTag:   ocpp.IdToken(parentIDTag),
		ReservationId: reservationID,
	})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.ReservationStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}

	return true, nil
}

// CancelReservation -
func (cp *CPointInterface) CancelReservation(reservationID int32) (bool, error) {

	cp.preSetActionHeaders("/CancelReservation")
	resp, err := cp.CPService.CancelReservation(&ocpp.CancelReservationRequest{
		ReservationId: reservationID,
	})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.CancelReservationStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}

	return true, nil
}

// UnlockConnector -
func (cp *CPointInterface) UnlockConnector(connectorID int32) (bool, error) {

	cp.preSetActionHeaders("/UnlockConnector")
	resp, err := cp.CPService.UnlockConnector(&ocpp.UnlockConnectorRequest{
		ConnectorId: connectorID,
	})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.UnlockStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}

	return true, nil
}

// DataTransfer -
func (cp *CPointInterface) DataTransfer(vendorID, messageID, data string) (bool, error) {

	cp.preSetActionHeaders("/DataTransfer")
	resp, err := cp.CPService.DataTransfer(&ocpp.DataTransferRequest{
		VendorId:  vendorID,
		MessageId: messageID,
		Data:      data,
	})
	if err != nil {
		return false, err
	}
	if resp.Status != ocpp.DataTransferStatusAccepted {
		return false, fmt.Errorf("Invalid response status: `%s`", resp.Status)
	}
	if resp.Status != ocpp.DataTransferStatus("Accepted") {
		return false, fmt.Errorf("Response: %s", resp.Status)
	}

	return true, nil
}
