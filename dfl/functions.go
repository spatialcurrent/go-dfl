package dfl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unicode"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-reader/reader"
)

func prefix(ctx interface{}, args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to prefix.")
	}

	switch lv := args[0].(type) {
	case *reader.Cache:
		switch prefix := args[1].(type) {
		case []byte:
			data, err := lv.ReadRange(0, len(prefix)-1)
			if err != nil {
				return false, nil
			}
			for i, c := range prefix {
				if data[i] != c {
					return false, nil
				}
			}
			return true, nil
		case string:
			data, err := lv.ReadRange(0, len(prefix)-1)
			if err != nil {
				return false, nil
			}
			s := []rune(string(data))
			if len(s) < len(prefix) {
				return false, nil
			}
			for i, c := range prefix {
				if s[i] != c {
					return false, nil
				}
			}
			return true, nil
		}
		return Null{}, errors.New("Invalid arguments for prefix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	case []byte:
		switch prefix := args[1].(type) {
		case []byte:
			if len(prefix) > len(lv) {
				return false, nil
			}
			for i, c := range prefix {
				if lv[i] != c {
					return false, nil
				}
			}
			return true, nil
		case string:
			prefix_bytes := []byte(prefix)
			if len(prefix_bytes) > len(lv) {
				return false, nil
			}
			for i, c := range prefix_bytes {
				if lv[i] != c {
					return false, nil
				}
			}
			return true, nil
		}
		return Null{}, errors.New("Invalid arguments for prefix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	case string:
		switch prefix := args[1].(type) {
		case string:
			return strings.HasPrefix(lv, prefix), nil
		}
		return Null{}, errors.New("Invalid arguments for prefix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	}

	return 0, errors.New("Invalid arguments for prefix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())

}

func suffix(ctx interface{}, args []interface{}) (interface{}, error) {

	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to suffix.")
	}

	switch lv := args[0].(type) {
	case *reader.Cache:
		switch suffix := args[1].(type) {
		case []byte:
			data, err := lv.ReadAll()
			if err != nil {
				return false, nil
			}
			if len(suffix) > len(data) {
				return false, nil
			}
			for i, _ := range suffix {
				if data[len(data)-i-1] != suffix[len(suffix)-i-1] {
					return false, nil
				}
			}
			return true, nil
		case string:
			data, err := lv.ReadAll()
			if err != nil {
				return false, nil
			}
			//s := []rune(string(data))
			s := string(data)
			if len(suffix) > len(s) {
				return false, nil
			}
			for i, _ := range suffix {
				if s[len(s)-i-1] != suffix[len(suffix)-i-1] {
					return false, nil
				}
			}
			return true, nil
		}
		return Null{}, errors.New("Invalid arguments for suffix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	case []byte:
		switch suffix := args[1].(type) {
		case []byte:
			if len(suffix) > len(lv) {
				return false, nil
			}
			for i, _ := range suffix {
				if lv[len(lv)-i-1] != suffix[len(suffix)-i-1] {
					return false, nil
				}
			}
			return true, nil
		}
		return Null{}, errors.New("Invalid arguments for suffix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	case string:
		switch suffix := args[1].(type) {
		case string:
			return strings.HasSuffix(lv, suffix), nil
		}
		return Null{}, errors.New("Invalid arguments for suffix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	}

	return 0, errors.New("Invalid arguments for suffix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())

}

func mapArray(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to map.")
	}

	switch key := args[1].(type) {
	case string:
		switch a := args[0].(type) {
		case []map[string]interface{}:
			values := make([]interface{}, 0, len(a))
			for _, value := range a {
				values = append(values, value[key])
			}
			return values, nil
		case []map[string]string:
			values := make([]string, 0, len(a))
			for _, value := range a {
				values = append(values, value[key])
			}
			return values, nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	return 0, errors.New("Invalid key for map function")

}

func splitString(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to split.")
	}
	return strings.Split(fmt.Sprint(args[0]), fmt.Sprint(args[1])), nil
}

func trimString(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to split.")
	}

	switch a := args[0].(type) {
	case string:
		return strings.TrimSpace(a), nil
	case []byte:
		return []byte(strings.TrimSpace(string(a))), nil
	case *reader.Cache:
		b, err := a.ReadAll()
		if err != nil {
			return make([]byte, 0), errors.Wrap(err, "error reading all bytes from *reader.Cache")
		}
		return []byte(strings.TrimSpace(string(b))), nil
	}

	return "", errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func trimStringLeft(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to ltrim.")
	}

	switch a := args[0].(type) {
	case string:
		return strings.TrimLeftFunc(a, unicode.IsSpace), nil
	case []byte:
		return []byte(strings.TrimLeftFunc(string(a), unicode.IsSpace)), nil
	case *reader.Cache:
		i := 0
		for i = 0; ; i++ {
			b, err := a.ReadAt(i)
			if err != nil {
				if err == io.EOF {
					return make([]byte, 0), nil
				} else {
					return make([]byte, 0), errors.Wrap(err, "error reading byte at position "+fmt.Sprint(i)+" in trimStringLeft")
				}
			}
			if !unicode.IsSpace(bytes.Runes([]byte{b})[0]) {
				break
			}
		}
		return reader.NewCacheWithContent(a.Reader, a.Content, i), nil
	}

	return "", errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func trimStringRight(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to rtrim.")
	}

	switch a := args[0].(type) {
	case string:
		return strings.TrimRightFunc(a, unicode.IsSpace), nil
	case []byte:
		return []byte(strings.TrimRightFunc(string(a), unicode.IsSpace)), nil
	case *reader.Cache:
		b, err := a.ReadAll()
		if err != nil {
			return make([]byte, 0), errors.Wrap(err, "error reading all bytes from *reader.Cache")
		}
		return []byte(strings.TrimRightFunc(string(b), unicode.IsSpace)), nil
	}
	return "", errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func getLength(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to len.")
	}

	switch a := args[0].(type) {
	case string:
		return len(a), nil
	case []int:
		return len(a), nil
	case []string:
		return len(a), nil
	case []uint8:
		return len(a), nil
	case []float64:
		return len(a), nil
	case []interface{}:
		return len(a), nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToBytes(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to len.")
	}

	switch a := args[0].(type) {
	case string:
		return []byte(a), nil
	case byte:
		return []byte{a}, nil
	case []byte:
		return a, nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToInt16(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to int16.")
	}
	switch a := args[0].(type) {
	case int:
		return int16(a), nil
	case int8:
		return int16(a), nil
	case int16:
		return a, nil
	case int32:
		return int16(a), nil
	case int64:
		return int16(a), nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToInt32(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to int32.")
	}
	switch a := args[0].(type) {
	case int:
		return int32(a), nil
	case int16:
		return int32(a), nil
	case int32:
		return a, nil
	case int64:
		return int32(a), nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToBigEndian(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to big-endian.")
	}
	switch a := args[0].(type) {
	case int:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, int64(a))
		return buf.Bytes(), err
	case int16:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, a)
		return buf.Bytes(), err
	case int32:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, a)
		return buf.Bytes(), err
	case int64:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, a)
		return buf.Bytes(), err
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToLittleEndian(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to little-endian.")
	}
	switch a := args[0].(type) {
	case int:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, int64(a))
		return buf.Bytes(), err
	case int16:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, a)
		return buf.Bytes(), err
	case int32:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, a)
		return buf.Bytes(), err
	case int64:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, a)
		return buf.Bytes(), err
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func repeat(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to repeat.")
	}
	switch value := args[0].(type) {
	case string:
		switch count := args[1].(type) {
		case int:
			return strings.Repeat(value, count), nil
		}
	case []byte:
		switch count := args[1].(type) {
		case int:
			return bytes.Repeat(value, count), nil
		}
	case byte:
		switch count := args[1].(type) {
		case int:
			return bytes.Repeat([]byte{value}, count), nil
		}
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToString(ctx interface{}, args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to len.")
	}
	switch a := args[0].(type) {
	case string:
		return a, nil
	case []byte:
		return string(a), nil
	case byte:
		return string([]byte{a}), nil
	case *reader.Cache:
		value, err := a.ReadAll()
		if err != nil {
			return "", errors.Wrap(err, "error reading all content from *reader.Cache in covertToString")
		}
		return string(value), nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}
