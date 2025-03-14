# ory_kratos_client.model.SessionAuthenticationMethod

## Load the model package
```dart
import 'package:ory_kratos_client/api.dart';
```

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**aal** | [**AuthenticatorAssuranceLevel**](AuthenticatorAssuranceLevel.md) |  | [optional] 
**completedAt** | [**DateTime**](DateTime.md) | When the authentication challenge was completed. | [optional] 
**deviceTrustBased** | **bool** | DeviceTrustBased indicates that this authentication method was added due to device trust | [optional] 
**method** | **String** |  | [optional] 
**organization** | **String** | The Organization id used for authentication | [optional] 
**provider** | **String** | OIDC or SAML provider id used for authentication | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


