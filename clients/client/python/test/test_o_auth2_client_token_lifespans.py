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

from ory_client.models.o_auth2_client_token_lifespans import OAuth2ClientTokenLifespans

class TestOAuth2ClientTokenLifespans(unittest.TestCase):
    """OAuth2ClientTokenLifespans unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> OAuth2ClientTokenLifespans:
        """Test OAuth2ClientTokenLifespans
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `OAuth2ClientTokenLifespans`
        """
        model = OAuth2ClientTokenLifespans()
        if include_optional:
            return OAuth2ClientTokenLifespans(
                authorization_code_grant_access_token_lifespan = '4ms',
                authorization_code_grant_id_token_lifespan = '4ms',
                authorization_code_grant_refresh_token_lifespan = '4ms',
                client_credentials_grant_access_token_lifespan = '4ms',
                implicit_grant_access_token_lifespan = '4ms',
                implicit_grant_id_token_lifespan = '4ms',
                jwt_bearer_grant_access_token_lifespan = '4ms',
                refresh_token_grant_access_token_lifespan = '4ms',
                refresh_token_grant_id_token_lifespan = '4ms',
                refresh_token_grant_refresh_token_lifespan = '4ms'
            )
        else:
            return OAuth2ClientTokenLifespans(
        )
        """

    def testOAuth2ClientTokenLifespans(self):
        """Test OAuth2ClientTokenLifespans"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
