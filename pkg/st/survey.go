package st

import "time"

// A resource survey of a waypoint, detailing a specific extraction location
// and the types of resources that can be found there.
type Survey struct {

	// A unique signature for the location of this survey. This signature is
	// verified when attempting an extraction using this survey.
	//
	//   >= 1 characters
	//
	// (required)
	Signature string `json:"signature"`

	// The symbol of the waypoint that this survey is for.
	//
	//   >= 1 characters
	//
	// (required)
	Symbol string `json:"symbol"`

	// A list of deposits that can be found at this location. A ship will
	// extract one of these deposits when using this survey in an extraction
	// request. If multiple deposits of the same type are present, the chance
	// of extracting that deposit is increased.
	//
	// (required)
	Deposits []SurveyDeposit `json:"deposits"`

	// The date and time when the survey expires. After this date and time, the
	// survey will no longer be available for extraction.
	//
	// (required)
	Expiration time.Time `json:"expiration"`

	// The size of the deposit. This value indicates how much can be extracted
	// from the survey before it is exhausted.
	//
	// (required)
	Size SurveySize `json:"size"`
}

// A surveyed deposit of a mineral or resource available for extraction.
type SurveyDeposit struct {
	// The symbol of the deposit.
	Symbol string `json:"symbol"`
}

type SurveySize string

const (
	SurveySizeSmall    SurveySize = "SMALL"
	SurveySizeModerate SurveySize = "MODERATE"
	SurveySizeLarge    SurveySize = "LARGE"
)
