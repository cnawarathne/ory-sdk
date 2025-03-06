# SessionDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**authentication_methods** | Option<[**Vec<models::SessionAuthenticationMethod>**](sessionAuthenticationMethod.md)> | A list of authenticators which were used to authenticate the session. | [optional]
**id** | **String** | Device record ID | 
**ip_address** | Option<**String**> | IPAddress of the client | [optional]
**location** | Option<**String**> | Geo Location corresponding to the IP Address | [optional]
**trusted** | Option<**bool**> | Is this device trusted? (only matters if this device submitted aal2+ credentials) | [optional]
**user_agent** | Option<**String**> | UserAgent of the client | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


