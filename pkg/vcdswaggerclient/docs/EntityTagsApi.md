# \EntityTagsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntityTags**](EntityTagsApi.md#GetEntityTags) | **Get** /1.0.0/entity/{id}/tags | Retrieves the tag information for a specific entity.
[**GetEntityTagsSummaries**](EntityTagsApi.md#GetEntityTagsSummaries) | **Get** /1.0.0/tags/entityTagsSummaries | Retrieves the list of entities that have at least one tag assigned to it.
[**UpdateEntityTags**](EntityTagsApi.md#UpdateEntityTags) | **Put** /1.0.0/entity/{id}/tags | Update the tag information for a specific entity.


# **GetEntityTags**
> EntityTags GetEntityTags(ctx, id)
Retrieves the tag information for a specific entity.

Retrieves the tag information for a specific entity. If user has view right to the entity, user can view its tags. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| the URN of the entity to manage tag for. | 

### Return type

[**EntityTags**](EntityTags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntityTagsSummaries**
> EntityTagsSummaries GetEntityTagsSummaries(ctx, page, pageSize, optional)
Retrieves the list of entities that have at least one tag assigned to it.

Retrieves the list of entities that have at least one tag assigned to it. The list of tags is limited to the first 10 found. For the full list, use the API <code>/1.0.0/entity/{id}/tags</code>, which also documents the list of all supported \"entityType\" value in the sample query below:   <code>filter=(tag.value==Web;tag.scope==Security;entityType==urn:vcloud:vm)</code> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***EntityTagsApiGetEntityTagsSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntityTagsApiGetEntityTagsSummariesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EntityTagsSummaries**](EntityTagsSummaries.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEntityTags**
> EntityTags UpdateEntityTags(ctx, entityTags, id)
Update the tag information for a specific entity.

Update the tag information for a specific entity. An empty list of tags means to delete all dags for the entity. If user has edit permission on the entity, user can edit its tags. Only entities of specific types are supported. Supported types which can be determined based on a VCD Entity's URN ID.   <ul>   <li>urn:vcloud:vm   </ul> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityTags** | [**EntityTags**](EntityTags.md)| the list of tags to update. | 
  **id** | **string**| the URN of the entity to manage tag for. | 

### Return type

[**EntityTags**](EntityTags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

