# Product Collection API v1.0.0

Provides CRUD operations on the Product Collection resource.

* Host ``

* Base Path ``

## Create a product collection

Creates a new product collection and returns the object.

```sh
curl -X POST \
	/v1/banks/{BankID}/product-collections/{CollectionCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"collection_code": "string",
		"product_collection": {
			"collection_code": "string",
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
	}'
```
{{snippet createproductcollection []}}

### HTTP Request

`POST /v1/banks/{BankID}/product-collections/{CollectionCode}`

### Query Parameters

| Name           | Type   | Description |
|----------------|--------|-------------|
| BankID         | string |             |
| CollectionCode | string |             |

### Body Parameters

| Name              | Type              | Description |
|-------------------|-------------------|-------------|
| BankID            | string            |             |
| CollectionCode    | string            |             |
| ProductCollection | ProductCollection |             |

##### Objects

###### ProductCollection

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CollectionCode | string    |             |
| Products       | []Product |             |

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

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CollectionCode | string    |             |
| Products       | []Product |             |

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
  "collection_code": "string",
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
| 201    | ProductCollection created successfully.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Delete a product collection

Permanently delete a product collection.

```sh
curl -X DELETE \
	/v1/banks/{BankID}/product-collections/{CollectionCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet deleteproductcollection []}}

### HTTP Request

`DELETE /v1/banks/{BankID}/product-collections/{CollectionCode}`

### Query Parameters

| Name           | Type   | Description |
|----------------|--------|-------------|
| BankID         | string |             |
| CollectionCode | string |             |

### Responses

#### Response codes

| Status | Description                                                                            |
|--------|----------------------------------------------------------------------------------------|
| 204    | ProductCollection successfully deleted.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |

## Retrieve product collection information

Retrieve information about the product specified by the ID

```sh
curl -X GET \
	/v1/banks/{BankID}/product-collections/{CollectionCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN'
```
{{snippet getproductcollection []}}

### HTTP Request

`GET /v1/banks/{BankID}/product-collections/{CollectionCode}`

### Query Parameters

| Name           | Type   | Description |
|----------------|--------|-------------|
| BankID         | string |             |
| CollectionCode | string |             |

### Responses

#### Response body

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CollectionCode | string    |             |
| Products       | []Product |             |

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
  "collection_code": "string",
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

## Update a product collection

Updates a product collection's information

```sh
curl -X PUT \
	/v1/banks/{BankID}/product-collections/{CollectionCode} \
	-H 'Authorization: Bearer USE_YOUR_TOKEN' \
	-d '{
		"bank_id": "string",
		"collection_code": "string",
		"product_collection": {
			"collection_code": "string",
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
	}'
```
{{snippet updateproductcollection []}}

### HTTP Request

`PUT /v1/banks/{BankID}/product-collections/{CollectionCode}`

### Query Parameters

| Name           | Type   | Description |
|----------------|--------|-------------|
| BankID         | string |             |
| CollectionCode | string |             |

### Body Parameters

| Name              | Type              | Description |
|-------------------|-------------------|-------------|
| BankID            | string            |             |
| CollectionCode    | string            |             |
| ProductCollection | ProductCollection |             |

##### Objects

###### ProductCollection

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CollectionCode | string    |             |
| Products       | []Product |             |

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

| Name           | Type      | Description |
|----------------|-----------|-------------|
| CollectionCode | string    |             |
| Products       | []Product |             |

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
  "collection_code": "string",
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
| 204    | ProductCollection successfully updated.                                                |
| 400    | Returned when the request body is malformatted or does not match the expected request. |
| 401    | Returned when the request does not contains the user's credentials.                    |
| 403    | Returned when the user does not have permission to access the resource.                |
| 404    | Returned when the resource is not found.                                               |
| 500    | Returned when an unexpected error occured on the server side.                          |
