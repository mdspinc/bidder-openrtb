// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: content.go

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Content) MarshalJSON() ([]byte, error) {
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
func (j *Content) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(j.ID) != 0 {
		buf.WriteString(`"id":`)
		fflib.WriteJsonString(buf, string(j.ID))
		buf.WriteByte(',')
	}
	if j.Episode != 0 {
		buf.WriteString(`"episode":`)
		fflib.FormatBits2(buf, uint64(j.Episode), 10, j.Episode < 0)
		buf.WriteByte(',')
	}
	if len(j.Title) != 0 {
		buf.WriteString(`"title":`)
		fflib.WriteJsonString(buf, string(j.Title))
		buf.WriteByte(',')
	}
	if len(j.Series) != 0 {
		buf.WriteString(`"series":`)
		fflib.WriteJsonString(buf, string(j.Series))
		buf.WriteByte(',')
	}
	if len(j.Season) != 0 {
		buf.WriteString(`"season":`)
		fflib.WriteJsonString(buf, string(j.Season))
		buf.WriteByte(',')
	}
	if len(j.Artist) != 0 {
		buf.WriteString(`"artist":`)
		fflib.WriteJsonString(buf, string(j.Artist))
		buf.WriteByte(',')
	}
	if len(j.Genre) != 0 {
		buf.WriteString(`"genre":`)
		fflib.WriteJsonString(buf, string(j.Genre))
		buf.WriteByte(',')
	}
	if len(j.Album) != 0 {
		buf.WriteString(`"album":`)
		fflib.WriteJsonString(buf, string(j.Album))
		buf.WriteByte(',')
	}
	if len(j.ISRC) != 0 {
		buf.WriteString(`"isrc":`)
		fflib.WriteJsonString(buf, string(j.ISRC))
		buf.WriteByte(',')
	}
	if j.Producer != nil {
		if true {
			buf.WriteString(`"producer":`)

			{

				err = j.Producer.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if j.Prodq != 0 {
		buf.WriteString(`"prodq":`)
		fflib.FormatBits2(buf, uint64(j.Prodq), 10, j.Prodq < 0)
		buf.WriteByte(',')
	}
	if j.VideoQuality != 0 {
		buf.WriteString(`"videoquality":`)
		fflib.FormatBits2(buf, uint64(j.VideoQuality), 10, j.VideoQuality < 0)
		buf.WriteByte(',')
	}
	if j.Context != 0 {
		buf.WriteString(`"context":`)
		fflib.FormatBits2(buf, uint64(j.Context), 10, j.Context < 0)
		buf.WriteByte(',')
	}
	if len(j.ContentRating) != 0 {
		buf.WriteString(`"contentrating":`)
		fflib.WriteJsonString(buf, string(j.ContentRating))
		buf.WriteByte(',')
	}
	if len(j.UserRating) != 0 {
		buf.WriteString(`"userrating":`)
		fflib.WriteJsonString(buf, string(j.UserRating))
		buf.WriteByte(',')
	}
	if j.QAGMediaRating != 0 {
		buf.WriteString(`"qagmediarating":`)
		fflib.FormatBits2(buf, uint64(j.QAGMediaRating), 10, j.QAGMediaRating < 0)
		buf.WriteByte(',')
	}
	if len(j.Keywords) != 0 {
		buf.WriteString(`"keywords":`)
		fflib.WriteJsonString(buf, string(j.Keywords))
		buf.WriteByte(',')
	}
	if j.LiveStream != 0 {
		buf.WriteString(`"livestream":`)
		fflib.FormatBits2(buf, uint64(j.LiveStream), 10, j.LiveStream < 0)
		buf.WriteByte(',')
	}
	if j.SourceRelationship != 0 {
		buf.WriteString(`"sourcerelationship":`)
		fflib.FormatBits2(buf, uint64(j.SourceRelationship), 10, j.SourceRelationship < 0)
		buf.WriteByte(',')
	}
	if j.Len != 0 {
		buf.WriteString(`"len":`)
		fflib.FormatBits2(buf, uint64(j.Len), 10, j.Len < 0)
		buf.WriteByte(',')
	}
	if len(j.Language) != 0 {
		buf.WriteString(`"language":`)
		fflib.WriteJsonString(buf, string(j.Language))
		buf.WriteByte(',')
	}
	if j.Embeddable != 0 {
		buf.WriteString(`"embeddable":`)
		fflib.FormatBits2(buf, uint64(j.Embeddable), 10, j.Embeddable < 0)
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
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtContentbase = iota
	ffjtContentnosuchkey

	ffjtContentID

	ffjtContentEpisode

	ffjtContentTitle

	ffjtContentSeries

	ffjtContentSeason

	ffjtContentArtist

	ffjtContentGenre

	ffjtContentAlbum

	ffjtContentISRC

	ffjtContentProducer

	ffjtContentProdq

	ffjtContentVideoQuality

	ffjtContentContext

	ffjtContentContentRating

	ffjtContentUserRating

	ffjtContentQAGMediaRating

	ffjtContentKeywords

	ffjtContentLiveStream

	ffjtContentSourceRelationship

	ffjtContentLen

	ffjtContentLanguage

	ffjtContentEmbeddable

	ffjtContentExt
)

var ffjKeyContentID = []byte("id")

var ffjKeyContentEpisode = []byte("episode")

var ffjKeyContentTitle = []byte("title")

var ffjKeyContentSeries = []byte("series")

var ffjKeyContentSeason = []byte("season")

var ffjKeyContentArtist = []byte("artist")

var ffjKeyContentGenre = []byte("genre")

var ffjKeyContentAlbum = []byte("album")

var ffjKeyContentISRC = []byte("isrc")

var ffjKeyContentProducer = []byte("producer")

var ffjKeyContentProdq = []byte("prodq")

var ffjKeyContentVideoQuality = []byte("videoquality")

var ffjKeyContentContext = []byte("context")

var ffjKeyContentContentRating = []byte("contentrating")

var ffjKeyContentUserRating = []byte("userrating")

var ffjKeyContentQAGMediaRating = []byte("qagmediarating")

var ffjKeyContentKeywords = []byte("keywords")

var ffjKeyContentLiveStream = []byte("livestream")

var ffjKeyContentSourceRelationship = []byte("sourcerelationship")

var ffjKeyContentLen = []byte("len")

var ffjKeyContentLanguage = []byte("language")

var ffjKeyContentEmbeddable = []byte("embeddable")

var ffjKeyContentExt = []byte("ext")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Content) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Content) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtContentbase
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
				currentKey = ffjtContentnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyContentArtist, kn) {
						currentKey = ffjtContentArtist
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentAlbum, kn) {
						currentKey = ffjtContentAlbum
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyContentContext, kn) {
						currentKey = ffjtContentContext
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentContentRating, kn) {
						currentKey = ffjtContentContentRating
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyContentEpisode, kn) {
						currentKey = ffjtContentEpisode
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentEmbeddable, kn) {
						currentKey = ffjtContentEmbeddable
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentExt, kn) {
						currentKey = ffjtContentExt
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'g':

					if bytes.Equal(ffjKeyContentGenre, kn) {
						currentKey = ffjtContentGenre
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyContentID, kn) {
						currentKey = ffjtContentID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentISRC, kn) {
						currentKey = ffjtContentISRC
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'k':

					if bytes.Equal(ffjKeyContentKeywords, kn) {
						currentKey = ffjtContentKeywords
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyContentLiveStream, kn) {
						currentKey = ffjtContentLiveStream
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentLen, kn) {
						currentKey = ffjtContentLen
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentLanguage, kn) {
						currentKey = ffjtContentLanguage
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyContentProducer, kn) {
						currentKey = ffjtContentProducer
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentProdq, kn) {
						currentKey = ffjtContentProdq
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'q':

					if bytes.Equal(ffjKeyContentQAGMediaRating, kn) {
						currentKey = ffjtContentQAGMediaRating
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyContentSeries, kn) {
						currentKey = ffjtContentSeries
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentSeason, kn) {
						currentKey = ffjtContentSeason
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyContentSourceRelationship, kn) {
						currentKey = ffjtContentSourceRelationship
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyContentTitle, kn) {
						currentKey = ffjtContentTitle
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffjKeyContentUserRating, kn) {
						currentKey = ffjtContentUserRating
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffjKeyContentVideoQuality, kn) {
						currentKey = ffjtContentVideoQuality
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentExt, kn) {
					currentKey = ffjtContentExt
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentEmbeddable, kn) {
					currentKey = ffjtContentEmbeddable
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentLanguage, kn) {
					currentKey = ffjtContentLanguage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentLen, kn) {
					currentKey = ffjtContentLen
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyContentSourceRelationship, kn) {
					currentKey = ffjtContentSourceRelationship
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyContentLiveStream, kn) {
					currentKey = ffjtContentLiveStream
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyContentKeywords, kn) {
					currentKey = ffjtContentKeywords
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentQAGMediaRating, kn) {
					currentKey = ffjtContentQAGMediaRating
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyContentUserRating, kn) {
					currentKey = ffjtContentUserRating
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentContentRating, kn) {
					currentKey = ffjtContentContentRating
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentContext, kn) {
					currentKey = ffjtContentContext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentVideoQuality, kn) {
					currentKey = ffjtContentVideoQuality
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentProdq, kn) {
					currentKey = ffjtContentProdq
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentProducer, kn) {
					currentKey = ffjtContentProducer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyContentISRC, kn) {
					currentKey = ffjtContentISRC
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentAlbum, kn) {
					currentKey = ffjtContentAlbum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentGenre, kn) {
					currentKey = ffjtContentGenre
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyContentArtist, kn) {
					currentKey = ffjtContentArtist
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyContentSeason, kn) {
					currentKey = ffjtContentSeason
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyContentSeries, kn) {
					currentKey = ffjtContentSeries
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentTitle, kn) {
					currentKey = ffjtContentTitle
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyContentEpisode, kn) {
					currentKey = ffjtContentEpisode
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyContentID, kn) {
					currentKey = ffjtContentID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtContentnosuchkey
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

				case ffjtContentID:
					goto handle_ID

				case ffjtContentEpisode:
					goto handle_Episode

				case ffjtContentTitle:
					goto handle_Title

				case ffjtContentSeries:
					goto handle_Series

				case ffjtContentSeason:
					goto handle_Season

				case ffjtContentArtist:
					goto handle_Artist

				case ffjtContentGenre:
					goto handle_Genre

				case ffjtContentAlbum:
					goto handle_Album

				case ffjtContentISRC:
					goto handle_ISRC

				case ffjtContentProducer:
					goto handle_Producer

				case ffjtContentProdq:
					goto handle_Prodq

				case ffjtContentVideoQuality:
					goto handle_VideoQuality

				case ffjtContentContext:
					goto handle_Context

				case ffjtContentContentRating:
					goto handle_ContentRating

				case ffjtContentUserRating:
					goto handle_UserRating

				case ffjtContentQAGMediaRating:
					goto handle_QAGMediaRating

				case ffjtContentKeywords:
					goto handle_Keywords

				case ffjtContentLiveStream:
					goto handle_LiveStream

				case ffjtContentSourceRelationship:
					goto handle_SourceRelationship

				case ffjtContentLen:
					goto handle_Len

				case ffjtContentLanguage:
					goto handle_Language

				case ffjtContentEmbeddable:
					goto handle_Embeddable

				case ffjtContentExt:
					goto handle_Ext

				case ffjtContentnosuchkey:
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

	/* handler: j.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Episode:

	/* handler: j.Episode type=int kind=int quoted=false*/

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

			j.Episode = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Title:

	/* handler: j.Title type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Title = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Series:

	/* handler: j.Series type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Series = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Season:

	/* handler: j.Season type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Season = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Artist:

	/* handler: j.Artist type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Artist = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Genre:

	/* handler: j.Genre type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Genre = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Album:

	/* handler: j.Album type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Album = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ISRC:

	/* handler: j.ISRC type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.ISRC = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Producer:

	/* handler: j.Producer type=openrtb.Producer kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Producer = nil

		} else {

			if j.Producer == nil {
				j.Producer = new(Producer)
			}

			err = j.Producer.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Prodq:

	/* handler: j.Prodq type=int kind=int quoted=false*/

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

			j.Prodq = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VideoQuality:

	/* handler: j.VideoQuality type=int kind=int quoted=false*/

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

			j.VideoQuality = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Context:

	/* handler: j.Context type=int kind=int quoted=false*/

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

			j.Context = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ContentRating:

	/* handler: j.ContentRating type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.ContentRating = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_UserRating:

	/* handler: j.UserRating type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.UserRating = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_QAGMediaRating:

	/* handler: j.QAGMediaRating type=int kind=int quoted=false*/

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

			j.QAGMediaRating = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Keywords:

	/* handler: j.Keywords type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Keywords = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LiveStream:

	/* handler: j.LiveStream type=openrtb.NumberOrBool kind=int quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LiveStream.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SourceRelationship:

	/* handler: j.SourceRelationship type=openrtb.NumberOrBool kind=int quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.SourceRelationship.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Len:

	/* handler: j.Len type=int kind=int quoted=false*/

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

			j.Len = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Language:

	/* handler: j.Language type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Language = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Embeddable:

	/* handler: j.Embeddable type=openrtb.NumberOrBool kind=int quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Embeddable.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
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
