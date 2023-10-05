package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var KeyMsgSubmitTxMaxMessages = []byte("MsgSubmitTxMaxMessages")

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		MsgSubmitTxMaxMessages: 16,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMsgSubmitTxMaxMessages, &p.MsgSubmitTxMaxMessages, validateMsgSubmitTxMaxMessages),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validateMsgSubmitTxMaxMessages(p.GetMsgSubmitTxMaxMessages())
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateMsgSubmitTxMaxMessages(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("MsgSubmitTxMaxMessages must be greater than zero")
	}

	return nil
}
