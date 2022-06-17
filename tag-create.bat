set version=v8.0.5
git tag -a addrmgr/%version% -m "%version%"
git tag -a bech32/%version% -m "%version%"
git tag -a blockchain/stake/%version% -m "%version%"
git tag -a blockchain/standalone/%version% -m "%version%"
git tag -a blockchain/%version% -m "%version%"
git tag -a chaincfg/%version% -m "%version%"
git tag -a chaincfg/chainhash/%version% -m "%version%"
git tag -a certgen/%version% -m "%version%"
git tag -a connmgr/%version% -m "%version%"
git tag -a txscript/version4/%version% -m "%version%"
git tag -a txscript/version2/%version% -m "%version%"
git tag -a rpcclient/%version% -m "%version%"
git tag -a wire/%version% -m "%version%"
git tag -a database/%version% -m "%version%"
git tag -a dcrec/%version% -m "%version%"
git tag -a dcrjson/%version% -m "%version%"
git tag -a dcrutil/version2/%version% -m "%version%"
git tag -a fees/%version% -m "%version%"
git tag -a gcs/version2/%version% -m "%version%"
git tag -a gcs/version1/%version% -m "%version%"
git tag -a dcrec/edwards/%version% -m "%version%"
git tag -a dcrec/secp256k1/version2/%version% -m "%version%"
git tag -a dcrec/secp256k1/version4/%version% -m "%version%"
git tag -a hdkeychain/version2/%version% -m "%version%"
git tag -a hdkeychain/version3/%version% -m "%version%"
git tag -a lru/%version% -m "%version%"
git tag -a mempool/%version% -m "%version%"
git tag -a peer/%version% -m "%version%"
git tag -a rpc/jsonrpc/types/version2/%version% -m "%version%"
git tag -a rpc/jsonrpc/types/version1/%version% -m "%version%"
git tag -a crypto/blake256/%version% -m "%version%"
git tag -a crypto/ripemd160/%version% -m "%version%"
git tag -a mining/%version% -m "%version%"
git push origin --tags