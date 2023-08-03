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
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

var (
	testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

	// testAddr is the Ethereum address of the tester account.
	// 0x027d980Ba4cAfe888C5Cb93c7BCdc259E61cd252
	testAddr = crypto.PubkeyToAddress(testKey.PublicKey)
)

func TestGetDeveloperStorageSlot(t *testing.T) {

	actual := GetDevelopersStorageSlot(testAddr).String()

	expected := "0x1528d588dcc99fc10a9e28e0b813aa11f9b871a66f2cceda2695cd13cddf0463"
	if actual != expected {
		t.Errorf("expect developer storage slot for %v to be %v but %v was calculated", testAddr, expected, actual)
	}
}
