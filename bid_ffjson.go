// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: bid.go
// DO NOT EDIT!

package openrtb

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Bid) MarshalJSON() ([]byte, error) {
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
func (mj *Bid) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
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
	buf.WriteString(`,"impid":`)
	fflib.WriteJsonString(buf, string(mj.ImpID))
	buf.WriteString(`,"price":`)
	fflib.AppendFloat(buf, float64(mj.Price), 'g', -1, 64)
	buf.WriteByte(',')
	if len(mj.AdID) != 0 {
		buf.WriteString(`"adid":`)
		fflib.WriteJsonString(buf, string(mj.AdID))
		buf.WriteByte(',')
	}
	if len(mj.NURL) != 0 {
		buf.WriteString(`"nurl":`)
		fflib.WriteJsonString(buf, string(mj.NURL))
		buf.WriteByte(',')
	}
	if len(mj.BURL) != 0 {
		buf.WriteString(`"burl":`)
		fflib.WriteJsonString(buf, string(mj.BURL))
		buf.WriteByte(',')
	}
	if len(mj.LURL) != 0 {
		buf.WriteString(`"lurl":`)
		fflib.WriteJsonString(buf, string(mj.LURL))
		buf.WriteByte(',')
	}
	if len(mj.AdMarkup) != 0 {
		buf.WriteString(`"adm":`)
		fflib.WriteJsonString(buf, string(mj.AdMarkup))
		buf.WriteByte(',')
	}
	if len(mj.AdvDomain) != 0 {
		buf.WriteString(`"adomain":`)
		if mj.AdvDomain != nil {
			buf.WriteString(`[`)
			for i, v := range mj.AdvDomain {
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
	if len(mj.IURL) != 0 {
		buf.WriteString(`"iurl":`)
		fflib.WriteJsonString(buf, string(mj.IURL))
		buf.WriteByte(',')
	}
	if len(mj.Tactic) != 0 {
		buf.WriteString(`"tactic":`)
		fflib.WriteJsonString(buf, string(mj.Tactic))
		buf.WriteByte(',')
	}
	if len(mj.Bundle) != 0 {
		buf.WriteString(`"bundle":`)
		fflib.WriteJsonString(buf, string(mj.Bundle))
		buf.WriteByte(',')
	}
	if len(mj.CampaignID) != 0 {
		buf.WriteString(`"cid":`)
		fflib.WriteJsonString(buf, string(mj.CampaignID))
		buf.WriteByte(',')
	}
	if len(mj.CreativeID) != 0 {
		buf.WriteString(`"crid":`)
		fflib.WriteJsonString(buf, string(mj.CreativeID))
		buf.WriteByte(',')
	}
	if mj.Qagmediarating != 0 {
		buf.WriteString(`"qagmediarating":`)
		fflib.FormatBits2(buf, uint64(mj.Qagmediarating), 10, mj.Qagmediarating < 0)
		buf.WriteByte(',')
	}
	if len(mj.Language) != 0 {
		buf.WriteString(`"language":`)
		fflib.WriteJsonString(buf, string(mj.Language))
		buf.WriteByte(',')
	}
	if len(mj.Cat) != 0 {
		buf.WriteString(`"cat":`)
		if mj.Cat != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Cat {
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
	if len(mj.Attr) != 0 {
		buf.WriteString(`"attr":`)
		if mj.Attr != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Attr {
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
	if len(mj.API) != 0 {
		buf.WriteString(`"api":`)
		fflib.WriteJsonString(buf, string(mj.API))
		buf.WriteByte(',')
	}
	if mj.Protocol != 0 {
		buf.WriteString(`"protocol":`)
		fflib.FormatBits2(buf, uint64(mj.Protocol), 10, mj.Protocol < 0)
		buf.WriteByte(',')
	}
	if len(mj.DealID) != 0 {
		buf.WriteString(`"dealid":`)
		fflib.WriteJsonString(buf, string(mj.DealID))
		buf.WriteByte(',')
	}
	if mj.H != 0 {
		buf.WriteString(`"h":`)
		fflib.FormatBits2(buf, uint64(mj.H), 10, mj.H < 0)
		buf.WriteByte(',')
	}
	if mj.W != 0 {
		buf.WriteString(`"w":`)
		fflib.FormatBits2(buf, uint64(mj.W), 10, mj.W < 0)
		buf.WriteByte(',')
	}
	if mj.Wratio != 0 {
		buf.WriteString(`"wratio":`)
		fflib.FormatBits2(buf, uint64(mj.Wratio), 10, mj.Wratio < 0)
		buf.WriteByte(',')
	}
	if mj.Hratio != 0 {
		buf.WriteString(`"hratio":`)
		fflib.FormatBits2(buf, uint64(mj.Hratio), 10, mj.Hratio < 0)
		buf.WriteByte(',')
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
	ffj_t_Bidbase = iota
	ffj_t_Bidno_such_key

	ffj_t_Bid_ID

	ffj_t_Bid_ImpID

	ffj_t_Bid_Price

	ffj_t_Bid_AdID

	ffj_t_Bid_NURL

	ffj_t_Bid_BURL

	ffj_t_Bid_LURL

	ffj_t_Bid_AdMarkup

	ffj_t_Bid_AdvDomain

	ffj_t_Bid_IURL

	ffj_t_Bid_Tactic

	ffj_t_Bid_Bundle

	ffj_t_Bid_CampaignID

	ffj_t_Bid_CreativeID

	ffj_t_Bid_Qagmediarating

	ffj_t_Bid_Language

	ffj_t_Bid_Cat

	ffj_t_Bid_Attr

	ffj_t_Bid_API

	ffj_t_Bid_Protocol

	ffj_t_Bid_DealID

	ffj_t_Bid_H

	ffj_t_Bid_W

	ffj_t_Bid_Wratio

	ffj_t_Bid_Hratio

	ffj_t_Bid_Exp

	ffj_t_Bid_Ext
)

var ffj_key_Bid_ID = []byte("id")

var ffj_key_Bid_ImpID = []byte("impid")

var ffj_key_Bid_Price = []byte("price")

var ffj_key_Bid_AdID = []byte("adid")

var ffj_key_Bid_NURL = []byte("nurl")

var ffj_key_Bid_BURL = []byte("burl")

var ffj_key_Bid_LURL = []byte("lurl")

var ffj_key_Bid_AdMarkup = []byte("adm")

var ffj_key_Bid_AdvDomain = []byte("adomain")

var ffj_key_Bid_IURL = []byte("iurl")

var ffj_key_Bid_Tactic = []byte("tactic")

var ffj_key_Bid_Bundle = []byte("bundle")

var ffj_key_Bid_CampaignID = []byte("cid")

var ffj_key_Bid_CreativeID = []byte("crid")

var ffj_key_Bid_Qagmediarating = []byte("qagmediarating")

var ffj_key_Bid_Language = []byte("language")

var ffj_key_Bid_Cat = []byte("cat")

var ffj_key_Bid_Attr = []byte("attr")

var ffj_key_Bid_API = []byte("api")

var ffj_key_Bid_Protocol = []byte("protocol")

var ffj_key_Bid_DealID = []byte("dealid")

var ffj_key_Bid_H = []byte("h")

var ffj_key_Bid_W = []byte("w")

var ffj_key_Bid_Wratio = []byte("wratio")

var ffj_key_Bid_Hratio = []byte("hratio")

var ffj_key_Bid_Exp = []byte("exp")

var ffj_key_Bid_Ext = []byte("ext")

func (uj *Bid) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Bid) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Bidbase
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
				currentKey = ffj_t_Bidno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_Bid_AdID, kn) {
						currentKey = ffj_t_Bid_AdID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_AdMarkup, kn) {
						currentKey = ffj_t_Bid_AdMarkup
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_AdvDomain, kn) {
						currentKey = ffj_t_Bid_AdvDomain
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_Attr, kn) {
						currentKey = ffj_t_Bid_Attr
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_API, kn) {
						currentKey = ffj_t_Bid_API
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffj_key_Bid_BURL, kn) {
						currentKey = ffj_t_Bid_BURL
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_Bundle, kn) {
						currentKey = ffj_t_Bid_Bundle
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffj_key_Bid_CampaignID, kn) {
						currentKey = ffj_t_Bid_CampaignID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_CreativeID, kn) {
						currentKey = ffj_t_Bid_CreativeID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_Cat, kn) {
						currentKey = ffj_t_Bid_Cat
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_Bid_DealID, kn) {
						currentKey = ffj_t_Bid_DealID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Bid_Exp, kn) {
						currentKey = ffj_t_Bid_Exp
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_Ext, kn) {
						currentKey = ffj_t_Bid_Ext
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_Bid_H, kn) {
						currentKey = ffj_t_Bid_H
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_Hratio, kn) {
						currentKey = ffj_t_Bid_Hratio
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_Bid_ID, kn) {
						currentKey = ffj_t_Bid_ID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_ImpID, kn) {
						currentKey = ffj_t_Bid_ImpID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_IURL, kn) {
						currentKey = ffj_t_Bid_IURL
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_Bid_LURL, kn) {
						currentKey = ffj_t_Bid_LURL
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_Language, kn) {
						currentKey = ffj_t_Bid_Language
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffj_key_Bid_NURL, kn) {
						currentKey = ffj_t_Bid_NURL
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Bid_Price, kn) {
						currentKey = ffj_t_Bid_Price
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_Protocol, kn) {
						currentKey = ffj_t_Bid_Protocol
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'q':

					if bytes.Equal(ffj_key_Bid_Qagmediarating, kn) {
						currentKey = ffj_t_Bid_Qagmediarating
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Bid_Tactic, kn) {
						currentKey = ffj_t_Bid_Tactic
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_Bid_W, kn) {
						currentKey = ffj_t_Bid_W
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Bid_Wratio, kn) {
						currentKey = ffj_t_Bid_Wratio
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Ext, kn) {
					currentKey = ffj_t_Bid_Ext
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Exp, kn) {
					currentKey = ffj_t_Bid_Exp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Hratio, kn) {
					currentKey = ffj_t_Bid_Hratio
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Wratio, kn) {
					currentKey = ffj_t_Bid_Wratio
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_W, kn) {
					currentKey = ffj_t_Bid_W
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_H, kn) {
					currentKey = ffj_t_Bid_H
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_DealID, kn) {
					currentKey = ffj_t_Bid_DealID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Protocol, kn) {
					currentKey = ffj_t_Bid_Protocol
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_API, kn) {
					currentKey = ffj_t_Bid_API
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Attr, kn) {
					currentKey = ffj_t_Bid_Attr
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Cat, kn) {
					currentKey = ffj_t_Bid_Cat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Language, kn) {
					currentKey = ffj_t_Bid_Language
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Qagmediarating, kn) {
					currentKey = ffj_t_Bid_Qagmediarating
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_CreativeID, kn) {
					currentKey = ffj_t_Bid_CreativeID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_CampaignID, kn) {
					currentKey = ffj_t_Bid_CampaignID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Bundle, kn) {
					currentKey = ffj_t_Bid_Bundle
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Tactic, kn) {
					currentKey = ffj_t_Bid_Tactic
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_IURL, kn) {
					currentKey = ffj_t_Bid_IURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_AdvDomain, kn) {
					currentKey = ffj_t_Bid_AdvDomain
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_AdMarkup, kn) {
					currentKey = ffj_t_Bid_AdMarkup
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_LURL, kn) {
					currentKey = ffj_t_Bid_LURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_BURL, kn) {
					currentKey = ffj_t_Bid_BURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_NURL, kn) {
					currentKey = ffj_t_Bid_NURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_AdID, kn) {
					currentKey = ffj_t_Bid_AdID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_Price, kn) {
					currentKey = ffj_t_Bid_Price
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_ImpID, kn) {
					currentKey = ffj_t_Bid_ImpID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Bid_ID, kn) {
					currentKey = ffj_t_Bid_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Bidno_such_key
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

				case ffj_t_Bid_ID:
					goto handle_ID

				case ffj_t_Bid_ImpID:
					goto handle_ImpID

				case ffj_t_Bid_Price:
					goto handle_Price

				case ffj_t_Bid_AdID:
					goto handle_AdID

				case ffj_t_Bid_NURL:
					goto handle_NURL

				case ffj_t_Bid_BURL:
					goto handle_BURL

				case ffj_t_Bid_LURL:
					goto handle_LURL

				case ffj_t_Bid_AdMarkup:
					goto handle_AdMarkup

				case ffj_t_Bid_AdvDomain:
					goto handle_AdvDomain

				case ffj_t_Bid_IURL:
					goto handle_IURL

				case ffj_t_Bid_Tactic:
					goto handle_Tactic

				case ffj_t_Bid_Bundle:
					goto handle_Bundle

				case ffj_t_Bid_CampaignID:
					goto handle_CampaignID

				case ffj_t_Bid_CreativeID:
					goto handle_CreativeID

				case ffj_t_Bid_Qagmediarating:
					goto handle_Qagmediarating

				case ffj_t_Bid_Language:
					goto handle_Language

				case ffj_t_Bid_Cat:
					goto handle_Cat

				case ffj_t_Bid_Attr:
					goto handle_Attr

				case ffj_t_Bid_API:
					goto handle_API

				case ffj_t_Bid_Protocol:
					goto handle_Protocol

				case ffj_t_Bid_DealID:
					goto handle_DealID

				case ffj_t_Bid_H:
					goto handle_H

				case ffj_t_Bid_W:
					goto handle_W

				case ffj_t_Bid_Wratio:
					goto handle_Wratio

				case ffj_t_Bid_Hratio:
					goto handle_Hratio

				case ffj_t_Bid_Exp:
					goto handle_Exp

				case ffj_t_Bid_Ext:
					goto handle_Ext

				case ffj_t_Bidno_such_key:
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

handle_ImpID:

	/* handler: uj.ImpID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ImpID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Price:

	/* handler: uj.Price type=float64 kind=float64 quoted=false*/

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

			uj.Price = float64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AdID:

	/* handler: uj.AdID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.AdID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NURL:

	/* handler: uj.NURL type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.NURL = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BURL:

	/* handler: uj.BURL type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.BURL = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LURL:

	/* handler: uj.LURL type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.LURL = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AdMarkup:

	/* handler: uj.AdMarkup type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.AdMarkup = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AdvDomain:

	/* handler: uj.AdvDomain type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.AdvDomain = nil
		} else {

			uj.AdvDomain = []string{}

			wantVal := true

			for {

				var tmp_uj__AdvDomain string

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

				/* handler: tmp_uj__AdvDomain type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__AdvDomain = string(string(outBuf))

					}
				}

				uj.AdvDomain = append(uj.AdvDomain, tmp_uj__AdvDomain)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IURL:

	/* handler: uj.IURL type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.IURL = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Tactic:

	/* handler: uj.Tactic type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Tactic = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Bundle:

	/* handler: uj.Bundle type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Bundle = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CampaignID:

	/* handler: uj.CampaignID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.CampaignID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CreativeID:

	/* handler: uj.CreativeID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.CreativeID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Qagmediarating:

	/* handler: uj.Qagmediarating type=int kind=int quoted=false*/

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

			uj.Qagmediarating = int(tval)

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

handle_Cat:

	/* handler: uj.Cat type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Cat = nil
		} else {

			uj.Cat = []string{}

			wantVal := true

			for {

				var tmp_uj__Cat string

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

				/* handler: tmp_uj__Cat type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmp_uj__Cat = string(string(outBuf))

					}
				}

				uj.Cat = append(uj.Cat, tmp_uj__Cat)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Attr:

	/* handler: uj.Attr type=[]int kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Attr = nil
		} else {

			uj.Attr = []int{}

			wantVal := true

			for {

				var tmp_uj__Attr int

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

				/* handler: tmp_uj__Attr type=int kind=int quoted=false*/

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

						tmp_uj__Attr = int(tval)

					}
				}

				uj.Attr = append(uj.Attr, tmp_uj__Attr)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_API:

	/* handler: uj.API type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.API = string(string(outBuf))

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

handle_DealID:

	/* handler: uj.DealID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.DealID = string(string(outBuf))

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

handle_Wratio:

	/* handler: uj.Wratio type=int kind=int quoted=false*/

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

			uj.Wratio = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Hratio:

	/* handler: uj.Hratio type=int kind=int quoted=false*/

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

			uj.Hratio = int(tval)

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
