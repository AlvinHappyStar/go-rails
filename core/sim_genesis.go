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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/vm"
	"math/big"
)

func GetSimulatedGenesisAccountFromAlloc(alloc GenesisAlloc) GenesisAccount {

	addresses := make([]common.Address, 0, len(alloc))
	for addr := range alloc {
		addresses = append(addresses, addr)
	}
	return GetSimulatedGenesisAccount(addresses)
}

func GetSimulatedGenesisAccount(addresses []common.Address) GenesisAccount {

	enabledHash := common.HexToHash("0x1")

	account := GenesisAccount{
		// Code here is irrelevant so just have a simple self-destruct contract
		Code:    []byte{byte(vm.PC), byte(vm.SELFDESTRUCT)},
		Balance: big.NewInt(0),
		Storage: map[common.Hash]common.Hash{},
	}

	for i := range addresses {
		account.Storage[ethash.GetDevelopersStorageSlot(addresses[i])] = enabledHash
	}

	return account
}
