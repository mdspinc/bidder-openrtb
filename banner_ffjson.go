// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: banner.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Banner) MarshalJSON() ([]byte, error) {
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
func (mj *Banner) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
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
	if mj.WMax != 0 {
		buf.WriteString(`"wmax":`)
		fflib.FormatBits2(buf, uint64(mj.WMax), 10, mj.WMax < 0)
		buf.WriteByte(',')
	}
	if mj.HMax != 0 {
		buf.WriteString(`"hmax":`)
		fflib.FormatBits2(buf, uint64(mj.HMax), 10, mj.HMax < 0)
		buf.WriteByte(',')
	}
	if mj.WMin != 0 {
		buf.WriteString(`"wmin":`)
		fflib.FormatBits2(buf, uint64(mj.WMin), 10, mj.WMin < 0)
		buf.WriteByte(',')
	}
	if mj.HMin != 0 {
		buf.WriteString(`"hmin":`)
		fflib.FormatBits2(buf, uint64(mj.HMin), 10, mj.HMin < 0)
		buf.WriteByte(',')
	}
	if len(mj.ID) != 0 {
		buf.WriteString(`"id":`)
		fflib.WriteJsonString(buf, string(mj.ID))
		buf.WriteByte(',')
	}
	if mj.Pos != 0 {
		buf.WriteString(`"pos":`)
		fflib.FormatBits2(buf, uint64(mj.Pos), 10, mj.Pos < 0)
		buf.WriteByte(',')
	}
	if len(mj.BType) != 0 {
		buf.WriteString(`"btype":`)
		if mj.BType != nil {
			buf.WriteString(`[`)
			for i, v := range mj.BType {
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
	if mj.TopFrame != 0 {
		buf.WriteString(`"topframe":`)
		fflib.FormatBits2(buf, uint64(mj.TopFrame), 10, mj.TopFrame < 0)
		buf.WriteByte(',')
	}
	if len(mj.ExpDir) != 0 {
		buf.WriteString(`"expdir":`)
		if mj.ExpDir != nil {
			buf.WriteString(`[`)
			for i, v := range mj.ExpDir {
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
	ffj_t_Bannerbase = iota
	ffj_t_Bannerno_such_key

	ffj_t_Banner_W

	ffj_t_Banner_H

	ffj_t_Banner_WMax

	ffj_t_Banner_HMax

	ffj_t_Banner_WMin

	ffj_t_Banner_HMin

	ffj_t_Banner_ID

	ffj_t_Banner_Pos

	ffj_t_Banner_BType

	ffj_t_Banner_BAttr

	ffj_t_Banner_Mimes

	ffj_t_Banner_TopFrame

	ffj_t_Banner_ExpDir

	ffj_t_Banner_Api

	ffj_t_Banner_Ext
)

var ffj_key_Banner_W = []byte("w")

var ffj_key_Banner_H = []byte("h")

var ffj_key_Banner_WMax = []byte("wmax")

var ffj_key_Banner_HMax = []byte("hmax")

var ffj_key_Banner_WMin = []byte("wmin")

var ffj_key_Banner_HMin = []byte("hmin")

var ffj_key_Banner_ID = []byte("id")

var ffj_key_Banner_Pos = []byte("pos")

var ffj_key_Banner_BType = []byte("btype")

var ffj_key_Banner_BAttr = []byte("battr")

var ffj_key_Banner_Mimes = []byte("mimes")

var ffj_key_Banner_TopFrame = []byte("topframe")

var ffj_key_Banner_ExpDir = []byte("expdir")

var ffj_key_Banner_Api = []byte("api")

var ffj_key_Banner_Ext = []byte("ext")

func (uj *Banner) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Banner) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Bannerbase
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
				currentKey = ffj_t_Bannerno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_Banner_Api, kn) {
						currentKey = ffj_t_Banner_Api
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffj_key_Banner_BType, kn) {
						currentKey = ffj_t_Banner_BType
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Banner_BAttr, kn) {
						currentKey = ffj_t_Banner_BAttr
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Banner_ExpDir, kn) {
						currentKey = ffj_t_Banner_ExpDir
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Banner_Ext, kn) {
						currentKey = ffj_t_Banner_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_Banner_H, kn) {
						currentKey = ffj_t_Banner_H
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Banner_HMax, kn) {
						currentKey = ffj_t_Banner_HMax
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Banner_HMin, kn) {
						currentKey = ffj_t_Banner_HMin
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_Banner_ID, kn) {
						currentKey = ffj_t_Banner_ID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_Banner_Mimes, kn) {
						currentKey = ffj_t_Banner_Mimes
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Banner_Pos, kn) {
						currentKey = ffj_t_Banner_Pos
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Banner_TopFrame, kn) {
						currentKey = ffj_t_Banner_TopFrame
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_Banner_W, kn) {
						currentKey = ffj_t_Banner_W
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Banner_WMax, kn) {
						currentKey = ffj_t_Banner_WMax
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Banner_WMin, kn) {
						currentKey = ffj_t_Banner_WMin
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_Ext, kn) {
					currentKey = ffj_t_Banner_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_Api, kn) {
					currentKey = ffj_t_Banner_Api
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_ExpDir, kn) {
					currentKey = ffj_t_Banner_ExpDir
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_TopFrame, kn) {
					currentKey = ffj_t_Banner_TopFrame
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Banner_Mimes, kn) {
					currentKey = ffj_t_Banner_Mimes
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_BAttr, kn) {
					currentKey = ffj_t_Banner_BAttr
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_BType, kn) {
					currentKey = ffj_t_Banner_BType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Banner_Pos, kn) {
					currentKey = ffj_t_Banner_Pos
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_ID, kn) {
					currentKey = ffj_t_Banner_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_HMin, kn) {
					currentKey = ffj_t_Banner_HMin
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_WMin, kn) {
					currentKey = ffj_t_Banner_WMin
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_HMax, kn) {
					currentKey = ffj_t_Banner_HMax
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_WMax, kn) {
					currentKey = ffj_t_Banner_WMax
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_H, kn) {
					currentKey = ffj_t_Banner_H
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Banner_W, kn) {
					currentKey = ffj_t_Banner_W
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Bannerno_such_key
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

				case ffj_t_Banner_W:
					goto handle_W

				case ffj_t_Banner_H:
					goto handle_H

				case ffj_t_Banner_WMax:
					goto handle_WMax

				case ffj_t_Banner_HMax:
					goto handle_HMax

				case ffj_t_Banner_WMin:
					goto handle_WMin

				case ffj_t_Banner_HMin:
					goto handle_HMin

				case ffj_t_Banner_ID:
					goto handle_ID

				case ffj_t_Banner_Pos:
					goto handle_Pos

				case ffj_t_Banner_BType:
					goto handle_BType

				case ffj_t_Banner_BAttr:
					goto handle_BAttr

				case ffj_t_Banner_Mimes:
					goto handle_Mimes

				case ffj_t_Banner_TopFrame:
					goto handle_TopFrame

				case ffj_t_Banner_ExpDir:
					goto handle_ExpDir

				case ffj_t_Banner_Api:
					goto handle_Api

				case ffj_t_Banner_Ext:
					goto handle_Ext

				case ffj_t_Bannerno_such_key:
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

handle_WMax:

	/* handler: uj.WMax type=int kind=int quoted=false*/

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

			uj.WMax = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HMax:

	/* handler: uj.HMax type=int kind=int quoted=false*/

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

			uj.HMax = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WMin:

	/* handler: uj.WMin type=int kind=int quoted=false*/

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

			uj.WMin = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HMin:

	/* handler: uj.HMin type=int kind=int quoted=false*/

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

			uj.HMin = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

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

handle_BType:

	/* handler: uj.BType type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.BType = nil
		} else {

			uj.BType = []int{}

			wantVal := true

			for {

				var tmp_uj__BType int

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

				/* handler: tmp_uj__BType type=int kind=int quoted=false*/

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

						tmp_uj__BType = int(tval)

					}
				}

				uj.BType = append(uj.BType, tmp_uj__BType)

				wantVal = false
			}
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

handle_TopFrame:

	/* handler: uj.TopFrame type=openrtb.NumberOrBool kind=int quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = uj.TopFrame.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ExpDir:

	/* handler: uj.ExpDir type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.ExpDir = nil
		} else {

			uj.ExpDir = []int{}

			wantVal := true

			for {

				var tmp_uj__ExpDir int

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

				/* handler: tmp_uj__ExpDir type=int kind=int quoted=false*/

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

						tmp_uj__ExpDir = int(tval)

					}
				}

				uj.ExpDir = append(uj.ExpDir, tmp_uj__ExpDir)

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
