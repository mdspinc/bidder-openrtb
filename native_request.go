package openrtb

import "encoding/json"

//go:generate ffjson -force-regenerate $GOFILE

type NativeRequestOuter struct {
	Native *NativeRequest `json:"native,omitempty"`
}

type NativeRequest struct {
	Ver      string                `json:"ver,omitempty"`
	Layout   int                   `json:"layout,omitempty"`
	Adunit   int                   `json:"adunit,omitempty"`
	Plcmtcnt int                   `json:"plcmtcnt,omitempty"`
	Seq      int                   `json:"seq,omitempty"`
	Assets   []*NativeRequestAsset `json:"assets,omitempty"`
	Ext      *json.RawMessage      `json:"ext,omitempty"`
}

type NativeRequestAsset struct {
	ID       int                 `json:"id,omitempty"`
	Required int                 `json:"required,omitempty"`
	Title    *NativeRequestTitle `json:"title,omitempty"`
	Image    *NativeRequestImage `json:"img,omitempty"`
	Video    *NativeRequestVideo `json:"video,omitempty"`
	Data     *NativeRequestData  `json:"data,omitempty"`
	Ext      *json.RawMessage    `json:"ext,omitempty"`
}

type NativeRequestData struct {
	Type int              `json:"type,omitempty"`
	Len  int              `json:"len,omitempty"`
	Ext  *json.RawMessage `json:"ext,omitempty"`
}

type NativeRequestImage struct {
	Type  int              `json:"type,omitempty"`
	H     int              `json:"h,omitempty"`
	W     int              `json:"w,omitempty"`
	HMin  int              `json:"hmin,omitempty"`
	WMin  int              `json:"wmin,omitempty"`
	Mimes []string         `json:"mimes,omitempty"`
	Ext   *json.RawMessage `json:"ext,omitempty"`
}

type NativeRequestTitle struct {
	Len int              `json:"len,omitempty"`
	Ext *json.RawMessage `json:"ext,omitempty"`
}

type NativeRequestVideo struct {
	Mimes       []string         `json:"mimes,omitempty"`       // Content MIME types supported.
	MinDuration int              `json:"minduration,omitempty"` // Minimum video ad duration in seconds
	MaxDuration int              `json:"maxduration,omitempty"` // Maximum video ad duration in seconds
	Protocols   []int            `json:"protocols,omitempty"`   // Video bid response protocols
	Ext         *json.RawMessage `json:"ext,omitempty"`
}
