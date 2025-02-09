# coding: utf-8

"""
    Ory Hydra API

    Documentation for all of Ory Hydra's APIs. 

    The version of the OpenAPI document: v2.4.0-alpha.1
    Contact: hi@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from pydantic import BaseModel, ConfigDict, Field, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from ory_hydra_client.models.o_auth2_client import OAuth2Client
from typing import Optional, Set
from typing_extensions import Self

class VerifyUserCodeRequest(BaseModel):
    """
    VerifyUserCodeRequest
    """ # noqa: E501
    challenge: Optional[StrictStr] = Field(default=None, description="ID is the identifier (\"device challenge\") of the device request. It is used to identify the session.")
    client: Optional[OAuth2Client] = None
    device_code_request_id: Optional[StrictStr] = None
    handled_at: Optional[datetime] = None
    request_url: Optional[StrictStr] = Field(default=None, description="RequestURL is the original Device Authorization URL requested.")
    requested_access_token_audience: Optional[List[StrictStr]] = None
    requested_scope: Optional[List[StrictStr]] = None
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["challenge", "client", "device_code_request_id", "handled_at", "request_url", "requested_access_token_audience", "requested_scope"]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of VerifyUserCodeRequest from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        * Fields in `self.additional_properties` are added to the output dict.
        """
        excluded_fields: Set[str] = set([
            "additional_properties",
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of client
        if self.client:
            _dict['client'] = self.client.to_dict()
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of VerifyUserCodeRequest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "challenge": obj.get("challenge"),
            "client": OAuth2Client.from_dict(obj["client"]) if obj.get("client") is not None else None,
            "device_code_request_id": obj.get("device_code_request_id"),
            "handled_at": obj.get("handled_at"),
            "request_url": obj.get("request_url"),
            "requested_access_token_audience": obj.get("requested_access_token_audience"),
            "requested_scope": obj.get("requested_scope")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


