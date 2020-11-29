# hash-challenge

**By**: Ednailson Junior

**Recruter**: July Demenjon

### Requirements

* [Docker-compose](https://docs.docker.com/compose/install/)

### Running

    docker-compose up

It will running the product-list service on the **port** `3333`

### Application

As required it has 2 microservices. The first one is an API that communicates with another microservice via gRPC asking
for the product discount. This last service calculates the discount according with the date (if it is the user birthday
or if it is black friday).

## Product List Service

This service is a REST API

#### Routes

##### List products

This route lists all the products with discounts. It can receive an `user_id` parameter, and if it is the user's birthday
it will calculate the equivalente discount for all products

Method `GET`

    /product

Parameters

* **user_id**: string

**Example**

Request

    curl --location --request GET 'http://localhost:3333/product?user_id=201'

Response

```json
[{
	"description": "A great notebook",
	"title": "Notebook Gamer",
	"price_in_cents": 2000,
	"id": "203",
	"discount": {
		"percentage": 5,
		"value_in_cents": 100
	}
}]
```

**Example 2**

Request

    curl --location --request GET 'http://localhost:3333/product'

Response

```json
[{
	"description": "A great notebook",
	"title": "Notebook Gamer",
	"price_in_cents": 2000,
	"id": "203",
	"discount": {
		"percentage": 0,
		"value_in_cents": 0
	}
}]
```
