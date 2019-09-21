package util

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
)

var (
	bytesT  = reflect.TypeOf(Bytes(nil))
	bigT    = reflect.TypeOf((*Big)(nil))
	uintT   = reflect.TypeOf(Uint(0))
	uint64T = reflect.TypeOf(Uint64(0))
)

type Bytes []byte

func (b Bytes) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b)
	return result, nil
}

func (b *Bytes) UnmarshalJSON(input []byte) error {
	if !isString(input) {
		return errNonString(bytesT)
	}
	return wrapTypeError(b.UnmarshalText(input[1:len(input)-1]), bytesT)
}

func (b *Bytes) UnmarshalText(input []byte) error {
	raw, err := checkText(input, true)
	if err != nil {
		return err
	}
	dec := make([]byte, len(raw)/2)
	if _, err = hex.Decode(dec, raw); err != nil {
		err = mapError(err)
	} else {
		*b = dec
	}
	return err
}

func (b Bytes) String() string {
	return Encode(b)
}

func (b Bytes) ImplementsGraphQLType(name string) bool { return name == "Bytes" }

func (b *Bytes) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		data, err := Decode(input)
		if err != nil {
			return err
		}
		*b = data
	default:
		err = fmt.Errorf("Unexpected type for Bytes: %v", input)
	}
	return err
}

func UnmarshalFixedJSON(typ reflect.Type, input, out []byte) error {
	if !isString(input) {
		return errNonString(typ)
	}
	return wrapTypeError(UnmarshalFixedText(typ.String(), input[1:len(input)-1], out), typ)
}

func UnmarshalFixedText(typname string, input, out []byte) error {
	raw, err := checkText(input, true)
	if err != nil {
		return err
	}
	if len(raw)/2 != len(out) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(out)*2, typname)
	}

	for _, b := range raw {
		if decodeNibble(b) == badNibble {
			return ErrSyntax
		}
	}
	hex.Decode(out, raw)
	return nil
}

func UnmarshalFixedUnprefixedText(typname string, input, out []byte) error {
	raw, err := checkText(input, false)
	if err != nil {
		return err
	}
	if len(raw)/2 != len(out) {
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(out)*2, typname)
	}

	for _, b := range raw {
		if decodeNibble(b) == badNibble {
			return ErrSyntax
		}
	}
	hex.Decode(out, raw)
	return nil
}

type Big big.Int

func (b Big) MarshalText() ([]byte, error) {
	return []byte(EncodeBig((*big.Int)(&b))), nil
}

func (b *Big) UnmarshalJSON(input []byte) error {
	if !isString(input) {
		return errNonString(bigT)
	}
	return wrapTypeError(b.UnmarshalText(input[1:len(input)-1]), bigT)
}

func (b *Big) UnmarshalText(input []byte) error {
	raw, err := checkNumberText(input)
	if err != nil {
		return err
	}
	if len(raw) > 64 {
		return ErrBig256Range
	}
	words := make([]big.Word, len(raw)/bigWordNibbles+1)
	end := len(raw)
	for i := range words {
		start := end - bigWordNibbles
		if start < 0 {
			start = 0
		}
		for ri := start; ri < end; ri++ {
			nib := decodeNibble(raw[ri])
			if nib == badNibble {
				return ErrSyntax
			}
			words[i] *= 16
			words[i] += big.Word(nib)
		}
		end = start
	}
	var dec big.Int
	dec.SetBits(words)
	*b = (Big)(dec)
	return nil
}

func (b *Big) ToInt() *big.Int {
	return (*big.Int)(b)
}

func (b *Big) String() string {
	return EncodeBig(b.ToInt())
}

func (b Big) ImplementsGraphQLType(name string) bool { return name == "BigInt" }

func (b *Big) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		return b.UnmarshalText([]byte(input))
	case int32:
		var num big.Int
		num.SetInt64(int64(input))
		*b = Big(num)
	default:
		err = fmt.Errorf("Unexpected type for BigInt: %v", input)
	}
	return err
}

type Uint64 uint64

func (b Uint64) MarshalText() ([]byte, error) {
	buf := make([]byte, 2, 10)
	copy(buf, `0x`)
	buf = strconv.AppendUint(buf, uint64(b), 16)
	return buf, nil
}

func (b *Uint64) UnmarshalJSON(input []byte) error {
	if !isString(input) {
		return errNonString(uint64T)
	}
	return wrapTypeError(b.UnmarshalText(input[1:len(input)-1]), uint64T)
}

func (b *Uint64) UnmarshalText(input []byte) error {
	raw, err := checkNumberText(input)
	if err != nil {
		return err
	}
	if len(raw) > 16 {
		return ErrUint64Range
	}
	var dec uint64
	for _, byte := range raw {
		nib := decodeNibble(byte)
		if nib == badNibble {
			return ErrSyntax
		}
		dec *= 16
		dec += nib
	}
	*b = Uint64(dec)
	return nil
}

func (b Uint64) String() string {
	return EncodeUint64(uint64(b))
}

func (b Uint64) ImplementsGraphQLType(name string) bool { return name == "Long" }

func (b *Uint64) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		return b.UnmarshalText([]byte(input))
	case int32:
		*b = Uint64(input)
	default:
		err = fmt.Errorf("Unexpected type for Long: %v", input)
	}
	return err
}

type Uint uint

func (b Uint) MarshalText() ([]byte, error) {
	return Uint64(b).MarshalText()
}

func (b *Uint) UnmarshalJSON(input []byte) error {
	if !isString(input) {
		return errNonString(uintT)
	}
	return wrapTypeError(b.UnmarshalText(input[1:len(input)-1]), uintT)
}

func (b *Uint) UnmarshalText(input []byte) error {
	var u64 Uint64
	err := u64.UnmarshalText(input)
	if u64 > Uint64(^uint(0)) || err == ErrUint64Range {
		return ErrUintRange
	} else if err != nil {
		return err
	}
	*b = Uint(u64)
	return nil
}

func (b Uint) String() string {
	return EncodeUint64(uint64(b))
}

func isString(input []byte) bool {
	return len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"'
}

func bytesHave0xPrefix(input []byte) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func checkText(input []byte, wantPrefix bool) ([]byte, error) {
	if len(input) == 0 {
		return nil, nil // empty strings are allowed
	}
	if bytesHave0xPrefix(input) {
		input = input[2:]
	} else if wantPrefix {
		return nil, ErrMissingPrefix
	}
	if len(input)%2 != 0 {
		return nil, ErrOddLength
	}
	return input, nil
}

func checkNumberText(input []byte) (raw []byte, err error) {
	if len(input) == 0 {
		return nil, nil // empty strings are allowed
	}
	if !bytesHave0xPrefix(input) {
		return nil, ErrMissingPrefix
	}
	input = input[2:]
	if len(input) == 0 {
		return nil, ErrEmptyNumber
	}
	if len(input) > 1 && input[0] == '0' {
		return nil, ErrLeadingZero
	}
	return input, nil
}

func wrapTypeError(err error, typ reflect.Type) error {
	if _, ok := err.(*decError); ok {
		return &json.UnmarshalTypeError{Value: err.Error(), Type: typ}
	}
	return err
}

func errNonString(typ reflect.Type) error {
	return &json.UnmarshalTypeError{Value: "non-string", Type: typ}
}

// JsonMap 存储任意数据的map
type JsonMap map[string]interface{}

// JsonEncode 将JsonMap转成[]byte
func JsonEncode(m JsonMap) ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		err = fmt.Errorf("JsonEncode err: %s", err.Error())
	}
	return b, err
}

// JsonDecode 将[]byte转成JsonMap
func JsonDecode(b []byte) (JsonMap, error) {
	m := make(JsonMap)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		err = fmt.Errorf("JsonDecode err: %s", err.Error())
	} else {
		m = f.(map[string]interface{})
	}
	return m, err
}

// JsonGetInt 从JsonMap中解析出一个int值
func (m JsonMap) JsonGetInt(key string) (int, bool) {
	if val, exists := m[key]; exists {
		switch val.(type) {
		case float64:
			return int(val.(float64)), true
		case int:
			return val.(int), true
		}
	}
	return 0, false
}

// JsonGetUint16 从JsonMap中解析出一个ushort值
func (m JsonMap) JsonGetUint16(key string) (uint16, bool) {
	if val, exists := m[key]; exists {
		return uint16(val.(float64)), true
	}
	return uint16(0), false
}

// JsonGetString 从JsonMap中解析出一个string值
func (m JsonMap) JsonGetString(key string) (string, bool) {
	if val, exists := m[key]; exists {
		return val.(string), true
	}
	return "", false
}

// JsonGetJsonMap 从JsonMap中解析出一个JsonMap值
func (this JsonMap) JsonGetJsonMap(key string) JsonMap {
	if val, exists := this[key]; exists {
		switch val.(type) {
		case map[string]interface{}:
			return JsonMap(val.(map[string]interface{}))
		case interface{}:
			return val.(JsonMap)
		}
	}
	return JsonMap{}
}

// InterfaceToJsonString 将任意类型的数据，转成json格式的字符串
func InterfaceToJsonString(s interface{}) (string, error) {
	byt, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
