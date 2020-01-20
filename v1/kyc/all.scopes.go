// Code generated by "scopegen"; DO NOT EDIT.
package kyc

var Scopes = map[string]string{
	"https://auth.bnk.to/kyc.read": "View kyc data",
	"https://auth.bnk.to/kyc.write": "Manage kyc data",
}

var AuthScopes = map[string][]string{
	"/kyc.KYCService/AddKYCCheck": []string{"https://auth.bnk.to/kyc.write"},
	"/kyc.KYCService/AddKYCDocument": []string{"https://auth.bnk.to/kyc.write"},
	"/kyc.KYCService/AddKYCMedia": []string{"https://auth.bnk.to/kyc.write"},
	"/kyc.KYCService/AddKYCStatus": []string{"https://auth.bnk.to/kyc.write"},
	"/kyc.KYCService/GetCustomerKYCCheck": []string{"https://auth.bnk.to/kyc.read"},
	"/kyc.KYCService/GetCustomerKYCDocument": []string{"https://auth.bnk.to/kyc.read"},
	"/kyc.KYCService/GetCustomerKYCStatus": []string{"https://auth.bnk.to/kyc.read"},
	"/kyc.KYCService/GetKYCMedia": []string{"https://auth.bnk.to/kyc.read"},
}
