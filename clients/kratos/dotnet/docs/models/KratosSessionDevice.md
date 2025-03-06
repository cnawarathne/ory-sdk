# Ory.Kratos.Client.Model.KratosSessionDevice
Device corresponding to a Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Device record ID | 
**AuthenticationMethods** | [**List&lt;KratosSessionAuthenticationMethod&gt;**](KratosSessionAuthenticationMethod.md) | A list of authenticators which were used to authenticate the session. | [optional] 
**IpAddress** | **string** | IPAddress of the client | [optional] 
**Location** | **string** | Geo Location corresponding to the IP Address | [optional] 
**Trusted** | **bool** | Is this device trusted? (only matters if this device submitted aal2+ credentials) | [optional] 
**UserAgent** | **string** | UserAgent of the client | [optional] 

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)

