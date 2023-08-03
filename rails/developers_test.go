// Copyright 2022 The go-rails Authors
// This file is part of the go-rails library.
// author: 0f0crypto <00ff00crypto@gmail.com>
//
//  ██████╗  ██████╗     ██████╗  █████╗ ██╗██╗     ███████╗
// ██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██║██║     ██╔════╝
// ██║  ███╗██║   ██║    ██████╔╝███████║██║██║     ███████╗
// ██║   ██║██║   ██║    ██╔══██╗██╔══██║██║██║     ╚════██║
// ╚██████╔╝╚██████╔╝    ██║  ██║██║  ██║██║███████╗███████║
//  ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚══════╝╚══════╝
//

package rails

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rails/contracts/developers"
	"math/big"
	"os"
	"strings"
	"testing"
)

var (
	genesisFile = "./testdata/genesis.json"
)

func TestDevelopersContract(t *testing.T) {

	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal("error generating private key")
	}
	address := crypto.PubkeyToAddress(key.PublicKey)

	client := getSimulatedClient(t, address)
	defer client.Close()

	instance, err := developers.NewContract(ethash.DevelopersAddress, client)
	if err != nil {
		t.Fatalf("error initializing the instance %v", err.Error())
	}

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(getNonce(t, client, address)))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice, _ = client.SuggestGasPrice(context.Background())

	_, err = instance.Initialize(auth, address)
	if err != nil {
		t.Fatalf("error initializing the smart contract %v", err.Error())
	}
	client.Commit()

	stateDb, err := client.Blockchain().State()
	if err != nil {
		t.Fatal("error getting the state db")
	}
	if ethash.IsDeveloperVerificationEnabled(stateDb) {
		t.Fatal("developer verification should be disabled by default")
	}

	auth.Nonce = big.NewInt(int64(getNonce(t, client, address)))
	_, err = instance.SetEnabled(auth, true)
	if err != nil {
		t.Fatalf("error enabling developers authorization %v", err.Error())
	}
	client.Commit()

	currentBlock := client.Blockchain().CurrentBlock()
	// this is based on the developers contract storage slot[0] layout, with the last two bytes specifying
	// the 'enabled' and 'initialized' boolean variables
	expectedSlot0Hash := fmt.Sprintf("0x00000000000000000000%s0101", strings.ToLower(address.String()[2:]))
	slot0, err := client.StorageAt(context.Background(), ethash.DevelopersAddress, common.Hash{}, currentBlock.Number())
	if err != nil {
		t.Fatalf("error reading the developers storage %v", err.Error())
	}
	actualSlot0Hash := common.BytesToHash(slot0).String()
	if expectedSlot0Hash != actualSlot0Hash {
		t.Fatalf("error at slot0 storage, expected %s but found %s", expectedSlot0Hash, actualSlot0Hash)
	}

	developerAddress := common.HexToAddress("0xC006c293676493759EBe7060b53c0E100DF451CE")

	auth.Nonce = big.NewInt(int64(getNonce(t, client, address)))
	_, err = instance.Approve(auth, developerAddress)
	if err != nil {
		t.Fatalf("error approving the developer %v", err.Error())
	}

	client.Commit()

	result, err := instance.IsApproved(nil, developerAddress)
	if result == false {
		t.Fatalf("error checking developer's authorization")
	}

	count, err := instance.GetCount(nil)
	if count.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("error counting authorized developers")
	}

}

// getSimulatedClient load the genesis file, add the address to the allocations
// and returns a simulated backend
func getSimulatedClient(t *testing.T, address common.Address) *backends.SimulatedBackend {
	genesis, err := loadGenesis(genesisFile)
	if err != nil {
		t.Fatalf("error loading the genesis file %v", err.Error())
	}
	genesisAlloc := genesis.Alloc
	genesisAlloc[address] = core.GenesisAccount{Balance: big.NewInt(6555444333222111000)}

	var gasLimit uint64 = 8000029
	return backends.NewSimulatedBackend(genesis.Alloc, gasLimit)
}

// loadGenesis taken from ./cmd/devp2p/internal/ethtest/chain.go
func loadGenesis(genesisFile string) (*core.Genesis, error) {
	chainConfig, err := os.ReadFile(genesisFile)
	if err != nil {
		return &core.Genesis{}, err
	}
	var genesis core.Genesis
	if err := json.Unmarshal(chainConfig, &genesis); err != nil {
		return &core.Genesis{}, err
	}
	return &genesis, nil
}

func getNonce(t *testing.T, client *backends.SimulatedBackend, address common.Address) uint64 {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		t.Fatalf("error getting the nonce of %v: %v", address, err.Error())
	}
	return nonce
}
