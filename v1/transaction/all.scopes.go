// Code generated by "scopegen"; DO NOT EDIT.
package transactions

type ServiceScope struct{}

var Scopes = map[string]string{
	"https://auth.bnk.to/transaction.read": "Allow application to view transaction history",
	"https://auth.bnk.to/transaction.write": "Allow application to execute a transaction",
}

var AuthScopes = map[string][]string{
	"/transactions.TransactionService/ApprovePayment": []string{"https://auth.bnk.to/transaction.write"},
	"/transactions.TransactionService/CreateTransaction": []string{"https://auth.bnk.to/transaction.write"},
	"/transactions.TransactionService/GetTransaction": []string{"https://auth.bnk.to/transaction.read"},
	"/transactions.TransactionService/GetTransactions": []string{"https://auth.bnk.to/transaction.read"},
	"/transactions.TransactionService/GetTransactionsByAccount": []string{"https://auth.bnk.to/transaction.read"},
	"/transactions.TransactionService/ResendTFA": []string{"https://auth.bnk.to/transaction.write"},
	"/transactions.TransactionService/TFA": []string{"https://auth.bnk.to/transaction.write"},
}

// Any allows a loose challenge, for claims containing any of the method scopes.
//
// Use All method as a default for OAuth2 style scopes.  Any is useful with more complex scope definitions.
func (svcSc *ServiceScope) Any(method string, claims []string) bool {
	ch := AuthScopes[method]
	for _, s := range ch {
		for _, c := range claims {
			if string(s) == c {
				return true
			}
		}
	}
	return len(ch) == 0
}

// All is the default OAuth2 challenge method, ensuring claims contains all method scopes.
func (svcSc *ServiceScope) All(method string, claims []string) bool {
	ch := AuthScopes[method]
	if len(ch) > len(claims) {
		return false
	}
scopes:
	for _, s := range ch {
		for _, c := range claims {
			if string(s) == c {
				continue scopes
			}
		}
		return false
	}
	return true
}
