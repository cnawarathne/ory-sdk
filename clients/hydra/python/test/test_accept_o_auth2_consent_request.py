# coding: utf-8

"""
    Ory Hydra API

    Documentation for all of Ory Hydra's APIs. 

    The version of the OpenAPI document: v2.4.0-alpha.1
    Contact: hi@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from ory_hydra_client.models.accept_o_auth2_consent_request import AcceptOAuth2ConsentRequest

class TestAcceptOAuth2ConsentRequest(unittest.TestCase):
    """AcceptOAuth2ConsentRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AcceptOAuth2ConsentRequest:
        """Test AcceptOAuth2ConsentRequest
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AcceptOAuth2ConsentRequest`
        """
        model = AcceptOAuth2ConsentRequest()
        if include_optional:
            return AcceptOAuth2ConsentRequest(
                context = None,
                grant_access_token_audience = [
                    ''
                    ],
                grant_scope = [
                    ''
                    ],
                handled_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                remember = True,
                remember_for = 56,
                session = ory_hydra_client.models.pass_session_data_to_a_consent_request/.Pass session data to a consent request.(
                    access_token = null, 
                    id_token = null, )
            )
        else:
            return AcceptOAuth2ConsentRequest(
        )
        """

    def testAcceptOAuth2ConsentRequest(self):
        """Test AcceptOAuth2ConsentRequest"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
