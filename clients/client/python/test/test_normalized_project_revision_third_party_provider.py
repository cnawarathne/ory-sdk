# coding: utf-8

"""
    Ory APIs

    # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

    The version of the OpenAPI document: v1.16.10
    Contact: support@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from ory_client.models.normalized_project_revision_third_party_provider import NormalizedProjectRevisionThirdPartyProvider

class TestNormalizedProjectRevisionThirdPartyProvider(unittest.TestCase):
    """NormalizedProjectRevisionThirdPartyProvider unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> NormalizedProjectRevisionThirdPartyProvider:
        """Test NormalizedProjectRevisionThirdPartyProvider
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `NormalizedProjectRevisionThirdPartyProvider`
        """
        model = NormalizedProjectRevisionThirdPartyProvider()
        if include_optional:
            return NormalizedProjectRevisionThirdPartyProvider(
                additional_id_token_audiences = [
                    ''
                    ],
                apple_private_key = '',
                apple_private_key_id = 'UX56C66723',
                apple_team_id = 'KP76DQS54M',
                auth_url = 'https://www.googleapis.com/oauth2/v2/auth',
                azure_tenant = 'contoso.onmicrosoft.com',
                claims_source = '',
                client_id = '',
                client_secret = '',
                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                fedcm_config_url = '',
                id = '',
                issuer_url = 'https://accounts.google.com',
                label = '',
                mapper_url = '',
                net_id_token_origin_header = '',
                organization_id = '',
                pkce = 'auto',
                project_revision_id = '',
                provider = 'google',
                provider_id = '',
                requested_claims = ory_client.models.json_raw_message_represents_a_json/raw_message_that_works_well_with_json,_sql,_and_swagger/.JSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger.(),
                scope = [
                    ''
                    ],
                state = 'enabled',
                subject_source = '',
                token_url = 'https://www.googleapis.com/oauth2/v4/token',
                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else:
            return NormalizedProjectRevisionThirdPartyProvider(
        )
        """

    def testNormalizedProjectRevisionThirdPartyProvider(self):
        """Test NormalizedProjectRevisionThirdPartyProvider"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
