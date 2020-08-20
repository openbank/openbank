Profile API v1.0.0
==================

Provides CRUD operations on the Profile resource.

* Host `https://`

* Base Path ``

Query profile {#method-get-getprofile}
--------------------------------------

Returns the profile and associated accounts

```sh
curl -X GET \
	https:///v1/profile \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getprofile}

`GET https:///v1/profile`

### Responses {#responses-method-get-getprofile}

#### Response body {#response-body-method-get-getprofile}

| Name     | Type       | Description |
|----------|------------|-------------|
| profile  | Profile    |             |
| accounts | \[]Account |             |

##### Objects {#objects-GetProfileResponse}

###### Profile

| Name                       | Type         | Description                                                |
|----------------------------|--------------|------------------------------------------------------------|
| profile_id                 | string       | ProfileID is the unique identifier of a profile.           |
| fullname                   | string       | Full name                                                  |
| username                   | string       | User name                                                  |
| birthdate                  | string       | Birth date                                                 |
| language                   | string       | Language code used                                         |
| country                    | string       | User country code (VN, US, ID, SG, ...).                   |
| email                      | string       | User email address                                         |
| email_verified             | bool         | True if email is verified, otherwise False                 |
| mobile                     | string       | Mobile number                                              |
| photo                      | string       | User profile photo url                                     |
| title                      | string       | Title                                                      |
| permanent_address          | Address      | Permanent address                                          |
| contact_address            | Address      | Contact address                                            |
| profile_number             | string       | profile number                                             |
| face_image_url             | string       | Face image of the customer                                 |
| face_image_date            | string       | Date when the face image was added/updated                 |
| relationship_status        | string       | RelationshipStatus. Ex: Single                             |
| dependents                 | int32        | Number of dependents                                       |
| dob_of_dependents          | \[]Timestamp | Date of birth of dependents                                |
| credit_rating              | CreditRating | Credit rating                                              |
| credit_limit               | Amount       | Credit Limit                                               |
| highest_education_attained | string       | Highest education such as bachelor, masters etc            |
| employment_status          | string       | Current employment status                                  |
| kyc_status                 | bool         | Know Your Customer status                                  |
| branchId                   | string       | Branch Identifier                                          |
| nameSuffix                 | string       | Name suffix                                                |
| first_name                 | string       | FirstName of the person                                    |
| middle_name                | string       | MiddleName or middle names (space separated) of the person |
| last_name                  | string       | LastName or last names (space separated) of the person     |
| contact_number             | string       | Contact number                                             |

###### Account

| Name                              | Type          | Description                                                                           |
|-----------------------------------|---------------|---------------------------------------------------------------------------------------|
| account_id                        | string        | AccountID is the unique identifier of an account.                                     |
| branch                            | string        | Branch is the branch code for the branch associated with the account.                 |
| branch_name                       | string        | BranchName is the long-form name of the branch associated with the account.           |
| status                            | string        | Status is the status of the account.                                                  |
| accrued_interest_at_maturity_date | Timestamp     | Interest accrues at an annual rate of interest that is fixed                          |
| amount_due                        | Amount        | Specify when payments are due on money borrowed                                       |
| available_balance                 | Amount        | AvailableBalance is the available balance of the account.                             |
| available_credit_limit            | string        | AvailableCreditLimit is the available credit limit for the account.                   |
| checking_interest_rate            | string        | CheckingInterestRate is the interest rate of the account if it is a checking account. |
| contract_date                     | Timestamp     | ContractDate is the date of the contract initialization.                              |
| credit_limit                      | string        | CreditLimit is the allowed credit limit.                                              |
| current_accrued_interest          | string        | Interest earned but not received                                                      |
| current_balance                   | Amount        | CurrentBalance is the current balance of the account.                                 |
| current_term                      | string        | CurrentTerm is the account validity period.                                           |
| due_date                          | Timestamp     | DueDate is the loan maturity date.                                                    |
| interest_rate                     | string        | InterestRate is the interest rate for the account.                                    |
| major_type                        | MajorType     | MajorType is the account type.                                                        |
| major_category                    | MajorCategory | MajorCategory is the account category.                                                |
| maturity_date                     | Timestamp     | MaturityDate is the maturity date, format is ISO 8601                                 |
| next_payment_due_date             | Timestamp     | Specify when payments are due on money borrowed                                       |
| owner_name                        | string        | OwnerName is the name of the account's owner.                                         |
| start_date                        | Timestamp     | Account opening date                                                                  |

###### Address

| Name         | Type   | Description                                     |
|--------------|--------|-------------------------------------------------|
| country_name | string | CountryName holds the country name information. |
| city_name    | string | CityName holds the city name information.       |
| state        | string | State holds the state information.              |
| line_1       | string | Street holds the street information.            |
| postal_code  | string | PostalCode holds the postal code information.   |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| rating | string |             |
| source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| cur  | string | Cur is the currency of the amount. |
| num  | string | Num is the value of the amount.    |

##### Enums {#enums-GetProfileResponse}

###### MajorType

MajorType describes the type of the account.

| Value            | Description                                 |
|------------------|---------------------------------------------|
| UnknownMajorType |                                             |
| Checking         | Checking account.                           |
| Saving           | Saving account.                             |
| TimeDeposit      | TimeDeposit for a time deposit account.     |
| CommercialLoan   | CommercialLoan for a business loan account. |
| MortgageLoan     | MortgageLoan for a home loan account.       |
| ConsumerLoan     | ConsumerLoan for a consumer loan account.   |

###### MajorCategory

MajorCategory describes the category of the account.

| Value                | Description               |
|----------------------|---------------------------|
| UnknownMajorCategory |                           |
| Dep                  | Dep for deposit category. |
| Loan                 | Loan for loan category.   |

Example:

```json
{
  "profile": {
    "profile_id": "string",
    "fullname": "string",
    "username": "string",
    "birthdate": "string",
    "language": "string",
    "country": "string",
    "email": "string",
    "email_verified": "bool",
    "mobile": "string",
    "photo": "string",
    "title": "string",
    "permanent_address": {
      "country_name": "string",
      "city_name": "string",
      "state": "string",
      "line_1": "string",
      "postal_code": "string"
    },
    "contact_address": {
      "country_name": "string",
      "city_name": "string",
      "state": "string",
      "line_1": "string",
      "postal_code": "string"
    },
    "profile_number": "string",
    "face_image_url": "string",
    "face_image_date": "string",
    "relationship_status": "string",
    "dependents": "int32",
    "dob_of_dependents": [
      {
        "seconds": "int64",
        "nanos": "int32"
      }
    ],
    "credit_rating": {
      "rating": "string",
      "source": "string"
    },
    "credit_limit": {
      "cur": "string",
      "num": "string"
    },
    "highest_education_attained": "string",
    "employment_status": "string",
    "kyc_status": "bool",
    "branchId": "string",
    "nameSuffix": "string",
    "first_name": "string",
    "middle_name": "string",
    "last_name": "string",
    "contact_number": "string"
  },
  "accounts": [
    {
      "account_id": "string",
      "branch": "string",
      "branch_name": "string",
      "status": "string",
      "accrued_interest_at_maturity_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "amount_due": {
        "cur": "string",
        "num": "string"
      },
      "available_balance": {
        "cur": "string",
        "num": "string"
      },
      "available_credit_limit": "string",
      "checking_interest_rate": "string",
      "contract_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "credit_limit": "string",
      "current_accrued_interest": "string",
      "current_balance": {
        "cur": "string",
        "num": "string"
      },
      "current_term": "string",
      "due_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "interest_rate": "string",
      "major_type": "MajorType",
      "major_category": "MajorCategory",
      "maturity_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "next_payment_due_date": {
        "seconds": "int64",
        "nanos": "int32"
      },
      "owner_name": "string",
      "start_date": {
        "seconds": "int64",
        "nanos": "int32"
      }
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getprofile}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Query cards {#method-get-getprofilecard}
----------------------------------------

Returns an array of ProfileCard associated with the account for the profile based on profile identifier

```sh
curl -X GET \
	https:///v1/profile/card \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getprofilecard}

`GET https:///v1/profile/card`

### Responses {#responses-method-get-getprofilecard}

#### Response body {#response-body-method-get-getprofilecard}

| Name  | Type           | Description |
|-------|----------------|-------------|
| cards | \[]ProfileCard |             |

##### Objects {#objects-GetProfileCardsResponse}

###### ProfileCard

| Name       | Type   | Description                                               |
|------------|--------|-----------------------------------------------------------|
| card_token | string | CardToken is the card number.                             |
| category   | string | Category is the card type.                                |
| last_four  | string | LastFour is the last four digits of the card.             |
| owner_name | string | OwnerName is the name of the card's owner.                |
| type       | string | Type is the type of the account associated with the card. |

Example:

```json
{
  "cards": [
    {
      "card_token": "string",
      "category": "string",
      "last_four": "string",
      "owner_name": "string",
      "type": "string"
    }
  ]
}
```

#### Response codes {#response-codes-method-get-getprofilecard}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
