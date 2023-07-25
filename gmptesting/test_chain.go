package gmp_testing

import (
	"encoding/json"
	"testing"

	ibctesting "github.com/cosmos/ibc-go/v4/testing"
	"github.com/jackalLabs/canine-chain/v3/app"
)

type TestChain struct {
	*ibctesting.TestChain
}

func SetupTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	var t *testing.T
	jackalApp := app.SetupWithEmptyStore(t)
	return jackalApp, app.NewDefaultGenesisState()
}

// GetJackalApp returns the current chain's app as a JackalApp
func (chain *TestChain) GetJackalApp() *app.JackalApp {
	v, _ := chain.App.(*app.JackalApp)
	return v
}
