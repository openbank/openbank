// Code generated by "scopegen"; DO NOT EDIT.
package accountpublic

var Scopes = map[string]string{
	"https://auth.bnk.to/accountpublic.read": "View accountpublic data",
	"https://auth.bnk.to/accountpublic.write": "Manage accountpublic data",
}

var AuthScopes = map[string][]string{
	"/accountpublic.AccountPublicService/GetBankPublicAccount": []string{"https://auth.bnk.to/accountpublic.read"},
	"/accountpublic.AccountPublicService/GetPublicAccountAtAllBanks": []string{"https://auth.bnk.to/accountpublic.read"},
	"/accountpublic.AccountPublicService/GetPublicAccountByID": []string{"https://auth.bnk.to/accountpublic.read"},
}
