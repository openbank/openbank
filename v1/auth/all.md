
# Auth API v1.0.0

In order to interact with openbank APIs, you must be authenticated. This API provides the needed endpoints to retrieve an authorization code, exchange it against an access token and refresh an access token.

*
Host ``
EOL

*
Base Path ``

## Create an access token

Then you can exchange the authorization code with an access token. This endpoint is also used to refresh your tokens.

```sh
curl -X POST \
	/v1/token \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"grant_type": "GrantType",
		"code": "string",
		"redirect_uri": "string",
		"refresh_token": "string"
	}'
```
{{snippet createaccesstoken []}}

### HTTP Request

`POST /v1/token`

### Body Parameters

| Name         | Type      | Description                                                                                                                                  |
|--------------|-----------|----------------------------------------------------------------------------------------------------------------------------------------------|
| GrantType    | GrantType | GrantType is an enum to define which operation to perform.                                                                                   |
| Code         | string    | Code is the authorization code previously received from the authorization server.                                                            |
| RedirectURI  | string    | RedirectURI must match RedirectURI provided for the autorization code request.Required only when GrantType is equal to "authorization_code". |
| RefreshToken | string    | RefreshToken is the refresh token previously issued to the client.Required only when GrantType is equal to "refresh_token"                   |

### Responses

#### Response body

| Name         | Type   | Description                                                                   |
|--------------|--------|-------------------------------------------------------------------------------|
| AccessToken  | string | AccessToken issued by the authorization server.                               |
| TokenType    | string | TokenType is the type of the token (which is always "bearer").                |
| ExpiresIn    | int32  | ExpiresIn is the duration in seconds that the access token will remain valid. |
| RefreshToken | string | RefreshToken is the token to renew an expired access token.                   |

Example:

```json
{
  "access_token": "string",
  "token_type": "string",
  "expires_in": "int32",
  "refresh_token": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve an authorization code

First you need to retrieve an authorization code.

```sh
curl -X GET \
	/v1/auth \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getauthorizationcode []}}

### HTTP Request

`GET /v1/auth`

### Responses

#### Response body

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Annex

####  GrantType

GrantType defines the grant type when requesting a token.

| Value             | Description                                                                                                |
|-------------------|------------------------------------------------------------------------------------------------------------|
| UnknownGrantType  |                                                                                                            |
| AuthorizationCode | GrantType_AuthorizationCode is the grant type used to exchange an authorization code with an access token. |
| RefreshToken      | GrantType_RefreshToken is the grant type used to refresh an access token.                                  |

