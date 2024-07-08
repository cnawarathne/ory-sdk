# coding: utf-8

"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

    The version of the OpenAPI document: v1.13.6
    Contact: support@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from ory_client.models.oidc_configuration import OidcConfiguration

class TestOidcConfiguration(unittest.TestCase):
    """OidcConfiguration unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> OidcConfiguration:
        """Test OidcConfiguration
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `OidcConfiguration`
        """
        model = OidcConfiguration()
        if include_optional:
            return OidcConfiguration(
                authorization_endpoint = 'https://playground.ory.sh/ory-hydra/public/oauth2/auth',
                backchannel_logout_session_supported = True,
                backchannel_logout_supported = True,
                claims_parameter_supported = True,
                claims_supported = [
                    ''
                    ],
                code_challenge_methods_supported = [
                    ''
                    ],
                credentials_endpoint_draft_00 = '',
                credentials_supported_draft_00 = [
                    ory_client.models.verifiable_credentials_metadata_(draft_00).Verifiable Credentials Metadata (Draft 00)(
                        cryptographic_binding_methods_supported = [
                            ''
                            ], 
                        cryptographic_suites_supported = [
                            ''
                            ], 
                        format = '', 
                        types = [
                            ''
                            ], )
                    ],
                end_session_endpoint = '',
                frontchannel_logout_session_supported = True,
                frontchannel_logout_supported = True,
                grant_types_supported = [
                    ''
                    ],
                id_token_signed_response_alg = [
                    ''
                    ],
                id_token_signing_alg_values_supported = [
                    ''
                    ],
                issuer = 'https://playground.ory.sh/ory-hydra/public/',
                jwks_uri = 'https://{slug}.projects.oryapis.com/.well-known/jwks.json',
                registration_endpoint = 'https://playground.ory.sh/ory-hydra/admin/client',
                request_object_signing_alg_values_supported = [
                    ''
                    ],
                request_parameter_supported = True,
                request_uri_parameter_supported = True,
                require_request_uri_registration = True,
                response_modes_supported = [
                    ''
                    ],
                response_types_supported = [
                    ''
                    ],
                revocation_endpoint = '',
                scopes_supported = [
                    ''
                    ],
                subject_types_supported = [
                    ''
                    ],
                token_endpoint = 'https://playground.ory.sh/ory-hydra/public/oauth2/token',
                token_endpoint_auth_methods_supported = [
                    ''
                    ],
                userinfo_endpoint = '',
                userinfo_signed_response_alg = [
                    ''
                    ],
                userinfo_signing_alg_values_supported = [
                    ''
                    ]
            )
        else:
            return OidcConfiguration(
                authorization_endpoint = 'https://playground.ory.sh/ory-hydra/public/oauth2/auth',
                id_token_signed_response_alg = [
                    ''
                    ],
                id_token_signing_alg_values_supported = [
                    ''
                    ],
                issuer = 'https://playground.ory.sh/ory-hydra/public/',
                jwks_uri = 'https://{slug}.projects.oryapis.com/.well-known/jwks.json',
                response_types_supported = [
                    ''
                    ],
                subject_types_supported = [
                    ''
                    ],
                token_endpoint = 'https://playground.ory.sh/ory-hydra/public/oauth2/token',
                userinfo_signed_response_alg = [
                    ''
                    ],
        )
        """

    def testOidcConfiguration(self):
        """Test OidcConfiguration"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
