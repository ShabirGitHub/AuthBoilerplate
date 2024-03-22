### Task of Go API

You have to implement following Rest APIs in this boilerplate.
Login API (Check user credentials and return Json web token if login successful)
Profile API (Should be protected with JWT. (Only logged in user can call and get his/her profile)
Document the API endpoints with swagger (as in boilerplate)
You should follow the structure in this boilerplate. Put any common code inside the pkg folder so that it can be reused. Please handle all the edge cases for above mentioned apis.

***
### Setup and usage
```
cd to the project directory
```

```go
go build main.go
./main
```

Navigate to http://localhost:8080/swagger/index.html in your browser and check out the Swagger UI.

***

### LogIn API Credential

```
{
    "username": "shabir",
    "password": "shabir123"
}
```
***