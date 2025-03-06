# OryKratosClient::SessionAuthenticationMethod

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **aal** | [**AuthenticatorAssuranceLevel**](AuthenticatorAssuranceLevel.md) |  | [optional] |
| **completed_at** | **Time** | When the authentication challenge was completed. | [optional] |
| **device_trust_based** | **Boolean** | DeviceTrustBased indicates that this authentication method was added due to device trust | [optional] |
| **method** | **String** |  | [optional] |
| **organization** | **String** | The Organization id used for authentication | [optional] |
| **provider** | **String** | OIDC or SAML provider id used for authentication | [optional] |

## Example

```ruby
require 'ory-kratos-client'

instance = OryKratosClient::SessionAuthenticationMethod.new(
  aal: null,
  completed_at: null,
  device_trust_based: null,
  method: null,
  organization: null,
  provider: null
)
```

