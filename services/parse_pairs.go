package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/beyondstorage/go-storage/v4/pairs"
	. "github.com/beyondstorage/go-storage/v4/types"
)

var (
	// ErrPairTypeNotParsable means the pair's type is not parseable.
	ErrPairTypeNotParsable = NewErrorCode("type not parseable")
	// ErrPairNotRegistered means the pair is not registered.
	ErrPairNotRegistered = NewErrorCode("pair not registered")
	// ErrPairValueInvalid means the pair's value is invalid.
	ErrPairValueInvalid = NewErrorCode("pair value invalid")
	// ErrConfigStringInvalid means the config string is invalid
	ErrConfigStringInvalid = NewErrorCode("config string is invalid")
)

// <type>://[<name>][<work_dir>][?key1=value1&...&keyN=valueN]
func parseString(config string) (ty string, ps []Pair, err error) {
	colon := strings.Index(config, ":")
	if colon == -1 {
		err = fmt.Errorf("%w: %s, %s", ErrConfigStringInvalid, "service type missing", config)
		return
	}
	ty = config[:colon]
	rest := config[colon+1:]
	m, ok := servicePairMaps[ty]
	if !ok {
		err = ErrServiceNotRegistered
		return
	}

	if !strings.HasPrefix(rest, "//") {
		err = fmt.Errorf("%w: %s", ErrConfigStringInvalid, config)
		return
	}
	rest = rest[2:]

	// [<name>][<work_dir>][?key1=value1&...&keyN=valueN]
	// <name> does not contain '/'
	// <work_dir> begins with '/'
	question := strings.Index(rest, "?")
	var path string
	if question == -1 {
		path = rest
		rest = ""
	} else {
		path = rest[:question]
		rest = rest[question+1:]
	}

	if len(path) == 0 {
		err = fmt.Errorf("%w: %s, %s", ErrConfigStringInvalid, "both <name> and <work_dir> missing", config)
		return
	} else {
		slash := strings.Index(path, "/")
		if slash == -1 {
			name := path
			ps = append(ps, pairs.WithName(name))
		} else if slash == 0 {
			work_dir := path
			ps = append(ps, pairs.WithWorkDir(work_dir))
		} else {
			name := path[:slash]
			work_dir := path[slash:]
			ps = append(ps, pairs.WithName(name), pairs.WithWorkDir(work_dir))
		}
	}

	for _, v := range strings.Split(rest, "&") {
		opt := strings.SplitN(v, "=", 2)
		if len(opt) != 2 {
			// && or &key& or &key=&, ignore
			continue
		}
		parse(m, opt[0], opt[1])
	}
	return
}

func parse(m map[string]string, k string, v string) (pair Pair, err error) {
	vType, ok := m[k]
	if !ok {
		err = fmt.Errorf("%w: %v", ErrPairNotRegistered, k)
		return Pair{}, err
	}

	pair.Key = k

	switch vType {
	case "string":
		pair.Value, err = v, nil
	case "bool":
		pair.Value, err = strconv.ParseBool(v)
	case "int":
		var i int64
		i, err = strconv.ParseInt(v, 0, 0)
		pair.Value = int(i)
	case "int64":
		pair.Value, err = strconv.ParseInt(v, 0, 64)
	case "[]byte":
		pair.Value, err = base64.RawStdEncoding.DecodeString(v)
	case "ListMode":
		pair.Value, err = parseListMode(v)
	default:
		return Pair{}, fmt.Errorf("%w: %v, %v", ErrPairTypeNotParsable, k, vType)
	}

	if err != nil {
		pair = Pair{}
		err = fmt.Errorf("%w: %v, %v, %v: %v", ErrPairValueInvalid, k, vType, v, err)
	}
	return
}

func parseListMode(s string) (ListMode, error) {
	if strings.TrimSpace(s) == "" {
		return ListMode(0), nil
	}
	modes := strings.Split(s, "|")
	var l ListMode
	for _, mode := range modes {
		mode = strings.TrimSpace(mode)
		switch mode {
		case ListModeBlock.String():
			l |= ListModeBlock
		case ListModeDir.String():
			l |= ListModeDir
		case ListModePart.String():
			l |= ListModePart
		case ListModePrefix.String():
			l |= ListModePrefix
		default:
			return ListMode(0), errors.New("invalid mode " + mode)
		}
	}
	return l, nil
}
