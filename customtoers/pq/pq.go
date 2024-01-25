package custompq

import (
	"github.com/lib/pq"
	"github.com/ntnn/dataparse"
)

func init() {
	dataparse.AddCustomToFunc(CustomTo)
}

func CustomTo(v dataparse.Value, other any) (any, bool, error) {
	switch other.(type) {
	case pq.StringArray:
		newValue, err := v.ListString()
		return newValue, true, err
	default:
		return nil, false, nil
	}
}
