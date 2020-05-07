package point

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Zarathos94/ocpp-service/services/ocpp"
)

// GetRemoteConfiguration -
func (cp *CPointInterface) GetRemoteConfiguration() (map[string]string, error) {

	payload := strings.NewReader(GetConfigurationRequest)
	req, err := http.NewRequest("POST", cp.Config.URL, payload)

	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Error response: %v", res.Status)
	}
	var resp GetConfigurationResponse
	if err := xml.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	respConf := make(map[string]string)
	for _, li := range resp.Body.GetConfigurationResponse.ConfigurationKey {
		if li.Key == "STATIONID" {
			cp.Config.ChargeBoxIdentity = li.Value
		}
		respConf[li.Key] = li.Value
	}
	return respConf, nil
}

// ChangeRemoteConfiguration -
func (cp *CPointInterface) ChangeRemoteConfiguration(key string, value string) (bool, error) {
	requestChange := ChangeConfigurationRequest
	requestChange = strings.Replace(requestChange, "{chargeBoxIdentity}", cp.Config.ChargeBoxIdentity, -1)
	payload := strings.NewReader(requestChange)

	req, err := http.NewRequest("POST", cp.Config.URL, payload)

	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, fmt.Errorf("Error response: %v", res.Status)
	}
	var resp ChangeConfigurationResponse
	if err := xml.Unmarshal(body, &resp); err != nil {
		return false, err
	}
	if resp.Body.ChangeConfigurationResponse.Status != "Accepted" {
		return false, errors.New("Not Accepted")
	}
	return true, nil
}

// ClearCache -
func (cp *CPointInterface) ClearCache() (bool, error) {

	requestChange := ClearCacheRequest
	requestChange = strings.Replace(requestChange, "{chargeBoxIdentity}", cp.Config.ChargeBoxIdentity, -1)
	payload := strings.NewReader(requestChange)

	req, err := http.NewRequest("POST", cp.Config.URL, payload)

	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, fmt.Errorf("Error response: %v", res.Status)
	}
	var resp ClearCacheResponse
	if err := xml.Unmarshal(body, &resp); err != nil {
		return false, err
	}
	if resp.Body.ClearCacheResponse.Status != "Accepted" {
		return false, errors.New("Not Accepted")
	}
	return true, nil
}

// Reset -
func (cp *CPointInterface) Reset(resetType ocpp.ResetType) (bool, error) {

	requestChange := ResetRequest
	requestChange = strings.Replace(requestChange, "{chargeBoxIdentity}", cp.Config.ChargeBoxIdentity, -1)
	requestChange = strings.Replace(requestChange, "{type}", string(resetType), -1)
	payload := strings.NewReader(requestChange)

	req, err := http.NewRequest("POST", cp.Config.URL, payload)

	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, fmt.Errorf("Error response: %v", res.Status)
	}
	var resp ResetResponse
	if err := xml.Unmarshal(body, &resp); err != nil {
		return false, err
	}
	if resp.Body.ResetResponse.Status != "Accepted" {
		return false, errors.New("Not Accepted")
	}
	return true, nil
}

// RemoteStartTransaction -
func (cp *CPointInterface) RemoteStartTransaction(tag string, connectorID int32) (bool, error) {

	requestChange := RemoteStartTransaction
	requestChange = strings.Replace(requestChange, "{chargeBoxIdentity}", cp.Config.ChargeBoxIdentity, -1)
	requestChange = strings.Replace(requestChange, "{idTag}", tag, -1)
	requestChange = strings.Replace(requestChange, "{connectorID}", fmt.Sprintf("%d", connectorID), -1)
	payload := strings.NewReader(requestChange)
	req, err := http.NewRequest("POST", cp.Config.URL, payload)

	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, fmt.Errorf("Error response: %v", res.Status)
	}
	var resp RemoteStartTransactionResponse
	if err := xml.Unmarshal(body, &resp); err != nil {
		return false, err
	}
	if resp.Body.RemoteStartTransactionResponse.Status != "Accepted" {
		return false, errors.New("Not Accepted")
	}
	return true, nil
}

// RemoteStopTransaction -
func (cp *CPointInterface) RemoteStopTransaction(transactionID int32) (bool, error) {

	requestChange := RemoteStopTransaction
	requestChange = strings.Replace(requestChange, "{chargeBoxIdentity}", cp.Config.ChargeBoxIdentity, -1)
	requestChange = strings.Replace(requestChange, "{transactionID}", fmt.Sprintf("%d", transactionID), -1)
	payload := strings.NewReader(requestChange)
	req, err := http.NewRequest("POST", cp.Config.URL, payload)

	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, fmt.Errorf("Error response: %v", res.Status)
	}
	var resp RemoteStopTransactionResponse
	if err := xml.Unmarshal(body, &resp); err != nil {
		return false, err
	}
	if resp.Body.RemoteStopTransactionResponse.Status != "Accepted" {
		return false, errors.New("Not Accepted")
	}
	return true, nil
}

// ChangeAvailability -
func (cp *CPointInterface) ChangeAvailability(connectorID int32, availabilityType ocpp.AvailabilityType) (bool, error) {

	requestChange := ChangeAvailabilityRequest
	requestChange = strings.Replace(requestChange, "{chargeBoxIdentity}", cp.Config.ChargeBoxIdentity, -1)
	requestChange = strings.Replace(requestChange, "{connectorID}", fmt.Sprintf("%d", connectorID), -1)
	requestChange = strings.Replace(requestChange, "{type}", string(availabilityType), -1)
	payload := strings.NewReader(requestChange)
	req, err := http.NewRequest("POST", cp.Config.URL, payload)

	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, fmt.Errorf("Error response: %v", res.Status)
	}
	var resp ChangeAvailabilityResponse
	if err := xml.Unmarshal(body, &resp); err != nil {
		return false, err
	}
	if resp.Body.ChangeAvailabilityResponse.Status != "Accepted" {
		return false, errors.New("Not Accepted")
	}
	return true, nil
}

// UnlockConnector -
func (cp *CPointInterface) UnlockConnector(connectorID int32) (bool, error) {

	requestChange := RemoteStopTransaction
	requestChange = strings.Replace(requestChange, "{chargeBoxIdentity}", cp.Config.ChargeBoxIdentity, -1)
	requestChange = strings.Replace(requestChange, "{connectorID}", fmt.Sprintf("%d", connectorID), -1)
	payload := strings.NewReader(requestChange)
	req, err := http.NewRequest("POST", cp.Config.URL, payload)

	req.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return false, fmt.Errorf("Error response: %v", res.Status)
	}
	var resp UnlockConnectorResponse
	if err := xml.Unmarshal(body, &resp); err != nil {
		return false, err
	}
	if resp.Body.UnlockConnectorResponse.Status != "Accepted" {
		return false, errors.New("Not Accepted")
	}
	return true, nil
}
