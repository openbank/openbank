// Code generated by "scopegen"; DO NOT EDIT.
package statement

var Scopes = map[string]string{
	"https://auth.bnk.to/statement.read": "View statement data",
	"https://auth.bnk.to/statement.write": "Manage statement data",
}

var AuthScopes = map[string][]string{
	"/statement.StatementService/GetStatement": []string{"https://auth.bnk.to/statement.read"},
	"/statement.StatementService/GetStatements": []string{"https://auth.bnk.to/statement.read"},
}
