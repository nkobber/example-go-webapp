# example-go-webapp

## API Endpoints

### GET `/api/companies`
Returns all the companies in an array

#### Example request
`curl -XGET -H "Content-type: application/json" 'http://192.168.99.100:8080/api/companies'`
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
`curl -XGET -H "Content-type: application/json" 'http://192.168.99.100:8080/api/companies/571d54cfd7759b00012d5070'`

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
