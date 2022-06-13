// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"math/big"
	"time"

	"github.com/Decred-Next/dcrnd/chaincfg/chainhash/v8"
	"github.com/Decred-Next/dcrnd/wire/v8"
)

// MainNetParams returns the network parameters for the main Decred network.
func MainNetParams() *Params {
	// mainPowLimit is the highest proof of work value a Decred block can have
	// for the main network.  It is the value 2^224 - 1.
	mainPowLimit := new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)

	// genesisBlock defines the genesis block of the block chain which serves as
	// the public transaction ledger for the main network.
	//
	// The genesis block for Decred mainnet, testnet, and simnet are not
	// evaluated for proof of work. The only values that are ever used elsewhere
	// in the blockchain from it are:
	// (1) The genesis block hash is used as the PrevBlock.
	// (2) The difficulty starts off at the value given by Bits.
	// (3) The stake difficulty starts off at the value given by SBits.
	// (4) The timestamp, which guides when blocks can be built on top of it
	//      and what the initial difficulty calculations come out to be.
	//
	// The genesis block is valid by definition and none of the fields within it
	// are validated for correctness.
	genesisBlock := wire.MsgBlock{
		Header: wire.BlockHeader{
			Version:   1,
			PrevBlock: chainhash.Hash{}, // All zero.
			// MerkleRoot: Calculated below.
			StakeRoot:    chainhash.Hash{},
			Timestamp:    time.Unix(1653386400, 0), // 2022/05/24 10:00:00 GMT
			Bits:         0x1b000a00,               // Difficulty 1677696
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
		Name:        "mainnet",
		Net:         wire.MainNet,
		DefaultPort: "9108",
		DNSSeeds: []DNSSeed{
			{"mainnet1.dcrn.xyz", true},
			{"mainnet2.dcrn.xyz", true},
			{"mainnet3.dcrn.xyz", true},
			{"mainnet4.dcrn.xyz", true},
			{"mainnet5.dcrn.xyz", true},
			{"mainnet6.dcrn.xyz", true},
		},

		// Chain parameters
		GenesisBlock:             &genesisBlock,
		GenesisHash:              genesisBlock.BlockHash(),
		PowLimit:                 mainPowLimit,
		PowLimitBits:             0x1d00ffff,
		ReduceMinDifficulty:      false,
		MinDiffReductionTime:     0, // Does not apply since ReduceMinDifficulty false
		GenerateSupported:        false,
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
		NetworkAddressPrefix: "D",
		PubKeyAddrID:         [2]byte{0x13, 0x86}, // starts with Dk
		PubKeyHashAddrID:     [2]byte{0x07, 0x3f}, // starts with Ds
		PKHEdwardsAddrID:     [2]byte{0x07, 0x1f}, // starts with De
		PKHSchnorrAddrID:     [2]byte{0x07, 0x01}, // starts with DS
		ScriptHashAddrID:     [2]byte{0x07, 0x1a}, // starts with Dc
		PrivateKeyID:         [2]byte{0x22, 0xde}, // starts with Pm

		// BIP32 hierarchical deterministic extended key magics
		HDPrivateKeyID: [4]byte{0x02, 0xfd, 0xa4, 0xe8}, // starts with dprv
		HDPublicKeyID:  [4]byte{0x02, 0xfd, 0xa9, 0x26}, // starts with dpub

		// BIP44 coin type used in the hierarchical deterministic path for
		// address generation.
		SLIP0044CoinType: 42, // SLIP0044, Decred
		LegacyCoinType:   20, // for backwards compatibility

		// Decred PoS parameters
		MinimumStakeDiff:        2 * 1e8, // 2 Coin
		TicketPoolSize:          8192,
		TicketsPerBlock:         5,
		TicketMaturity:          256,
		TicketExpiry:            40960, // 5*TicketPoolSize
		CoinbaseMaturity:        256,
		SStxChangeMaturity:      1,
		TicketPoolSizeWeight:    4,
		StakeDiffAlpha:          1, // Minimal
		StakeDiffWindowSize:     144,
		StakeDiffWindows:        20,
		StakeVersionInterval:    144 * 2 * 7, // ~1 week
		MaxFreshStakePerBlock:   20,          // 4*TicketsPerBlock
		StakeEnabledHeight:      256 + 256,   // CoinbaseMaturity + TicketMaturity
		StakeValidationHeight:   2048,        // ~7 days
		StakeBaseSigScript:      []byte{0x00, 0x00},
		StakeMajorityMultiplier: 3,
		StakeMajorityDivisor:    4,

		// Decred organization related parameters
		// Organization address is DsUnEWbLXrqMWTEMkEGsUC9R1mP6p3FDCRN
		OrganizationPkScript:        hexDecode("76a91429dbcdffde4bf456b9df612b0916cc331ad4720288ac"),
		OrganizationPkScriptVersion: 0,
		BlockOneLedger:              tokenPayouts_MainNetParams(),
		AirdropBlockOffset:          4096,
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
