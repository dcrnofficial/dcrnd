package chaincfg

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed subsidy.txt
var subsidy string

func init() {
	sc := bufio.NewScanner(strings.NewReader(subsidy))
	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		elems := strings.Split(line, ",")
		if amount, err := strconv.Atoi(elems[1]); err == nil {
			mainnetAirdropPayouts = append(mainnetAirdropPayouts, TokenPayout{
				ScriptVersion: 0,
				Script:        hexDecode(elems[0]),
				Amount:        int64(amount),
			})
		} else {
			panic("failed to read subsidy " + err.Error())
		}
	}
}

var mainnetAirdropPayouts []TokenPayout

func tokenPayouts_MainNetParams() []TokenPayout {
	return mainnetAirdropPayouts
}

func tokenPayouts_TestNet3Params() []TokenPayout {
	return mainnetAirdropPayouts
}

func tokenPayouts_SimNetParams() []TokenPayout {
	return mainnetAirdropPayouts
}

func tokenPayouts_RegNetParams() []TokenPayout {
	return mainnetAirdropPayouts
}
