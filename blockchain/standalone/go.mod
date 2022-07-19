module github.com/Decred-Next/dcrnd/blockchain/standalone/v8

go 1.11

require (
	github.com/Decred-Next/dcrnd/chaincfg/chainhash/v8 v8.0.12
	github.com/Decred-Next/dcrnd/wire/v8 v8.0.12
)

replace(
	github.com/Decred-Next/dcrnd/wire/v8 => ../../wire
	github.com/Decred-Next/dcrnd/chaincfg/chainhash/v8 => ../../chaincfg/chainhash
)