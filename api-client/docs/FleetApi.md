# \FleetApi

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChart**](FleetApi.md#CreateChart) | **Post** /my/ships/{shipSymbol}/chart | Create Chart
[**CreateShipShipScan**](FleetApi.md#CreateShipShipScan) | **Post** /my/ships/{shipSymbol}/scan/ships | Scan Ships
[**CreateShipSystemScan**](FleetApi.md#CreateShipSystemScan) | **Post** /my/ships/{shipSymbol}/scan/systems | Scan Systems
[**CreateShipWaypointScan**](FleetApi.md#CreateShipWaypointScan) | **Post** /my/ships/{shipSymbol}/scan/waypoints | Scan Waypoints
[**CreateSurvey**](FleetApi.md#CreateSurvey) | **Post** /my/ships/{shipSymbol}/survey | Create Survey
[**DockShip**](FleetApi.md#DockShip) | **Post** /my/ships/{shipSymbol}/dock | Dock Ship
[**ExtractResources**](FleetApi.md#ExtractResources) | **Post** /my/ships/{shipSymbol}/extract | Extract Resources
[**GetMyShip**](FleetApi.md#GetMyShip) | **Get** /my/ships/{shipSymbol} | Get Ship
[**GetMyShipCargo**](FleetApi.md#GetMyShipCargo) | **Get** /my/ships/{shipSymbol}/cargo | Get Ship Cargo
[**GetMyShips**](FleetApi.md#GetMyShips) | **Get** /my/ships | List Ships
[**GetShipCooldown**](FleetApi.md#GetShipCooldown) | **Get** /my/ships/{shipSymbol}/cooldown | Get Ship Cooldown
[**GetShipNav**](FleetApi.md#GetShipNav) | **Get** /my/ships/{shipSymbol}/nav | Get Ship Nav
[**Jettison**](FleetApi.md#Jettison) | **Post** /my/ships/{shipSymbol}/jettison | Jettison Cargo
[**JumpShip**](FleetApi.md#JumpShip) | **Post** /my/ships/{shipSymbol}/jump | Jump Ship
[**NavigateShip**](FleetApi.md#NavigateShip) | **Post** /my/ships/{shipSymbol}/navigate | Navigate Ship
[**OrbitShip**](FleetApi.md#OrbitShip) | **Post** /my/ships/{shipSymbol}/orbit | Orbit Ship
[**PatchShipNav**](FleetApi.md#PatchShipNav) | **Patch** /my/ships/{shipSymbol}/nav | Patch Ship Nav
[**PurchaseCargo**](FleetApi.md#PurchaseCargo) | **Post** /my/ships/{shipSymbol}/purchase | Purchase Cargo
[**PurchaseShip**](FleetApi.md#PurchaseShip) | **Post** /my/ships | Purchase Ship
[**RefuelShip**](FleetApi.md#RefuelShip) | **Post** /my/ships/{shipSymbol}/refuel | Refuel Ship
[**SellCargo**](FleetApi.md#SellCargo) | **Post** /my/ships/{shipSymbol}/sell | Sell Cargo
[**ShipRefine**](FleetApi.md#ShipRefine) | **Post** /my/ships/{shipSymbol}/refine | Ship Refine
[**TransferCargo**](FleetApi.md#TransferCargo) | **Post** /my/ships/{shipSymbol}/transfer | Transfer Cargo
[**WarpShip**](FleetApi.md#WarpShip) | **Post** /my/ships/{shipSymbol}/warp | Warp Ship



## CreateChart

> CreateChart201Response CreateChart(ctx, shipSymbol).Execute()

Create Chart



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
    shipSymbol := "shipSymbol_example" // string | The symbol of the ship

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.CreateChart(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.CreateChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChart`: CreateChart201Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.CreateChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateChart201Response**](CreateChart201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShipShipScan

> CreateShipShipScan201Response CreateShipShipScan(ctx, shipSymbol).Execute()

Scan Ships



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
    shipSymbol := "shipSymbol_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.CreateShipShipScan(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.CreateShipShipScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShipShipScan`: CreateShipShipScan201Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.CreateShipShipScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShipShipScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateShipShipScan201Response**](CreateShipShipScan201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShipSystemScan

> CreateShipSystemScan201Response CreateShipSystemScan(ctx, shipSymbol).Execute()

Scan Systems



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
    shipSymbol := "shipSymbol_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.CreateShipSystemScan(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.CreateShipSystemScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShipSystemScan`: CreateShipSystemScan201Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.CreateShipSystemScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShipSystemScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateShipSystemScan201Response**](CreateShipSystemScan201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShipWaypointScan

> CreateShipWaypointScan201Response CreateShipWaypointScan(ctx, shipSymbol).Execute()

Scan Waypoints



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
    shipSymbol := "shipSymbol_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.CreateShipWaypointScan(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.CreateShipWaypointScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShipWaypointScan`: CreateShipWaypointScan201Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.CreateShipWaypointScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShipWaypointScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateShipWaypointScan201Response**](CreateShipWaypointScan201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSurvey

> CreateSurvey201Response CreateSurvey(ctx, shipSymbol).Execute()

Create Survey



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
    shipSymbol := "shipSymbol_example" // string | The symbol of the ship

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.CreateSurvey(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.CreateSurvey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSurvey`: CreateSurvey201Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.CreateSurvey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSurveyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateSurvey201Response**](CreateSurvey201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DockShip

> DockShip200Response DockShip(ctx, shipSymbol).Execute()

Dock Ship



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
    shipSymbol := "shipSymbol_example" // string | The symbol of the ship

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.DockShip(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.DockShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DockShip`: DockShip200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.DockShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship | 

### Other Parameters

Other parameters are passed through a pointer to a apiDockShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DockShip200Response**](DockShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtractResources

> ExtractResources201Response ExtractResources(ctx, shipSymbol).ExtractResourcesRequest(extractResourcesRequest).Execute()

Extract Resources



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
    shipSymbol := "shipSymbol_example" // string | The ship symbol
    extractResourcesRequest := *openapiclient.NewExtractResourcesRequest() // ExtractResourcesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.ExtractResources(context.Background(), shipSymbol).ExtractResourcesRequest(extractResourcesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.ExtractResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtractResources`: ExtractResources201Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.ExtractResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtractResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extractResourcesRequest** | [**ExtractResourcesRequest**](ExtractResourcesRequest.md) |  | 

### Return type

[**ExtractResources201Response**](ExtractResources201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyShip

> GetMyShip200Response GetMyShip(ctx, shipSymbol).Execute()

Get Ship



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
    shipSymbol := "shipSymbol_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.GetMyShip(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.GetMyShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMyShip`: GetMyShip200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.GetMyShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMyShip200Response**](GetMyShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyShipCargo

> GetMyShipCargo200Response GetMyShipCargo(ctx, shipSymbol).Execute()

Get Ship Cargo



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
    shipSymbol := "shipSymbol_example" // string | The symbol of the ship

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.GetMyShipCargo(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.GetMyShipCargo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMyShipCargo`: GetMyShipCargo200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.GetMyShipCargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyShipCargoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMyShipCargo200Response**](GetMyShipCargo200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyShips

> GetMyShips200Response GetMyShips(ctx).Page(page).Limit(limit).Execute()

List Ships



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
    resp, r, err := apiClient.FleetApi.GetMyShips(context.Background()).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.GetMyShips``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMyShips`: GetMyShips200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.GetMyShips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyShipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | What entry offset to request | 
 **limit** | **int32** | How many entries to return per page | 

### Return type

[**GetMyShips200Response**](GetMyShips200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipCooldown

> GetShipCooldown200Response GetShipCooldown(ctx, shipSymbol).Execute()

Get Ship Cooldown



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
    shipSymbol := "shipSymbol_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.GetShipCooldown(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.GetShipCooldown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShipCooldown`: GetShipCooldown200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.GetShipCooldown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipCooldownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetShipCooldown200Response**](GetShipCooldown200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipNav

> GetShipNav200Response GetShipNav(ctx, shipSymbol).Execute()

Get Ship Nav



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
    shipSymbol := "shipSymbol_example" // string | The ship symbol

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.GetShipNav(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.GetShipNav``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShipNav`: GetShipNav200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.GetShipNav`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipNavRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetShipNav200Response**](GetShipNav200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Jettison

> Jettison200Response Jettison(ctx, shipSymbol).JettisonRequest(jettisonRequest).Execute()

Jettison Cargo



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
    shipSymbol := "shipSymbol_example" // string | 
    jettisonRequest := *openapiclient.NewJettisonRequest("Symbol_example", int32(123)) // JettisonRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.Jettison(context.Background(), shipSymbol).JettisonRequest(jettisonRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.Jettison``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Jettison`: Jettison200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.Jettison`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJettisonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jettisonRequest** | [**JettisonRequest**](JettisonRequest.md) |  | 

### Return type

[**Jettison200Response**](Jettison200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JumpShip

> JumpShip200Response JumpShip(ctx, shipSymbol).JumpShipRequest(jumpShipRequest).Execute()

Jump Ship



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
    shipSymbol := "shipSymbol_example" // string | 
    jumpShipRequest := *openapiclient.NewJumpShipRequest("SystemSymbol_example") // JumpShipRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.JumpShip(context.Background(), shipSymbol).JumpShipRequest(jumpShipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.JumpShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JumpShip`: JumpShip200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.JumpShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJumpShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jumpShipRequest** | [**JumpShipRequest**](JumpShipRequest.md) |  | 

### Return type

[**JumpShip200Response**](JumpShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NavigateShip

> NavigateShip200Response NavigateShip(ctx, shipSymbol).NavigateShipRequest(navigateShipRequest).Execute()

Navigate Ship



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
    shipSymbol := "shipSymbol_example" // string | The ship symbol
    navigateShipRequest := *openapiclient.NewNavigateShipRequest("WaypointSymbol_example") // NavigateShipRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.NavigateShip(context.Background(), shipSymbol).NavigateShipRequest(navigateShipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.NavigateShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NavigateShip`: NavigateShip200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.NavigateShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiNavigateShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **navigateShipRequest** | [**NavigateShipRequest**](NavigateShipRequest.md) |  | 

### Return type

[**NavigateShip200Response**](NavigateShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrbitShip

> OrbitShip200Response OrbitShip(ctx, shipSymbol).Execute()

Orbit Ship



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
    shipSymbol := "shipSymbol_example" // string | The symbol of the ship

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.OrbitShip(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.OrbitShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrbitShip`: OrbitShip200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.OrbitShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrbitShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrbitShip200Response**](OrbitShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchShipNav

> GetShipNav200Response PatchShipNav(ctx, shipSymbol).PatchShipNavRequest(patchShipNavRequest).Execute()

Patch Ship Nav



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
    shipSymbol := "shipSymbol_example" // string | The ship symbol
    patchShipNavRequest := *openapiclient.NewPatchShipNavRequest() // PatchShipNavRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.PatchShipNav(context.Background(), shipSymbol).PatchShipNavRequest(patchShipNavRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.PatchShipNav``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchShipNav`: GetShipNav200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.PatchShipNav`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchShipNavRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchShipNavRequest** | [**PatchShipNavRequest**](PatchShipNavRequest.md) |  | 

### Return type

[**GetShipNav200Response**](GetShipNav200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseCargo

> PurchaseCargo201Response PurchaseCargo(ctx, shipSymbol).PurchaseCargoRequest(purchaseCargoRequest).Execute()

Purchase Cargo



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
    shipSymbol := "shipSymbol_example" // string | 
    purchaseCargoRequest := *openapiclient.NewPurchaseCargoRequest("Symbol_example", int32(123)) // PurchaseCargoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.PurchaseCargo(context.Background(), shipSymbol).PurchaseCargoRequest(purchaseCargoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.PurchaseCargo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurchaseCargo`: PurchaseCargo201Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.PurchaseCargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseCargoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseCargoRequest** | [**PurchaseCargoRequest**](PurchaseCargoRequest.md) |  | 

### Return type

[**PurchaseCargo201Response**](PurchaseCargo201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseShip

> PurchaseShip201Response PurchaseShip(ctx).PurchaseShipRequest(purchaseShipRequest).Execute()

Purchase Ship



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
    purchaseShipRequest := *openapiclient.NewPurchaseShipRequest(openapiclient.ShipType("SHIP_PROBE"), "WaypointSymbol_example") // PurchaseShipRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.PurchaseShip(context.Background()).PurchaseShipRequest(purchaseShipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.PurchaseShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurchaseShip`: PurchaseShip201Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.PurchaseShip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseShipRequest** | [**PurchaseShipRequest**](PurchaseShipRequest.md) |  | 

### Return type

[**PurchaseShip201Response**](PurchaseShip201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefuelShip

> RefuelShip200Response RefuelShip(ctx, shipSymbol).Execute()

Refuel Ship



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
    shipSymbol := "shipSymbol_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.RefuelShip(context.Background(), shipSymbol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.RefuelShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefuelShip`: RefuelShip200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.RefuelShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefuelShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefuelShip200Response**](RefuelShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SellCargo

> SellCargo201Response SellCargo(ctx, shipSymbol).SellCargoRequest(sellCargoRequest).Execute()

Sell Cargo



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
    shipSymbol := "shipSymbol_example" // string | 
    sellCargoRequest := *openapiclient.NewSellCargoRequest("Symbol_example", int32(123)) // SellCargoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.SellCargo(context.Background(), shipSymbol).SellCargoRequest(sellCargoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.SellCargo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SellCargo`: SellCargo201Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.SellCargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSellCargoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sellCargoRequest** | [**SellCargoRequest**](SellCargoRequest.md) |  | 

### Return type

[**SellCargo201Response**](SellCargo201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShipRefine

> ShipRefine200Response ShipRefine(ctx, shipSymbol).ShipRefineRequest(shipRefineRequest).Execute()

Ship Refine



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
    shipSymbol := "shipSymbol_example" // string | The symbol of the ship
    shipRefineRequest := *openapiclient.NewShipRefineRequest("Produce_example") // ShipRefineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.ShipRefine(context.Background(), shipSymbol).ShipRefineRequest(shipRefineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.ShipRefine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShipRefine`: ShipRefine200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.ShipRefine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipRefineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipRefineRequest** | [**ShipRefineRequest**](ShipRefineRequest.md) |  | 

### Return type

[**ShipRefine200Response**](ShipRefine200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferCargo

> TransferCargo200Response TransferCargo(ctx, shipSymbol).TransferCargoRequest(transferCargoRequest).Execute()

Transfer Cargo



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
    shipSymbol := "shipSymbol_example" // string | 
    transferCargoRequest := *openapiclient.NewTransferCargoRequest("TradeSymbol_example", int32(123), "ShipSymbol_example") // TransferCargoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.TransferCargo(context.Background(), shipSymbol).TransferCargoRequest(transferCargoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.TransferCargo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferCargo`: TransferCargo200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.TransferCargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferCargoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferCargoRequest** | [**TransferCargoRequest**](TransferCargoRequest.md) |  | 

### Return type

[**TransferCargo200Response**](TransferCargo200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarpShip

> NavigateShip200Response WarpShip(ctx, shipSymbol).NavigateShipRequest(navigateShipRequest).Execute()

Warp Ship



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
    shipSymbol := "shipSymbol_example" // string | 
    navigateShipRequest := *openapiclient.NewNavigateShipRequest("WaypointSymbol_example") // NavigateShipRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FleetApi.WarpShip(context.Background(), shipSymbol).NavigateShipRequest(navigateShipRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FleetApi.WarpShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WarpShip`: NavigateShip200Response
    fmt.Fprintf(os.Stdout, "Response from `FleetApi.WarpShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarpShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **navigateShipRequest** | [**NavigateShipRequest**](NavigateShipRequest.md) |  | 

### Return type

[**NavigateShip200Response**](NavigateShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

