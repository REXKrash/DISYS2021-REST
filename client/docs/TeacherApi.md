# \TeacherApi

All URIs are relative to *https://rest.itu.dk/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeacher**](TeacherApi.md#AddTeacher) | **Post** /teacher | Add a new teacher to the store
[**DeleteTeacher**](TeacherApi.md#DeleteTeacher) | **Delete** /teacher/{teacherId} | Deletes a teacher
[**GetTeacherById**](TeacherApi.md#GetTeacherById) | **Get** /teacher/{teacherId} | Find teacher by ID
[**ScoreTeacher**](TeacherApi.md#ScoreTeacher) | **Put** /teacher/{teacherId}/score/{score} | Set score for a specific teacher
[**UpdateTeacher**](TeacherApi.md#UpdateTeacher) | **Put** /teacher | Update an existing teacher
[**UpdateTeacherWithId**](TeacherApi.md#UpdateTeacherWithId) | **Post** /teacher/{teacherId} | Updates a teacher using specific ID


# **AddTeacher**
> AddTeacher(ctx, body)
Add a new teacher to the store



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Teacher**](Teacher.md)| Teacher object that needs to be added to the school | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTeacher**
> DeleteTeacher(ctx, teacherId, optional)
Deletes a teacher



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teacherId** | **int64**| teacher id to delete | 
 **optional** | ***TeacherApiDeleteTeacherOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeacherApiDeleteTeacherOpts struct

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

# **GetTeacherById**
> Teacher GetTeacherById(ctx, teacherId)
Find teacher by ID

Returns a single teacher

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teacherId** | **int64**| ID of teacher to return | 

### Return type

[**Teacher**](Teacher.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScoreTeacher**
> ScoreTeacher(ctx, teacherId, score)
Set score for a specific teacher



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teacherId** | **int64**| ID of teacher to score | 
  **score** | **int64**| Score for the teacher | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeacher**
> UpdateTeacher(ctx, body)
Update an existing teacher



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Teacher**](Teacher.md)| Teacher object that needs to be updated | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeacherWithId**
> UpdateTeacherWithId(ctx, teacherId, body)
Updates a teacher using specific ID



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teacherId** | **int64**| ID of teacher that needs to be updated | 
  **body** | [**Teacher**](Teacher.md)| order placed for purchasing the teacher | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

