package point

import "encoding/xml"

// GetConfigurationResponse -
type GetConfigurationResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text                     string `xml:",chardata"`
		GetConfigurationResponse struct {
			Text             string `xml:",chardata"`
			Xmlns            string `xml:"xmlns,attr"`
			ConfigurationKey []struct {
				Text     string `xml:",chardata"`
				Key      string `xml:"key"`
				Readonly string `xml:"readonly"`
				Value    string `xml:"value"`
			} `xml:"configurationKey"`
		} `xml:"getConfigurationResponse"`
	} `xml:"Body"`
}

// ChangeConfigurationResponse -
type ChangeConfigurationResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text                        string `xml:",chardata"`
		ChangeConfigurationResponse struct {
			Text   string `xml:",chardata"`
			Xmlns  string `xml:"xmlns,attr"`
			Status string `xml:"status"`
		} `xml:"changeConfigurationResponse"`
	} `xml:"Body"`
}

// ResetResponse -
type ResetResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text          string `xml:",chardata"`
		ResetResponse struct {
			Text   string `xml:",chardata"`
			Xmlns  string `xml:"xmlns,attr"`
			Status string `xml:"status"`
		} `xml:"resetResponse"`
	} `xml:"Body"`
}

// RemoteStartTransactionResponse -
type RemoteStartTransactionResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text                           string `xml:",chardata"`
		RemoteStartTransactionResponse struct {
			Text   string `xml:",chardata"`
			Xmlns  string `xml:"xmlns,attr"`
			Status string `xml:"status"`
		} `xml:"remoteStartTransactionResponse"`
	} `xml:"Body"`
}

// RemoteStopTransactionResponse -
type RemoteStopTransactionResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text                          string `xml:",chardata"`
		RemoteStopTransactionResponse struct {
			Text   string `xml:",chardata"`
			Xmlns  string `xml:"xmlns,attr"`
			Status string `xml:"status"`
		} `xml:"remoteStopTransactionResponse"`
	} `xml:"Body"`
}

// ChangeAvailabilityResponse -
type ChangeAvailabilityResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text                       string `xml:",chardata"`
		ChangeAvailabilityResponse struct {
			Text   string `xml:",chardata"`
			Xmlns  string `xml:"xmlns,attr"`
			Status string `xml:"status"`
		} `xml:"changeAvailabilityResponse"`
	} `xml:"Body"`
}

// UnlockConnectorResponse -
type UnlockConnectorResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text                    string `xml:",chardata"`
		UnlockConnectorResponse struct {
			Text   string `xml:",chardata"`
			Xmlns  string `xml:"xmlns,attr"`
			Status string `xml:"status"`
		} `xml:"unlockConnectorResponse"`
	} `xml:"Body"`
}

// ClearCacheResponse -
type ClearCacheResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text               string `xml:",chardata"`
		ClearCacheResponse struct {
			Text   string `xml:",chardata"`
			Xmlns  string `xml:"xmlns,attr"`
			Status string `xml:"status"`
		} `xml:"clearCacheResponse"`
	} `xml:"Body"`
}
