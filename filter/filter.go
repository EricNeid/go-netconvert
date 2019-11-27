package filter

import (
	"fmt"
	"strings"

	"github.com/EricNeid/go-netconvert/osm"
)

// ConditionNode defines a filter which is applied
// to each node during the filtering process.
type ConditionNode = func(node osm.Node) bool

// ConditionWay defines a filter which is applied
// to each way during the filtering process.
type ConditionWay = func(way osm.Way) bool

// Nodes applies the given condition on all nodes and
// returns slice of nodes, satisfying the condition.
func Nodes(nodes []osm.Node, filter ConditionNode) []osm.Node {
	var result []osm.Node
	for _, n := range nodes {
		if filter(n) {
			result = append(result, n)
		}
	}
	return result
}

// Ways applies the given condition on all ways and
// returns slice of ways, satisfying the condition.
func Ways(ways []osm.Way, filter ConditionWay) []osm.Way {
	var result []osm.Way
	for _, w := range ways {
		if filter(w) {
			result = append(result, w)
		}
	}
	return result
}

// Operand is a operator which is applied on a specific tag to
// check if the value of a tag has  specific value.
type Operand string

const (
	// EQ is equivalent to = for int types or string equals
	EQ = "="
	// LT is equivalent to < for int types and noop for strings
	LT = "<"
	// GT is equivalent to > for int types and noop for strings
	GT = ">"
	// NOP means ignore the value and only check for matching name
	NOP = ""
)

// Filter represents a single filter criteria for a tag.
type Filter struct {
	Name    string
	Value   string
	Operand Operand
}

// ToFilter parses given string in slice of Filter commands.
// expected format of input string is:
// => a>4,b=6,c
// valid operators are:
// =,<,> and no operator
func ToFilter(filter string) ([]Filter, error) {
	var result []Filter

	tagStatements := strings.Split(filter, ",")
	for _, stm := range tagStatements {
		var operand Operand

		if strings.Contains(stm, EQ) {
			operand = EQ
		} else if strings.Contains(stm, LT) {
			operand = LT
		} else if strings.Contains(stm, GT) {
			operand = GT
		} else {
			operand = NOP
		}

		if operand == NOP {
			// just filter by tag and no operand or value
			result = append(result, Filter{
				Name:    strings.TrimSpace(stm),
				Value:   "",
				Operand: operand,
			})
		} else {
			// parse statement into key and value
			name, value, err := getKeyValue(stm, operand)
			if err != nil {
				return result, err
			}
			result = append(result, Filter{
				Name:    name,
				Value:   value,
				Operand: operand,
			})
		}
	}

	return result, nil
}

func getKeyValue(statement string, operand Operand) (string, string, error) {
	var key string
	var value string
	keyValue := strings.Split(statement, string(operand))

	if len(keyValue) != 2 {
		return key, value, fmt.Errorf("Invalid statement found. Expected format to be a%sb, got %s", operand, statement)
	}

	key = keyValue[0]
	value = keyValue[1]

	return strings.TrimSpace(key), strings.TrimSpace(value), nil
}
