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

// Package rails provides a repository of strings used throughout the application
package rails

type Data struct {
	Client    Client
	Console   Console
	Consensus Consensus
	Miner     Miner
}

type Client struct {
	Copyright   string
	Name        string
	Usage       string
	InfoNetwork InfoNetwork
}
type Console struct {
	Welcome string
}
type Consensus struct {
}
type Miner struct {
	Name string
}
type InfoNetwork struct {
	Mainnet string
}

const (
	Name = "grails"
)

var (
	Strings = &Data{
		Client: Client{
			//Copyright: "Copyright 2013-2022 The go-ethereum Authors",
			Copyright: "Copyright 2021-2022 The go-rails Authors",
			//Name: "geth",
			Name: Name,
			//Usage: "the go-ethereum command line interface",
			Usage: "the go-rails command line interface",
			InfoNetwork: InfoNetwork{
				//Mainnet: "Starting Geth on Ethereum mainnet...",
				Mainnet: "Starting Grails on Rails mainnet...",
			},
		},
		Console: Console{
			//Welcome: "Welcome to the Geth JavaScript console!\n\n",
			Welcome: "Welcome to the Grails JavaScript console!\n\n",
		},
		Consensus: Consensus{},
		Miner: Miner{
			Name: Name,
		},
	}
)
