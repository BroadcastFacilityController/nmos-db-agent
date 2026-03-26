package schema

import (
	"github.com/guregu/null/v6"
	"github.com/guregu/null/v6/zero"
)

func strPtrOrNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func strNilPtrOrNil(str null.String) *string {
	if !str.Valid {
		return nil
	}
	return strPtrOrNil(str.ValueOrZero())
}

func boolZeroPtrOrNil(b zero.Bool) *bool {
	if !b.Valid {
		return nil
	}
	return &b.Bool
}
