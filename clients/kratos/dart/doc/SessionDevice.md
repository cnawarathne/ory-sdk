# ory_kratos_client.model.SessionDevice

## Load the model package
```dart
import 'package:ory_kratos_client/api.dart';
```

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**authenticationMethods** | [**BuiltList&lt;SessionAuthenticationMethod&gt;**](SessionAuthenticationMethod.md) | A list of authenticators which were used to authenticate the session. | [optional] 
**id** | **String** | Device record ID | 
**ipAddress** | **String** | IPAddress of the client | [optional] 
**location** | **String** | Geo Location corresponding to the IP Address | [optional] 
**trusted** | **bool** | Is this device trusted? (only matters if this device submitted aal2+ credentials) | [optional] 
**userAgent** | **String** | UserAgent of the client | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


