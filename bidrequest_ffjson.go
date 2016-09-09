// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: bidrequest.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *BidRequest) MarshalJSON() ([]byte, error) {
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
func (mj *BidRequest) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
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
	if len(mj.Imp) != 0 {
		buf.WriteString(`"imp":`)
		if mj.Imp != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Imp {
				if i != 0 {
					buf.WriteString(`,`)
				}
				/* Struct fall back. type=openrtb.Impression kind=struct */
				err = buf.Encode(&v)
				if err != nil {
					return err
				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Site != nil {
		if true {
			/* Struct fall back. type=openrtb.Site kind=struct */
			buf.WriteString(`"site":`)
			err = buf.Encode(mj.Site)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if mj.App != nil {
		if true {
			/* Struct fall back. type=openrtb.App kind=struct */
			buf.WriteString(`"app":`)
			err = buf.Encode(mj.App)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if mj.Device != nil {
		if true {
			/* Struct fall back. type=openrtb.Device kind=struct */
			buf.WriteString(`"device":`)
			err = buf.Encode(mj.Device)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if mj.User != nil {
		if true {
			/* Struct fall back. type=openrtb.User kind=struct */
			buf.WriteString(`"user":`)
			err = buf.Encode(mj.User)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.WriteString(`"at":`)
	fflib.FormatBits2(buf, uint64(mj.AuctionType), 10, mj.AuctionType < 0)
	buf.WriteByte(',')
	if mj.TMax != 0 {
		buf.WriteString(`"tmax":`)
		fflib.FormatBits2(buf, uint64(mj.TMax), 10, mj.TMax < 0)
		buf.WriteByte(',')
	}
	if len(mj.WSeat) != 0 {
		buf.WriteString(`"wseat":`)
		if mj.WSeat != nil {
			buf.WriteString(`[`)
			for i, v := range mj.WSeat {
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
	if mj.AllImps != 0 {
		buf.WriteString(`"allimps":`)
		fflib.FormatBits2(buf, uint64(mj.AllImps), 10, mj.AllImps < 0)
		buf.WriteByte(',')
	}
	if len(mj.Cur) != 0 {
		buf.WriteString(`"cur":`)
		if mj.Cur != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Cur {
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
	if len(mj.Bcat) != 0 {
		buf.WriteString(`"bcat":`)
		if mj.Bcat != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Bcat {
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
	if len(mj.BAdv) != 0 {
		buf.WriteString(`"badv":`)
		if mj.BAdv != nil {
			buf.WriteString(`[`)
			for i, v := range mj.BAdv {
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
	if mj.Regs != nil {
		if true {
			/* Struct fall back. type=openrtb.Regulations kind=struct */
			buf.WriteString(`"regs":`)
			err = buf.Encode(mj.Regs)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
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
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_BidRequestbase = iota
	ffj_t_BidRequestno_such_key

	ffj_t_BidRequest_ID

	ffj_t_BidRequest_Imp

	ffj_t_BidRequest_Site

	ffj_t_BidRequest_App

	ffj_t_BidRequest_Device

	ffj_t_BidRequest_User

	ffj_t_BidRequest_AuctionType

	ffj_t_BidRequest_TMax

	ffj_t_BidRequest_WSeat

	ffj_t_BidRequest_AllImps

	ffj_t_BidRequest_Cur

	ffj_t_BidRequest_Bcat

	ffj_t_BidRequest_BAdv

	ffj_t_BidRequest_Regs

	ffj_t_BidRequest_Ext

	ffj_t_BidRequest_Pmp
)

var ffj_key_BidRequest_ID = []byte("id")

var ffj_key_BidRequest_Imp = []byte("imp")

var ffj_key_BidRequest_Site = []byte("site")

var ffj_key_BidRequest_App = []byte("app")

var ffj_key_BidRequest_Device = []byte("device")

var ffj_key_BidRequest_User = []byte("user")

var ffj_key_BidRequest_AuctionType = []byte("at")

var ffj_key_BidRequest_TMax = []byte("tmax")

var ffj_key_BidRequest_WSeat = []byte("wseat")

var ffj_key_BidRequest_AllImps = []byte("allimps")

var ffj_key_BidRequest_Cur = []byte("cur")

var ffj_key_BidRequest_Bcat = []byte("bcat")

var ffj_key_BidRequest_BAdv = []byte("badv")

var ffj_key_BidRequest_Regs = []byte("regs")

var ffj_key_BidRequest_Ext = []byte("ext")

var ffj_key_BidRequest_Pmp = []byte("pmp")

func (uj *BidRequest) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *BidRequest) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_BidRequestbase
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
				currentKey = ffj_t_BidRequestno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_BidRequest_App, kn) {
						currentKey = ffj_t_BidRequest_App
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_BidRequest_AuctionType, kn) {
						currentKey = ffj_t_BidRequest_AuctionType
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_BidRequest_AllImps, kn) {
						currentKey = ffj_t_BidRequest_AllImps
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffj_key_BidRequest_Bcat, kn) {
						currentKey = ffj_t_BidRequest_Bcat
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_BidRequest_BAdv, kn) {
						currentKey = ffj_t_BidRequest_BAdv
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffj_key_BidRequest_Cur, kn) {
						currentKey = ffj_t_BidRequest_Cur
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_BidRequest_Device, kn) {
						currentKey = ffj_t_BidRequest_Device
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_BidRequest_Ext, kn) {
						currentKey = ffj_t_BidRequest_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_BidRequest_ID, kn) {
						currentKey = ffj_t_BidRequest_ID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_BidRequest_Imp, kn) {
						currentKey = ffj_t_BidRequest_Imp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_BidRequest_Pmp, kn) {
						currentKey = ffj_t_BidRequest_Pmp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffj_key_BidRequest_Regs, kn) {
						currentKey = ffj_t_BidRequest_Regs
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_BidRequest_Site, kn) {
						currentKey = ffj_t_BidRequest_Site
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_BidRequest_TMax, kn) {
						currentKey = ffj_t_BidRequest_TMax
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffj_key_BidRequest_User, kn) {
						currentKey = ffj_t_BidRequest_User
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_BidRequest_WSeat, kn) {
						currentKey = ffj_t_BidRequest_WSeat
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_Pmp, kn) {
					currentKey = ffj_t_BidRequest_Pmp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_Ext, kn) {
					currentKey = ffj_t_BidRequest_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_BidRequest_Regs, kn) {
					currentKey = ffj_t_BidRequest_Regs
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_BAdv, kn) {
					currentKey = ffj_t_BidRequest_BAdv
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_Bcat, kn) {
					currentKey = ffj_t_BidRequest_Bcat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_Cur, kn) {
					currentKey = ffj_t_BidRequest_Cur
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_BidRequest_AllImps, kn) {
					currentKey = ffj_t_BidRequest_AllImps
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_BidRequest_WSeat, kn) {
					currentKey = ffj_t_BidRequest_WSeat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_TMax, kn) {
					currentKey = ffj_t_BidRequest_TMax
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_AuctionType, kn) {
					currentKey = ffj_t_BidRequest_AuctionType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_BidRequest_User, kn) {
					currentKey = ffj_t_BidRequest_User
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_Device, kn) {
					currentKey = ffj_t_BidRequest_Device
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_App, kn) {
					currentKey = ffj_t_BidRequest_App
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_BidRequest_Site, kn) {
					currentKey = ffj_t_BidRequest_Site
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_Imp, kn) {
					currentKey = ffj_t_BidRequest_Imp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_BidRequest_ID, kn) {
					currentKey = ffj_t_BidRequest_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_BidRequestno_such_key
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

				case ffj_t_BidRequest_ID:
					goto handle_ID

				case ffj_t_BidRequest_Imp:
					goto handle_Imp

				case ffj_t_BidRequest_Site:
					goto handle_Site

				case ffj_t_BidRequest_App:
					goto handle_App

				case ffj_t_BidRequest_Device:
					goto handle_Device

				case ffj_t_BidRequest_User:
					goto handle_User

				case ffj_t_BidRequest_AuctionType:
					goto handle_AuctionType

				case ffj_t_BidRequest_TMax:
					goto handle_TMax

				case ffj_t_BidRequest_WSeat:
					goto handle_WSeat

				case ffj_t_BidRequest_AllImps:
					goto handle_AllImps

				case ffj_t_BidRequest_Cur:
					goto handle_Cur

				case ffj_t_BidRequest_Bcat:
					goto handle_Bcat

				case ffj_t_BidRequest_BAdv:
					goto handle_BAdv

				case ffj_t_BidRequest_Regs:
					goto handle_Regs

				case ffj_t_BidRequest_Ext:
					goto handle_Ext

				case ffj_t_BidRequest_Pmp:
					goto handle_Pmp

				case ffj_t_BidRequestno_such_key:
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

handle_Imp:

	/* handler: uj.Imp type=[]openrtb.Impression kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Imp = nil
		} else {

			uj.Imp = []Impression{}

			wantVal := true

			for {

				var tmp_uj__Imp Impression

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

				/* handler: tmp_uj__Imp type=openrtb.Impression kind=struct quoted=false*/

				{
					/* Falling back. type=openrtb.Impression kind=struct */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmp_uj__Imp)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				uj.Imp = append(uj.Imp, tmp_uj__Imp)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Site:

	/* handler: uj.Site type=openrtb.Site kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Site kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Site)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_App:

	/* handler: uj.App type=openrtb.App kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.App kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.App)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Device:

	/* handler: uj.Device type=openrtb.Device kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Device kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Device)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_User:

	/* handler: uj.User type=openrtb.User kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.User kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.User)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AuctionType:

	/* handler: uj.AuctionType type=int kind=int quoted=false*/

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

			uj.AuctionType = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TMax:

	/* handler: uj.TMax type=int kind=int quoted=false*/

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

			uj.TMax = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WSeat:

	/* handler: uj.WSeat type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.WSeat = nil
		} else {

			uj.WSeat = []string{}

			wantVal := true

			for {

				var tmp_uj__WSeat string

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

				/* handler: tmp_uj__WSeat type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__WSeat = string(string(outBuf))

					}
				}

				uj.WSeat = append(uj.WSeat, tmp_uj__WSeat)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AllImps:

	/* handler: uj.AllImps type=int kind=int quoted=false*/

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

			uj.AllImps = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Cur:

	/* handler: uj.Cur type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Cur = nil
		} else {

			uj.Cur = []string{}

			wantVal := true

			for {

				var tmp_uj__Cur string

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

				/* handler: tmp_uj__Cur type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__Cur = string(string(outBuf))

					}
				}

				uj.Cur = append(uj.Cur, tmp_uj__Cur)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Bcat:

	/* handler: uj.Bcat type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Bcat = nil
		} else {

			uj.Bcat = []string{}

			wantVal := true

			for {

				var tmp_uj__Bcat string

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

				/* handler: tmp_uj__Bcat type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__Bcat = string(string(outBuf))

					}
				}

				uj.Bcat = append(uj.Bcat, tmp_uj__Bcat)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BAdv:

	/* handler: uj.BAdv type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.BAdv = nil
		} else {

			uj.BAdv = []string{}

			wantVal := true

			for {

				var tmp_uj__BAdv string

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

				/* handler: tmp_uj__BAdv type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__BAdv = string(string(outBuf))

					}
				}

				uj.BAdv = append(uj.BAdv, tmp_uj__BAdv)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Regs:

	/* handler: uj.Regs type=openrtb.Regulations kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Regulations kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Regs)
		if err != nil {
			return fs.WrapErr(err)
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
