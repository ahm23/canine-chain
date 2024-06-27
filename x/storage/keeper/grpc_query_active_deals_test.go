package keeper_test

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	jklapp "github.com/jackalLabs/canine-chain/v4/app"
	"github.com/jackalLabs/canine-chain/v4/testutil"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/tendermint/libs/log"
)

func genApp(withGenesis bool, invCheckPeriod uint) (*jklapp.JackalApp, jklapp.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := jklapp.MakeEncodingConfig()
	app := jklapp.NewJackalApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		simapp.DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		jklapp.GetEnabledProposals(),
		simapp.EmptyAppOptions{},
		jklapp.GetWasmOpts(simapp.EmptyAppOptions{}),
	)

	if withGenesis {
		return app, jklapp.NewDefaultGenesisState()
	}

	return app, jklapp.GenesisState{}
}

func setup(isCheckTx bool) *jklapp.JackalApp {
	app, genesisState := genApp(!isCheckTx, 5)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func (suite *KeeperTestSuite) SetupTest() {
	app := setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, app.StorageKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.ctx = ctx

	suite.queryClient = queryClient
}

func (suite *KeeperTestSuite) TestAllFiles() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]
	depoAccount := testAddresses[1]

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	testAcc, _ := sdk.AccAddressFromBech32(testAccount)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAcc, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     8,
		CollateralPrice:        2,
		CheckWindow:            10,
	})

	merkle := []byte("merkle")

	suite.storageKeeper.SetFile(suite.ctx, types.UnifiedFile{
		Merkle:        merkle,
		Owner:         testAccount,
		Start:         0,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 400,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          "test",
	})

	pg := query.PageRequest{
		Offset:  0,
		Reverse: false,
	}

	res, err := suite.queryClient.AllFiles(context.Background(), &types.QueryAllFiles{
		Pagination: &pg,
	})
	suite.Require().NoError(err)

	suite.Require().Equal(1, len(res.Files))

	mres, err := suite.queryClient.AllFilesByMerkle(context.Background(), &types.QueryAllFilesByMerkle{
		Pagination: &pg,
		Merkle:     merkle,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(mres.Files))

	suite.reset()
}

func (suite *KeeperTestSuite) TestFileNotes() {
	suite.SetupSuite()

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]
	depoAccount := testAddresses[1]

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	testAcc, _ := sdk.AccAddressFromBech32(testAccount)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAcc, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     8,
		CollateralPrice:        2,
		CheckWindow:            10,
		PolRatio:               40,
		ReferralCommission:     25,
	})

	merkle := []byte("merkle")

	k := "myKey"
	v := "myValue"
	m := make(map[string]string)
	m[k] = v
	b, err := json.Marshal(m)
	suite.Require().NoError(err)

	suite.storageKeeper.SetFile(suite.ctx, types.UnifiedFile{
		Merkle:        merkle,
		Owner:         testAccount,
		Start:         0,
		Expires:       0,
		FileSize:      1024,
		ProofInterval: 400,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          string(b),
	})

	bk := "terribleKey"
	bv := sdk.NewDec(46)
	bm := make(map[string]sdk.Dec)
	bm[bk] = bv
	bb, err := json.Marshal(bm)
	suite.Require().NoError(err)

	bmerkle := []byte("badmerkle")

	suite.storageKeeper.SetFile(suite.ctx, types.UnifiedFile{
		Merkle:        bmerkle,
		Owner:         testAccount,
		Start:         0,
		Expires:       0,
		FileSize:      2048,
		ProofInterval: 400,
		ProofType:     0,
		Proofs:        make([]string, 0),
		MaxProofs:     3,
		Note:          string(bb),
	})

	pg := query.PageRequest{
		Offset:  0,
		Reverse: false,
	}

	res, err := suite.queryClient.AllFiles(context.Background(), &types.QueryAllFiles{
		Pagination: &pg,
	})
	suite.Require().NoError(err)

	suite.Require().Equal(2, len(res.Files))

	mres, err := suite.queryClient.AllFilesByMerkle(context.Background(), &types.QueryAllFilesByMerkle{
		Pagination: &pg,
		Merkle:     merkle,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(mres.Files))

	nres, err := suite.queryClient.FilesFromNote(context.Background(), &types.QueryFilesFromNote{
		Pagination: &pg,
		Key:        k,
		Value:      v,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(nres.Files))

	nres, err = suite.queryClient.FilesFromNote(context.Background(), &types.QueryFilesFromNote{
		Pagination: &pg,
		Key:        k,
		Value:      "bad-value",
	})
	suite.Require().NoError(err)
	suite.Require().Equal(0, len(nres.Files))

	nres, err = suite.queryClient.FilesFromNote(context.Background(), &types.QueryFilesFromNote{
		Pagination: &pg,
		Key:        "bad-key",
		Value:      v,
	})
	suite.Require().NoError(err)
	suite.Require().Equal(0, len(nres.Files))

	suite.reset()
}
