# \OrgVdcNetworkDhcpApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworkDhcpConfig**](OrgVdcNetworkDhcpApi.md#DeleteNetworkDhcpConfig) | **Delete** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp | Removes Dhcp configuration on a specific Org vDC network.
[**GetNetworkDhcpConfig**](OrgVdcNetworkDhcpApi.md#GetNetworkDhcpConfig) | **Get** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp | Retrieves Dhcp configuration of a specific Org vDC network.
[**UpdateNetworkDhcpConfig**](OrgVdcNetworkDhcpApi.md#UpdateNetworkDhcpConfig) | **Put** /1.0.0/orgVdcNetworks/{vdcNetworkId}/dhcp | Updates Dhcp configuration of a specific Org vDC network.


# **DeleteNetworkDhcpConfig**
> DeleteNetworkDhcpConfig(ctx, vdcNetworkId)
Removes Dhcp configuration on a specific Org vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkDhcpConfig**
> VdcNetworkDhcpConfig GetNetworkDhcpConfig(ctx, vdcNetworkId)
Retrieves Dhcp configuration of a specific Org vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vdcNetworkId** | **string**|  | 

### Return type

[**VdcNetworkDhcpConfig**](VdcNetworkDhcpConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkDhcpConfig**
> UpdateNetworkDhcpConfig(ctx, dhcpConfig, vdcNetworkId)
Updates Dhcp configuration of a specific Org vDC network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dhcpConfig** | [**VdcNetworkDhcpConfig**](VdcNetworkDhcpConfig.md)|  | 
  **vdcNetworkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

