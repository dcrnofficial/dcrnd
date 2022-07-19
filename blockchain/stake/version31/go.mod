module github.com/Decred-Next/dcrnd/blockchain/stake/version31/v8

go 1.11

require (
	github.com/Decred-Next/dcrnd/chaincfg/chainhash/v8 v8.0.12
	github.com/Decred-Next/dcrnd/chaincfg/v8 v8.0.12
	github.com/Decred-Next/dcrnd/database/v8 v8.0.12
	github.com/Decred-Next/dcrnd/dcrec/v8 v8.0.12
	github.com/Decred-Next/dcrnd/dcrutil/version31/v8 v8.0.12
	github.com/Decred-Next/dcrnd/txscript/version31/v8 v8.0.12
	github.com/Decred-Next/dcrnd/wire/v8 v8.0.12
	github.com/Decred-Next/slog/v8 v8.0.1
)

replace(
	github.com/Decred-Next/dcrnd/wire/v8 => ../../../wire
	github.com/Decred-Next/dcrnd/chaincfg/chainhash/v8 => ../../../chaincfg/chainhash
	github.com/Decred-Next/dcrnd/chaincfg/v8 => ../../../chaincfg
)