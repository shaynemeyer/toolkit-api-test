# Toolkit API Test

This example project demonstrates how to use <https://github.com/shaynemeyer/toolkit/tree/main/v2>, which is a toolkit for use with Golang web applications.

## Installing Toolkit V2

```bash
go get https://github.com/shaynemeyer/toolkit/v2
```

## Example Uses

### Read a request and respond with JSON

```go
func login(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	var payload struct {
		Email string `json:"username"`
		Password string `json:"password"`
	}

	err := tools.ReadJSON(w, r, &payload)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}

	var respPayload toolkit.JSONResponse

	if payload.Email == "me@here.com" && payload.Password == "verysecret" {
		respPayload.Error = false
		respPayload.Message = "Logged in"
		_ = tools.WriteJSON(w, http.StatusAccepted, respPayload)
		return
	}

	respPayload.Error = true
	respPayload.Message = "invalid credentials"
	_ = tools.WriteJSON(w, http.StatusUnauthorized, respPayload)
}
```

### Writing JSON

```go
func logout(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	payload := toolkit.JSONResponse{
		Message: "Logged out",
	}

	_ = tools.WriteJSON(w, http.StatusAccepted, payload)
}
```

## Run the server

```bash
go run .
```

## Try it out in the browser

<http://localhost:8080>

For a successful case:

- email = "me@here.com"
- password = "verysecret"

Response from the server should be:

```json
{
  "error": false,
  "message": "Logged in"
}
```

Change the email to "me@there.com" and the response should be:

```json
{
  "error": true,
  "message": "invalid credentials"
}
```
