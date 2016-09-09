// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: content.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Content) MarshalJSON() ([]byte, error) {
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
func (mj *Content) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
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
	if mj.Episode != 0 {
		buf.WriteString(`"episode":`)
		fflib.FormatBits2(buf, uint64(mj.Episode), 10, mj.Episode < 0)
		buf.WriteByte(',')
	}
	if len(mj.Title) != 0 {
		buf.WriteString(`"title":`)
		fflib.WriteJsonString(buf, string(mj.Title))
		buf.WriteByte(',')
	}
	if len(mj.Series) != 0 {
		buf.WriteString(`"series":`)
		fflib.WriteJsonString(buf, string(mj.Series))
		buf.WriteByte(',')
	}
	if len(mj.Season) != 0 {
		buf.WriteString(`"season":`)
		fflib.WriteJsonString(buf, string(mj.Season))
		buf.WriteByte(',')
	}
	if mj.Producer != nil {
		if true {
			/* Struct fall back. type=openrtb.Producer kind=struct */
			buf.WriteString(`"producer":`)
			err = buf.Encode(mj.Producer)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if mj.VideoQuality != 0 {
		buf.WriteString(`"videoquality":`)
		fflib.FormatBits2(buf, uint64(mj.VideoQuality), 10, mj.VideoQuality < 0)
		buf.WriteByte(',')
	}
	if mj.Context != 0 {
		buf.WriteString(`"context":`)
		fflib.FormatBits2(buf, uint64(mj.Context), 10, mj.Context < 0)
		buf.WriteByte(',')
	}
	if len(mj.ContentRating) != 0 {
		buf.WriteString(`"contentrating":`)
		fflib.WriteJsonString(buf, string(mj.ContentRating))
		buf.WriteByte(',')
	}
	if len(mj.UserRating) != 0 {
		buf.WriteString(`"userrating":`)
		fflib.WriteJsonString(buf, string(mj.UserRating))
		buf.WriteByte(',')
	}
	if mj.QAGMediaRating != 0 {
		buf.WriteString(`"qagmediarating":`)
		fflib.FormatBits2(buf, uint64(mj.QAGMediaRating), 10, mj.QAGMediaRating < 0)
		buf.WriteByte(',')
	}
	if len(mj.Keywords) != 0 {
		buf.WriteString(`"keywords":`)
		fflib.WriteJsonString(buf, string(mj.Keywords))
		buf.WriteByte(',')
	}
	if mj.LiveStream != 0 {
		buf.WriteString(`"livestream":`)
		fflib.FormatBits2(buf, uint64(mj.LiveStream), 10, mj.LiveStream < 0)
		buf.WriteByte(',')
	}
	if mj.SourceRelationship != 0 {
		buf.WriteString(`"sourcerelationship":`)
		fflib.FormatBits2(buf, uint64(mj.SourceRelationship), 10, mj.SourceRelationship < 0)
		buf.WriteByte(',')
	}
	if mj.Len != 0 {
		buf.WriteString(`"len":`)
		fflib.FormatBits2(buf, uint64(mj.Len), 10, mj.Len < 0)
		buf.WriteByte(',')
	}
	if len(mj.Language) != 0 {
		buf.WriteString(`"language":`)
		fflib.WriteJsonString(buf, string(mj.Language))
		buf.WriteByte(',')
	}
	if mj.Embeddable != 0 {
		buf.WriteString(`"embeddable":`)
		fflib.FormatBits2(buf, uint64(mj.Embeddable), 10, mj.Embeddable < 0)
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
	ffj_t_Contentbase = iota
	ffj_t_Contentno_such_key

	ffj_t_Content_ID

	ffj_t_Content_Episode

	ffj_t_Content_Title

	ffj_t_Content_Series

	ffj_t_Content_Season

	ffj_t_Content_Producer

	ffj_t_Content_VideoQuality

	ffj_t_Content_Context

	ffj_t_Content_ContentRating

	ffj_t_Content_UserRating

	ffj_t_Content_QAGMediaRating

	ffj_t_Content_Keywords

	ffj_t_Content_LiveStream

	ffj_t_Content_SourceRelationship

	ffj_t_Content_Len

	ffj_t_Content_Language

	ffj_t_Content_Embeddable

	ffj_t_Content_Ext
)

var ffj_key_Content_ID = []byte("id")

var ffj_key_Content_Episode = []byte("episode")

var ffj_key_Content_Title = []byte("title")

var ffj_key_Content_Series = []byte("series")

var ffj_key_Content_Season = []byte("season")

var ffj_key_Content_Producer = []byte("producer")

var ffj_key_Content_VideoQuality = []byte("videoquality")

var ffj_key_Content_Context = []byte("context")

var ffj_key_Content_ContentRating = []byte("contentrating")

var ffj_key_Content_UserRating = []byte("userrating")

var ffj_key_Content_QAGMediaRating = []byte("qagmediarating")

var ffj_key_Content_Keywords = []byte("keywords")

var ffj_key_Content_LiveStream = []byte("livestream")

var ffj_key_Content_SourceRelationship = []byte("sourcerelationship")

var ffj_key_Content_Len = []byte("len")

var ffj_key_Content_Language = []byte("language")

var ffj_key_Content_Embeddable = []byte("embeddable")

var ffj_key_Content_Ext = []byte("ext")

func (uj *Content) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Content) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Contentbase
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
				currentKey = ffj_t_Contentno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffj_key_Content_Context, kn) {
						currentKey = ffj_t_Content_Context
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Content_ContentRating, kn) {
						currentKey = ffj_t_Content_ContentRating
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Content_Episode, kn) {
						currentKey = ffj_t_Content_Episode
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Content_Embeddable, kn) {
						currentKey = ffj_t_Content_Embeddable
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Content_Ext, kn) {
						currentKey = ffj_t_Content_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_Content_ID, kn) {
						currentKey = ffj_t_Content_ID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'k':

					if bytes.Equal(ffj_key_Content_Keywords, kn) {
						currentKey = ffj_t_Content_Keywords
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_Content_LiveStream, kn) {
						currentKey = ffj_t_Content_LiveStream
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Content_Len, kn) {
						currentKey = ffj_t_Content_Len
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Content_Language, kn) {
						currentKey = ffj_t_Content_Language
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Content_Producer, kn) {
						currentKey = ffj_t_Content_Producer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'q':

					if bytes.Equal(ffj_key_Content_QAGMediaRating, kn) {
						currentKey = ffj_t_Content_QAGMediaRating
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_Content_Series, kn) {
						currentKey = ffj_t_Content_Series
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Content_Season, kn) {
						currentKey = ffj_t_Content_Season
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Content_SourceRelationship, kn) {
						currentKey = ffj_t_Content_SourceRelationship
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Content_Title, kn) {
						currentKey = ffj_t_Content_Title
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffj_key_Content_UserRating, kn) {
						currentKey = ffj_t_Content_UserRating
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffj_key_Content_VideoQuality, kn) {
						currentKey = ffj_t_Content_VideoQuality
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_Ext, kn) {
					currentKey = ffj_t_Content_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_Embeddable, kn) {
					currentKey = ffj_t_Content_Embeddable
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_Language, kn) {
					currentKey = ffj_t_Content_Language
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_Len, kn) {
					currentKey = ffj_t_Content_Len
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Content_SourceRelationship, kn) {
					currentKey = ffj_t_Content_SourceRelationship
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Content_LiveStream, kn) {
					currentKey = ffj_t_Content_LiveStream
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Content_Keywords, kn) {
					currentKey = ffj_t_Content_Keywords
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_QAGMediaRating, kn) {
					currentKey = ffj_t_Content_QAGMediaRating
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Content_UserRating, kn) {
					currentKey = ffj_t_Content_UserRating
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_ContentRating, kn) {
					currentKey = ffj_t_Content_ContentRating
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_Context, kn) {
					currentKey = ffj_t_Content_Context
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_VideoQuality, kn) {
					currentKey = ffj_t_Content_VideoQuality
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_Producer, kn) {
					currentKey = ffj_t_Content_Producer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Content_Season, kn) {
					currentKey = ffj_t_Content_Season
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Content_Series, kn) {
					currentKey = ffj_t_Content_Series
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_Title, kn) {
					currentKey = ffj_t_Content_Title
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Content_Episode, kn) {
					currentKey = ffj_t_Content_Episode
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Content_ID, kn) {
					currentKey = ffj_t_Content_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Contentno_such_key
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

				case ffj_t_Content_ID:
					goto handle_ID

				case ffj_t_Content_Episode:
					goto handle_Episode

				case ffj_t_Content_Title:
					goto handle_Title

				case ffj_t_Content_Series:
					goto handle_Series

				case ffj_t_Content_Season:
					goto handle_Season

				case ffj_t_Content_Producer:
					goto handle_Producer

				case ffj_t_Content_VideoQuality:
					goto handle_VideoQuality

				case ffj_t_Content_Context:
					goto handle_Context

				case ffj_t_Content_ContentRating:
					goto handle_ContentRating

				case ffj_t_Content_UserRating:
					goto handle_UserRating

				case ffj_t_Content_QAGMediaRating:
					goto handle_QAGMediaRating

				case ffj_t_Content_Keywords:
					goto handle_Keywords

				case ffj_t_Content_LiveStream:
					goto handle_LiveStream

				case ffj_t_Content_SourceRelationship:
					goto handle_SourceRelationship

				case ffj_t_Content_Len:
					goto handle_Len

				case ffj_t_Content_Language:
					goto handle_Language

				case ffj_t_Content_Embeddable:
					goto handle_Embeddable

				case ffj_t_Content_Ext:
					goto handle_Ext

				case ffj_t_Contentno_such_key:
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

handle_Episode:

	/* handler: uj.Episode type=int kind=int quoted=false*/

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

			uj.Episode = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Title:

	/* handler: uj.Title type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Title = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Series:

	/* handler: uj.Series type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Series = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Season:

	/* handler: uj.Season type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Season = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Producer:

	/* handler: uj.Producer type=openrtb.Producer kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Producer kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Producer)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VideoQuality:

	/* handler: uj.VideoQuality type=int kind=int quoted=false*/

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

			uj.VideoQuality = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Context:

	/* handler: uj.Context type=int kind=int quoted=false*/

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

			uj.Context = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ContentRating:

	/* handler: uj.ContentRating type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ContentRating = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_UserRating:

	/* handler: uj.UserRating type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.UserRating = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_QAGMediaRating:

	/* handler: uj.QAGMediaRating type=int kind=int quoted=false*/

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

			uj.QAGMediaRating = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Keywords:

	/* handler: uj.Keywords type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Keywords = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LiveStream:

	/* handler: uj.LiveStream type=int kind=int quoted=false*/

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

			uj.LiveStream = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SourceRelationship:

	/* handler: uj.SourceRelationship type=int kind=int quoted=false*/

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

			uj.SourceRelationship = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Len:

	/* handler: uj.Len type=int kind=int quoted=false*/

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

			uj.Len = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Language:

	/* handler: uj.Language type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Language = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Embeddable:

	/* handler: uj.Embeddable type=int kind=int quoted=false*/

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

			uj.Embeddable = int(tval)

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
