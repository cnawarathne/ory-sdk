"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v1.4.0
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_client
from ory_client.model.attributes_count_datapoint import AttributesCountDatapoint
globals()['AttributesCountDatapoint'] = AttributesCountDatapoint
from ory_client.model.get_attributes_count_response import GetAttributesCountResponse


class TestGetAttributesCountResponse(unittest.TestCase):
    """GetAttributesCountResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testGetAttributesCountResponse(self):
        """Test GetAttributesCountResponse"""
        # FIXME: construct object with mandatory attributes with example values
        # model = GetAttributesCountResponse()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
