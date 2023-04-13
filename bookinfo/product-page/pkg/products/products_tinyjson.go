// Code generated by tinyjson for marshaling/unmarshaling. DO NOT EDIT.

package products

import (
	tinyjson "github.com/CosmWasm/tinyjson"
	jlexer "github.com/CosmWasm/tinyjson/jlexer"
	jwriter "github.com/CosmWasm/tinyjson/jwriter"
)

// suppress unused package warning
var (
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ tinyjson.Marshaler
)

func tinyjson797f9fe8DecodeGithubComProductPagePkgProducts(in *jlexer.Lexer, out *Review) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "reviewer":
			out.Reviewer = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "rating":
			(out.Rating).UnmarshalTinyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson797f9fe8EncodeGithubComProductPagePkgProducts(out *jwriter.Writer, in Review) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"reviewer\":"
		out.RawString(prefix[1:])
		out.String(string(in.Reviewer))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		(in.Rating).MarshalTinyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Review) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v Review) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Review) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *Review) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts(l, v)
}
func tinyjson797f9fe8DecodeGithubComProductPagePkgProducts1(in *jlexer.Lexer, out *Rating) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "stars":
			if in.IsNull() {
				in.Skip()
				out.Stars = nil
			} else {
				if out.Stars == nil {
					out.Stars = new(int)
				}
				*out.Stars = int(in.Int())
			}
		case "color":
			if in.IsNull() {
				in.Skip()
				out.Color = nil
			} else {
				if out.Color == nil {
					out.Color = new(string)
				}
				*out.Color = string(in.String())
			}
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(string)
				}
				*out.Error = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson797f9fe8EncodeGithubComProductPagePkgProducts1(out *jwriter.Writer, in Rating) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"stars\":"
		out.RawString(prefix[1:])
		if in.Stars == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.Stars))
		}
	}
	{
		const prefix string = ",\"color\":"
		out.RawString(prefix)
		if in.Color == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Color))
		}
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		if in.Error == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Error))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Rating) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v Rating) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Rating) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts1(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *Rating) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts1(l, v)
}
func tinyjson797f9fe8DecodeGithubComProductPagePkgProducts2(in *jlexer.Lexer, out *Products) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Products, 0, 1)
			} else {
				*out = Products{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Product
			(v1).UnmarshalTinyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson797f9fe8EncodeGithubComProductPagePkgProducts2(out *jwriter.Writer, in Products) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalTinyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Products) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v Products) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Products) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts2(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *Products) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts2(l, v)
}
func tinyjson797f9fe8DecodeGithubComProductPagePkgProducts3(in *jlexer.Lexer, out *ProductReviews) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "podname":
			out.PodName = string(in.String())
		case "clustername":
			out.ClusterName = string(in.String())
		case "reviews":
			if in.IsNull() {
				in.Skip()
				out.Reviews = nil
			} else {
				in.Delim('[')
				if out.Reviews == nil {
					if !in.IsDelim(']') {
						out.Reviews = make([]Review, 0, 1)
					} else {
						out.Reviews = []Review{}
					}
				} else {
					out.Reviews = (out.Reviews)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Review
					(v4).UnmarshalTinyJSON(in)
					out.Reviews = append(out.Reviews, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson797f9fe8EncodeGithubComProductPagePkgProducts3(out *jwriter.Writer, in ProductReviews) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"podname\":"
		out.RawString(prefix)
		out.String(string(in.PodName))
	}
	{
		const prefix string = ",\"clustername\":"
		out.RawString(prefix)
		out.String(string(in.ClusterName))
	}
	{
		const prefix string = ",\"reviews\":"
		out.RawString(prefix)
		if in.Reviews == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Reviews {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalTinyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProductReviews) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v ProductReviews) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProductReviews) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts3(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *ProductReviews) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts3(l, v)
}
func tinyjson797f9fe8DecodeGithubComProductPagePkgProducts4(in *jlexer.Lexer, out *ProductRatings) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "ratings":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Ratings = make(map[string]int)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 int
					v7 = int(in.Int())
					(out.Ratings)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson797f9fe8EncodeGithubComProductPagePkgProducts4(out *jwriter.Writer, in ProductRatings) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"ratings\":"
		out.RawString(prefix)
		if in.Ratings == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v8First := true
			for v8Name, v8Value := range in.Ratings {
				if v8First {
					v8First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v8Name))
				out.RawByte(':')
				out.Int(int(v8Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProductRatings) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v ProductRatings) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProductRatings) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts4(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *ProductRatings) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts4(l, v)
}
func tinyjson797f9fe8DecodeGithubComProductPagePkgProducts5(in *jlexer.Lexer, out *ProductDetails) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "author":
			out.Author = string(in.String())
		case "year":
			out.Year = uint16(in.Uint16())
		case "type":
			out.Type = string(in.String())
		case "pages":
			out.Pages = uint16(in.Uint16())
		case "publisher":
			out.Publisher = string(in.String())
		case "language":
			out.Language = string(in.String())
		case "ISBN-10":
			out.ISBN10 = string(in.String())
		case "ISBN-13":
			out.ISBN13 = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson797f9fe8EncodeGithubComProductPagePkgProducts5(out *jwriter.Writer, in ProductDetails) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"year\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.Year))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"pages\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.Pages))
	}
	{
		const prefix string = ",\"publisher\":"
		out.RawString(prefix)
		out.String(string(in.Publisher))
	}
	{
		const prefix string = ",\"language\":"
		out.RawString(prefix)
		out.String(string(in.Language))
	}
	{
		const prefix string = ",\"ISBN-10\":"
		out.RawString(prefix)
		out.String(string(in.ISBN10))
	}
	{
		const prefix string = ",\"ISBN-13\":"
		out.RawString(prefix)
		out.String(string(in.ISBN13))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProductDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v ProductDetails) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProductDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts5(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *ProductDetails) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts5(l, v)
}
func tinyjson797f9fe8DecodeGithubComProductPagePkgProducts6(in *jlexer.Lexer, out *Product) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "descriptionHtml":
			out.DescriptionHtml = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson797f9fe8EncodeGithubComProductPagePkgProducts6(out *jwriter.Writer, in Product) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"descriptionHtml\":"
		out.RawString(prefix)
		out.String(string(in.DescriptionHtml))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Product) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v Product) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson797f9fe8EncodeGithubComProductPagePkgProducts6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Product) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts6(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *Product) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson797f9fe8DecodeGithubComProductPagePkgProducts6(l, v)
}
