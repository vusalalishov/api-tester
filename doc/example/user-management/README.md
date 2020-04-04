# This folder contains tests for the simple user management api.

### API request/responses info:

```
POST /create
{
    "name": "John Doe",
    "age": "30"
}
```

Success:
```json
{
    "errorCode": "0",
    "errorMessage": "",
    "data": {
        "userId": "15"
    }
}
```
 
Failure:
```json
{
    "errorCode": "-100",
    "errorMessage": "Something went wrong"
}
```

```
GET /get-user/<user-id>
{
    "errorCode": "0",
    "errorMessage": "",
    "data": {
        "userId": "15",
        "name": "John Doe",
        "age": "30"
    }
}
```