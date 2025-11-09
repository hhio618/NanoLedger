package libledger

import (
	"bytes"
	"encoding/json"
	"fmt"

	ledger "github.com/howeyc/ledger"
	"gopkg.in/yaml.v3"
)

func getBalances(data []byte) ([]*ledger.Account, error) {
	// b := bytes.NewBufferString(data)
	reader := bytes.NewReader(data)
	transactions, err := ledger.ParseLedger(reader)
	bals := ledger.GetBalances(transactions, []string{})
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	return bals, nil
}

// Exported function
func GetBalances(data []byte) ([]byte, error) {

	bals, err := getBalances(data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	got, _ := json.Marshal(bals)
	return got, nil
}

// Exported function
func GetBalancesYaml(data []byte) ([]byte, error) {
	bals, err := getBalances(data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	got, _ := yaml.Marshal(bals)
	return got, nil
}
