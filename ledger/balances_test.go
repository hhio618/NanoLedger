package libledger

import (
	"encoding/json"
	"testing"

	ledger "github.com/howeyc/ledger"
	"github.com/howeyc/ledger/decimal"
)

type testBalCase struct {
	name     string
	data     string
	balances []ledger.Account
	err      error
}

var testBalCases = []testBalCase{
	{
		"simple case",
		`1970/01/01 Payee
	Expense/test  (123 * 3)
	Assets

1970/01/01 Payee
	Expense/test   123
	Assets
`,
		[]ledger.Account{
			{
				Name:    "Assets",
				Balance: decimal.NewFromFloat(-4 * 123),
			},
			{
				Name:    "Expense/test",
				Balance: decimal.NewFromFloat(123 + 369),
			},
		},
		nil,
	},
	{
		"heirarchy",
		`1970/01/01 Payee
	Expense:test  (123 * 3)
	Assets

1970/01/01 Payee
	Expense:foo   123
	Assets
`,
		[]ledger.Account{
			{
				Name:    "Assets",
				Balance: decimal.NewFromFloat(-4 * 123),
			},
			{
				Name:    "Expense",
				Balance: decimal.NewFromFloat(123 + 369),
			},
			{
				Name:    "Expense:foo",
				Balance: decimal.NewFromFloat(123),
			},
			{
				Name:    "Expense:test",
				Balance: decimal.NewFromFloat(369),
			},
		},
		nil,
	},
}

// TestAdd is a test function for the Add function

func TestBalanceLedger(t *testing.T) {
	for _, tc := range testBalCases {
		bals, err := GetBalances([]byte(tc.data))
		if (err != nil && tc.err == nil) || (err != nil && tc.err != nil && err.Error() != tc.err.Error()) {
			t.Errorf("Error: expected `%s`, got `%s`", tc.err, err)
		}
		exp, _ := json.Marshal(tc.balances)
		if string(exp) != string(bals) {
			t.Errorf("Error(%s): expected \n`%s`, \ngot \n`%s`", tc.name, exp, bals)
		}
	}
}
