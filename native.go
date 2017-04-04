package openrtb

//go:generate ffjson -force-regenerate $GOFILE

import "encoding/json"

// This object represents a native type impression. Native ad units are intended to blend seamlessly into
// the surrounding content (e.g., a sponsored Twitter or Facebook post). As such, the response must be
// well-structured to afford the publisher fine-grained control over rendering.
// The presence of a Native as a subordinate of the Imp object indicates that this impression is offered as
// a native type impression. At the publisher’s discretion, that same impression may also be offered as
// banner and/or video by also including as Imp subordinates the Banner and/or Video objects,
// respectively. However, any given bid for the impression must conform to one of the offered types.
type Native struct {
	Request string           `json:"request"`         // Request payload complying with the Native Ad Specification.
	Ver     interface{}      `json:"ver,omitempty"`   // Version of the Native Ad Specification to which request complies; highly recommended for efficient parsing.
	API     []int            `json:"api,omitempty"`   // List of supported API frameworks for this impression.
	BAttr   []int            `json:"battr,omitempty"` // Blocked creative attributes
	Ext     *json.RawMessage `json:"ext,omitempty"`

	ParsedRequest *NativeRequest `json:"_parsed_request,omitempty"` // custom field for parsed request payload
}

type jsonNative Native

// MarshalJSON custom marshalling with parsing request
func (v *Native) MarshalJSON() ([]byte, error) {
	if v.ParsedRequest != nil {
		req, err := json.Marshal(v.ParsedRequest)
		if err != nil {
			return []byte{}, err
		}
		v.Request = string(req)
	}
	return json.Marshal((*jsonNative)(v))
}

// UnmarshalJSON custom unmarshalling with parsing request
func (v *Native) UnmarshalJSON(data []byte) error {
	var h jsonNative
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	var nrOuter NativeRequestOuter
	if err := json.Unmarshal([]byte(h.Request), &nrOuter); err != nil {
		return err
	}

	if nrOuter.Native != nil {
		h.ParsedRequest = nrOuter.Native
	} else {
		if err := json.Unmarshal([]byte(h.Request), &h.ParsedRequest); err != nil {
			return err
		}
	}

	*v = (Native)(h)
	return nil
}
