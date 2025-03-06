# # SessionDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**authenticationMethods** | [**\Ory\Kratos\Client\Model\SessionAuthenticationMethod[]**](SessionAuthenticationMethod.md) | A list of authenticators which were used to authenticate the session. | [optional]
**id** | **string** | Device record ID |
**ipAddress** | **string** | IPAddress of the client | [optional]
**location** | **string** | Geo Location corresponding to the IP Address | [optional]
**trusted** | **bool** | Is this device trusted? (only matters if this device submitted aal2+ credentials) | [optional]
**userAgent** | **string** | UserAgent of the client | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
