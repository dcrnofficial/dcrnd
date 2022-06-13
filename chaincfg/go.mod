module github.com/Decred-Next/dcrnd/chaincfg/v8

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/Decred-Next/dcrnd/chaincfg/chainhash/v8 v8.0.0
	github.com/Decred-Next/dcrnd/dcrec/edwards/v8 v8.0.0
	github.com/Decred-Next/dcrnd/dcrec/secp256k1/v8 v8.0.0
	github.com/Decred-Next/dcrnd/wire v1.3.0
)
replace (
	github.com/Decred-Next/dcrnd/wire v1.3.0 => ../wire
)
