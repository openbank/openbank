# Product API v1.0.0

Provides CRUD operations on the Product resource.

* Host ``

* Base Path ``

## Create a product

Creates a new product and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/products/{ProductCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"product_code": "string",
		"product": {
			"bank_id": "string",
			"code": "string",
			"parent_product_code": "string",
			"name": "string",
			"category": "string",
			"family": "string",
			"super_family": "string",
			"more_info_url": "string",
			"details": "string",
			"description": "string",
			"meta": {
				"license": "string"
			},
			"product_attributes": [
				{
					"product_code": "string",
					"product_attribute_id": "string",
					"name": "string",
					"type": "AttributeType",
					"value": "string"
				}
			]
		}
	}'
```
### HTTP Request

`POST /v1/banks/{BankID}/products/{ProductCode}`

### Query Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| BankID      | string |             |
| ProductCode | string |             |

### Body Parameters

| Name        | Type    | Description |
|-------------|---------|-------------|
| BankID      | string  |             |
| ProductCode | string  |             |
| Product     | Product |             |

##### Objects

###### Product

| Name              | Type               | Description |
|-------------------|--------------------|-------------|
| BankID            | string             |             |
| Code              | string             |             |
| ParentProductCode | string             |             |
| Name              | string             |             |
| Category          | string             |             |
| Family            | string             |             |
| SuperFamily       | string             |             |
| MoreInfoURL       | string             |             |
| Details           | string             |             |
| Description       | string             |             |
| Meta              | Metadata           |             |
| ProductAttributes | []ProductAttribute |             |

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| License | string |             |

###### ProductAttribute

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

### Responses

#### Response body

| Name              | Type               | Description |
|-------------------|--------------------|-------------|
| BankID            | string             |             |
| Code              | string             |             |
| ParentProductCode | string             |             |
| Name              | string             |             |
| Category          | string             |             |
| Family            | string             |             |
| SuperFamily       | string             |             |
| MoreInfoURL       | string             |             |
| Details           | string             |             |
| Description       | string             |             |
| Meta              | Metadata           |             |
| ProductAttributes | []ProductAttribute |             |

##### Objects

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| License | string |             |

###### ProductAttribute

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

Example:

```json
{
  "bank_id": "string",
  "code": "string",
  "parent_product_code": "string",
  "name": "string",
  "category": "string",
  "family": "string",
  "super_family": "string",
  "more_info_url": "string",
  "details": "string",
  "description": "string",
  "meta": {
    "license": "string"
  },
  "product_attributes": [
    {
      "product_code": "string",
      "product_attribute_id": "string",
      "name": "string",
      "type": "AttributeType",
      "value": "string"
    }
  ]
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | Product created successfully.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Create a product attribute

Creates a new product attribute and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"product_code": "string",
		"product_attribute_id": "string",
		"product_attribute": {
			"product_code": "string",
			"product_attribute_id": "string",
			"name": "string",
			"type": "AttributeType",
			"value": "string"
		}
	}'
```
{{snippet createproductattribute []}}

### HTTP Request

`POST /v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID}`

### Query Parameters

| Name               | Type   | Description |
|--------------------|--------|-------------|
| BankID             | string |             |
| ProductCode        | string |             |
| ProductAttributeID | string |             |

### Body Parameters

| Name               | Type             | Description |
|--------------------|------------------|-------------|
| BankID             | string           |             |
| ProductCode        | string           |             |
| ProductAttributeID | string           |             |
| ProductAttribute   | ProductAttribute |             |

##### Objects

###### ProductAttribute

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

### Responses

#### Response body

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

Example:

```json
{
  "product_code": "string",
  "product_attribute_id": "string",
  "name": "string",
  "type": "AttributeType",
  "value": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 201    | ProductAttribute created successfully.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a product

Permanently delete a product.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/products/{ProductCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`DELETE /v1/banks/{BankID}/products/{ProductCode}`

### Query Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| BankID      | string |             |
| ProductCode | string |             |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Product successfully deleted.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a product attribute

Permanently delete a product attribute.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deleteproductattribute []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID}`

### Query Parameters

| Name               | Type   | Description |
|--------------------|--------|-------------|
| BankID             | string |             |
| ProductCode        | string |             |
| ProductAttributeID | string |             |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ProductAttribute successfully deleted.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve product information

Retrieve information about the product specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/products/{ProductCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/banks/{BankID}/products/{ProductCode}`

### Query Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| BankID      | string |             |
| ProductCode | string |             |

### Responses

#### Response body

| Name              | Type               | Description |
|-------------------|--------------------|-------------|
| BankID            | string             |             |
| Code              | string             |             |
| ParentProductCode | string             |             |
| Name              | string             |             |
| Category          | string             |             |
| Family            | string             |             |
| SuperFamily       | string             |             |
| MoreInfoURL       | string             |             |
| Details           | string             |             |
| Description       | string             |             |
| Meta              | Metadata           |             |
| ProductAttributes | []ProductAttribute |             |

##### Objects

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| License | string |             |

###### ProductAttribute

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

Example:

```json
{
  "bank_id": "string",
  "code": "string",
  "parent_product_code": "string",
  "name": "string",
  "category": "string",
  "family": "string",
  "super_family": "string",
  "more_info_url": "string",
  "details": "string",
  "description": "string",
  "meta": {
    "license": "string"
  },
  "product_attributes": [
    {
      "product_code": "string",
      "product_attribute_id": "string",
      "name": "string",
      "type": "AttributeType",
      "value": "string"
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

## Retrieve product attribute information

Retrieve information about the product specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getproductattribute []}}

### HTTP Request

`GET /v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID}`

### Query Parameters

| Name               | Type   | Description |
|--------------------|--------|-------------|
| BankID             | string |             |
| ProductCode        | string |             |
| ProductAttributeID | string |             |

### Responses

#### Response body

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

Example:

```json
{
  "product_code": "string",
  "product_attribute_id": "string",
  "name": "string",
  "type": "AttributeType",
  "value": "string"
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

## Retrieve all available products

Retrieve information regarding all available products.

```sh
curl -X GET \
	/v1/banks/{BankID}/products \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
### HTTP Request

`GET /v1/banks/{BankID}/products`

### Query Parameters

| Name   | Type   | Description |
|--------|--------|-------------|
| BankID | string |             |

### Responses

#### Response body

| Name     | Type      | Description |
|----------|-----------|-------------|
| Products | []Product |             |

##### Objects

###### Product

| Name              | Type               | Description |
|-------------------|--------------------|-------------|
| BankID            | string             |             |
| Code              | string             |             |
| ParentProductCode | string             |             |
| Name              | string             |             |
| Category          | string             |             |
| Family            | string             |             |
| SuperFamily       | string             |             |
| MoreInfoURL       | string             |             |
| Details           | string             |             |
| Description       | string             |             |
| Meta              | Metadata           |             |
| ProductAttributes | []ProductAttribute |             |

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| License | string |             |

###### ProductAttribute

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

Example:

```json
{
  "products": [
    {
      "bank_id": "string",
      "code": "string",
      "parent_product_code": "string",
      "name": "string",
      "category": "string",
      "family": "string",
      "super_family": "string",
      "more_info_url": "string",
      "details": "string",
      "description": "string",
      "meta": {
        "license": "string"
      },
      "product_attributes": [
        {
          "product_code": "string",
          "product_attribute_id": "string",
          "name": "string",
          "type": "AttributeType",
          "value": "string"
        }
      ]
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

## Update a product

Updates a product's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/products/{ProductCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"product_code": "string",
		"product": {
			"bank_id": "string",
			"code": "string",
			"parent_product_code": "string",
			"name": "string",
			"category": "string",
			"family": "string",
			"super_family": "string",
			"more_info_url": "string",
			"details": "string",
			"description": "string",
			"meta": {
				"license": "string"
			},
			"product_attributes": [
				{
					"product_code": "string",
					"product_attribute_id": "string",
					"name": "string",
					"type": "AttributeType",
					"value": "string"
				}
			]
		}
	}'
```
### HTTP Request

`PUT /v1/banks/{BankID}/products/{ProductCode}`

### Query Parameters

| Name        | Type   | Description |
|-------------|--------|-------------|
| BankID      | string |             |
| ProductCode | string |             |

### Body Parameters

| Name        | Type    | Description |
|-------------|---------|-------------|
| BankID      | string  |             |
| ProductCode | string  |             |
| Product     | Product |             |

##### Objects

###### Product

| Name              | Type               | Description |
|-------------------|--------------------|-------------|
| BankID            | string             |             |
| Code              | string             |             |
| ParentProductCode | string             |             |
| Name              | string             |             |
| Category          | string             |             |
| Family            | string             |             |
| SuperFamily       | string             |             |
| MoreInfoURL       | string             |             |
| Details           | string             |             |
| Description       | string             |             |
| Meta              | Metadata           |             |
| ProductAttributes | []ProductAttribute |             |

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| License | string |             |

###### ProductAttribute

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

### Responses

#### Response body

| Name              | Type               | Description |
|-------------------|--------------------|-------------|
| BankID            | string             |             |
| Code              | string             |             |
| ParentProductCode | string             |             |
| Name              | string             |             |
| Category          | string             |             |
| Family            | string             |             |
| SuperFamily       | string             |             |
| MoreInfoURL       | string             |             |
| Details           | string             |             |
| Description       | string             |             |
| Meta              | Metadata           |             |
| ProductAttributes | []ProductAttribute |             |

##### Objects

###### Metadata

| Name    | Type   | Description |
|---------|--------|-------------|
| License | string |             |

###### ProductAttribute

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

Example:

```json
{
  "bank_id": "string",
  "code": "string",
  "parent_product_code": "string",
  "name": "string",
  "category": "string",
  "family": "string",
  "super_family": "string",
  "more_info_url": "string",
  "details": "string",
  "description": "string",
  "meta": {
    "license": "string"
  },
  "product_attributes": [
    {
      "product_code": "string",
      "product_attribute_id": "string",
      "name": "string",
      "type": "AttributeType",
      "value": "string"
    }
  ]
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | Product successfully updated.                                                          |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Update a product attribute

Updates a product attribute's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"product_code": "string",
		"product_attribute_id": "string",
		"product_attribute": {
			"product_code": "string",
			"product_attribute_id": "string",
			"name": "string",
			"type": "AttributeType",
			"value": "string"
		}
	}'
```
{{snippet updateproductattribute []}}

### HTTP Request

`PUT /v1/banks/{BankID}/products/{ProductCode}/attributes/{ProductAttributeID}`

### Query Parameters

| Name               | Type   | Description |
|--------------------|--------|-------------|
| BankID             | string |             |
| ProductCode        | string |             |
| ProductAttributeID | string |             |

### Body Parameters

| Name               | Type             | Description |
|--------------------|------------------|-------------|
| BankID             | string           |             |
| ProductCode        | string           |             |
| ProductAttributeID | string           |             |
| ProductAttribute   | ProductAttribute |             |

##### Objects

###### ProductAttribute

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

### Responses

#### Response body

| Name               | Type          | Description |
|--------------------|---------------|-------------|
| ProductCode        | string        |             |
| ProductAttributeID | string        |             |
| Name               | string        |             |
| Type               | AttributeType |             |
| Value              | string        |             |

Example:

```json
{
  "product_code": "string",
  "product_attribute_id": "string",
  "name": "string",
  "type": "AttributeType",
  "value": "string"
}
```
#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ProductAttribute successfully updated.                                                 |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Annex

####  AttributeType

| Value       | Description |
|-------------|-------------|
| UnknownType |             |
| STRING      |             |
| INTEGER     |             |
| DOUBLE      |             |
| DAY         |             |
