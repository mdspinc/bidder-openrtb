// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: impression.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Impression) MarshalJSON() ([]byte, error) {
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
func (mj *Impression) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "id":`)
	fflib.WriteJsonString(buf, string(mj.ID))
	buf.WriteByte(',')
	if mj.Banner != nil {
		if true {
			buf.WriteString(`"banner":`)

			{

				err = mj.Banner.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.Video != nil {
		if true {
			buf.WriteString(`"video":`)

			{

				obj, err = mj.Video.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	if mj.Native != nil {
		if true {
			buf.WriteString(`"native":`)

			{

				obj, err = mj.Native.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	if len(mj.DisplayManager) != 0 {
		buf.WriteString(`"displaymanager":`)
		fflib.WriteJsonString(buf, string(mj.DisplayManager))
		buf.WriteByte(',')
	}
	if len(mj.DisplayManagerVer) != 0 {
		buf.WriteString(`"displaymanagerver":`)
		fflib.WriteJsonString(buf, string(mj.DisplayManagerVer))
		buf.WriteByte(',')
	}
	if mj.Instl != 0 {
		buf.WriteString(`"instl":`)
		fflib.FormatBits2(buf, uint64(mj.Instl), 10, mj.Instl < 0)
		buf.WriteByte(',')
	}
	if len(mj.TagID) != 0 {
		buf.WriteString(`"tagid":`)
		fflib.WriteJsonString(buf, string(mj.TagID))
		buf.WriteByte(',')
	}
	if mj.BidFloor != 0 {
		buf.WriteString(`"bidfloor":`)
		fflib.AppendFloat(buf, float64(mj.BidFloor), 'g', -1, 64)
		buf.WriteByte(',')
	}
	if len(mj.BidFloorCurrency) != 0 {
		buf.WriteString(`"bidfloorcur":`)
		fflib.WriteJsonString(buf, string(mj.BidFloorCurrency))
		buf.WriteByte(',')
	}
	if mj.Secure != 0 {
		buf.WriteString(`"secure":`)
		fflib.FormatBits2(buf, uint64(mj.Secure), 10, mj.Secure < 0)
		buf.WriteByte(',')
	}
	if len(mj.IFrameBuster) != 0 {
		buf.WriteString(`"iframebuster":`)
		if mj.IFrameBuster != nil {
			buf.WriteString(`[`)
			for i, v := range mj.IFrameBuster {
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
	if mj.Pmp != nil {
		if true {
			/* Struct fall back. type=openrtb.Pmp kind=struct */
			buf.WriteString(`"pmp":`)
			err = buf.Encode(mj.Pmp)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if mj.Exp != 0 {
		buf.WriteString(`"exp":`)
		fflib.FormatBits2(buf, uint64(mj.Exp), 10, mj.Exp < 0)
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
	ffj_t_Impressionbase = iota
	ffj_t_Impressionno_such_key

	ffj_t_Impression_ID

	ffj_t_Impression_Banner

	ffj_t_Impression_Video

	ffj_t_Impression_Native

	ffj_t_Impression_DisplayManager

	ffj_t_Impression_DisplayManagerVer

	ffj_t_Impression_Instl

	ffj_t_Impression_TagID

	ffj_t_Impression_BidFloor

	ffj_t_Impression_BidFloorCurrency

	ffj_t_Impression_Secure

	ffj_t_Impression_IFrameBuster

	ffj_t_Impression_Pmp

	ffj_t_Impression_Exp

	ffj_t_Impression_Ext
)

var ffj_key_Impression_ID = []byte("id")

var ffj_key_Impression_Banner = []byte("banner")

var ffj_key_Impression_Video = []byte("video")

var ffj_key_Impression_Native = []byte("native")

var ffj_key_Impression_DisplayManager = []byte("displaymanager")

var ffj_key_Impression_DisplayManagerVer = []byte("displaymanagerver")

var ffj_key_Impression_Instl = []byte("instl")

var ffj_key_Impression_TagID = []byte("tagid")

var ffj_key_Impression_BidFloor = []byte("bidfloor")

var ffj_key_Impression_BidFloorCurrency = []byte("bidfloorcur")

var ffj_key_Impression_Secure = []byte("secure")

var ffj_key_Impression_IFrameBuster = []byte("iframebuster")

var ffj_key_Impression_Pmp = []byte("pmp")

var ffj_key_Impression_Exp = []byte("exp")

var ffj_key_Impression_Ext = []byte("ext")

func (uj *Impression) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Impression) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Impressionbase
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
				currentKey = ffj_t_Impressionno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffj_key_Impression_Banner, kn) {
						currentKey = ffj_t_Impression_Banner
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Impression_BidFloor, kn) {
						currentKey = ffj_t_Impression_BidFloor
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Impression_BidFloorCurrency, kn) {
						currentKey = ffj_t_Impression_BidFloorCurrency
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_Impression_DisplayManager, kn) {
						currentKey = ffj_t_Impression_DisplayManager
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Impression_DisplayManagerVer, kn) {
						currentKey = ffj_t_Impression_DisplayManagerVer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Impression_Exp, kn) {
						currentKey = ffj_t_Impression_Exp
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Impression_Ext, kn) {
						currentKey = ffj_t_Impression_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_Impression_ID, kn) {
						currentKey = ffj_t_Impression_ID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Impression_Instl, kn) {
						currentKey = ffj_t_Impression_Instl
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Impression_IFrameBuster, kn) {
						currentKey = ffj_t_Impression_IFrameBuster
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffj_key_Impression_Native, kn) {
						currentKey = ffj_t_Impression_Native
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Impression_Pmp, kn) {
						currentKey = ffj_t_Impression_Pmp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_Impression_Secure, kn) {
						currentKey = ffj_t_Impression_Secure
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Impression_TagID, kn) {
						currentKey = ffj_t_Impression_TagID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffj_key_Impression_Video, kn) {
						currentKey = ffj_t_Impression_Video
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_Ext, kn) {
					currentKey = ffj_t_Impression_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_Exp, kn) {
					currentKey = ffj_t_Impression_Exp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_Pmp, kn) {
					currentKey = ffj_t_Impression_Pmp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Impression_IFrameBuster, kn) {
					currentKey = ffj_t_Impression_IFrameBuster
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Impression_Secure, kn) {
					currentKey = ffj_t_Impression_Secure
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_BidFloorCurrency, kn) {
					currentKey = ffj_t_Impression_BidFloorCurrency
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_BidFloor, kn) {
					currentKey = ffj_t_Impression_BidFloor
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_TagID, kn) {
					currentKey = ffj_t_Impression_TagID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Impression_Instl, kn) {
					currentKey = ffj_t_Impression_Instl
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Impression_DisplayManagerVer, kn) {
					currentKey = ffj_t_Impression_DisplayManagerVer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Impression_DisplayManager, kn) {
					currentKey = ffj_t_Impression_DisplayManager
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_Native, kn) {
					currentKey = ffj_t_Impression_Native
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_Video, kn) {
					currentKey = ffj_t_Impression_Video
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_Banner, kn) {
					currentKey = ffj_t_Impression_Banner
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Impression_ID, kn) {
					currentKey = ffj_t_Impression_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Impressionno_such_key
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

				case ffj_t_Impression_ID:
					goto handle_ID

				case ffj_t_Impression_Banner:
					goto handle_Banner

				case ffj_t_Impression_Video:
					goto handle_Video

				case ffj_t_Impression_Native:
					goto handle_Native

				case ffj_t_Impression_DisplayManager:
					goto handle_DisplayManager

				case ffj_t_Impression_DisplayManagerVer:
					goto handle_DisplayManagerVer

				case ffj_t_Impression_Instl:
					goto handle_Instl

				case ffj_t_Impression_TagID:
					goto handle_TagID

				case ffj_t_Impression_BidFloor:
					goto handle_BidFloor

				case ffj_t_Impression_BidFloorCurrency:
					goto handle_BidFloorCurrency

				case ffj_t_Impression_Secure:
					goto handle_Secure

				case ffj_t_Impression_IFrameBuster:
					goto handle_IFrameBuster

				case ffj_t_Impression_Pmp:
					goto handle_Pmp

				case ffj_t_Impression_Exp:
					goto handle_Exp

				case ffj_t_Impression_Ext:
					goto handle_Ext

				case ffj_t_Impressionno_such_key:
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

handle_ID:

	/* handler: uj.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Banner:

	/* handler: uj.Banner type=openrtb.Banner kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Banner = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Banner == nil {
			uj.Banner = new(Banner)
		}

		err = uj.Banner.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Video:

	/* handler: uj.Video type=openrtb.Video kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Video = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		if uj.Video == nil {
			uj.Video = new(Video)
		}

		err = uj.Video.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Native:

	/* handler: uj.Native type=openrtb.Native kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Native = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		if uj.Native == nil {
			uj.Native = new(Native)
		}

		err = uj.Native.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DisplayManager:

	/* handler: uj.DisplayManager type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.DisplayManager = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DisplayManagerVer:

	/* handler: uj.DisplayManagerVer type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.DisplayManagerVer = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Instl:

	/* handler: uj.Instl type=int kind=int quoted=false*/

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

			uj.Instl = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TagID:

	/* handler: uj.TagID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.TagID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BidFloor:

	/* handler: uj.BidFloor type=float64 kind=float64 quoted=false*/

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.BidFloor = float64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BidFloorCurrency:

	/* handler: uj.BidFloorCurrency type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.BidFloorCurrency = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Secure:

	/* handler: uj.Secure type=int kind=int quoted=false*/

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

			uj.Secure = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IFrameBuster:

	/* handler: uj.IFrameBuster type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.IFrameBuster = nil
		} else {

			uj.IFrameBuster = []string{}

			wantVal := true

			for {

				var tmp_uj__IFrameBuster string

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

				/* handler: tmp_uj__IFrameBuster type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__IFrameBuster = string(string(outBuf))

					}
				}

				uj.IFrameBuster = append(uj.IFrameBuster, tmp_uj__IFrameBuster)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Pmp:

	/* handler: uj.Pmp type=openrtb.Pmp kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Pmp kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Pmp)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Exp:

	/* handler: uj.Exp type=int kind=int quoted=false*/

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

			uj.Exp = int(tval)

		}
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
