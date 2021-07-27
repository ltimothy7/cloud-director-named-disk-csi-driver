# \VmTagsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVmTags**](VmTagsApi.md#GetVmTags) | **Get** /1.0.0/vms/{vmId}/tags | Retrieves all the tags for the VM.
[**UpdateVmTags**](VmTagsApi.md#UpdateVmTags) | **Put** /1.0.0/vms/{vmId}/tags | Update the tag for the VM.


# **GetVmTags**
> EntityTags GetVmTags(ctx, vmId)
Retrieves all the tags for the VM.

Retrieves all the tags for the VM. User can view the tags if user can view the VM. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **string**| the URN of the VM to manage tag for. | 

### Return type

[**EntityTags**](EntityTags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVmTags**
> EntityTags UpdateVmTags(ctx, entityTags, vmId)
Update the tag for the VM.

Update the tag for the VM. An empty list of tags means to delete all dags for the VM. User can edit the tags if user can modify the VM. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityTags** | [**EntityTags**](EntityTags.md)| the list of tags to update. | 
  **vmId** | **string**| the URN of the VM to manage tag for. | 

### Return type

[**EntityTags**](EntityTags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

