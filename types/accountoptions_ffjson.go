// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: accountoptions.go

package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AccountOptions) MarshalJSON() ([]byte, error) {
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
func (j *AccountOptions) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"memo_key":`)

	{

		obj, err = j.MemoKey.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"voting_account":`)

	{

		obj, err = j.VotingAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"num_witness":`)
	fflib.FormatBits2(buf, uint64(j.NumWitness), 10, false)
	buf.WriteString(`,"num_committee":`)
	fflib.FormatBits2(buf, uint64(j.NumCommittee), 10, false)
	buf.WriteString(`,"votes":`)
	if j.Votes != nil {
		buf.WriteString(`[`)
		for i, v := range j.Votes {
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
	buf.WriteString(`,"extensions":`)
	if j.Extensions != nil {
		buf.WriteString(`[`)
		for i, v := range j.Extensions {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Interface types must use runtime reflection. type=interface {} kind=interface */
			err = buf.Encode(v)
			if err != nil {
				return err
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAccountOptionsbase = iota
	ffjtAccountOptionsnosuchkey

	ffjtAccountOptionsMemoKey

	ffjtAccountOptionsVotingAccount

	ffjtAccountOptionsNumWitness

	ffjtAccountOptionsNumCommittee

	ffjtAccountOptionsVotes

	ffjtAccountOptionsExtensions
)

var ffjKeyAccountOptionsMemoKey = []byte("memo_key")

var ffjKeyAccountOptionsVotingAccount = []byte("voting_account")

var ffjKeyAccountOptionsNumWitness = []byte("num_witness")

var ffjKeyAccountOptionsNumCommittee = []byte("num_committee")

var ffjKeyAccountOptionsVotes = []byte("votes")

var ffjKeyAccountOptionsExtensions = []byte("extensions")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AccountOptions) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AccountOptions) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAccountOptionsbase
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
				currentKey = ffjtAccountOptionsnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffjKeyAccountOptionsExtensions, kn) {
						currentKey = ffjtAccountOptionsExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyAccountOptionsMemoKey, kn) {
						currentKey = ffjtAccountOptionsMemoKey
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyAccountOptionsNumWitness, kn) {
						currentKey = ffjtAccountOptionsNumWitness
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountOptionsNumCommittee, kn) {
						currentKey = ffjtAccountOptionsNumCommittee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffjKeyAccountOptionsVotingAccount, kn) {
						currentKey = ffjtAccountOptionsVotingAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountOptionsVotes, kn) {
						currentKey = ffjtAccountOptionsVotes
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyAccountOptionsExtensions, kn) {
					currentKey = ffjtAccountOptionsExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountOptionsVotes, kn) {
					currentKey = ffjtAccountOptionsVotes
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAccountOptionsNumCommittee, kn) {
					currentKey = ffjtAccountOptionsNumCommittee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountOptionsNumWitness, kn) {
					currentKey = ffjtAccountOptionsNumWitness
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAccountOptionsVotingAccount, kn) {
					currentKey = ffjtAccountOptionsVotingAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountOptionsMemoKey, kn) {
					currentKey = ffjtAccountOptionsMemoKey
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAccountOptionsnosuchkey
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

				case ffjtAccountOptionsMemoKey:
					goto handle_MemoKey

				case ffjtAccountOptionsVotingAccount:
					goto handle_VotingAccount

				case ffjtAccountOptionsNumWitness:
					goto handle_NumWitness

				case ffjtAccountOptionsNumCommittee:
					goto handle_NumCommittee

				case ffjtAccountOptionsVotes:
					goto handle_Votes

				case ffjtAccountOptionsExtensions:
					goto handle_Extensions

				case ffjtAccountOptionsnosuchkey:
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

handle_MemoKey:

	/* handler: j.MemoKey type=types.PublicKey kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MemoKey.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VotingAccount:

	/* handler: j.VotingAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.VotingAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NumWitness:

	/* handler: j.NumWitness type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.NumWitness.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NumCommittee:

	/* handler: j.NumCommittee type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.NumCommittee.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Votes:

	/* handler: j.Votes type=types.Votes kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Votes", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Votes = nil
		} else {

			j.Votes = []VoteID{}

			wantVal := true

			for {

				var tmpJVotes VoteID

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

				/* handler: tmpJVotes type=types.VoteID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJVotes.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.Votes = append(j.Votes, tmpJVotes)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Extensions", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Extensions = nil
		} else {

			j.Extensions = []interface{}{}

			wantVal := true

			for {

				var tmpJExtensions interface{}

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

				/* handler: tmpJExtensions type=interface {} kind=interface quoted=false*/

				{
					/* Falling back. type=interface {} kind=interface */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJExtensions)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.Extensions = append(j.Extensions, tmpJExtensions)

				wantVal = false
			}
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
