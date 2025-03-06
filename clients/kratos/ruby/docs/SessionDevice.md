# OryKratosClient::SessionDevice

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **authentication_methods** | [**Array&lt;SessionAuthenticationMethod&gt;**](SessionAuthenticationMethod.md) | A list of authenticators which were used to authenticate the session. | [optional] |
| **id** | **String** | Device record ID |  |
| **ip_address** | **String** | IPAddress of the client | [optional] |
| **location** | **String** | Geo Location corresponding to the IP Address | [optional] |
| **trusted** | **Boolean** | Is this device trusted? (only matters if this device submitted aal2+ credentials) | [optional] |
| **user_agent** | **String** | UserAgent of the client | [optional] |

## Example

```ruby
require 'ory-kratos-client'

instance = OryKratosClient::SessionDevice.new(
  authentication_methods: null,
  id: null,
  ip_address: null,
  location: null,
  trusted: null,
  user_agent: null
)
```

