# \ContractsApi

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptContract**](ContractsApi.md#AcceptContract) | **Post** /my/contracts/{contractId}/accept | Accept Contract
[**DeliverContract**](ContractsApi.md#DeliverContract) | **Post** /my/contracts/{contractId}/deliver | Deliver Contract
[**FulfillContract**](ContractsApi.md#FulfillContract) | **Post** /my/contracts/{contractId}/fulfill | Fulfill Contract
[**GetContract**](ContractsApi.md#GetContract) | **Get** /my/contracts/{contractId} | Get Contract
[**GetContracts**](ContractsApi.md#GetContracts) | **Get** /my/contracts | List Contracts



## AcceptContract

> AcceptContract200Response AcceptContract(ctx, contractId).Execute()

Accept Contract



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
    contractId := "contractId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.AcceptContract(context.Background(), contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.AcceptContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptContract`: AcceptContract200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.AcceptContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AcceptContract200Response**](AcceptContract200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliverContract

> DeliverContract200Response DeliverContract(ctx, contractId).DeliverContractRequest(deliverContractRequest).Execute()

Deliver Contract



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
    contractId := "contractId_example" // string | The ID of the contract
    deliverContractRequest := *openapiclient.NewDeliverContractRequest("ShipSymbol_example", "TradeSymbol_example", int32(123)) // DeliverContractRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.DeliverContract(context.Background(), contractId).DeliverContractRequest(deliverContractRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.DeliverContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliverContract`: DeliverContract200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.DeliverContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | The ID of the contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliverContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deliverContractRequest** | [**DeliverContractRequest**](DeliverContractRequest.md) |  | 

### Return type

[**DeliverContract200Response**](DeliverContract200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FulfillContract

> FulfillContract200Response FulfillContract(ctx, contractId).Execute()

Fulfill Contract



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
    contractId := "contractId_example" // string | The ID of the contract

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.FulfillContract(context.Background(), contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.FulfillContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FulfillContract`: FulfillContract200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.FulfillContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | The ID of the contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiFulfillContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FulfillContract200Response**](FulfillContract200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContract

> GetContract200Response GetContract(ctx, contractId).Execute()

Get Contract



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
    contractId := "contractId_example" // string | The contract ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.GetContract(context.Background(), contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContract`: GetContract200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | The contract ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContract200Response**](GetContract200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContracts

> GetContracts200Response GetContracts(ctx).Page(page).Limit(limit).Execute()

List Contracts



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
    page := int32(56) // int32 | What entry offset to request (optional)
    limit := int32(56) // int32 | How many entries to return per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsApi.GetContracts(context.Background()).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContracts`: GetContracts200Response
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | What entry offset to request | 
 **limit** | **int32** | How many entries to return per page | 

### Return type

[**GetContracts200Response**](GetContracts200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

