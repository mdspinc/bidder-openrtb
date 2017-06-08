// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: video.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *jsonVideo) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *jsonVideo) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.Mimes) != 0 {
		buf.WriteString(`"mimes":`)
		if mj.Mimes != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Mimes {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.MinDuration != 0 {
		buf.WriteString(`"minduration":`)
		fflib.FormatBits2(buf, uint64(mj.MinDuration), 10, mj.MinDuration < 0)
		buf.WriteByte(',')
	}
	if mj.MaxDuration != 0 {
		buf.WriteString(`"maxduration":`)
		fflib.FormatBits2(buf, uint64(mj.MaxDuration), 10, mj.MaxDuration < 0)
		buf.WriteByte(',')
	}
	if mj.Protocol != 0 {
		buf.WriteString(`"protocol":`)
		fflib.FormatBits2(buf, uint64(mj.Protocol), 10, mj.Protocol < 0)
		buf.WriteByte(',')
	}
	if len(mj.Protocols) != 0 {
		buf.WriteString(`"protocols":`)
		if mj.Protocols != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Protocols {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.W != 0 {
		buf.WriteString(`"w":`)
		fflib.FormatBits2(buf, uint64(mj.W), 10, mj.W < 0)
		buf.WriteByte(',')
	}
	if mj.H != 0 {
		buf.WriteString(`"h":`)
		fflib.FormatBits2(buf, uint64(mj.H), 10, mj.H < 0)
		buf.WriteByte(',')
	}
	if mj.StartDelay != 0 {
		buf.WriteString(`"startdelay":`)
		fflib.FormatBits2(buf, uint64(mj.StartDelay), 10, mj.StartDelay < 0)
		buf.WriteByte(',')
	}
	if mj.Linearity != 0 {
		buf.WriteString(`"linearity":`)
		fflib.FormatBits2(buf, uint64(mj.Linearity), 10, mj.Linearity < 0)
		buf.WriteByte(',')
	}
	if mj.Sequence != 0 {
		buf.WriteString(`"sequence":`)
		fflib.FormatBits2(buf, uint64(mj.Sequence), 10, mj.Sequence < 0)
		buf.WriteByte(',')
	}
	if len(mj.BAttr) != 0 {
		buf.WriteString(`"battr":`)
		if mj.BAttr != nil {
			buf.WriteString(`[`)
			for i, v := range mj.BAttr {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.MaxExtended != 0 {
		buf.WriteString(`"maxextended":`)
		fflib.FormatBits2(buf, uint64(mj.MaxExtended), 10, mj.MaxExtended < 0)
		buf.WriteByte(',')
	}
	if mj.MinBitrate != 0 {
		buf.WriteString(`"minbitrate":`)
		fflib.FormatBits2(buf, uint64(mj.MinBitrate), 10, mj.MinBitrate < 0)
		buf.WriteByte(',')
	}
	if mj.MaxBitrate != 0 {
		buf.WriteString(`"maxbitrate":`)
		fflib.FormatBits2(buf, uint64(mj.MaxBitrate), 10, mj.MaxBitrate < 0)
		buf.WriteByte(',')
	}
	if mj.BoxingAllowed != nil {
		if true {
			buf.WriteString(`"boxingallowed":`)
			fflib.FormatBits2(buf, uint64(*mj.BoxingAllowed), 10, *mj.BoxingAllowed < 0)
			buf.WriteByte(',')
		}
	}
	if len(mj.PlaybackMethod) != 0 {
		buf.WriteString(`"playbackmethod":`)
		if mj.PlaybackMethod != nil {
			buf.WriteString(`[`)
			for i, v := range mj.PlaybackMethod {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.Delivery) != 0 {
		buf.WriteString(`"delivery":`)
		if mj.Delivery != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Delivery {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Pos != 0 {
		buf.WriteString(`"pos":`)
		fflib.FormatBits2(buf, uint64(mj.Pos), 10, mj.Pos < 0)
		buf.WriteByte(',')
	}
	if len(mj.CompanionAd) != 0 {
		buf.WriteString(`"companionad":`)
		if mj.CompanionAd != nil {
			buf.WriteString(`[`)
			for i, v := range mj.CompanionAd {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.Api) != 0 {
		buf.WriteString(`"api":`)
		if mj.Api != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Api {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.CompanionType) != 0 {
		buf.WriteString(`"companiontype":`)
		if mj.CompanionType != nil {
			buf.WriteString(`[`)
			for i, v := range mj.CompanionType {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.FormatBits2(buf, uint64(v), 10, v < 0)
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Skip != 0 {
		buf.WriteString(`"skip":`)
		fflib.FormatBits2(buf, uint64(mj.Skip), 10, mj.Skip < 0)
		buf.WriteByte(',')
	}
	if mj.Ext != nil {
		if true {
			buf.WriteString(`"ext":`)

			{

				obj, err = mj.Ext.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_jsonVideobase = iota
	ffj_t_jsonVideono_such_key

	ffj_t_jsonVideo_Mimes

	ffj_t_jsonVideo_MinDuration

	ffj_t_jsonVideo_MaxDuration

	ffj_t_jsonVideo_Protocol

	ffj_t_jsonVideo_Protocols

	ffj_t_jsonVideo_W

	ffj_t_jsonVideo_H

	ffj_t_jsonVideo_StartDelay

	ffj_t_jsonVideo_Linearity

	ffj_t_jsonVideo_Sequence

	ffj_t_jsonVideo_BAttr

	ffj_t_jsonVideo_MaxExtended

	ffj_t_jsonVideo_MinBitrate

	ffj_t_jsonVideo_MaxBitrate

	ffj_t_jsonVideo_BoxingAllowed

	ffj_t_jsonVideo_PlaybackMethod

	ffj_t_jsonVideo_Delivery

	ffj_t_jsonVideo_Pos

	ffj_t_jsonVideo_CompanionAd

	ffj_t_jsonVideo_Api

	ffj_t_jsonVideo_CompanionType

	ffj_t_jsonVideo_Skip

	ffj_t_jsonVideo_Ext
)

var ffj_key_jsonVideo_Mimes = []byte("mimes")

var ffj_key_jsonVideo_MinDuration = []byte("minduration")

var ffj_key_jsonVideo_MaxDuration = []byte("maxduration")

var ffj_key_jsonVideo_Protocol = []byte("protocol")

var ffj_key_jsonVideo_Protocols = []byte("protocols")

var ffj_key_jsonVideo_W = []byte("w")

var ffj_key_jsonVideo_H = []byte("h")

var ffj_key_jsonVideo_StartDelay = []byte("startdelay")

var ffj_key_jsonVideo_Linearity = []byte("linearity")

var ffj_key_jsonVideo_Sequence = []byte("sequence")

var ffj_key_jsonVideo_BAttr = []byte("battr")

var ffj_key_jsonVideo_MaxExtended = []byte("maxextended")

var ffj_key_jsonVideo_MinBitrate = []byte("minbitrate")

var ffj_key_jsonVideo_MaxBitrate = []byte("maxbitrate")

var ffj_key_jsonVideo_BoxingAllowed = []byte("boxingallowed")

var ffj_key_jsonVideo_PlaybackMethod = []byte("playbackmethod")

var ffj_key_jsonVideo_Delivery = []byte("delivery")

var ffj_key_jsonVideo_Pos = []byte("pos")

var ffj_key_jsonVideo_CompanionAd = []byte("companionad")

var ffj_key_jsonVideo_Api = []byte("api")

var ffj_key_jsonVideo_CompanionType = []byte("companiontype")

var ffj_key_jsonVideo_Skip = []byte("skip")

var ffj_key_jsonVideo_Ext = []byte("ext")

func (uj *jsonVideo) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *jsonVideo) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_jsonVideobase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_jsonVideono_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_jsonVideo_Api, kn) {
						currentKey = ffj_t_jsonVideo_Api
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffj_key_jsonVideo_BAttr, kn) {
						currentKey = ffj_t_jsonVideo_BAttr
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_BoxingAllowed, kn) {
						currentKey = ffj_t_jsonVideo_BoxingAllowed
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffj_key_jsonVideo_CompanionAd, kn) {
						currentKey = ffj_t_jsonVideo_CompanionAd
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_CompanionType, kn) {
						currentKey = ffj_t_jsonVideo_CompanionType
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_jsonVideo_Delivery, kn) {
						currentKey = ffj_t_jsonVideo_Delivery
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_jsonVideo_Ext, kn) {
						currentKey = ffj_t_jsonVideo_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_jsonVideo_H, kn) {
						currentKey = ffj_t_jsonVideo_H
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_jsonVideo_Linearity, kn) {
						currentKey = ffj_t_jsonVideo_Linearity
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_jsonVideo_Mimes, kn) {
						currentKey = ffj_t_jsonVideo_Mimes
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_MinDuration, kn) {
						currentKey = ffj_t_jsonVideo_MinDuration
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_MaxDuration, kn) {
						currentKey = ffj_t_jsonVideo_MaxDuration
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_MaxExtended, kn) {
						currentKey = ffj_t_jsonVideo_MaxExtended
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_MinBitrate, kn) {
						currentKey = ffj_t_jsonVideo_MinBitrate
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_MaxBitrate, kn) {
						currentKey = ffj_t_jsonVideo_MaxBitrate
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_jsonVideo_Protocol, kn) {
						currentKey = ffj_t_jsonVideo_Protocol
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_Protocols, kn) {
						currentKey = ffj_t_jsonVideo_Protocols
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_PlaybackMethod, kn) {
						currentKey = ffj_t_jsonVideo_PlaybackMethod
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_Pos, kn) {
						currentKey = ffj_t_jsonVideo_Pos
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_jsonVideo_StartDelay, kn) {
						currentKey = ffj_t_jsonVideo_StartDelay
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_Sequence, kn) {
						currentKey = ffj_t_jsonVideo_Sequence
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonVideo_Skip, kn) {
						currentKey = ffj_t_jsonVideo_Skip
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_jsonVideo_W, kn) {
						currentKey = ffj_t_jsonVideo_W
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_Ext, kn) {
					currentKey = ffj_t_jsonVideo_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_jsonVideo_Skip, kn) {
					currentKey = ffj_t_jsonVideo_Skip
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_CompanionType, kn) {
					currentKey = ffj_t_jsonVideo_CompanionType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_Api, kn) {
					currentKey = ffj_t_jsonVideo_Api
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_CompanionAd, kn) {
					currentKey = ffj_t_jsonVideo_CompanionAd
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_jsonVideo_Pos, kn) {
					currentKey = ffj_t_jsonVideo_Pos
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_Delivery, kn) {
					currentKey = ffj_t_jsonVideo_Delivery
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_jsonVideo_PlaybackMethod, kn) {
					currentKey = ffj_t_jsonVideo_PlaybackMethod
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_BoxingAllowed, kn) {
					currentKey = ffj_t_jsonVideo_BoxingAllowed
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_MaxBitrate, kn) {
					currentKey = ffj_t_jsonVideo_MaxBitrate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_MinBitrate, kn) {
					currentKey = ffj_t_jsonVideo_MinBitrate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_MaxExtended, kn) {
					currentKey = ffj_t_jsonVideo_MaxExtended
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_BAttr, kn) {
					currentKey = ffj_t_jsonVideo_BAttr
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_jsonVideo_Sequence, kn) {
					currentKey = ffj_t_jsonVideo_Sequence
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_Linearity, kn) {
					currentKey = ffj_t_jsonVideo_Linearity
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_jsonVideo_StartDelay, kn) {
					currentKey = ffj_t_jsonVideo_StartDelay
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_H, kn) {
					currentKey = ffj_t_jsonVideo_H
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_W, kn) {
					currentKey = ffj_t_jsonVideo_W
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_jsonVideo_Protocols, kn) {
					currentKey = ffj_t_jsonVideo_Protocols
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_Protocol, kn) {
					currentKey = ffj_t_jsonVideo_Protocol
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_MaxDuration, kn) {
					currentKey = ffj_t_jsonVideo_MaxDuration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonVideo_MinDuration, kn) {
					currentKey = ffj_t_jsonVideo_MinDuration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_jsonVideo_Mimes, kn) {
					currentKey = ffj_t_jsonVideo_Mimes
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_jsonVideono_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_jsonVideo_Mimes:
					goto handle_Mimes

				case ffj_t_jsonVideo_MinDuration:
					goto handle_MinDuration

				case ffj_t_jsonVideo_MaxDuration:
					goto handle_MaxDuration

				case ffj_t_jsonVideo_Protocol:
					goto handle_Protocol

				case ffj_t_jsonVideo_Protocols:
					goto handle_Protocols

				case ffj_t_jsonVideo_W:
					goto handle_W

				case ffj_t_jsonVideo_H:
					goto handle_H

				case ffj_t_jsonVideo_StartDelay:
					goto handle_StartDelay

				case ffj_t_jsonVideo_Linearity:
					goto handle_Linearity

				case ffj_t_jsonVideo_Sequence:
					goto handle_Sequence

				case ffj_t_jsonVideo_BAttr:
					goto handle_BAttr

				case ffj_t_jsonVideo_MaxExtended:
					goto handle_MaxExtended

				case ffj_t_jsonVideo_MinBitrate:
					goto handle_MinBitrate

				case ffj_t_jsonVideo_MaxBitrate:
					goto handle_MaxBitrate

				case ffj_t_jsonVideo_BoxingAllowed:
					goto handle_BoxingAllowed

				case ffj_t_jsonVideo_PlaybackMethod:
					goto handle_PlaybackMethod

				case ffj_t_jsonVideo_Delivery:
					goto handle_Delivery

				case ffj_t_jsonVideo_Pos:
					goto handle_Pos

				case ffj_t_jsonVideo_CompanionAd:
					goto handle_CompanionAd

				case ffj_t_jsonVideo_Api:
					goto handle_Api

				case ffj_t_jsonVideo_CompanionType:
					goto handle_CompanionType

				case ffj_t_jsonVideo_Skip:
					goto handle_Skip

				case ffj_t_jsonVideo_Ext:
					goto handle_Ext

				case ffj_t_jsonVideono_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Mimes:

	/* handler: uj.Mimes type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Mimes = nil
		} else {

			uj.Mimes = []string{}

			wantVal := true

			for {

				var tmp_uj__Mimes string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__Mimes type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__Mimes = string(string(outBuf))

					}
				}

				uj.Mimes = append(uj.Mimes, tmp_uj__Mimes)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MinDuration:

	/* handler: uj.MinDuration type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.MinDuration = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MaxDuration:

	/* handler: uj.MaxDuration type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.MaxDuration = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Protocol:

	/* handler: uj.Protocol type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Protocol = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Protocols:

	/* handler: uj.Protocols type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Protocols = nil
		} else {

			uj.Protocols = []int{}

			wantVal := true

			for {

				var tmp_uj__Protocols int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__Protocols type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmp_uj__Protocols = int(tval)

					}
				}

				uj.Protocols = append(uj.Protocols, tmp_uj__Protocols)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_W:

	/* handler: uj.W type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.W = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_H:

	/* handler: uj.H type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.H = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_StartDelay:

	/* handler: uj.StartDelay type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.StartDelay = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Linearity:

	/* handler: uj.Linearity type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Linearity = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Sequence:

	/* handler: uj.Sequence type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Sequence = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BAttr:

	/* handler: uj.BAttr type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.BAttr = nil
		} else {

			uj.BAttr = []int{}

			wantVal := true

			for {

				var tmp_uj__BAttr int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__BAttr type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmp_uj__BAttr = int(tval)

					}
				}

				uj.BAttr = append(uj.BAttr, tmp_uj__BAttr)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MaxExtended:

	/* handler: uj.MaxExtended type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.MaxExtended = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MinBitrate:

	/* handler: uj.MinBitrate type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.MinBitrate = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MaxBitrate:

	/* handler: uj.MaxBitrate type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.MaxBitrate = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BoxingAllowed:

	/* handler: uj.BoxingAllowed type=openrtb.NumberOrBool kind=int quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.BoxingAllowed = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		if uj.BoxingAllowed == nil {
			uj.BoxingAllowed = new(NumberOrBool)
		}

		err = uj.BoxingAllowed.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PlaybackMethod:

	/* handler: uj.PlaybackMethod type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.PlaybackMethod = nil
		} else {

			uj.PlaybackMethod = []int{}

			wantVal := true

			for {

				var tmp_uj__PlaybackMethod int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__PlaybackMethod type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmp_uj__PlaybackMethod = int(tval)

					}
				}

				uj.PlaybackMethod = append(uj.PlaybackMethod, tmp_uj__PlaybackMethod)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Delivery:

	/* handler: uj.Delivery type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Delivery = nil
		} else {

			uj.Delivery = []int{}

			wantVal := true

			for {

				var tmp_uj__Delivery int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__Delivery type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmp_uj__Delivery = int(tval)

					}
				}

				uj.Delivery = append(uj.Delivery, tmp_uj__Delivery)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Pos:

	/* handler: uj.Pos type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Pos = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CompanionAd:

	/* handler: uj.CompanionAd type=[]openrtb.Banner kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.CompanionAd = nil
		} else {

			uj.CompanionAd = []Banner{}

			wantVal := true

			for {

				var tmp_uj__CompanionAd Banner

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__CompanionAd type=openrtb.Banner kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					err = tmp_uj__CompanionAd.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.CompanionAd = append(uj.CompanionAd, tmp_uj__CompanionAd)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Api:

	/* handler: uj.Api type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Api = nil
		} else {

			uj.Api = []int{}

			wantVal := true

			for {

				var tmp_uj__Api int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__Api type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmp_uj__Api = int(tval)

					}
				}

				uj.Api = append(uj.Api, tmp_uj__Api)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CompanionType:

	/* handler: uj.CompanionType type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.CompanionType = nil
		} else {

			uj.CompanionType = []int{}

			wantVal := true

			for {

				var tmp_uj__CompanionType int

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__CompanionType type=int kind=int quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmp_uj__CompanionType = int(tval)

					}
				}

				uj.CompanionType = append(uj.CompanionType, tmp_uj__CompanionType)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Skip:

	/* handler: uj.Skip type=openrtb.NumberOrBool kind=int quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = uj.Skip.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ext:

	/* handler: uj.Ext type=json.RawMessage kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Ext = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		if uj.Ext == nil {
			uj.Ext = new(json.RawMessage)
		}

		err = uj.Ext.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
