package point

// Request definitions
const (
	ResetRequest = `<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
	<s12:Header>
	  <ns1:chargeBoxIdentity xmlns:ns1='urn://Ocpp/Cp/2012/06/'>{chargeBoxIdentity}</ns1:chargeBoxIdentity>
	</s12:Header>
	<s12:Body>
	  <ns1:resetRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/'>
  <!-- possible value: Hard, possible value: Soft -->
		<ns1:type>{type}</ns1:type>
	  </ns1:resetRequest>
	</s12:Body>
  </s12:Envelope>`
	ClearCacheRequest = `<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
	<s12:Header>
	  <ns1:chargeBoxIdentity xmlns:ns1='urn://Ocpp/Cp/2012/06/'>{chargeBoxIdentity}</ns1:chargeBoxIdentity>
	</s12:Header>
	<s12:Body>
	  <ns1:clearCacheRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/' />
	</s12:Body>
  </s12:Envelope>`

	ChangeAvailabilityRequest = `<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
	<s12:Header>
	  <ns1:chargeBoxIdentity xmlns:ns1='urn://Ocpp/Cp/2012/06/'>{chargeBoxIdentity}</ns1:chargeBoxIdentity>
	</s12:Header>
	<s12:Body>
	  <ns1:changeAvailabilityRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/'>
		<ns1:connectorId>{connectorID}</ns1:connectorId>
  <!-- possible value: Inoperative, possible value: Operative -->
		<ns1:type>{type}</ns1:type>
	  </ns1:changeAvailabilityRequest>
	</s12:Body>
  </s12:Envelope>`

	ChangeConfigurationRequest = `<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
	<s12:Header>
	  <ns1:chargeBoxIdentity xmlns:ns1='urn://Ocpp/Cp/2012/06/'>{chargeBoxIdentity}</ns1:chargeBoxIdentity>
	</s12:Header>
	<s12:Body>
	  <ns1:changeConfigurationRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/'>
		<ns1:key>{key}</ns1:key>
		<ns1:value>{value}</ns1:value>
	  </ns1:changeConfigurationRequest>
	</s12:Body>
  </s12:Envelope>`

	RemoteStartTransaction = `<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
  <s12:Header>
    <ns1:chargeBoxIdentity xmlns:ns1='urn://Ocpp/Cp/2012/06/'>{chargeBoxIdentity}</ns1:chargeBoxIdentity>
  </s12:Header>
  <s12:Body>
    <ns1:remoteStartTransactionRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/'>
<!-- Max Length: 20 -->
      <ns1:idTag>{idTag}</ns1:idTag>
<!-- optional -->
      <ns1:connectorId>{connectorID}</ns1:connectorId>
    </ns1:remoteStartTransactionRequest>
  </s12:Body>
</s12:Envelope>`

	RemoteStopTransaction = `<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
	<s12:Header>
	  <ns1:chargeBoxIdentity xmlns:ns1='urn://Ocpp/Cp/2012/06/'>{chargeBoxIdentity}</ns1:chargeBoxIdentity>
	</s12:Header>
	<s12:Body>
	  <ns1:remoteStopTransactionRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/'>
		<ns1:transactionId>{transactionID}</ns1:transactionId>
	  </ns1:remoteStopTransactionRequest>
	</s12:Body>
  </s12:Envelope>`
	CancelReservation = `<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
	<s12:Header>
	  <ns1:chargeBoxIdentity xmlns:ns1='urn://Ocpp/Cp/2012/06/'>{chargeBoxIdentity}</ns1:chargeBoxIdentity>
	</s12:Header>
	<s12:Body>
	  <ns1:cancelReservationRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/'>
		<ns1:reservationId>{reservationID}</ns1:reservationId>
	  </ns1:cancelReservationRequest>
	</s12:Body>
  </s12:Envelope>`
	GetConfigurationRequest = `<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
	<s12:Body>
	  <ns1:getConfigurationRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/'>
	  </ns1:getConfigurationRequest>
	</s12:Body>
  </s12:Envelope>`
	UnlockConnectorRequest = `<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
	<s12:Header>
		<ns1:chargeBoxIdentity xmlns:ns1='urn://Ocpp/Cp/2012/06/'>{chargeBoxIdentity}</ns1:chargeBoxIdentity>
	</s12:Header>
	<s12:Body>
		<ns1:unlockConnectorRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/'>
		<ns1:connectorId>{connectorID}</ns1:connectorId>
		</ns1:unlockConnectorRequest>
	</s12:Body>
	</s12:Envelope>`
)

/*
ChangeConfigurationRequest -
<s12:Envelope xmlns:s12='http://www.w3.org/2003/05/soap-envelope'>
  <s12:Header>
    <ns1:chargeBoxIdentity xmlns:ns1='urn://Ocpp/Cp/2012/06/'>?XXX?</ns1:chargeBoxIdentity>
  </s12:Header>
  <s12:Body>
    <ns1:changeConfigurationRequest xmlns:ns1='urn://Ocpp/Cp/2012/06/'>
      <ns1:key>?XXX?</ns1:key>
      <ns1:value>?XXX?</ns1:value>
    </ns1:changeConfigurationRequest>
  </s12:Body>
</s12:Envelope>

*/
