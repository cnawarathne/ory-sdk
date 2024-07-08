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

from ory_client.models.update_settings_flow_with_passkey_method import UpdateSettingsFlowWithPasskeyMethod

class TestUpdateSettingsFlowWithPasskeyMethod(unittest.TestCase):
    """UpdateSettingsFlowWithPasskeyMethod unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> UpdateSettingsFlowWithPasskeyMethod:
        """Test UpdateSettingsFlowWithPasskeyMethod
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `UpdateSettingsFlowWithPasskeyMethod`
        """
        model = UpdateSettingsFlowWithPasskeyMethod()
        if include_optional:
            return UpdateSettingsFlowWithPasskeyMethod(
                csrf_token = '',
                method = '',
                passkey_remove = '',
                passkey_settings_register = ''
            )
        else:
            return UpdateSettingsFlowWithPasskeyMethod(
                method = '',
        )
        """

    def testUpdateSettingsFlowWithPasskeyMethod(self):
        """Test UpdateSettingsFlowWithPasskeyMethod"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
