# \FleetAPI

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChart**](FleetAPI.md#CreateChart) | **Post** /my/ships/{shipSymbol}/chart | Create Chart
[**CreateShipShipScan**](FleetAPI.md#CreateShipShipScan) | **Post** /my/ships/{shipSymbol}/scan/ships | Scan Ships
[**CreateShipSystemScan**](FleetAPI.md#CreateShipSystemScan) | **Post** /my/ships/{shipSymbol}/scan/systems | Scan Systems
[**CreateShipWaypointScan**](FleetAPI.md#CreateShipWaypointScan) | **Post** /my/ships/{shipSymbol}/scan/waypoints | Scan Waypoints
[**CreateSurvey**](FleetAPI.md#CreateSurvey) | **Post** /my/ships/{shipSymbol}/survey | Create Survey
[**DockShip**](FleetAPI.md#DockShip) | **Post** /my/ships/{shipSymbol}/dock | Dock Ship
[**ExtractResources**](FleetAPI.md#ExtractResources) | **Post** /my/ships/{shipSymbol}/extract | Extract Resources
[**ExtractResourcesWithSurvey**](FleetAPI.md#ExtractResourcesWithSurvey) | **Post** /my/ships/{shipSymbol}/extract/survey | Extract Resources with Survey
[**GetMounts**](FleetAPI.md#GetMounts) | **Get** /my/ships/{shipSymbol}/mounts | Get Mounts
[**GetMyShip**](FleetAPI.md#GetMyShip) | **Get** /my/ships/{shipSymbol} | Get Ship
[**GetMyShipCargo**](FleetAPI.md#GetMyShipCargo) | **Get** /my/ships/{shipSymbol}/cargo | Get Ship Cargo
[**GetMyShips**](FleetAPI.md#GetMyShips) | **Get** /my/ships | List Ships
[**GetRepairShip**](FleetAPI.md#GetRepairShip) | **Get** /my/ships/{shipSymbol}/repair | Get Repair Ship
[**GetScrapShip**](FleetAPI.md#GetScrapShip) | **Get** /my/ships/{shipSymbol}/scrap | Get Scrap Ship
[**GetShipCooldown**](FleetAPI.md#GetShipCooldown) | **Get** /my/ships/{shipSymbol}/cooldown | Get Ship Cooldown
[**GetShipNav**](FleetAPI.md#GetShipNav) | **Get** /my/ships/{shipSymbol}/nav | Get Ship Nav
[**InstallMount**](FleetAPI.md#InstallMount) | **Post** /my/ships/{shipSymbol}/mounts/install | Install Mount
[**Jettison**](FleetAPI.md#Jettison) | **Post** /my/ships/{shipSymbol}/jettison | Jettison Cargo
[**JumpShip**](FleetAPI.md#JumpShip) | **Post** /my/ships/{shipSymbol}/jump | Jump Ship
[**NavigateShip**](FleetAPI.md#NavigateShip) | **Post** /my/ships/{shipSymbol}/navigate | Navigate Ship
[**NegotiateContract**](FleetAPI.md#NegotiateContract) | **Post** /my/ships/{shipSymbol}/negotiate/contract | Negotiate Contract
[**OrbitShip**](FleetAPI.md#OrbitShip) | **Post** /my/ships/{shipSymbol}/orbit | Orbit Ship
[**PatchShipNav**](FleetAPI.md#PatchShipNav) | **Patch** /my/ships/{shipSymbol}/nav | Patch Ship Nav
[**PurchaseCargo**](FleetAPI.md#PurchaseCargo) | **Post** /my/ships/{shipSymbol}/purchase | Purchase Cargo
[**PurchaseShip**](FleetAPI.md#PurchaseShip) | **Post** /my/ships | Purchase Ship
[**RefuelShip**](FleetAPI.md#RefuelShip) | **Post** /my/ships/{shipSymbol}/refuel | Refuel Ship
[**RemoveMount**](FleetAPI.md#RemoveMount) | **Post** /my/ships/{shipSymbol}/mounts/remove | Remove Mount
[**RepairShip**](FleetAPI.md#RepairShip) | **Post** /my/ships/{shipSymbol}/repair | Repair Ship
[**ScrapShip**](FleetAPI.md#ScrapShip) | **Post** /my/ships/{shipSymbol}/scrap | Scrap Ship
[**SellCargo**](FleetAPI.md#SellCargo) | **Post** /my/ships/{shipSymbol}/sell | Sell Cargo
[**ShipRefine**](FleetAPI.md#ShipRefine) | **Post** /my/ships/{shipSymbol}/refine | Ship Refine
[**SiphonResources**](FleetAPI.md#SiphonResources) | **Post** /my/ships/{shipSymbol}/siphon | Siphon Resources
[**TransferCargo**](FleetAPI.md#TransferCargo) | **Post** /my/ships/{shipSymbol}/transfer | Transfer Cargo
[**WarpShip**](FleetAPI.md#WarpShip) | **Post** /my/ships/{shipSymbol}/warp | Warp Ship



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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The symbol of the ship.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.CreateChart(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CreateChart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChart`: CreateChart201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CreateChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.CreateShipShipScan(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CreateShipShipScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShipShipScan`: CreateShipShipScan201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CreateShipShipScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.CreateShipSystemScan(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CreateShipSystemScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShipSystemScan`: CreateShipSystemScan201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CreateShipSystemScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.CreateShipWaypointScan(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CreateShipWaypointScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShipWaypointScan`: CreateShipWaypointScan201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CreateShipWaypointScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The symbol of the ship.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.CreateSurvey(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.CreateSurvey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSurvey`: CreateSurvey201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.CreateSurvey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The symbol of the ship.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.DockShip(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.DockShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockShip`: DockShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.DockShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.
	extractResourcesRequest := *openapiclient.NewExtractResourcesRequest() // ExtractResourcesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.ExtractResources(context.Background(), shipSymbol).ExtractResourcesRequest(extractResourcesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.ExtractResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtractResources`: ExtractResources201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.ExtractResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

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


## ExtractResourcesWithSurvey

> ExtractResources201Response ExtractResourcesWithSurvey(ctx, shipSymbol).Survey(survey).Execute()

Extract Resources with Survey



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.
	survey := *openapiclient.NewSurvey("Signature_example", "Symbol_example", []openapiclient.SurveyDeposit{*openapiclient.NewSurveyDeposit("Symbol_example")}, time.Now(), "Size_example") // Survey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.ExtractResourcesWithSurvey(context.Background(), shipSymbol).Survey(survey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.ExtractResourcesWithSurvey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtractResourcesWithSurvey`: ExtractResources201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.ExtractResourcesWithSurvey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtractResourcesWithSurveyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **survey** | [**Survey**](Survey.md) |  | 

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


## GetMounts

> GetMounts200Response GetMounts(ctx, shipSymbol).Execute()

Get Mounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship's symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetMounts(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetMounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMounts`: GetMounts200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetMounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship&#39;s symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMounts200Response**](GetMounts200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The symbol of the ship.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetMyShip(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetMyShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyShip`: GetMyShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetMyShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The symbol of the ship.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetMyShipCargo(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetMyShipCargo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyShipCargo`: GetMyShipCargo200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetMyShipCargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	page := int32(56) // int32 | What entry offset to request (optional) (default to 1)
	limit := int32(56) // int32 | How many entries to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetMyShips(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetMyShips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyShips`: GetMyShips200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetMyShips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyShipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | What entry offset to request | [default to 1]
 **limit** | **int32** | How many entries to return per page | [default to 10]

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


## GetRepairShip

> GetRepairShip200Response GetRepairShip(ctx, shipSymbol).Execute()

Get Repair Ship



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetRepairShip(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetRepairShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepairShip`: GetRepairShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetRepairShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepairShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRepairShip200Response**](GetRepairShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScrapShip

> GetScrapShip200Response GetScrapShip(ctx, shipSymbol).Execute()

Get Scrap Ship



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetScrapShip(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetScrapShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScrapShip`: GetScrapShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetScrapShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScrapShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetScrapShip200Response**](GetScrapShip200Response.md)

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The symbol of the ship.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetShipCooldown(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetShipCooldown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShipCooldown`: GetShipCooldown200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetShipCooldown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.GetShipNav(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.GetShipNav``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShipNav`: GetShipNav200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.GetShipNav`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

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


## InstallMount

> InstallMount201Response InstallMount(ctx, shipSymbol).InstallMountRequest(installMountRequest).Execute()

Install Mount



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship's symbol.
	installMountRequest := *openapiclient.NewInstallMountRequest("Symbol_example") // InstallMountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.InstallMount(context.Background(), shipSymbol).InstallMountRequest(installMountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.InstallMount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallMount`: InstallMount201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.InstallMount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship&#39;s symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallMountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **installMountRequest** | [**InstallMountRequest**](InstallMountRequest.md) |  | 

### Return type

[**InstallMount201Response**](InstallMount201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.
	jettisonRequest := *openapiclient.NewJettisonRequest(openapiclient.TradeSymbol("PRECIOUS_STONES"), int32(123)) // JettisonRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.Jettison(context.Background(), shipSymbol).JettisonRequest(jettisonRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.Jettison``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Jettison`: Jettison200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.Jettison`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.
	jumpShipRequest := *openapiclient.NewJumpShipRequest("WaypointSymbol_example") // JumpShipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.JumpShip(context.Background(), shipSymbol).JumpShipRequest(jumpShipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.JumpShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JumpShip`: JumpShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.JumpShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.
	navigateShipRequest := *openapiclient.NewNavigateShipRequest("WaypointSymbol_example") // NavigateShipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.NavigateShip(context.Background(), shipSymbol).NavigateShipRequest(navigateShipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.NavigateShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NavigateShip`: NavigateShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.NavigateShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

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


## NegotiateContract

> NegotiateContract200Response NegotiateContract(ctx, shipSymbol).Execute()

Negotiate Contract



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship's symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.NegotiateContract(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.NegotiateContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NegotiateContract`: NegotiateContract200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.NegotiateContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship&#39;s symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNegotiateContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NegotiateContract200Response**](NegotiateContract200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The symbol of the ship.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.OrbitShip(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.OrbitShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrbitShip`: OrbitShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.OrbitShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.
	patchShipNavRequest := *openapiclient.NewPatchShipNavRequest() // PatchShipNavRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.PatchShipNav(context.Background(), shipSymbol).PatchShipNavRequest(patchShipNavRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.PatchShipNav``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchShipNav`: GetShipNav200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.PatchShipNav`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship's symbol.
	purchaseCargoRequest := *openapiclient.NewPurchaseCargoRequest(openapiclient.TradeSymbol("PRECIOUS_STONES"), int32(123)) // PurchaseCargoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.PurchaseCargo(context.Background(), shipSymbol).PurchaseCargoRequest(purchaseCargoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.PurchaseCargo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseCargo`: PurchaseCargo201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.PurchaseCargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship&#39;s symbol. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	purchaseShipRequest := *openapiclient.NewPurchaseShipRequest(openapiclient.ShipType("SHIP_PROBE"), "WaypointSymbol_example") // PurchaseShipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.PurchaseShip(context.Background()).PurchaseShipRequest(purchaseShipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.PurchaseShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseShip`: PurchaseShip201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.PurchaseShip`: %v\n", resp)
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

> RefuelShip200Response RefuelShip(ctx, shipSymbol).RefuelShipRequest(refuelShipRequest).Execute()

Refuel Ship



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.
	refuelShipRequest := *openapiclient.NewRefuelShipRequest() // RefuelShipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.RefuelShip(context.Background(), shipSymbol).RefuelShipRequest(refuelShipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.RefuelShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefuelShip`: RefuelShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.RefuelShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefuelShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refuelShipRequest** | [**RefuelShipRequest**](RefuelShipRequest.md) |  | 

### Return type

[**RefuelShip200Response**](RefuelShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMount

> RemoveMount201Response RemoveMount(ctx, shipSymbol).RemoveMountRequest(removeMountRequest).Execute()

Remove Mount



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship's symbol.
	removeMountRequest := *openapiclient.NewRemoveMountRequest("Symbol_example") // RemoveMountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.RemoveMount(context.Background(), shipSymbol).RemoveMountRequest(removeMountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.RemoveMount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveMount`: RemoveMount201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.RemoveMount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship&#39;s symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeMountRequest** | [**RemoveMountRequest**](RemoveMountRequest.md) |  | 

### Return type

[**RemoveMount201Response**](RemoveMount201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepairShip

> RepairShip200Response RepairShip(ctx, shipSymbol).Execute()

Repair Ship



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.RepairShip(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.RepairShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepairShip`: RepairShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.RepairShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepairShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepairShip200Response**](RepairShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScrapShip

> ScrapShip200Response ScrapShip(ctx, shipSymbol).Execute()

Scrap Ship



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.ScrapShip(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.ScrapShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScrapShip`: ScrapShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.ScrapShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScrapShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScrapShip200Response**](ScrapShip200Response.md)

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | Symbol of a ship.
	sellCargoRequest := *openapiclient.NewSellCargoRequest(openapiclient.TradeSymbol("PRECIOUS_STONES"), int32(123)) // SellCargoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.SellCargo(context.Background(), shipSymbol).SellCargoRequest(sellCargoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.SellCargo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SellCargo`: SellCargo201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.SellCargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | Symbol of a ship. | 

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

> ShipRefine201Response ShipRefine(ctx, shipSymbol).ShipRefineRequest(shipRefineRequest).Execute()

Ship Refine



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The symbol of the ship.
	shipRefineRequest := *openapiclient.NewShipRefineRequest("Produce_example") // ShipRefineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.ShipRefine(context.Background(), shipSymbol).ShipRefineRequest(shipRefineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.ShipRefine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShipRefine`: ShipRefine201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.ShipRefine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The symbol of the ship. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShipRefineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipRefineRequest** | [**ShipRefineRequest**](ShipRefineRequest.md) |  | 

### Return type

[**ShipRefine201Response**](ShipRefine201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiphonResources

> SiphonResources201Response SiphonResources(ctx, shipSymbol).Execute()

Siphon Resources



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.SiphonResources(context.Background(), shipSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.SiphonResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SiphonResources`: SiphonResources201Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.SiphonResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiphonResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiphonResources201Response**](SiphonResources201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The transferring ship's symbol.
	transferCargoRequest := *openapiclient.NewTransferCargoRequest(openapiclient.TradeSymbol("PRECIOUS_STONES"), int32(123), "ShipSymbol_example") // TransferCargoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.TransferCargo(context.Background(), shipSymbol).TransferCargoRequest(transferCargoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.TransferCargo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferCargo`: TransferCargo200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.TransferCargo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The transferring ship&#39;s symbol. | 

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

> WarpShip200Response WarpShip(ctx, shipSymbol).NavigateShipRequest(navigateShipRequest).Execute()

Warp Ship



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	shipSymbol := "shipSymbol_example" // string | The ship symbol.
	navigateShipRequest := *openapiclient.NewNavigateShipRequest("WaypointSymbol_example") // NavigateShipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FleetAPI.WarpShip(context.Background(), shipSymbol).NavigateShipRequest(navigateShipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FleetAPI.WarpShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarpShip`: WarpShip200Response
	fmt.Fprintf(os.Stdout, "Response from `FleetAPI.WarpShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipSymbol** | **string** | The ship symbol. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarpShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **navigateShipRequest** | [**NavigateShipRequest**](NavigateShipRequest.md) |  | 

### Return type

[**WarpShip200Response**](WarpShip200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

