# \TagApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScopes**](TagApi.md#GetScopes) | **Get** /1.0.0/tags/scopes | Retrieves the list of unique scopes associated with tags.
[**GetTags**](TagApi.md#GetTags) | **Get** /1.0.0/tags/ | Retrieves the list of tags.


# **GetScopes**
> Scopes GetScopes(ctx, page, pageSize, optional)
Retrieves the list of unique scopes associated with tags.

Retrieves the list of unique scopes associated with tags. A tag can have multiple scopes, which are different contexts in which a tag is used. An example is that a tag of \"web\" can be use in different scopes: security, compute. All scopes associated with all entities in the organization will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***TagApiGetScopesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiGetScopesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Scopes**](Scopes.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTags**
> Tags GetTags(ctx, page, pageSize, optional)
Retrieves the list of tags.

Retrieves the list of tags.  The tag/scope pair is considered a unique item in the list. All tags associated with all entities in the organization and all system tags will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***TagApiGetTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiGetTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Tags**](Tags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

