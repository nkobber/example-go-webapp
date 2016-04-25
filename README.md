# example-go-webapp
Example webapp using Go, MongoDB and Docker

## API Endpoints

### GET `/api/companies`
Returns all the companies in an array

#### Example request
```bash
curl -XGET -H "Content-type: application/json" 'http://192.168.99.100:8080/api/companies'
```
#### Example response

```json
[
  {
    "id":"571d54cfd7759b00012d5070",
    "name":"Apple",
    "address":"1 Infinite Loop",
    "city":"Cupertino",
    "zipcode":"CA 95014",
    "country":"USA",
    "email":"apple@apple.com",
    "phone":"(408) 996–1010",
    "owners":[],
    "directors":[],
    "revisions":[]
  },
  {
    "id":"571d61142a017600018ebc0a",
    "name":"Microsoft Corporation",
    "address":"One Microsoft Way",
    "city":"Redmond",
    "zipcode":"WA 98052-7329",
    "country":"USA",
    "email":"microsoft@microsoft.com",
    "phone":"(425) 882-8080",
    "owners":["Bill Gates"],
    "directors":["Satya Nadella"],
    "revisions":[]
    }
]
```

### GET `/api/companies/<id>`
Returns information about the specific company

#### Example request
```bash
curl -XGET -H "Content-type: application/json" 'http://192.168.99.100:8080/api/companies/571d54cfd7759b00012d5070'
```

#### Example response
```json
{
  "id":"571d54cfd7759b00012d5070",
  "name":"Apple",
  "address":"1 Infinite Loop",
  "city":"Cupertino",
  "zipcode":"CA 95014",
  "country":"USA",
  "email":"apple@apple.com",
  "phone":"(408) 996–1010",
  "owners":[],
  "directors":[],
  "revisions":[]
}
```

### POST `/api/companies`
Creates a new company

#### Example request
```bash
curl -XPOST -H "Content-type: application/json" -d '{"name":"Apple","address":"1 Infinite Loop","city":"Cupertino","zipcode":"CA 95014","country":"USA","email":"apple@apple.com","phone":"(408) 996–1010","owners":["Steve Jobs"],"directors":["Tim Cook"]}' 'http://192.168.99.100:8080/api/companies'
```
#### Example response
```json
{
  "id":"571d62ee2a017600018ebc0b",
  "name":"Apple",
  "address":"1 Infinite Loop",
  "city":"Cupertino",
  "zipcode":"CA 95014",
  "country":"USA",
  "email":"apple@apple.com",
  "phone":"(408) 996–1010",
  "owners":["Steve Jobs"],
  "directors":["Tim Cook"],
  "revisions":null
}
```

### POST `/api/companies/<id>`
Updates an existing company

#### Example request
```bash
{"id":"571d62ee2a017600018ebc0b","name":"Apple Inc","address":"1 Infinite Loop","city":"Cupertino","zipcode":"CA 95014","country":"USA","email":"apple@apple.com","phone":"(408) 996–1010","owners":["Steve Jobs"],"directors":["Tim Cook"],"revisions":[{"id":"571d62ee2a017600018ebc0b","name":"Apple","address":"1 Infinite Loop","city":"Cupertino","zipcode":"CA 95014","country":"USA","email":"apple@apple.com","phone":"(408) 996–1010","owners":["Steve Jobs"],"directors":["Tim Cook"],"revisions":[]}]}

```
#### Example response
```json
{
  "id":"571d62ee2a017600018ebc0b",
  "name":"Apple Inc",
  "address":"1 Infinite Loop",
  "city":"Cupertino",
  "zipcode":"CA 95014",
  "country":"USA",
  "email":"apple@apple.com",
  "phone":"(408) 996–1010",
  "owners":["Steve Jobs"],
  "directors":["Tim Cook"],
  "revisions":[]
}
```

### DELETE `/api/companies/<id>`
Deletes a company

#### Example request
```bash
curl -XDELETE -H "Content-type: application/json" 'http://192.168.99.100:8080/api/companies/571d62ee2a017600018ebc0b'
```
#### Example response
```json
{
  "status": "OK"
}
```






