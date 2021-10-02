# \StudentApi

All URIs are relative to *https://rest.itu.dk/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStudent**](StudentApi.md#AddStudent) | **Post** /student | Add a new student to the store
[**DeleteStudent**](StudentApi.md#DeleteStudent) | **Delete** /student/{studentId} | Deletes a student
[**GetStudentById**](StudentApi.md#GetStudentById) | **Get** /student/{studentId} | Find student by ID
[**UpdateStudent**](StudentApi.md#UpdateStudent) | **Put** /student | Update an existing student
[**UpdateStudentWithId**](StudentApi.md#UpdateStudentWithId) | **Post** /student/{studentId} | Updates a student using specific ID


# **AddStudent**
> AddStudent(ctx, body)
Add a new student to the store



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Student**](Student.md)| Student object that needs to be added to the school | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStudent**
> DeleteStudent(ctx, studentId, optional)
Deletes a student



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **studentId** | **int64**| student id to delete | 
 **optional** | ***StudentApiDeleteStudentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StudentApiDeleteStudentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStudentById**
> Student GetStudentById(ctx, studentId)
Find student by ID

Returns a single student

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **studentId** | **int64**| ID of student to return | 

### Return type

[**Student**](Student.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStudent**
> UpdateStudent(ctx, body)
Update an existing student



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Student**](Student.md)| Student object that needs to be updated | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStudentWithId**
> UpdateStudentWithId(ctx, studentId, body)
Updates a student using specific ID



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **studentId** | **int64**| ID of student that needs to be updated | 
  **body** | [**Student**](Student.md)| order placed for purchasing the student | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

