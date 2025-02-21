package types11

import (
	"github.com/MinterTeam/minter-go-node/core/types"
)

type AppState struct {
	Note         string             `json:"note"`
	StartHeight  uint64             `json:"start_height"`
	Validators   []Validator        `json:"validators,omitempty"`
	Candidates   []types.Candidate  `json:"candidates,omitempty"`
	Accounts     []types.Account    `json:"accounts,omitempty"`
	Coins        []Coin             `json:"coins,omitempty"`
	FrozenFunds  []types.FrozenFund `json:"frozen_funds,omitempty"`
	UsedChecks   []types.UsedCheck  `json:"used_checks,omitempty"`
	MaxGas       uint64             `json:"max_gas"`
	TotalSlashed string             `json:"total_slashed"`
}

type Validator struct {
	TotalBipStake string          `json:"total_bip_stake"`
	PubKey        types.Pubkey    `json:"pub_key"`
	AccumReward   string          `json:"accum_reward"`
	AbsentTimes   *types.BitArray `json:"absent_times"`
}

type Coin struct {
	Name      string           `json:"name"`
	Symbol    types.CoinSymbol `json:"symbol"`
	Volume    string           `json:"volume"`
	Crr       uint             `json:"crr"`
	Reserve   string           `json:"reserve"`
	MaxSupply string           `json:"max_supply"`
}
