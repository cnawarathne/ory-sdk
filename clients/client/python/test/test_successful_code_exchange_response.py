# coding: utf-8

"""
    Ory APIs

    # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

    The version of the OpenAPI document: v1.17.1
    Contact: support@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from ory_client.models.successful_code_exchange_response import SuccessfulCodeExchangeResponse

class TestSuccessfulCodeExchangeResponse(unittest.TestCase):
    """SuccessfulCodeExchangeResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SuccessfulCodeExchangeResponse:
        """Test SuccessfulCodeExchangeResponse
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SuccessfulCodeExchangeResponse`
        """
        model = SuccessfulCodeExchangeResponse()
        if include_optional:
            return SuccessfulCodeExchangeResponse(
                session = ory_client.models.session.session(
                    active = True, 
                    authenticated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    authentication_methods = [
                        ory_client.models.authentication_method_identifies_an_authentication_method.AuthenticationMethod identifies an authentication method(
                            aal = 'aal0', 
                            completed_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            method = 'link_recovery', 
                            organization = '', 
                            provider = '', )
                        ], 
                    authenticator_assurance_level = 'aal0', 
                    devices = [
                        ory_client.models.session_device.sessionDevice(
                            id = '', 
                            ip_address = '', 
                            location = '', 
                            user_agent = '', )
                        ], 
                    expires_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    id = '', 
                    identity = ory_client.models.identity_represents_an_ory_kratos_identity.Identity represents an Ory Kratos identity(
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        credentials = {
                            'key' : ory_client.models.identity_credentials.identityCredentials(
                                config = ory_client.models.json_raw_message_represents_a_json/raw_message_that_works_well_with_json,_sql,_and_swagger/.JSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger.(), 
                                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                identifiers = [
                                    ''
                                    ], 
                                type = 'password', 
                                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                version = 56, )
                            }, 
                        id = '', 
                        metadata_admin = ory_client.models.null_json_raw_message.nullJsonRawMessage(), 
                        metadata_public = ory_client.models.null_json_raw_message.nullJsonRawMessage(), 
                        organization_id = '', 
                        recovery_addresses = [
                            ory_client.models.recovery_identity_address.recoveryIdentityAddress(
                                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                id = '', 
                                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                value = '', 
                                via = '', )
                            ], 
                        schema_id = '', 
                        schema_url = '', 
                        state = 'active', 
                        state_changed_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        traits = null, 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        verifiable_addresses = [
                            ory_client.models.verifiable_identity_address.verifiableIdentityAddress(
                                created_at = '2014-01-01T23:28:56.782Z', 
                                id = '', 
                                status = '', 
                                updated_at = '2014-01-01T23:28:56.782Z', 
                                value = '', 
                                verified = True, 
                                verified_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                via = 'email', )
                            ], ), 
                    issued_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    tokenized = '', ),
                session_token = ''
            )
        else:
            return SuccessfulCodeExchangeResponse(
                session = ory_client.models.session.session(
                    active = True, 
                    authenticated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    authentication_methods = [
                        ory_client.models.authentication_method_identifies_an_authentication_method.AuthenticationMethod identifies an authentication method(
                            aal = 'aal0', 
                            completed_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            method = 'link_recovery', 
                            organization = '', 
                            provider = '', )
                        ], 
                    authenticator_assurance_level = 'aal0', 
                    devices = [
                        ory_client.models.session_device.sessionDevice(
                            id = '', 
                            ip_address = '', 
                            location = '', 
                            user_agent = '', )
                        ], 
                    expires_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    id = '', 
                    identity = ory_client.models.identity_represents_an_ory_kratos_identity.Identity represents an Ory Kratos identity(
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        credentials = {
                            'key' : ory_client.models.identity_credentials.identityCredentials(
                                config = ory_client.models.json_raw_message_represents_a_json/raw_message_that_works_well_with_json,_sql,_and_swagger/.JSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger.(), 
                                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                identifiers = [
                                    ''
                                    ], 
                                type = 'password', 
                                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                version = 56, )
                            }, 
                        id = '', 
                        metadata_admin = ory_client.models.null_json_raw_message.nullJsonRawMessage(), 
                        metadata_public = ory_client.models.null_json_raw_message.nullJsonRawMessage(), 
                        organization_id = '', 
                        recovery_addresses = [
                            ory_client.models.recovery_identity_address.recoveryIdentityAddress(
                                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                id = '', 
                                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                value = '', 
                                via = '', )
                            ], 
                        schema_id = '', 
                        schema_url = '', 
                        state = 'active', 
                        state_changed_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        traits = null, 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        verifiable_addresses = [
                            ory_client.models.verifiable_identity_address.verifiableIdentityAddress(
                                created_at = '2014-01-01T23:28:56.782Z', 
                                id = '', 
                                status = '', 
                                updated_at = '2014-01-01T23:28:56.782Z', 
                                value = '', 
                                verified = True, 
                                verified_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                via = 'email', )
                            ], ), 
                    issued_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    tokenized = '', ),
        )
        """

    def testSuccessfulCodeExchangeResponse(self):
        """Test SuccessfulCodeExchangeResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
