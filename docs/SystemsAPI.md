# \SystemsAPI

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConstruction**](SystemsAPI.md#GetConstruction) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/construction | Get Construction Site
[**GetJumpGate**](SystemsAPI.md#GetJumpGate) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/jump-gate | Get Jump Gate
[**GetMarket**](SystemsAPI.md#GetMarket) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/market | Get Market
[**GetShipyard**](SystemsAPI.md#GetShipyard) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/shipyard | Get Shipyard
[**GetSystem**](SystemsAPI.md#GetSystem) | **Get** /systems/{systemSymbol} | Get System
[**GetSystemWaypoints**](SystemsAPI.md#GetSystemWaypoints) | **Get** /systems/{systemSymbol}/waypoints | List Waypoints in System
[**GetSystems**](SystemsAPI.md#GetSystems) | **Get** /systems | List Systems
[**GetWaypoint**](SystemsAPI.md#GetWaypoint) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol} | Get Waypoint
[**SupplyConstruction**](SystemsAPI.md#SupplyConstruction) | **Post** /systems/{systemSymbol}/waypoints/{waypointSymbol}/construction/supply | Supply Construction Site



## GetConstruction

> GetConstruction200Response GetConstruction(ctx, systemSymbol, waypointSymbol).Execute()

Get Construction Site



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
	systemSymbol := "systemSymbol_example" // string | The system symbol
	waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.GetConstruction(context.Background(), systemSymbol, waypointSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.GetConstruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConstruction`: GetConstruction200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.GetConstruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemSymbol** | **string** | The system symbol | 
**waypointSymbol** | **string** | The waypoint symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConstructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetConstruction200Response**](GetConstruction200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	systemSymbol := "systemSymbol_example" // string | The system symbol
	waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.GetJumpGate(context.Background(), systemSymbol, waypointSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.GetJumpGate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJumpGate`: GetJumpGate200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.GetJumpGate`: %v\n", resp)
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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	systemSymbol := "systemSymbol_example" // string | The system symbol
	waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.GetMarket(context.Background(), systemSymbol, waypointSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.GetMarket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarket`: GetMarket200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.GetMarket`: %v\n", resp)
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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	systemSymbol := "systemSymbol_example" // string | The system symbol
	waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.GetShipyard(context.Background(), systemSymbol, waypointSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.GetShipyard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShipyard`: GetShipyard200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.GetShipyard`: %v\n", resp)
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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	systemSymbol := "systemSymbol_example" // string | The system symbol (default to "X1-OE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.GetSystem(context.Background(), systemSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.GetSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystem`: GetSystem200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.GetSystem`: %v\n", resp)
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

> GetSystemWaypoints200Response GetSystemWaypoints(ctx, systemSymbol).Page(page).Limit(limit).Type_(type_).Traits(traits).Execute()

List Waypoints in System



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
	systemSymbol := "systemSymbol_example" // string | The system symbol
	page := int32(56) // int32 | What entry offset to request (optional) (default to 1)
	limit := int32(56) // int32 | How many entries to return per page (optional) (default to 10)
	type_ := openapiclient.WaypointType("PLANET") // WaypointType | Filter waypoints by type. (optional)
	traits := openapiclient.get_system_waypoints_traits_parameter{WaypointTraitSymbol: openapiclient.WaypointTraitSymbol("UNCHARTED")} // GetSystemWaypointsTraitsParameter | Filter waypoints by one or more traits. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.GetSystemWaypoints(context.Background(), systemSymbol).Page(page).Limit(limit).Type_(type_).Traits(traits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.GetSystemWaypoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWaypoints`: GetSystemWaypoints200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.GetSystemWaypoints`: %v\n", resp)
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

 **page** | **int32** | What entry offset to request | [default to 1]
 **limit** | **int32** | How many entries to return per page | [default to 10]
 **type_** | [**WaypointType**](WaypointType.md) | Filter waypoints by type. | 
 **traits** | [**GetSystemWaypointsTraitsParameter**](GetSystemWaypointsTraitsParameter.md) | Filter waypoints by one or more traits. | 

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	page := int32(56) // int32 | What entry offset to request (optional) (default to 1)
	limit := int32(56) // int32 | How many entries to return per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.GetSystems(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.GetSystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystems`: GetSystems200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.GetSystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | What entry offset to request | [default to 1]
 **limit** | **int32** | How many entries to return per page | [default to 10]

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
	openapiclient "github.com/nothub/spacetraders"
)

func main() {
	systemSymbol := "systemSymbol_example" // string | The system symbol
	waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.GetWaypoint(context.Background(), systemSymbol, waypointSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.GetWaypoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWaypoint`: GetWaypoint200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.GetWaypoint`: %v\n", resp)
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


## SupplyConstruction

> SupplyConstruction201Response SupplyConstruction(ctx, systemSymbol, waypointSymbol).SupplyConstructionRequest(supplyConstructionRequest).Execute()

Supply Construction Site



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
	systemSymbol := "systemSymbol_example" // string | The system symbol
	waypointSymbol := "waypointSymbol_example" // string | The waypoint symbol
	supplyConstructionRequest := *openapiclient.NewSupplyConstructionRequest("ShipSymbol_example", "TradeSymbol_example", int32(123)) // SupplyConstructionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemsAPI.SupplyConstruction(context.Background(), systemSymbol, waypointSymbol).SupplyConstructionRequest(supplyConstructionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemsAPI.SupplyConstruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplyConstruction`: SupplyConstruction201Response
	fmt.Fprintf(os.Stdout, "Response from `SystemsAPI.SupplyConstruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemSymbol** | **string** | The system symbol | 
**waypointSymbol** | **string** | The waypoint symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupplyConstructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supplyConstructionRequest** | [**SupplyConstructionRequest**](SupplyConstructionRequest.md) |  | 

### Return type

[**SupplyConstruction201Response**](SupplyConstruction201Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

