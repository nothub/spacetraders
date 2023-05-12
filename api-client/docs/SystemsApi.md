# \SystemsApi

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJumpGate**](SystemsApi.md#GetJumpGate) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/jump-gate | Get Jump Gate
[**GetMarket**](SystemsApi.md#GetMarket) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/market | Get Market
[**GetShipyard**](SystemsApi.md#GetShipyard) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/shipyard | Get Shipyard
[**GetSystem**](SystemsApi.md#GetSystem) | **Get** /systems/{systemSymbol} | Get System
[**GetSystemWaypoints**](SystemsApi.md#GetSystemWaypoints) | **Get** /systems/{systemSymbol}/waypoints | List Waypoints
[**GetSystems**](SystemsApi.md#GetSystems) | **Get** /systems | List Systems
[**GetWaypoint**](SystemsApi.md#GetWaypoint) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol} | Get Waypoint



## GetJumpGate

> GetJumpGate200Response GetJumpGate(ctx, systemSymbol, waypointSymbol).Execute()

Get Jump Gate



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
    systemSymbol := "systemSymbol_example" // string | The system symbol
    waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemsApi.GetJumpGate(context.Background(), systemSymbol, waypointSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetJumpGate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJumpGate`: GetJumpGate200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetJumpGate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemSymbol** | **string** | The system symbol | 
**waypointSymbol** | **string** | The waypoint symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJumpGateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetJumpGate200Response**](GetJumpGate200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarket

> GetMarket200Response GetMarket(ctx, systemSymbol, waypointSymbol).Execute()

Get Market



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
    systemSymbol := "systemSymbol_example" // string | The system symbol
    waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemsApi.GetMarket(context.Background(), systemSymbol, waypointSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMarket`: GetMarket200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetMarket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemSymbol** | **string** | The system symbol | 
**waypointSymbol** | **string** | The waypoint symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetMarket200Response**](GetMarket200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipyard

> GetShipyard200Response GetShipyard(ctx, systemSymbol, waypointSymbol).Execute()

Get Shipyard



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
    systemSymbol := "systemSymbol_example" // string | The system symbol
    waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemsApi.GetShipyard(context.Background(), systemSymbol, waypointSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetShipyard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShipyard`: GetShipyard200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetShipyard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemSymbol** | **string** | The system symbol | 
**waypointSymbol** | **string** | The waypoint symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipyardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetShipyard200Response**](GetShipyard200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystem

> GetSystem200Response GetSystem(ctx, systemSymbol).Execute()

Get System



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
    systemSymbol := "systemSymbol_example" // string | The system symbol (default to "X1-OE")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemsApi.GetSystem(context.Background(), systemSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystem`: GetSystem200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemSymbol** | **string** | The system symbol | [default to &quot;X1-OE&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSystem200Response**](GetSystem200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWaypoints

> GetSystemWaypoints200Response GetSystemWaypoints(ctx, systemSymbol).Page(page).Limit(limit).Execute()

List Waypoints



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
    systemSymbol := "systemSymbol_example" // string | The system symbol
    page := int32(56) // int32 | What entry offset to request (optional)
    limit := int32(56) // int32 | How many entries to return per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemsApi.GetSystemWaypoints(context.Background(), systemSymbol).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetSystemWaypoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemWaypoints`: GetSystemWaypoints200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetSystemWaypoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemSymbol** | **string** | The system symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWaypointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | What entry offset to request | 
 **limit** | **int32** | How many entries to return per page | 

### Return type

[**GetSystemWaypoints200Response**](GetSystemWaypoints200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystems

> GetSystems200Response GetSystems(ctx).Page(page).Limit(limit).Execute()

List Systems



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
    resp, r, err := apiClient.SystemsApi.GetSystems(context.Background()).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystems`: GetSystems200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetSystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | What entry offset to request | 
 **limit** | **int32** | How many entries to return per page | 

### Return type

[**GetSystems200Response**](GetSystems200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWaypoint

> GetWaypoint200Response GetWaypoint(ctx, systemSymbol, waypointSymbol).Execute()

Get Waypoint



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
    systemSymbol := "systemSymbol_example" // string | The system symbol
    waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemsApi.GetWaypoint(context.Background(), systemSymbol, waypointSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemsApi.GetWaypoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWaypoint`: GetWaypoint200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemsApi.GetWaypoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemSymbol** | **string** | The system symbol | 
**waypointSymbol** | **string** | The waypoint symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWaypointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetWaypoint200Response**](GetWaypoint200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

