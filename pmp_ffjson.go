// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: pmp.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Pmp) MarshalJSON() ([]byte, error) {
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
func (mj *Pmp) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.Private != 0 {
		buf.WriteString(`"private_auction":`)
		fflib.FormatBits2(buf, uint64(mj.Private), 10, mj.Private < 0)
		buf.WriteByte(',')
	}
	if len(mj.Deals) != 0 {
		buf.WriteString(`"deals":`)
		if mj.Deals != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Deals {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					obj, err = v.MarshalJSON()
					if err != nil {
						return err
					}
					buf.Write(obj)

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
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
	ffj_t_Pmpbase = iota
	ffj_t_Pmpno_such_key

	ffj_t_Pmp_Private

	ffj_t_Pmp_Deals

	ffj_t_Pmp_Ext
)

var ffj_key_Pmp_Private = []byte("private_auction")

var ffj_key_Pmp_Deals = []byte("deals")

var ffj_key_Pmp_Ext = []byte("ext")

func (uj *Pmp) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Pmp) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Pmpbase
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
				currentKey = ffj_t_Pmpno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffj_key_Pmp_Deals, kn) {
						currentKey = ffj_t_Pmp_Deals
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Pmp_Ext, kn) {
						currentKey = ffj_t_Pmp_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Pmp_Private, kn) {
						currentKey = ffj_t_Pmp_Private
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Pmp_Ext, kn) {
					currentKey = ffj_t_Pmp_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Pmp_Deals, kn) {
					currentKey = ffj_t_Pmp_Deals
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_Pmp_Private, kn) {
					currentKey = ffj_t_Pmp_Private
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Pmpno_such_key
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

				case ffj_t_Pmp_Private:
					goto handle_Private

				case ffj_t_Pmp_Deals:
					goto handle_Deals

				case ffj_t_Pmp_Ext:
					goto handle_Ext

				case ffj_t_Pmpno_such_key:
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

handle_Private:

	/* handler: uj.Private type=openrtb.NumberOrBool kind=int quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = uj.Private.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Deals:

	/* handler: uj.Deals type=[]openrtb.Deal kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Deals = nil
		} else {

			uj.Deals = []Deal{}

			wantVal := true

			for {

				var tmp_uj__Deals Deal

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

				/* handler: tmp_uj__Deals type=openrtb.Deal kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = tmp_uj__Deals.UnmarshalJSON(tbuf)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
				}

				uj.Deals = append(uj.Deals, tmp_uj__Deals)

				wantVal = false
			}
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

func (mj *jsonDeal) MarshalJSON() ([]byte, error) {
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
func (mj *jsonDeal) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.ID) != 0 {
		buf.WriteString(`"id":`)
		fflib.WriteJsonString(buf, string(mj.ID))
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
	if len(mj.WAdvDomain) != 0 {
		buf.WriteString(`"wadomain":`)
		if mj.WAdvDomain != nil {
			buf.WriteString(`[`)
			for i, v := range mj.WAdvDomain {
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
	if mj.AuctionType != 0 {
		buf.WriteString(`"at":`)
		fflib.FormatBits2(buf, uint64(mj.AuctionType), 10, mj.AuctionType < 0)
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
	if len(mj.Seats) != 0 {
		buf.WriteString(`"seats":`)
		if mj.Seats != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Seats {
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
	if mj.Type != 0 {
		buf.WriteString(`"type":`)
		fflib.FormatBits2(buf, uint64(mj.Type), 10, mj.Type < 0)
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_jsonDealbase = iota
	ffj_t_jsonDealno_such_key

	ffj_t_jsonDeal_ID

	ffj_t_jsonDeal_BidFloor

	ffj_t_jsonDeal_BidFloorCurrency

	ffj_t_jsonDeal_WSeat

	ffj_t_jsonDeal_WAdvDomain

	ffj_t_jsonDeal_AuctionType

	ffj_t_jsonDeal_Ext

	ffj_t_jsonDeal_Seats

	ffj_t_jsonDeal_Type
)

var ffj_key_jsonDeal_ID = []byte("id")

var ffj_key_jsonDeal_BidFloor = []byte("bidfloor")

var ffj_key_jsonDeal_BidFloorCurrency = []byte("bidfloorcur")

var ffj_key_jsonDeal_WSeat = []byte("wseat")

var ffj_key_jsonDeal_WAdvDomain = []byte("wadomain")

var ffj_key_jsonDeal_AuctionType = []byte("at")

var ffj_key_jsonDeal_Ext = []byte("ext")

var ffj_key_jsonDeal_Seats = []byte("seats")

var ffj_key_jsonDeal_Type = []byte("type")

func (uj *jsonDeal) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *jsonDeal) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_jsonDealbase
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
				currentKey = ffj_t_jsonDealno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_jsonDeal_AuctionType, kn) {
						currentKey = ffj_t_jsonDeal_AuctionType
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffj_key_jsonDeal_BidFloor, kn) {
						currentKey = ffj_t_jsonDeal_BidFloor
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonDeal_BidFloorCurrency, kn) {
						currentKey = ffj_t_jsonDeal_BidFloorCurrency
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_jsonDeal_Ext, kn) {
						currentKey = ffj_t_jsonDeal_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_jsonDeal_ID, kn) {
						currentKey = ffj_t_jsonDeal_ID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_jsonDeal_Seats, kn) {
						currentKey = ffj_t_jsonDeal_Seats
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_jsonDeal_Type, kn) {
						currentKey = ffj_t_jsonDeal_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_jsonDeal_WSeat, kn) {
						currentKey = ffj_t_jsonDeal_WSeat
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_jsonDeal_WAdvDomain, kn) {
						currentKey = ffj_t_jsonDeal_WAdvDomain
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonDeal_Type, kn) {
					currentKey = ffj_t_jsonDeal_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_jsonDeal_Seats, kn) {
					currentKey = ffj_t_jsonDeal_Seats
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonDeal_Ext, kn) {
					currentKey = ffj_t_jsonDeal_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonDeal_AuctionType, kn) {
					currentKey = ffj_t_jsonDeal_AuctionType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonDeal_WAdvDomain, kn) {
					currentKey = ffj_t_jsonDeal_WAdvDomain
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_jsonDeal_WSeat, kn) {
					currentKey = ffj_t_jsonDeal_WSeat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonDeal_BidFloorCurrency, kn) {
					currentKey = ffj_t_jsonDeal_BidFloorCurrency
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonDeal_BidFloor, kn) {
					currentKey = ffj_t_jsonDeal_BidFloor
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_jsonDeal_ID, kn) {
					currentKey = ffj_t_jsonDeal_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_jsonDealno_such_key
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

				case ffj_t_jsonDeal_ID:
					goto handle_ID

				case ffj_t_jsonDeal_BidFloor:
					goto handle_BidFloor

				case ffj_t_jsonDeal_BidFloorCurrency:
					goto handle_BidFloorCurrency

				case ffj_t_jsonDeal_WSeat:
					goto handle_WSeat

				case ffj_t_jsonDeal_WAdvDomain:
					goto handle_WAdvDomain

				case ffj_t_jsonDeal_AuctionType:
					goto handle_AuctionType

				case ffj_t_jsonDeal_Ext:
					goto handle_Ext

				case ffj_t_jsonDeal_Seats:
					goto handle_Seats

				case ffj_t_jsonDeal_Type:
					goto handle_Type

				case ffj_t_jsonDealno_such_key:
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

handle_WAdvDomain:

	/* handler: uj.WAdvDomain type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.WAdvDomain = nil
		} else {

			uj.WAdvDomain = []string{}

			wantVal := true

			for {

				var tmp_uj__WAdvDomain string

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

				/* handler: tmp_uj__WAdvDomain type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__WAdvDomain = string(string(outBuf))

					}
				}

				uj.WAdvDomain = append(uj.WAdvDomain, tmp_uj__WAdvDomain)

				wantVal = false
			}
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

handle_Seats:

	/* handler: uj.Seats type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Seats = nil
		} else {

			uj.Seats = []string{}

			wantVal := true

			for {

				var tmp_uj__Seats string

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

				/* handler: tmp_uj__Seats type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__Seats = string(string(outBuf))

					}
				}

				uj.Seats = append(uj.Seats, tmp_uj__Seats)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Type:

	/* handler: uj.Type type=int kind=int quoted=false*/

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

			uj.Type = int(tval)

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
