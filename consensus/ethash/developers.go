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

package ethash

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	DevelopersMappingSlot = 2
)

var (
	DevelopersAddress = common.HexToAddress("0x0000000000000000000000000000000000001002")
)

// CanDeploy checks whether the specified address (wallet) is approved for deploying smart
// contracts to the blockchain
func (ethash *Ethash) CanDeploy(state consensus.StateReader, addr common.Address) bool {

	// TODO (0f0crypto)
	if !IsDeveloperVerificationEnabled(state) {
		return true
	}

	// TODO (0f0crypto)
	// check if the addr is not the backup admin, e.g.
	// if addr.Hash == ethash.config.BackupAdminHash { return true }

	slot := GetDevelopersStorageSlot(addr)
	valueHash := state.GetState(DevelopersAddress, slot)

	return valueHash.Big().Sign() > 0

}

func GetDevelopersStorageSlot(addr common.Address) common.Hash {

	position := make([]byte, common.HashLength)

	// write the slot number (uint16 = 0 - 65535) into the last two indexes in the position array
	// e.g. slot = 2 becomes [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2]
	binary.BigEndian.PutUint16(position[common.HashLength-2:], uint16(DevelopersMappingSlot))

	// hash the input address with the position, solidity equivalent of web3.sha3(address + position, {"encoding":"hex"})
	// where both address and position strings are left-padded with 0s to 64 chars
	return crypto.Keccak256Hash(addr.Hash().Bytes(), position)
}

func IsDeveloperVerificationEnabled(state consensus.StateReader) bool {

	slot0 := state.GetState(DevelopersAddress, common.Hash{})
	enabledByte := slot0.Bytes()[common.HashLength-2]
	return enabledByte == 0x01
}
