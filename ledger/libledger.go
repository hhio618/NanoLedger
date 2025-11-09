package libledger

import (
	"bytes"
	"encoding/json"
	"fmt"

	ledger "github.com/howeyc/ledger"
)

// Exported function
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Exported function
func Balance(data []byte) string {
	// b := bytes.NewBufferString(data)
	reader := bytes.NewReader(data)
	transactions, err := ledger.ParseLedger(reader)
	bals := ledger.GetBalances(transactions, []string{})
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	got, _ := json.Marshal(bals)
	return string(got)
}
