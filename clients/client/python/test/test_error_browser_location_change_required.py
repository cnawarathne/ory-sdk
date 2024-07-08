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

from ory_client.models.error_browser_location_change_required import ErrorBrowserLocationChangeRequired

class TestErrorBrowserLocationChangeRequired(unittest.TestCase):
    """ErrorBrowserLocationChangeRequired unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ErrorBrowserLocationChangeRequired:
        """Test ErrorBrowserLocationChangeRequired
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ErrorBrowserLocationChangeRequired`
        """
        model = ErrorBrowserLocationChangeRequired()
        if include_optional:
            return ErrorBrowserLocationChangeRequired(
                error = ory_client.models.json_api_error_response.JSON API Error Response(
                    error = ory_client.models.generic_error_content.genericErrorContent(
                        debug = 'The database adapter was unable to find the element', 
                        error_description = 'Object with ID 12345 does not exist', 
                        message = '', 
                        status_code = 404, ), ),
                redirect_browser_to = ''
            )
        else:
            return ErrorBrowserLocationChangeRequired(
        )
        """

    def testErrorBrowserLocationChangeRequired(self):
        """Test ErrorBrowserLocationChangeRequired"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
