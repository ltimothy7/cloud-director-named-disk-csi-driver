# \EdgeClustersApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEdgeCluster**](EdgeClustersApi.md#CreateEdgeCluster) | **Post** /1.0.0/edgeClusters | Create a new Edge Cluster
[**GetEdgeClusters**](EdgeClustersApi.md#GetEdgeClusters) | **Get** /1.0.0/edgeClusters | Get all Edge Clusters in the system


# **CreateEdgeCluster**
> CreateEdgeCluster(ctx, edgeCluster)
Create a new Edge Cluster

Create a new Edge Cluster. An Edge Cluster is defined by a Resouce Pool and Storage Profile for deploying Edge Gateways. It can subsequently be assigned to a given vDC Network Profile as a Primary or Secondary Edge Cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeCluster** | [**EdgeCluster**](EdgeCluster.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeClusters**
> EdgeClusters GetEdgeClusters(ctx, page, pageSize, optional)
Get all Edge Clusters in the system

Retrieves all Edge Clusters. Results can be filtered by id, name and vCenter (resourcePool.vcId). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
 **optional** | ***EdgeClustersApiGetEdgeClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EdgeClustersApiGetEdgeClustersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Filter for a query.  FIQL format. | 
 **sortAsc** | **optional.String**| Field to use for ascending sort | 
 **sortDesc** | **optional.String**| Field to use for descending sort | 

### Return type

[**EdgeClusters**](EdgeClusters.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=35.0

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

