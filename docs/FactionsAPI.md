# \FactionsAPI

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFaction**](FactionsAPI.md#GetFaction) | **Get** /factions/{factionSymbol} | Get Faction
[**GetFactions**](FactionsAPI.md#GetFactions) | **Get** /factions | List Factions



## GetFaction

> GetFaction200Response GetFaction(ctx, factionSymbol).Execute()

Get Faction



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
	factionSymbol := "COSMIC" // string | The faction symbol

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FactionsAPI.GetFaction(context.Background(), factionSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FactionsAPI.GetFaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFaction`: GetFaction200Response
	fmt.Fprintf(os.Stdout, "Response from `FactionsAPI.GetFaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**factionSymbol** | **string** | The faction symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFaction200Response**](GetFaction200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFactions

> GetFactions200Response GetFactions(ctx).Page(page).Limit(limit).Execute()

List Factions



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
	resp, r, err := apiClient.FactionsAPI.GetFactions(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FactionsAPI.GetFactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFactions`: GetFactions200Response
	fmt.Fprintf(os.Stdout, "Response from `FactionsAPI.GetFactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | What entry offset to request | [default to 1]
 **limit** | **int32** | How many entries to return per page | [default to 10]

### Return type

[**GetFactions200Response**](GetFactions200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

