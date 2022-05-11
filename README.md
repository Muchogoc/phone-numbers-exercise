# phone-numbers-exercise

Lists and categorizes country phone numbers.Phone numbers are categorized by country, state (valid or not valid), country code and number.It uses a provided sqlite database and regular expressions to validate the numbers.

| Country    | Country Code       | Regex                            |
| ---------- | ------------------ | -------------------------------- |
| Cameroon   | Country code: +237 | Regex = \(237\)\ ?[2368]\d{7,8}$ |
| Ethiopia   | Country code: +251 | Regex = \(251\)\ ?[1-59]\d{8}$   |
| Morocco    | Country code: +212 | Regex = \(212\)\ ?[5-9]\d{8}$    |
| Mozambique | Country code: +258 | Regex = \(258\)\ ?[28]\d{7,8}$   |
| Uganda     | Country code: +256 | Regex = \(256\)\ ?\d{9}$         |

## Repository overview

Contains a backend (service) an frontend. The backend is implemented in Go and the frontend is a Single Page Application in React.

The backend follows a [Hexagonal Architecture](<'https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)>) which allows decoupling of application e.g data sources, business logic, external services.

- domain - holds the common entities
- infrastructure - holds interactions with external dependencies e.g databases
- usecases - holds the business logic
- presentation - holds implementations for how the service is accessed e.g REST APIs, gRPC

```bash
├── Dockerfile
├── frontend
│   ├── build
│   ├── public
│   ├── src
│   ├── package.json
│   ├── package-lock.json
│   └── README.md
├── service
|   ├── domain
|   ├── infrastructure
|   ├── presentation
|   └── usecases
├── go.mod
├── go.sum
├── sample.db
├── server.go
└── README.md
```

## Running

A Dockerfile is included to build and run the application. The frontend is already pre-build in the `frontend/build` directory.

1. Build a docker image

```bash
docker build . -t phone-numbers
```

2. Run the image

The server runs internally on port 8080 which needs to be exposed

```bash
docker run -p 8080:8080 phone-numbers
```

### Routes

- SPA --> `http://localhost:8080/`
- REST API --> `http://localhost:8080/api/v1/customers`

### REST API

Used to access the running backend service. It only supports `GET` requests

#### Query Parameters

- Filtering

The API supports the following filters:

1. `country` - filters records for a specific supported country e.g Cameroon

Example:

```bash
curl --location --request GET 'http://localhost:8080/api/v1/customers?country=Cameroon'
```

2. `state` - filters records where the phone number is valid/invalid. It is a boolean i.e true/false

Example:

```bash
curl --location --request GET 'http://localhost:8080/api/v1/customers?state=true'
```

- Pagination

The API uses offset-base pagination using the following parameters

1. `offset`
2. `limit`

Example:

```bash
curl --location --request GET 'http://localhost:8080/api/v1/customers?offset=0&limit=10'
```
