# \CourseApi

All URIs are relative to *https://rest.itu.dk/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCourse**](CourseApi.md#AddCourse) | **Post** /course | Add a new course to the store
[**DeleteCourse**](CourseApi.md#DeleteCourse) | **Delete** /course/{courseId} | Deletes a course
[**EnrollStudent**](CourseApi.md#EnrollStudent) | **Put** /course/{courseId}/enroll/{studentId} | Enroll student to a specific course
[**GetCourseById**](CourseApi.md#GetCourseById) | **Get** /course/{courseId} | Find course by ID
[**RateCourse**](CourseApi.md#RateCourse) | **Put** /course/{courseId}/rating/{rating} | Set rating for a specific course
[**UpdateCourse**](CourseApi.md#UpdateCourse) | **Put** /course | Update an existing course
[**UpdateCourseWithId**](CourseApi.md#UpdateCourseWithId) | **Post** /course/{courseId} | Updates a course using specific ID


# **AddCourse**
> AddCourse(ctx, body)
Add a new course to the store



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Course**](Course.md)| Course object that needs to be added to the school | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCourse**
> DeleteCourse(ctx, courseId, optional)
Deletes a course



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int64**| course id to delete | 
 **optional** | ***CourseApiDeleteCourseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CourseApiDeleteCourseOpts struct

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

# **EnrollStudent**
> EnrollStudent(ctx, courseId, studentId)
Enroll student to a specific course



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int64**| ID of course to enroll student in | 
  **studentId** | **int64**| ID of student to enroll | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCourseById**
> Course GetCourseById(ctx, courseId)
Find course by ID

Returns a single course

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int64**| ID of course to return | 

### Return type

[**Course**](Course.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RateCourse**
> RateCourse(ctx, courseId, rating)
Set rating for a specific course



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int64**| ID of course to rate | 
  **rating** | **int64**| Rating for the course | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCourse**
> UpdateCourse(ctx, body)
Update an existing course



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Course**](Course.md)| Course object that needs to be updated | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCourseWithId**
> UpdateCourseWithId(ctx, courseId, body)
Updates a course using specific ID



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int64**| ID of course that needs to be updated | 
  **body** | [**Course**](Course.md)| order placed for purchasing the course | 

### Return type

 (empty response body)

### Authorization

[school_auth](../README.md#school_auth)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

