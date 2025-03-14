# coding: utf-8

"""
    Ory Identities API

    This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 

    The version of the OpenAPI document: v1.4.0-alpha.0
    Contact: office@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from ory_kratos_client.models.o_auth2_consent_request_open_id_connect_context import OAuth2ConsentRequestOpenIDConnectContext

class TestOAuth2ConsentRequestOpenIDConnectContext(unittest.TestCase):
    """OAuth2ConsentRequestOpenIDConnectContext unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> OAuth2ConsentRequestOpenIDConnectContext:
        """Test OAuth2ConsentRequestOpenIDConnectContext
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `OAuth2ConsentRequestOpenIDConnectContext`
        """
        model = OAuth2ConsentRequestOpenIDConnectContext()
        if include_optional:
            return OAuth2ConsentRequestOpenIDConnectContext(
                acr_values = [
                    ''
                    ],
                display = '',
                id_token_hint_claims = {
                    'key' : null
                    },
                login_hint = '',
                ui_locales = [
                    ''
                    ]
            )
        else:
            return OAuth2ConsentRequestOpenIDConnectContext(
        )
        """

    def testOAuth2ConsentRequestOpenIDConnectContext(self):
        """Test OAuth2ConsentRequestOpenIDConnectContext"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
