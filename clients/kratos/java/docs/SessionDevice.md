

# SessionDevice

Device corresponding to a Session

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**authenticationMethods** | [**List&lt;SessionAuthenticationMethod&gt;**](SessionAuthenticationMethod.md) | A list of authenticators which were used to authenticate the session. |  [optional] |
|**id** | **String** | Device record ID |  |
|**ipAddress** | **String** | IPAddress of the client |  [optional] |
|**location** | **String** | Geo Location corresponding to the IP Address |  [optional] |
|**trusted** | **Boolean** | Is this device trusted? (only matters if this device submitted aal2+ credentials) |  [optional] |
|**userAgent** | **String** | UserAgent of the client |  [optional] |



