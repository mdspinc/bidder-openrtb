package openrtb

//go:generate ffjson -force-regenerate $GOFILE

import "encoding/json"

// The "banner" object must be included directly in the impression object if the impression offered
// for auction is display or rich media, or it may be optionally embedded in the video object to
// describe the companion banners available for the linear or non-linear video ad.  The banner
// object may include a unique identifier; this can be useful if these IDs can be leveraged in the
// VAST response to dictate placement of the companion creatives when multiple companion ad
// opportunities of the same size are available on a page.
type Banner struct {
	Format   []Format         `json:"format,omitempty"`
	W        int              `json:"w,omitempty"`        // Width
	H        int              `json:"h,omitempty"`        // Height
	WMax     int              `json:"wmax,omitempty"`     // Width maximum
	HMax     int              `json:"hmax,omitempty"`     // Height maximum
	WMin     int              `json:"wmin,omitempty"`     // Width minimum
	HMin     int              `json:"hmin,omitempty"`     // Height minimum
	ID       string           `json:"id,omitempty"`       // A unique identifier
	Pos      int              `json:"pos,omitempty"`      // Ad Position
	BType    []int            `json:"btype,omitempty"`    // Blocked creative types
	BAttr    []int            `json:"battr,omitempty"`    // Blocked creative attributes
	Mimes    []string         `json:"mimes,omitempty"`    // Whitelist of content MIME types supported
	TopFrame NumberOrBool     `json:"topframe,omitempty"` // Default: 0 ("1": Delivered in top frame, "0": Elsewhere)
	ExpDir   []int            `json:"expdir,omitempty"`   // Specify properties for an expandable ad
	Api      []int            `json:"api,omitempty"`      // List of supported API frameworks
	Vcm      int              `json:"vcm,omitempty"`      // 0: concurrent, 1: end-card
	Ext      *json.RawMessage `json:"ext,omitempty"`
}

type Format struct {
	W      int              `json:"w,omitempty"`      // Width in DIPS
	H      int              `json:"h,omitempty"`      // Height in DIPS
	WRatio int              `json:"wratio,omitempty"` // Relative width when expressing size as a ratio.
	HRatio int              `json:"hratio,omitempty"` // Relative height when expressing size as a ratio.
	WMin   int              `json:"wmin,omitempty"`   // The minimum width in device independent pixels (DIPS) at which the ad will be displayed the size is expressed as a ratio.
	Ext    *json.RawMessage `json:"ext,omitempty"`
}
