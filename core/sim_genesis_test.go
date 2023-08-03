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

package core

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"math/big"
	"testing"
)

func TestUnauthorizedDeveloperErrorMessage(t *testing.T) {
	address := "0xd25A13529908f162E7e3Abf820BaDe73d03B87a0"
	from := common.HexToAddress(address)
	actual := fmt.Sprintf("%s: sender %s", ErrUnauthorizedDeveloper, from.Hex())
	expected := "unauthorized developer: sender 0xd25A13529908f162E7e3Abf820BaDe73d03B87a0"
	if actual != expected {
		t.Error("expect error message formatting to be correct")
	}
}

func TestGetDevelopersSimulatedGenesisAccountFromAlloc(t *testing.T) {
	address := common.HexToAddress("0xd25A13529908f162E7e3Abf820BaDe73d03B87a0")

	alloc := GenesisAlloc{
		address: {Balance: big.NewInt(10000000000000000)},
	}
	slot := ethash.GetDevelopersStorageSlot(address)
	account := GetSimulatedGenesisAccountFromAlloc(alloc)

	if account.Storage == nil {
		t.Errorf("expect account.Storage not to be nil")
	}

	value := account.Storage[slot]
	if value != common.HexToHash("0x1") {
		t.Errorf("expect account storage value to be 0x1, but found: %v", value)
	}
}
