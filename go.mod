module github.com/Decred-Next/dcrnd

go 1.11

require (
	github.com/Decred-Next/base58/v8 v8.0.11
	github.com/Decred-Next/dcrnd/addrmgr/v8 v8.0.11
	github.com/Decred-Next/dcrnd/blockchain/stake/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/blockchain/standalone/v8 v8.0.11
	github.com/Decred-Next/dcrnd/blockchain/v8 v8.0.11
	github.com/Decred-Next/dcrnd/certgen/v8 v8.0.11
	github.com/Decred-Next/dcrnd/chaincfg/chainhash/v8 v8.0.11
	github.com/Decred-Next/dcrnd/chaincfg/v8 v8.0.11
	github.com/Decred-Next/dcrnd/connmgr/v8 v8.0.11
	github.com/Decred-Next/dcrnd/crypto/ripemd160/v8 v8.0.11
	github.com/Decred-Next/dcrnd/database/v8 v8.0.11
	github.com/Decred-Next/dcrnd/dcrec/secp256k1/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/dcrec/v8 v8.0.11
	github.com/Decred-Next/dcrnd/dcrjson/version3/v8 v8.0.11
	github.com/Decred-Next/dcrnd/dcrutil/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/fees/v8 v8.0.11
	github.com/Decred-Next/dcrnd/gcs/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/hdkeychain/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/lru/v8 v8.0.11
	github.com/Decred-Next/dcrnd/mempool/v8 v8.0.11
	github.com/Decred-Next/dcrnd/mining/v8 v8.0.11
	github.com/Decred-Next/dcrnd/peer/v8 v8.0.11
	github.com/Decred-Next/dcrnd/rpc/jsonrpc/types/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/rpcclient/version5/v8 v8.0.11
	github.com/Decred-Next/dcrnd/txscript/version2/v8 v8.0.11
	github.com/Decred-Next/dcrnd/wire/v8 v8.0.11
	github.com/Decred-Next/dcrnwallet/rpc/jsonrpc/types/version14/v8 v8.0.11
	github.com/Decred-Next/go-socks/v8 v8.0.1
	github.com/Decred-Next/slog/v8 v8.0.1
	github.com/btcsuite/winsvc v1.0.0
	github.com/decred/base58 v1.0.4 // indirect
	github.com/gorilla/websocket v1.5.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/jrick/bitset v1.0.0
	github.com/jrick/logrotate v1.0.0
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f

)

replace (
	github.com/Decred-Next/dcrnd/addrmgr/v8 => ./addrmgr
	github.com/Decred-Next/dcrnd/blockchain/stake/version2/v8 => ./blockchain/stake/version2
	github.com/Decred-Next/dcrnd/blockchain/standalone/v8 => ./blockchain/standalone
	github.com/Decred-Next/dcrnd/blockchain/v8 => ./blockchain
	github.com/Decred-Next/dcrnd/certgen/v8 => ./certgen
	github.com/Decred-Next/dcrnd/chaincfg/chainhash/v8 => ./chaincfg/chainhash
	github.com/Decred-Next/dcrnd/chaincfg/v8 => ./chaincfg
	github.com/Decred-Next/dcrnd/connmgr/v8 => ./connmgr
	github.com/Decred-Next/dcrnd/crypto/ripemd160/v8 => ./crypto/ripemd160
	github.com/Decred-Next/dcrnd/database/v8 => ./database
	github.com/Decred-Next/dcrnd/dcrec/secp256k1/version2/v8 => ./dcrec/secp256k1/version2
	github.com/Decred-Next/dcrnd/dcrec/v8 => ./dcrec
	github.com/Decred-Next/dcrnd/dcrjson/version3/v8 => ./dcrjson/version3
	github.com/Decred-Next/dcrnd/dcrutil/version2/v8 => ./dcrutil/version2
	github.com/Decred-Next/dcrnd/fees/v8 => ./fees
	github.com/Decred-Next/dcrnd/gcs/version2/v8 => ./gcs/version2
	github.com/Decred-Next/dcrnd/hdkeychain/version2/v8 => ./hdkeychain/version2
	github.com/Decred-Next/dcrnd/lru/v8 => ./lru
	github.com/Decred-Next/dcrnd/mempool/v8 => ./mempool
	github.com/Decred-Next/dcrnd/mining/v8 => ./mining
	github.com/Decred-Next/dcrnd/peer/v8 => ./peer
	github.com/Decred-Next/dcrnd/rpc/jsonrpc/types/version2/v8 => ./rpc/jsonrpc/types/version2
	github.com/Decred-Next/dcrnd/rpcclient/version5/v8 => ./rpcclient/version5
	github.com/Decred-Next/dcrnd/txscript/version2/v8 => ./txscript/version2
	github.com/Decred-Next/dcrnd/wire/v8 => ./wire
)
