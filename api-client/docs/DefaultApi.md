# \DefaultApi

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Register**](DefaultApi.md#Register) | **Post** /register | Register New Agent



## Register

> Register201Response Register(ctx).RegisterRequest(registerRequest).Execute()

Register New Agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nothub/spacetraders/api"
)

func main() {
    registerRequest := *openapiclient.NewRegisterRequest("Faction_example", "BADGER") // RegisterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Register(context.Background()).RegisterRequest(registerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Register``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Register`: Register201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Register`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerRequest** | [**RegisterRequest**](RegisterRequest.md) |  | 

### Return type

[**Register201Response**](Register201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

