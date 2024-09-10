package parse

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

type Tag string

const (
	TagJson Tag = "json"
)

var tagRegex = regexp.MustCompile(`^(\w+):"(.*)"`)

type TagMap map[Tag]string

func (m TagMap) JSON() string {
	return m[TagJson]
}

func parseTags(field *ast.Field) (TagMap, error) {
	if field.Tag == nil {
		return nil, nil
	}

	result := make(TagMap)
	tags := strings.Split(strings.Trim(field.Tag.Value, "`"), " ")

	for _, tag := range tags {
		matches := tagRegex.FindStringSubmatch(tag)
		if len(matches) != 3 {
			return nil, fmt.Errorf("invalid tag: %s", tag)
		}
		key := Tag(matches[1])
		if _, exists := result[key]; exists {
			return nil, fmt.Errorf("duplicate tag: %s", key)
		}
		result[key] = matches[2]
	}

	return result, nil
}
