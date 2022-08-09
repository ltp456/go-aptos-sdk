package go_apots_sdk

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type Params map[string]interface{}

func (p Params) SetValue(key string, value interface{}) {
	p[key] = value
}

func (p Params) Encode() string {
	if p == nil {
		return ""
	}

	var buf strings.Builder
	keys := make([]string, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		vs := p[k]
		keyEscaped := url.QueryEscape(k)
		if i == len(keys)-1 {
			buf.WriteString(fmt.Sprintf("%v=%v", keyEscaped, vs))
		} else {
			buf.WriteString(fmt.Sprintf("%v=%v&", keyEscaped, vs))
		}

	}
	return buf.String()
}

type Value struct {
	Key   string
	Value interface{}
}

type Values []Value

func (ov *Values) Add(key string, value interface{}) {
	*ov = append(*ov, Value{key, value})
}
