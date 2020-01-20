package openbank

import (
	accounts "github.com/openbank/openbank/v1/account"
	"github.com/openbank/openbank/v1/accountapplication"
	"github.com/openbank/openbank/v1/accountpublic"
	"github.com/openbank/openbank/v1/apierror"
	"github.com/openbank/openbank/v1/atm"
	"github.com/openbank/openbank/v1/auth"
	"github.com/openbank/openbank/v1/bank"
	"github.com/openbank/openbank/v1/branch"
	"github.com/openbank/openbank/v1/card"
	"github.com/openbank/openbank/v1/consent"
	"github.com/openbank/openbank/v1/consumer"
	"github.com/openbank/openbank/v1/counterparty"
	"github.com/openbank/openbank/v1/counterpartymetadata"
	"github.com/openbank/openbank/v1/customer"
	"github.com/openbank/openbank/v1/entitlement"
	"github.com/openbank/openbank/v1/fx"
	"github.com/openbank/openbank/v1/kyc"
	"github.com/openbank/openbank/v1/product"
	"github.com/openbank/openbank/v1/productcollection"
	"github.com/openbank/openbank/v1/profile"
	"github.com/openbank/openbank/v1/statement"
	transactions "github.com/openbank/openbank/v1/transaction"
	"github.com/openbank/openbank/v1/transactionmetadata"
	"github.com/openbank/openbank/v1/transactionrequest"
	"github.com/openbank/openbank/v1/types"
)

// AllScopes returns all OAuth2 scopes and their descriptions.
func AllScopes() map[string]string {
	var (
		allScopes = []map[string]string{
			accounts.Scopes,
			accountapplication.Scopes,
			accountpublic.Scopes,
			apierror.Scopes,
			atm.Scopes,
			auth.Scopes,
			bank.Scopes,
			branch.Scopes,
			card.Scopes,
			consent.Scopes,
			consumer.Scopes,
			counterparty.Scopes,
			counterpartymetadata.Scopes,
			customer.Scopes,
			entitlement.Scopes,
			fx.Scopes,
			kyc.Scopes,
			product.Scopes,
			productcollection.Scopes,
			profile.Scopes,
			statement.Scopes,
			transactions.Scopes,
			transactionmetadata.Scopes,
			transactionrequest.Scopes,
			types.Scopes,
		}
		result = make(map[string]string)
	)

	for _, scopes := range allScopes {
		for k, v := range scopes {
			result[k] = v
		}
	}
	return result
}
