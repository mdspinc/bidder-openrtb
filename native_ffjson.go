// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: native.go

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *jsonNative) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *jsonNative) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "request":`)
	fflib.WriteJsonString(buf, string(j.Request))
	buf.WriteByte(',')
	if len(j.Ver) != 0 {
		buf.WriteString(`"ver":`)
		fflib.WriteJsonString(buf, string(j.Ver))
		buf.WriteByte(',')
	}
	if len(j.API) != 0 {
		buf.WriteString(`"api":`)
		if j.API != nil {
			buf.WriteString(`[`)
			for i, v := range j.API {
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
	if len(j.BAttr) != 0 {
		buf.WriteString(`"battr":`)
		if j.BAttr != nil {
			buf.WriteString(`[`)
			for i, v := range j.BAttr {
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
	if j.Ext != nil {
		if true {
			buf.WriteString(`"ext":`)

			{

				obj, err = j.Ext.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	if j.ParsedRequest != nil {
		if true {
			buf.WriteString(`"_parsed_request":`)

			{

				err = j.ParsedRequest.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtjsonNativebase = iota
	ffjtjsonNativenosuchkey

	ffjtjsonNativeRequest

	ffjtjsonNativeVer

	ffjtjsonNativeAPI

	ffjtjsonNativeBAttr

	ffjtjsonNativeExt

	ffjtjsonNativeParsedRequest
)

var ffjKeyjsonNativeRequest = []byte("request")

var ffjKeyjsonNativeVer = []byte("ver")

var ffjKeyjsonNativeAPI = []byte("api")

var ffjKeyjsonNativeBAttr = []byte("battr")

var ffjKeyjsonNativeExt = []byte("ext")

var ffjKeyjsonNativeParsedRequest = []byte("_parsed_request")

// UnmarshalJSON umarshall json - template of ffjson
func (j *jsonNative) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *jsonNative) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtjsonNativebase
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
				currentKey = ffjtjsonNativenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case '_':

					if bytes.Equal(ffjKeyjsonNativeParsedRequest, kn) {
						currentKey = ffjtjsonNativeParsedRequest
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'a':

					if bytes.Equal(ffjKeyjsonNativeAPI, kn) {
						currentKey = ffjtjsonNativeAPI
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyjsonNativeBAttr, kn) {
						currentKey = ffjtjsonNativeBAttr
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyjsonNativeExt, kn) {
						currentKey = ffjtjsonNativeExt
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyjsonNativeRequest, kn) {
						currentKey = ffjtjsonNativeRequest
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffjKeyjsonNativeVer, kn) {
						currentKey = ffjtjsonNativeVer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyjsonNativeParsedRequest, kn) {
					currentKey = ffjtjsonNativeParsedRequest
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyjsonNativeExt, kn) {
					currentKey = ffjtjsonNativeExt
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyjsonNativeBAttr, kn) {
					currentKey = ffjtjsonNativeBAttr
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyjsonNativeAPI, kn) {
					currentKey = ffjtjsonNativeAPI
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyjsonNativeVer, kn) {
					currentKey = ffjtjsonNativeVer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyjsonNativeRequest, kn) {
					currentKey = ffjtjsonNativeRequest
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtjsonNativenosuchkey
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

				case ffjtjsonNativeRequest:
					goto handle_Request

				case ffjtjsonNativeVer:
					goto handle_Ver

				case ffjtjsonNativeAPI:
					goto handle_API

				case ffjtjsonNativeBAttr:
					goto handle_BAttr

				case ffjtjsonNativeExt:
					goto handle_Ext

				case ffjtjsonNativeParsedRequest:
					goto handle_ParsedRequest

				case ffjtjsonNativenosuchkey:
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

handle_Request:

	/* handler: j.Request type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Request = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ver:

	/* handler: j.Ver type=openrtb.StringOrNumber kind=string quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Ver.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_API:

	/* handler: j.API type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.API = nil
		} else {

			j.API = []int{}

			wantVal := true

			for {

				var tmpJAPI int

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

				/* handler: tmpJAPI type=int kind=int quoted=false*/

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

						tmpJAPI = int(tval)

					}
				}

				j.API = append(j.API, tmpJAPI)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BAttr:

	/* handler: j.BAttr type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BAttr = nil
		} else {

			j.BAttr = []int{}

			wantVal := true

			for {

				var tmpJBAttr int

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

				/* handler: tmpJBAttr type=int kind=int quoted=false*/

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

						tmpJBAttr = int(tval)

					}
				}

				j.BAttr = append(j.BAttr, tmpJBAttr)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ext:

	/* handler: j.Ext type=json.RawMessage kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Ext = nil

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			if j.Ext == nil {
				j.Ext = new(json.RawMessage)
			}

			err = j.Ext.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ParsedRequest:

	/* handler: j.ParsedRequest type=openrtb.NativeRequest kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.ParsedRequest = nil

		} else {

			if j.ParsedRequest == nil {
				j.ParsedRequest = new(NativeRequest)
			}

			err = j.ParsedRequest.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
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
