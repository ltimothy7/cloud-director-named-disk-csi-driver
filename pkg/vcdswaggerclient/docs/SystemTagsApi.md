# \SystemTagsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemTags**](SystemTagsApi.md#GetSystemTags) | **Get** /1.0.0/tags/systemTags | Retrieves the list of system tags.
[**UpdateSystemTags**](SystemTagsApi.md#UpdateSystemTags) | **Put** /1.0.0/tags/systemTags | Update the list of system tags.


# **GetSystemTags**
> SystemTags GetSystemTags(ctx, )
Retrieves the list of system tags.

Retrieves the list of system tags.  These tags are available for all users across all organizations to view (i.e. <code>/1.0.0/entity/{id}/tags</code>) even if there is no VCD entity assigned to the tag. These tags have no scope and user can set the scope during tag assignment to a VCD entity. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemTags**](SystemTags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSystemTags**
> SystemTags UpdateSystemTags(ctx, systemTags)
Update the list of system tags.

Update the list of system tags.  Note if any system tag is removed, VCD entities that use those tags are unaffected. This API is restricted to system admin providers. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemTags** | [**SystemTags**](SystemTags.md)| the list of tags to update. | 

### Return type

[**SystemTags**](SystemTags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

