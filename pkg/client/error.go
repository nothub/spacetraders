package st

// see: https://docs.spacetraders.io/api-guide/response-errors

type ErrorResponse struct {
	Error struct {
		Message string         `json:"message"`
		Code    ErrorCode      `json:"code"`
		Data    map[string]any `json:"data"`
	} `json:"error"`
}

type ErrorCode int

// General Error Codes
const (
	CooldownConflictError ErrorCode = 4000
	WaypointNoAccessError ErrorCode = 4001
)

// Account Error Codes
const (
	TokenEmptyError                  ErrorCode = 4100
	TokenMissingSubjectError         ErrorCode = 4101
	TokenInvalidSubjectError         ErrorCode = 4102
	MissingTokenRequestError         ErrorCode = 4103
	InvalidTokenRequestError         ErrorCode = 4104
	InvalidTokenSubjectError         ErrorCode = 4105
	AccountNotExistsError            ErrorCode = 4106
	AgentNotExistsError              ErrorCode = 4107
	AccountHasNoAgentError           ErrorCode = 4108
	RegisterAgentExistsError         ErrorCode = 4109
	RegisterAgentSymbolReservedError ErrorCode = 4110
	RegisterAgentConflictSymbolError ErrorCode = 4111
)

// Ship Error Codes
const (
	NavigateInTransitError                    ErrorCode = 4200
	NavigateInvalidDestinationError           ErrorCode = 4201
	NavigateOutsideSystemError                ErrorCode = 4202
	NavigateInsufficientFuelError             ErrorCode = 4203
	NavigateSameDestinationError              ErrorCode = 4204
	ShipExtractInvalidWaypointError           ErrorCode = 4205
	ShipExtractPermissionError                ErrorCode = 4206
	ShipJumpNoSystemError                     ErrorCode = 4207
	ShipJumpSameSystemError                   ErrorCode = 4208
	ShipJumpMissingModuleError                ErrorCode = 4210
	ShipJumpNoValidWaypointError              ErrorCode = 4211
	ShipJumpMissingAntimatterError            ErrorCode = 4212
	ShipInTransitError                        ErrorCode = 4214
	ShipMissingSensorArraysError              ErrorCode = 4215
	PurchaseShipCreditsError                  ErrorCode = 4216
	ShipCargoExceedsLimitError                ErrorCode = 4217
	ShipCargoMissingError                     ErrorCode = 4218
	ShipCargoUnitCountError                   ErrorCode = 4219
	ShipSurveyVerificationError               ErrorCode = 4220
	ShipSurveyExpirationError                 ErrorCode = 4221
	ShipSurveyWaypointTypeError               ErrorCode = 4222
	ShipSurveyOrbitError                      ErrorCode = 4223
	ShipSurveyExhaustedError                  ErrorCode = 4224
	ShipRefuelDockedError                     ErrorCode = 4225
	ShipRefuelInvalidWaypointError            ErrorCode = 4226
	ShipMissingMountsError                    ErrorCode = 4227
	ShipCargoFullError                        ErrorCode = 4228
	ShipJumpFromGateToGateError               ErrorCode = 4229
	WaypointChartedError                      ErrorCode = 4230
	ShipTransferShipNotFound                  ErrorCode = 4231
	ShipTransferAgentConflict                 ErrorCode = 4232
	ShipTransferSameShipConflict              ErrorCode = 4233
	ShipTransferLocationConflict              ErrorCode = 4234
	WarpInsideSystemError                     ErrorCode = 4235
	ShipNotInOrbitError                       ErrorCode = 4236
	ShipInvalidRefineryGoodError              ErrorCode = 4237
	ShipInvalidRefineryTypeError              ErrorCode = 4238
	ShipMissingRefineryError                  ErrorCode = 4239
	ShipMissingSurveyorError                  ErrorCode = 4240
	ShipMissingWarpDriveError                 ErrorCode = 4241
	ShipMissingMineralProcessorError          ErrorCode = 4242
	ShipMissingMiningLasersError              ErrorCode = 4243
	ShipNotDockedError                        ErrorCode = 4244
	PurchaseShipNotPresentError               ErrorCode = 4245
	ShipMountNoShipyardError                  ErrorCode = 4246
	ShipMissingMountError                     ErrorCode = 4247
	ShipMountInsufficientCreditsError         ErrorCode = 4248
	ShipMissingPowerError                     ErrorCode = 4249
	ShipMissingSlotsError                     ErrorCode = 4250
	ShipMissingMountsError2                   ErrorCode = 4251
	ShipMissingCrewError                      ErrorCode = 4252
	ShipExtractDestabilizedError              ErrorCode = 4253
	ShipJumpInvalidOriginError                ErrorCode = 4254
	ShipJumpInvalidWaypointError              ErrorCode = 4255
	ShipJumpOriginUnderConstructionError      ErrorCode = 4256
	ShipMissingGasProcessorError              ErrorCode = 4257
	ShipMissingGasSiphonsError                ErrorCode = 4258
	ShipSiphonInvalidWaypointError            ErrorCode = 4259
	ShipSiphonPermissionError                 ErrorCode = 4260
	WaypointNoYieldError                      ErrorCode = 4261
	ShipJumpDestinationUnderConstructionError ErrorCode = 4262
)

// Contract Error Codes
const (
	AcceptContractNotAuthorizedError ErrorCode = 4500
	AcceptContractConflictError      ErrorCode = 4501
	FulfillContractDeliveryError     ErrorCode = 4502
	ContractDeadlineError            ErrorCode = 4503
	ContractFulfilledError           ErrorCode = 4504
	ContractNotAcceptedError         ErrorCode = 4505
	ContractNotAuthorizedError       ErrorCode = 4506
	ShipDeliverTermsError            ErrorCode = 4508
	ShipDeliverFulfilledError        ErrorCode = 4509
	ShipDeliverInvalidLocationError  ErrorCode = 4510
	ExistingContractError            ErrorCode = 4511
)

// Market Error Codes
const (
	MarketTradeInsufficientCreditsError ErrorCode = 4600
	MarketTradeNoPurchaseError          ErrorCode = 4601
	MarketTradeNotSoldError             ErrorCode = 4602
	MarketNotFoundError                 ErrorCode = 4603
	MarketTradeUnitLimitError           ErrorCode = 4604
)

// Faction Error Codes
const (
	WaypointNoFactionError ErrorCode = 4700
)

// Construction Error Code
const (
	ConstructionMaterialNotRequired      ErrorCode = 4800
	ConstructionMaterialFulfilled        ErrorCode = 4801
	ShipConstructionInvalidLocationError ErrorCode = 4802
)
