Auth API v1.0.0
===============

In order to interact with openbank APIs, you must be authenticated. This API provides the needed endpoints to retrieve an authorization code, exchange it against an access token and refresh an access token.

* Host `https://`

* Base Path ``

Create an access token {#method-post-createaccesstoken}
-------------------------------------------------------

Then you can exchange the authorization code with an access token. This endpoint is also used to refresh your tokens.

```sh
curl -X POST \
	https:///v1/token \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"grant_type": "GrantType",
		"code": "string",
		"redirect_uri": "string",
		"refresh_token": "string"
	}'
```

### HTTP Request {#http-request-method-post-createaccesstoken}

`POST https:///v1/token`

### Body Parameters {#body-parameters-method-post-createaccesstoken}

| Name          | Type      | Description                                                                                                                                  |
|---------------|-----------|----------------------------------------------------------------------------------------------------------------------------------------------|
| grant_type    | GrantType | GrantType is an enum to define which operation to perform.                                                                                   |
| code          | string    | Code is the authorization code previously received from the authorization server.                                                            |
| redirect_uri  | string    | RedirectURI must match RedirectURI provided for the autorization code request.Required only when GrantType is equal to "authorization_code". |
| refresh_token | string    | RefreshToken is the refresh token previously issued to the client.Required only when GrantType is equal to "refresh_token"                   |

##### Enums {#enums-CreateAccessTokenRequest}

###### GrantType

GrantType defines the grant type when requesting a token.

| Value             | Description                                                                                      |
|-------------------|--------------------------------------------------------------------------------------------------|
| UnknownGrantType  |                                                                                                  |
| AuthorizationCode | AuthorizationCode is the grant type used to exchange an authorization code with an access token. |
| RefreshToken      | RefreshToken is the grant type used to refresh an access token.                                  |

### Responses {#responses-method-post-createaccesstoken}

#### Response body {#response-body-method-post-createaccesstoken}

| Name          | Type   | Description                                                                   |
|---------------|--------|-------------------------------------------------------------------------------|
| access_token  | string | AccessToken issued by the authorization server.                               |
| token_type    | string | TokenType is the type of the token (which is always "bearer").                |
| expires_in    | int32  | ExpiresIn is the duration in seconds that the access token will remain valid. |
| refresh_token | string | RefreshToken is the token to renew an expired access token.                   |

Example:

```json
{
  "access_token": "string",
  "token_type": "string",
  "expires_in": "int32",
  "refresh_token": "string"
}
```

#### Response codes {#response-codes-method-post-createaccesstoken}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

Retrieve an authorization code {#method-get-getauthorizationcode}
-----------------------------------------------------------------

First you need to retrieve an authorization code.

```sh
curl -X GET \
	https:///v1/auth \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```

### HTTP Request {#http-request-method-get-getauthorizationcode}

`GET https:///v1/auth`

### Responses {#responses-method-get-getauthorizationcode}

#### Response body {#response-body-method-get-getauthorizationcode}

| Name | Type | Description |
|------|------|-------------|

Example:

```json
{}
```

#### Response codes {#response-codes-method-get-getauthorizationcode}

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 200    | Request executed successfully.                                                         |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
