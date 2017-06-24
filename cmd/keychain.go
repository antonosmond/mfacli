package cmd

import (
	"fmt"

	keychain "github.com/keybase/go-keychain"
)

var kc = keychain.NewItem()

func init() {
	kc.SetSecClass(keychain.SecClassGenericPassword)
	kc.SetService("mfacli")
}

func delete(account string) error {
	kc.SetAccount(account)
	return keychain.DeleteItem(kc)
}

func exists(account string) bool {
	if _, err := load(account); err != nil {
		return false
	}
	return true
}

func save(account string, secret string) error {
	err := delete(account)
	if err != nil && err != keychain.ErrorItemNotFound {
		return err
	}
	kc.SetAccount(account)
	kc.SetLabel(account)
	kc.SetData([]byte(secret))
	kc.SetSynchronizable(keychain.SynchronizableNo)
	kc.SetAccessible(keychain.AccessibleWhenUnlocked)
	err = keychain.AddItem(kc)
	if err != nil {
		return err
	}
	return nil
}

func load(account string) (secret string, err error) {
	kc.SetAccount(account)
	kc.SetMatchLimit(keychain.MatchLimitOne)
	kc.SetReturnData(true)
	results, err := keychain.QueryItem(kc)
	if err != nil {
		return "", err
	}
	if len(results) != 1 {
		return "", fmt.Errorf("account '%s' not found", account)
	}
	return string(results[0].Data), nil
}

func query() (accounts []string, err error) {
	kc.SetMatchLimit(keychain.MatchLimitAll)
	kc.SetReturnAttributes(true)
	results, err := keychain.QueryItem(kc)
	if err != nil {
		return nil, err
	}
	for _, r := range results {
		accounts = append(accounts, r.Account)
	}
	return accounts, nil
}
