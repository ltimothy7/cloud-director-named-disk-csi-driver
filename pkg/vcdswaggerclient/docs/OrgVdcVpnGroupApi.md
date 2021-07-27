# \OrgVdcVpnGroupApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVdcVpnGroup**](OrgVdcVpnGroupApi.md#CreateVdcVpnGroup) | **Post** /1.0.0/orgVdcVpnGroups | creates an Org VDC VPN Group
[**DeleteVdcVpnGroup**](OrgVdcVpnGroupApi.md#DeleteVdcVpnGroup) | **Delete** /1.0.0/orgVdcVpnGroups/{vdcVpnGroupId} | deletes specified Org VDC VPN Grouping
[**GetVdcVpnGroup**](OrgVdcVpnGroupApi.md#GetVdcVpnGroup) | **Get** /1.0.0/orgVdcVpnGroups/{vdcVpnGroupId} | gets an Org VDC VPN Group
[**GetVdcVpnGroups**](OrgVdcVpnGroupApi.md#GetVdcVpnGroups) | **Get** /1.0.0/orgVdcVpnGroups | gets all Org VDC VPN Groups
[**UpdateVdcVpnGroup**](OrgVdcVpnGroupApi.md#UpdateVdcVpnGroup) | **Put** /1.0.0/orgVdcVpnGroups/{vdcVpnGroupId} | updates an Org VDC VPN Group


# **CreateVdcVpnGroup**
> CreateVdcVpnGroup(ctx, vdcVpnGroup)
creates an Org VDC VPN Group

Takes an Org VDC VPN Group object and establishes a VPN connection between each pair of VDCs (and the associated networks) within the object 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcVpnGroup** | [**VdcVpnGroup**](VdcVpnGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVdcVpnGroup**
> DeleteVdcVpnGroup(ctx, vdcVpnGroupId)
deletes specified Org VDC VPN Grouping

Deletes a singular Org VDC VPN Group with the given ID 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcVpnGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcVpnGroup**
> VdcVpnGroup GetVdcVpnGroup(ctx, vdcVpnGroupId)
gets an Org VDC VPN Group

Retrieves a singular Org VDC VPN Group with the given ID 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcVpnGroupId** | **string**|  | 

### Return type

[**VdcVpnGroup**](VdcVpnGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVdcVpnGroups**
> VdcVpnGroups GetVdcVpnGroups(ctx, page, pageSize, optional)
gets all Org VDC VPN Groups

Gets all the Org VDC VPN Groups available to the user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***OrgVdcVpnGroupApiGetVdcVpnGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgVdcVpnGroupApiGetVdcVpnGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**VdcVpnGroups**](VdcVpnGroups.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVdcVpnGroup**
> UpdateVdcVpnGroup(ctx, vdcVpnGroup, vdcVpnGroupId)
updates an Org VDC VPN Group

Updates a singular Org VDC VON Group with the given ID 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcVpnGroup** | [**VdcVpnGroup**](VdcVpnGroup.md)|  | 
  **vdcVpnGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

