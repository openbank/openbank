// Code generated by "scopegen"; DO NOT EDIT.
package accountapplication

type ServiceScope struct{}

var Scopes = map[string]string{
	"https://auth.bnk.to/accountapplication.read":  "View accountapplication data",
	"https://auth.bnk.to/accountapplication.write": "Manage accountapplication data",
}

var AuthScopes = map[string][]string{
	"/accountapplication.AccountApplicationService/CreateAccountApplication":       {"https://auth.bnk.to/accountapplication.write"},
	"/accountapplication.AccountApplicationService/GetAccountApplication":          {"https://auth.bnk.to/accountapplication.read"},
	"/accountapplication.AccountApplicationService/GetAccountApplications":         {"https://auth.bnk.to/accountapplication.read"},
	"/accountapplication.AccountApplicationService/UpdateAccountApplicationStatus": {"https://auth.bnk.to/accountapplication.write"},
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
