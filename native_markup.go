package openrtb

import "encoding/json"

//go:generate ffjson -force-regenerate $GOFILE

type NativeAdMarkup struct {
	Native NativeAdMarkupInner `json:"native,omitempty"`
}

type NativeAdMarkupInner struct {
	Ver         int              `json:"ver,omitempty"`
	Assets      []*NativeAsset   `json:"assets,omitempty"`
	Link        string           `json:"link,omitempty"`
	ImpTrackers []string         `json:"imptrackers,omitempty"`
	JsTracker   string           `json:"jstracker,omitempty"`
	Ext         *json.RawMessage `json:"ext,omitempty"`
}

type NativeAsset struct {
	ID       int              `json:"id,omitempty"`
	Required int              `json:"required,omitempty"`
	Title    *NativeTitle     `json:"title,omitempty"`
	Image    *NativeImage     `json:"img,omitempty"`
	Video    *NativeVideo     `json:"video,omitempty"`
	Data     *NativeData      `json:"data,omitempty"`
	Link     *NativeLink      `json:"link,omitempty"`
	Ext      *json.RawMessage `json:"ext,omitempty"`
}

type NativeData struct {
	Value string           `json:"value,omitempty"`
	Label string           `json:"label,omitempty"`
	Ext   *json.RawMessage `json:"ext,omitempty"`
}

type NativeImage struct {
	URL string           `json:"url,omitempty"`
	H   int              `json:"h,omitempty"`
	W   int              `json:"w,omitempty"`
	Ext *json.RawMessage `json:"ext,omitempty"`
}

type NativeLink struct {
	URL           string           `json:"url,omitempty"`
	ClickTrackers []string         `json:"clicktrackers,omitempty"`
	Fallback      string           `json:"fallback,omitempty"`
	Ext           *json.RawMessage `json:"ext,omitempty"`
}

type NativeTitle struct {
	Text string           `json:"text,omitempty"`
	Ext  *json.RawMessage `json:"ext,omitempty"`
}

type NativeVideo struct {
	VASTTag string `json:"vasttag,omitempty"`
}
