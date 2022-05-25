// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"math/big"
	"time"

	"github.com/decred/dcrd/chaincfg/chainhash"
	"github.com/decred/dcrd/wire"
)

// TestNet3Params return the network parameters for the test currency network.
// This network is sometimes simply called "testnet".
// This is the third public iteration of testnet.
func TestNet3Params() *Params {
	// testNetPowLimit is the highest proof of work value a Decred block
	// can have for the test network.  It is the value 2^232 - 1.
	testNetPowLimit := new(big.Int).Sub(new(big.Int).Lsh(bigOne, 232), bigOne)

	// genesisBlock defines the genesis block of the block chain which serves as
	// the public transaction ledger for the test network (version 3).
	genesisBlock := wire.MsgBlock{
		Header: wire.BlockHeader{
			Version:   1,
			PrevBlock: chainhash.Hash{}, // All zero.
			// MerkleRoot: Calculated below.
			StakeRoot:    chainhash.Hash{},
			Timestamp:    time.Unix(1653386400, 0), // 2022/05/24 10:00:00 GMT
			Bits:         0x1e00ffff,               // Difficulty 1
			SBits:        2 * 1e8,                  // 2 Coin
			Nonce:        0x00000000,
			StakeVersion: 0,
		},
		Transactions: []*wire.MsgTx{{
			SerType: wire.TxSerializeFull,
			Version: 1,
			TxIn: []*wire.TxIn{{
				// Fully null.
				PreviousOutPoint: wire.OutPoint{
					Hash:  chainhash.Hash{},
					Index: 0xffffffff,
					Tree:  0,
				},
				SignatureScript: hexDecode("0000"),
				Sequence:        0xffffffff,
				BlockHeight:     wire.NullBlockHeight,
				BlockIndex:      wire.NullBlockIndex,
				ValueIn:         wire.NullValueIn,
			}},
			TxOut: []*wire.TxOut{{
				Version:  0x0000,
				Value:    0x00000000,
				PkScript: hexDecode("80165584be68cb981cfc0b5eeee0e3cf6738de31d3970505"),
			}},
			LockTime: 0,
			Expiry:   0,
		}},
	}
	genesisBlock.Header.MerkleRoot = genesisBlock.Transactions[0].TxHashFull()

	return &Params{
		Name:        "testnet3",
		Net:         wire.TestNet3,
		DefaultPort: "19108",
		DNSSeeds: []DNSSeed{
			{"testnet1.dcrn.xyz", true},
			{"testnet2.dcrn.xyz", true},
			{"testnet3.dcrn.xyz", true},
			{"testnet4.dcrn.xyz", true},
		},

		// Chain parameters
		GenesisBlock:             &genesisBlock,
		GenesisHash:              genesisBlock.BlockHash(),
		PowLimit:                 testNetPowLimit,
		PowLimitBits:             0x1e00ffff,
		ReduceMinDifficulty:      false,
		MinDiffReductionTime:     0,
		GenerateSupported:        true,
		MaximumBlockSizes:        []int{393216},
		MaxTxSize:                393216,
		TargetTimePerBlock:       time.Minute * 5,
		WorkDiffAlpha:            1,
		WorkDiffWindowSize:       144,
		WorkDiffWindows:          20,
		TargetTimespan:           time.Minute * 5 * 144, // TimePerBlock * WindowSize
		RetargetAdjustmentFactor: 4,

		// Subsidy parameters.
		BaseSubsidy:              10 * 1e8,
		MulSubsidy:               100,
		DivSubsidy:               101,
		SubsidyReductionInterval: 6144,
		WorkRewardProportion:     6,
		StakeRewardProportion:    3,
		BlockTaxProportion:       1,

		// Checkpoints ordered from oldest to newest.
		Checkpoints: []Checkpoint{},

		// Consensus rule change deployments.
		//
		// The miner confirmation window is defined as:
		//   target proof of work timespan / target proof of work spacing
		RuleChangeActivationQuorum:     4032, // 10 % of RuleChangeActivationInterval * TicketsPerBlock
		RuleChangeActivationMultiplier: 3,    // 75%
		RuleChangeActivationDivisor:    4,
		RuleChangeActivationInterval:   2016 * 4, // 4 weeks
		Deployments:                    map[uint32][]ConsensusDeployment{},

		// Enforce current block version once majority of the network has
		// upgraded.
		// 75% (750 / 1000)
		//
		// Reject previous block versions once a majority of the network has
		// upgraded.
		// 95% (950 / 1000)
		BlockEnforceNumRequired: 750,
		BlockRejectNumRequired:  950,
		BlockUpgradeNumToCheck:  1000,

		// AcceptNonStdTxs is a mempool param to either accept and relay non
		// standard txs to the network or reject them
		AcceptNonStdTxs: false,

		// Address encoding magics
		NetworkAddressPrefix: "T",
		PubKeyAddrID:         [2]byte{0x28, 0xf7}, // starts with Tk
		PubKeyHashAddrID:     [2]byte{0x0f, 0x21}, // starts with Ts
		PKHEdwardsAddrID:     [2]byte{0x0f, 0x01}, // starts with Te
		PKHSchnorrAddrID:     [2]byte{0x0e, 0xe3}, // starts with TS
		ScriptHashAddrID:     [2]byte{0x0e, 0xfc}, // starts with Tc
		PrivateKeyID:         [2]byte{0x23, 0x0e}, // starts with Pt

		// BIP32 hierarchical deterministic extended key magics
		HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x97}, // starts with tprv
		HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xd1}, // starts with tpub

		// BIP44 coin type used in the hierarchical deterministic path for
		// address generation.
		SLIP0044CoinType: 1,  // SLIP0044, Testnet (all coins)
		LegacyCoinType:   11, // for backwards compatibility

		// Decred PoS parameters
		MinimumStakeDiff:        2 * 1e8, // 2 Coin
		TicketPoolSize:          8192,
		TicketsPerBlock:         5,
		TicketMaturity:          16,
		TicketExpiry:            40960, // 5*TicketPoolSize
		CoinbaseMaturity:        16,
		SStxChangeMaturity:      1,
		TicketPoolSizeWeight:    4,
		StakeDiffAlpha:          1, // Minimal
		StakeDiffWindowSize:     144,
		StakeDiffWindows:        20,
		StakeVersionInterval:    144 * 2 * 7, // ~1 week
		MaxFreshStakePerBlock:   20,          // 4*TicketsPerBlock
		StakeEnabledHeight:      16 + 16,     // CoinbaseMaturity + TicketMaturity
		StakeValidationHeight:   64,          // ~7 days
		StakeBaseSigScript:      []byte{0x00, 0x00},
		StakeMajorityMultiplier: 3,
		StakeMajorityDivisor:    4,

		// Decred organization related parameters.
		// Organization address is TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9
		OrganizationPkScript:        hexDecode("76a914f4f3e728ea60e5ce3fa4a32958925a67cf6c221c88ac"),
		OrganizationPkScriptVersion: 0,
		BlockOneLedger:              tokenPayouts_TestNet3Params(),
		AirdropBlockOffset:          16,
		DaoInitLedger: []TokenPayout{{
			ScriptVersion: 0,
			Script:        hexDecode("76a914be7e1c2739c62db0ab77399ddb76268801ad144288ac"),
			Amount:        100000 * 1e8,
		}, {
			ScriptVersion: 0,
			Script:        hexDecode("76a914dc53132f167ec93c608a13b33f1f1627a81ad0c588ac"),
			Amount:        100000 * 1e8,
		}, {
			ScriptVersion: 0,
			Script:        hexDecode("76a914dc9e9eea83db5244aee79f88a6c17a399e4c612088ac"),
			Amount:        100000 * 1e8,
		}, {
			ScriptVersion: 0,
			Script:        hexDecode("76a914a6b5056f2c9ed620f09b26cc61e5202eeb61169088ac"),
			Amount:        100000 * 1e8,
		}, {
			ScriptVersion: 0,
			Script:        hexDecode("76a914dec73de6806a2ce8aee9b2a2b5f75be689cf1b9888ac"),
			Amount:        100000 * 1e8,
		}, {
			ScriptVersion: 0,
			Script:        hexDecode("76a914531ea176a7164ada0563982d87ee2c3be40e1dff88ac"),
			Amount:        100000 * 1e8,
		}, {
			ScriptVersion: 0,
			Script:        hexDecode("76a91461dbfa0aa1099a2a51c1638f446eb417b19daf5a88ac"),
			Amount:        100000 * 1e8,
		}},
	}
}
