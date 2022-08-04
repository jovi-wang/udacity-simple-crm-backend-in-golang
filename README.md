# simple CRM backend API

## Overview

This is a simple CRM backend API service to perform CRUD customers operations.

## Prerequisites:

1. install golang(1.18)
2. install http client(Postman or cURL)

## Setup local dev/test environment

### Launch server by executing this command in Terminal

```
go run .
```

The dev server will listen on port `3000`

### Run test cases by executing this command in Terminal

```
go test
```

All test cases should pass and print out `PASS`

## File structure

```text

data.go            --> customer data model and mock data
routeHandlers.go   --> a list of route handlers
main.go            --> main server file, register all request endpoints
main_test.go       --> test cases
swagger.yml        --> OpenAPI 3.0 definition file
static\
  index.html       --> home page with OpenAPI definition
```

## Request endpoints

http://localhost:3000

| HTTP method | url             | handler name   |
| ----------- | --------------- | -------------- |
| GET         | /customers      | getCustomers   |
| GET         | /customers/{id} | getCustomer    |
| POST        | /customers/{id} | addCustomer    |
| PUT         | /customers/{id} | updateCustomer |
| DELETE      | /customers/{id} | deleteCustomer |

## Customer struct model

| field     | type    | example                  |
| --------- | ------- | ------------------------ |
| ID        | string  | iNov-6iQ9J               |
| Name      | string  | Urbanus                  |
| Role      | string  | Administrative Officer   |
| Email     | string  | uabriani0@vistaprint.com |
| Phone     | string  | 772-772-1211             |
| Contacted | boolean | false                    |

a mock database is used with some mock data

## Depencencies:

- [gorilla mux v1.8.0](https://pkg.go.dev/github.com/gorilla/mux@v1.8.0)
- [github.com/teris-io/shortid v1.0](https://github.com/teris-io/shortid)
