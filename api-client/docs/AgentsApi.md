# \AgentsApi

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyAgent**](AgentsApi.md#GetMyAgent) | **Get** /my/agent | My Agent Details



## GetMyAgent

> GetMyAgent200Response GetMyAgent(ctx).Execute()

My Agent Details



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.GetMyAgent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.GetMyAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMyAgent`: GetMyAgent200Response
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.GetMyAgent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyAgentRequest struct via the builder pattern


### Return type

[**GetMyAgent200Response**](GetMyAgent200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

