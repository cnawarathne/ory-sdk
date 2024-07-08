# coding: utf-8

"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

    The version of the OpenAPI document: v1.13.6
    Contact: support@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from pydantic import BaseModel, ConfigDict, Field, StrictStr, field_validator
from typing import Any, ClassVar, Dict, List, Optional
from ory_client.models.normalized_project_revision import NormalizedProjectRevision
from ory_client.models.workspace import Workspace
from typing import Optional, Set
from typing_extensions import Self

class NormalizedProject(BaseModel):
    """
    NormalizedProject
    """ # noqa: E501
    created_at: datetime = Field(description="The Project's Creation Date")
    current_revision: NormalizedProjectRevision
    environment: StrictStr = Field(description="The environment of the project. prod Production stage Staging dev Development")
    home_region: StrictStr = Field(description="The project's data home region. eu-central EUCentral us-east USEast us-west USWest global Global")
    hosts: List[StrictStr]
    id: StrictStr = Field(description="The project's ID.")
    slug: StrictStr = Field(description="The project's slug")
    state: StrictStr = Field(description="The state of the project. running Running halted Halted deleted Deleted")
    subscription_id: Optional[StrictStr] = None
    subscription_plan: Optional[StrictStr] = None
    updated_at: datetime = Field(description="Last Time Project was Updated")
    workspace: Optional[Workspace] = None
    workspace_id: Optional[StrictStr]
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["created_at", "current_revision", "environment", "home_region", "hosts", "id", "slug", "state", "subscription_id", "subscription_plan", "updated_at", "workspace", "workspace_id"]

    @field_validator('environment')
    def environment_validate_enum(cls, value):
        """Validates the enum"""
        if value not in set(['prod', 'stage', 'dev']):
            raise ValueError("must be one of enum values ('prod', 'stage', 'dev')")
        return value

    @field_validator('home_region')
    def home_region_validate_enum(cls, value):
        """Validates the enum"""
        if value not in set(['eu-central', 'us-east', 'us-west', 'global']):
            raise ValueError("must be one of enum values ('eu-central', 'us-east', 'us-west', 'global')")
        return value

    @field_validator('state')
    def state_validate_enum(cls, value):
        """Validates the enum"""
        if value not in set(['running', 'halted', 'deleted']):
            raise ValueError("must be one of enum values ('running', 'halted', 'deleted')")
        return value

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
        """Create an instance of NormalizedProject from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        * OpenAPI `readOnly` fields are excluded.
        * OpenAPI `readOnly` fields are excluded.
        * OpenAPI `readOnly` fields are excluded.
        * OpenAPI `readOnly` fields are excluded.
        * OpenAPI `readOnly` fields are excluded.
        * OpenAPI `readOnly` fields are excluded.
        * Fields in `self.additional_properties` are added to the output dict.
        """
        excluded_fields: Set[str] = set([
            "created_at",
            "home_region",
            "id",
            "slug",
            "state",
            "updated_at",
            "additional_properties",
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of current_revision
        if self.current_revision:
            _dict['current_revision'] = self.current_revision.to_dict()
        # override the default output from pydantic by calling `to_dict()` of workspace
        if self.workspace:
            _dict['workspace'] = self.workspace.to_dict()
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        # set to None if subscription_id (nullable) is None
        # and model_fields_set contains the field
        if self.subscription_id is None and "subscription_id" in self.model_fields_set:
            _dict['subscription_id'] = None

        # set to None if subscription_plan (nullable) is None
        # and model_fields_set contains the field
        if self.subscription_plan is None and "subscription_plan" in self.model_fields_set:
            _dict['subscription_plan'] = None

        # set to None if workspace_id (nullable) is None
        # and model_fields_set contains the field
        if self.workspace_id is None and "workspace_id" in self.model_fields_set:
            _dict['workspace_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of NormalizedProject from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "created_at": obj.get("created_at"),
            "current_revision": NormalizedProjectRevision.from_dict(obj["current_revision"]) if obj.get("current_revision") is not None else None,
            "environment": obj.get("environment"),
            "home_region": obj.get("home_region"),
            "hosts": obj.get("hosts"),
            "id": obj.get("id"),
            "slug": obj.get("slug"),
            "state": obj.get("state"),
            "subscription_id": obj.get("subscription_id"),
            "subscription_plan": obj.get("subscription_plan"),
            "updated_at": obj.get("updated_at"),
            "workspace": Workspace.from_dict(obj["workspace"]) if obj.get("workspace") is not None else None,
            "workspace_id": obj.get("workspace_id")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


