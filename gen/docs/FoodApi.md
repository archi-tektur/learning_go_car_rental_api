# FoodApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**foodGet**](FoodApi.md#foodGet) | **GET** /food | Get all food
[**foodIdBitePost**](FoodApi.md#foodIdBitePost) | **POST** /food/{id}/bite | Post food bite
[**foodPost**](FoodApi.md#foodPost) | **POST** /food | Post food


<a name="foodGet"></a>
# **foodGet**
> List&lt;ApplicationFoodDTO&gt; foodGet()

Get all food

Get all food

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.FoodApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080");

    FoodApi apiInstance = new FoodApi(defaultClient);
    try {
      List<ApplicationFoodDTO> result = apiInstance.foodGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling FoodApi#foodGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List&lt;ApplicationFoodDTO&gt;**](ApplicationFoodDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad request |  -  |
**404** | Not found |  -  |
**500** | Internal Server Error |  -  |

<a name="foodIdBitePost"></a>
# **foodIdBitePost**
> String foodIdBitePost(id)

Post food bite

Bite a selected food

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.FoodApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080");

    FoodApi apiInstance = new FoodApi(defaultClient);
    Integer id = 56; // Integer | Food ID
    try {
      String result = apiInstance.foodIdBitePost(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling FoodApi#foodIdBitePost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Integer**| Food ID |

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad request |  -  |
**404** | Not found |  -  |
**500** | Internal Server Error |  -  |

<a name="foodPost"></a>
# **foodPost**
> String foodPost()

Post food

Create new food item

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.FoodApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080");

    FoodApi apiInstance = new FoodApi(defaultClient);
    try {
      String result = apiInstance.foodPost();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling FoodApi#foodPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |
**400** | Bad request |  -  |
**404** | Not found |  -  |
**500** | Internal Server Error |  -  |

