package openrtb

//go:generate ffjson -force-regenerate $GOFILE

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidImpNoID        = errors.New("openrtb: impression ID missing")
	ErrInvalidImpNoAssets    = errors.New("openrtb: impression has no assets")       // neither Banner, nor Video, nor Native
	ErrInvalidImpMultiAssets = errors.New("openrtb: impression has multiple assets") // at least two out of Banner, Video, Native
)

// The "imp" object describes the ad position or impression being auctioned.  A single bid request
// can include multiple "imp" objects, a use case for which might be an exchange that supports
// selling all ad positions on a given page as a bundle.  Each "imp" object has a required ID so that
// bids can reference them individually.  An exchange can also conduct private auctions by
// restricting involvement to specific subsets of seats within bidders.
type Impression struct {
	ID                string           `json:"id"` // A unique identifier for this impression
	Banner            *Banner          `json:"banner,omitempty"`
	Video             *Video           `json:"video,omitempty"`
	Native            *Native          `json:"native,omitempty"`
	Audio             *json.RawMessage `json:"audio,omitempty"`
	DisplayManager    string           `json:"displaymanager,omitempty"`    // Name of ad mediation partner, SDK technology, etc
	DisplayManagerVer string           `json:"displaymanagerver,omitempty"` // Version of the above
	Instl             NumberOrBool     `json:"instl,omitempty"`             // Interstitial, Default: 0 ("1": Interstitial, "0": Something else)
	TagID             string           `json:"tagid,omitempty"`             // IDentifier for specific ad placement or ad tag
	BidFloor          float64          `json:"bidfloor,omitempty"`          // Bid floor for this impression in CPM
	BidFloorCurrency  string           `json:"bidfloorcur,omitempty"`       // Currency of bid floor
	ClickBrowser      int              `json:"clickbrowser,omitempty"`      // Indicates the type of browser opened upon clicking the creative in an app, where 0 = embedded, 1 = native.
	Secure            NumberOrBool     `json:"secure,omitempty"`            // Flag to indicate whether the impression requires secure HTTPS URL creative assets and markup.
	IFrameBuster      []string         `json:"iframebuster,omitempty"`      // Array of names for supportediframe busters.
	Pmp               *Pmp             `json:"pmp,omitempty"`               // A reference to the PMP object containing any Deals eligible for the impression object.
	Exp               int              `json:"exp,omitempty"`               // Advisory as to the number of seconds that may elapse between the auction and the actual impression.
	Metric            *json.RawMessage `json:"metric,omitempty"`
	Ext               *json.RawMessage `json:"ext,omitempty"`
}

func (imp *Impression) assetCount() int {
	n := 0
	if imp.Banner != nil {
		n++
	}
	if imp.Video != nil {
		n++
	}
	if imp.Native != nil {
		n++
	}
	return n
}

// Validates the `imp` object
func (imp *Impression) Validate() error {
	if imp.ID == "" {
		return ErrInvalidImpNoID
	}

	if count := imp.assetCount(); count == 0 {
		return ErrInvalidImpNoAssets
	} else if count > 1 {
		return ErrInvalidImpMultiAssets
	}

	if imp.Video != nil {
		if err := imp.Video.Validate(); err != nil {
			return err
		}
	}

	return nil
}
