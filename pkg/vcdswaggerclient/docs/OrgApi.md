# \OrgApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrg**](OrgApi.md#GetOrg) | **Get** /1.0.0/orgs/{orgUrn} | Get specified organization.
[**QueryOrgs**](OrgApi.md#QueryOrgs) | **Get** /1.0.0/orgs | Get list of all orgs accessible to the user.


# **GetOrg**
> Org GetOrg(ctx, orgUrn)
Get specified organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgUrn** | **string**| orgUrn | 

### Return type

[**Org**](Org.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryOrgs**
> Orgs QueryOrgs(ctx, page, pageSize, optional)
Get list of all orgs accessible to the user.

Get list of all orgs accessible to the user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***OrgApiQueryOrgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgApiQueryOrgsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**Orgs**](Orgs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

