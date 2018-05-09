package auth

import (
	bapp "github.com/cosmos/cosmos-sdk/baseapp"
)

/*

Usage:

var accounts types.AccountMapper

// Fetch all signer accounts.
addrs := tx.GetSigners()
signers := make([]types.Account, len(addrs))
for i, addr := range addrs {
	acc := accounts.GetAccount(ctx)
	signers[i] = acc
}
ctx = auth.SetSigners(ctx, signers)

// Get all signer accounts.
signers := auth.GetSigners(ctx)
for i, signer := range signers {
	signer.Address() == tx.GetSigners()[i]
}

*/

type contextKey int // local to the auth module

const (
	contextKeySigners contextKey = iota
)

// add the signers to the context
func WithSigners(ctx bapp.Context, accounts []Account) bapp.Context {
	return ctx.WithValue(contextKeySigners, accounts)
}

// get the signers from the context
func GetSigners(ctx bapp.Context) []Account {
	v := ctx.Value(contextKeySigners)
	if v == nil {
		return []Account{}
	}
	return v.([]Account)
}
