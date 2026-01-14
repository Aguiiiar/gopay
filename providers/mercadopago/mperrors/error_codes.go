package mperrors

// MPErrorCode is Mercado Pago's "error" field value.
type MPErrorCode string

const (
	ErrCollectorDoesNotComplyWithCurrentRegulation MPErrorCode = "collector_does_not_comply_with_current_regulation"
	ErrInvalidCollectorID                          MPErrorCode = "invalid_collector_id"
	ErrInvalidSponsorID                            MPErrorCode = "invalid_sponsor_id"
	ErrInvalidCollectorEmail                       MPErrorCode = "invalid_collector_email"
	ErrInvalidOperationType                        MPErrorCode = "invalid_operation_type"
	ErrInvalidExpirationDateTo                     MPErrorCode = "invalid_expiration_date_to"
	ErrInvalidDate                                 MPErrorCode = "invalid_date"
	ErrInvalidExpirationDateFrom                   MPErrorCode = "invalid_expiration_date_from"
	ErrInvalidItems                                MPErrorCode = "invalid_items"
	ErrInvalidBackURLs                             MPErrorCode = "invalid_back_urls"
	ErrInvalidPaymentMethods                       MPErrorCode = "invalid_payment_methods"
	ErrInvalidMarketplaceFee                       MPErrorCode = "invalid_marketplace_fee"
	ErrInvalidID                                   MPErrorCode = "invalid_id"
	ErrInvalidAccessToken                          MPErrorCode = "invalid_access_token"
	ErrInvalidShipments                            MPErrorCode = "invalid_shipments"
	ErrInvalidBinaryMode                           MPErrorCode = "invalid_binary_mode"
	ErrSponsorIDSiteMustBeSameAsCollectorIDSite    MPErrorCode = "sponsor_id site must be the same as collector_id"
)
