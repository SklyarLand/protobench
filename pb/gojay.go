package pb

import (
	"github.com/francoispqt/gojay"
)

func (u *BenchSmall) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("action", u.Action)
	enc.StringKey("key", u.Key)
}
func (u *BenchSmall) IsNil() bool {
	return u == nil
}

func (u *BenchSmall) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "action":
		return dec.String(&u.Action)
	case "key":
		return nil
	}
	return nil
}
func (u *BenchSmall) NKeys() int {
	return 2
}

func (u *BenchMedium) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("name", u.Name)
	enc.Int64Key("age", u.Age)
	enc.Float32Key("height", u.Height)
	enc.Float64Key("weight", u.Weight)
	enc.BoolKey("alive", u.Alive)
	enc.StringKey("desc", u.Desc)
}
func (u *BenchMedium) IsNil() bool {
	return u == nil
}

func (u *BenchMedium) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "name":
		return dec.String(&u.Name)
	case "age":
		return dec.Int64(&u.Age)
	case "height":
		return dec.Float32(&u.Height)
	case "weight":
		return dec.Float64(&u.Weight)
	case "alive":
		return dec.Bool(&u.Alive)
	case "desc":
		return dec.String(&u.Desc)
	}
	return nil
}
func (u *BenchMedium) NKeys() int {
	return 6
}

func (u *BenchLarge) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("name", u.Name)
	enc.Int64Key("age", u.Age)
	enc.Float32Key("height", u.Height)
	enc.Float64Key("weight", u.Weight)
	enc.BoolKey("alive", u.Alive)
	enc.StringKey("desc", u.Desc)
	enc.StringKey("nickname", u.Nickname)
	enc.Int64Key("num", u.Num)
	enc.Float32Key("flt", u.Flt)
	enc.StringKey("data", u.Data)
}
func (u *BenchLarge) IsNil() bool {
	return u == nil
}

func (u *BenchLarge) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "name":
		return dec.String(&u.Name)
	case "age":
		return dec.Int64(&u.Age)
	case "height":
		return dec.Float32(&u.Height)
	case "weight":
		return dec.Float64(&u.Weight)
	case "alive":
		return dec.Bool(&u.Alive)
	case "desc":
		return dec.String(&u.Desc)
	case "nickname":
		return dec.String(&u.Nickname)
	case "num":
		return dec.Int64(&u.Num)
	case "flt":
		return dec.Float32(&u.Flt)
	case "data":
		return dec.String(&u.Data)
	}
	return nil
}
func (u *BenchLarge) NKeys() int {
	return 10
}
