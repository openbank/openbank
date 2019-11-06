# Profile API v1.0.0

Provides CRUD operations on the Profile resource.

*
Host ``
EOL

*
Base Path ``

## Query profiles

Returns the profile and associated accounts

```sh
curl -X GET \
	/v1/profiles \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/profiles`

### Responses

#### Response body

| Name     | Type      | Description |
|----------|-----------|-------------|
| Profile  | Profile   |             |
| Accounts | []Account |             |

##### Objects

###### Profile

| Name                     | Type         | Description                                                |
|--------------------------|--------------|------------------------------------------------------------|
| ProfileID                | string       | ProfileID is the unique identifier of a profile.           |
| FullName                 | string       | Full name                                                  |
| UserName                 | string       | User name                                                  |
| BirthDate                | string       | Birth date                                                 |
| Language                 | string       | Language code used                                         |
| Country                  | string       | User country code (VN, US, ID, SG, ...).                   |
| Email                    | string       | User email address                                         |
| EmailVefified            | bool         | True if email is verified, otherwise False                 |
| Mobile                   | string       | Mobile number                                              |
| Photo                    | string       | User profile photo url                                     |
| Title                    | string       | Title                                                      |
| PermanentAddress         | Address      | Permanent address                                          |
| ContactAddress           | Address      | Contact address                                            |
| ProfileNUmber            | string       | profile number                                             |
| FaceImageUrl             | string       | Face image of the customer                                 |
| FaceImageDate            | string       | Date when the face image was added/updated                 |
| RelationshipStatus       | string       | RelationshipStatus. Ex: Single                             |
| Dependents               | int32        | Number of dependents                                       |
| DobOfDependents          | []Timestamp  | Date of birth of dependents                                |
| CreditRating             | CreditRating | Credit rating                                              |
| CreditLimit              | Amount       | Credit Limit                                               |
| HighestEducationAttained | string       | Highest education such as bachelor, masters etc            |
| EmploymentStatus         | string       | Current employment status                                  |
| KycStatus                | bool         | Know Your Customer status                                  |
| BranchID                 | string       | Branch Identifier                                          |
| NameSuffix               | string       | Name suffix                                                |
| FirstName                | string       | FirstName of the person                                    |
| MiddleName               | string       | MiddleName or middle names (space separated) of the person |
| LastName                 | string       | LastName or last names (space separated) of the person     |

###### Account

| Name                          | Type          | Description                                                                           |
|-------------------------------|---------------|---------------------------------------------------------------------------------------|
| AccountID                     | string        | AccountID is the unique identifier of an account.                                     |
| Branch                        | string        | Branch is the branch code for the branch associated with the account.                 |
| BranchName                    | string        | BranchName is the long-form name of the branch associated with the account.           |
| Status                        | string        | Status is the status of the account.                                                  |
| AccruedInterestAtMaturityDate | Timestamp     | TODO: add comment.                                                                    |
| AmountDue                     | Amount        | TODO: add comment.                                                                    |
| AvailableBalance              | Amount        | AvailableBalance is the available balance of the account.                             |
| AvailableCreditLimit          | string        | AvailableCreditLimit is the available credit limit for the account.                   |
| CheckingInterestRate          | string        | CheckingInterestRate is the interest rate of the account if it is a checking account. |
| ContractDate                  | Timestamp     | ContractDate is the date of the contract initialization.                              |
| CreditLimit                   | string        | CreditLimit is the allowed credit limit.                                              |
| CurrentAccruedInterest        | string        | TODO: add comment.                                                                    |
| CurrentBalance                | Amount        | CurrentBalance is the current balance of the account.                                 |
| CurrentTerm                   | string        | CurrentTerm is the account validity period.                                           |
| DueDate                       | Timestamp     | DueDate is the loan maturity date.                                                    |
| InterestRate                  | string        | InterestRate is the interest rate for the account.                                    |
| MajorType                     | MajorType     | MajorType is the account type.                                                        |
| MajorCategory                 | MajorCategory | MajorCategory is the account category.                                                |
| MaturityDate                  | Timestamp     | MaturityDate is the maturity date, format is ISO 8601                                 |
| NextPaymentDueDate            | Timestamp     | TODO: add comment.                                                                    |
| OwnerName                     | string        | OwnerName is the name of the account's owner.                                         |
| StartDate                     | Timestamp     | TODO: add comment.                                                                    |

###### Address

| Name        | Type   | Description                                     |
|-------------|--------|-------------------------------------------------|
| CountryName | string | CountryName holds the country name information. |
| CityName    | string | CityName holds the city name information.       |
| State       | string | State holds the state information.              |
| Street      | string | Street holds the street information.            |
| PostalCode  | string | PostalCode holds the postal code information.   |

###### Timestamp

| Name    | Type  | Description |
|---------|-------|-------------|
| seconds | int64 |             |
| nanos   | int32 |             |

###### CreditRating

| Name   | Type   | Description |
|--------|--------|-------------|
| Rating | string |             |
| Source | string |             |

###### Amount

| Name | Type   | Description                        |
|------|--------|------------------------------------|
| Cur  | string | Cur is the currency of the amount. |
| Num  | string | Num is the value of the amount.    |

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
    "last_name": "string"
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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Query cards

Returns an array of ProfileCard associated with the account for the profile based on profile identifier

```sh
curl -X GET \
	/v1/profiles/card \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getprofilecard []}}

### HTTP Request

`GET /v1/profiles/card`

### Responses

#### Response body

| Name  | Type          | Description |
|-------|---------------|-------------|
| Cards | []ProfileCard |             |

##### Objects

###### ProfileCard

| Name      | Type   | Description                                               |
|-----------|--------|-----------------------------------------------------------|
| CardToken | string | CardToken is the card number.                             |
| Category  | string | Category is the card type.                                |
| LastFour  | string | LastFour is the last four digits of the card.             |
| OwnerName | string | OwnerName is the name of the card's owner.                |
| Type      | string | Type is the type of the account associated with the card. |

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
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 404    | Returned when the resource is not found.                                               |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
