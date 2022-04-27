package chaincfg

import (
	_ "embed"
	"bufio"
	"strconv"
	"strings"
)

//go:embed subsidy.txt
var subsidy string

var mainnetAirdropPayouts []TokenPayout

func tokenPayouts_MainNetParams() []TokenPayout {
	return mainnetAirdropPayouts
}

func init() {
	sc := bufio.NewScanner(strings.NewReader(subsidy))
	for sc.Scan() {
		line := sc.Text()
		elems := strings.Split(line,",")
		if amount,err := strconv.Atoi(elems[1]); err == nil {
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

func tokenPayouts_TestNet3Params() []TokenPayout {
	payout := TokenPayout{
		0, hexDecode("76a914cf646ab09da6133415ec2c629aacfc7c5b850cb788ac"), 100_0000 * 1e8,
	}
	payouts := append(mainnetAirdropPayouts, payout)
	return payouts
}

func tokenPayouts_SimNetParams() []TokenPayout {
	payout := TokenPayout{
		0, hexDecode("76a914cf646ab09da6133415ec2c629aacfc7c5b850cb788ac"), 100_0000 * 1e8,
	}
	payouts := append(mainnetAirdropPayouts, payout)
	return payouts
}

func tokenPayouts_RegNetParams() []TokenPayout {
	payout := TokenPayout{
		0, hexDecode("76a914cf646ab09da6133415ec2c629aacfc7c5b850cb788ac"), 100_0000 * 1e8,
	}
	payouts := append(mainnetAirdropPayouts, payout)
	return payouts
}