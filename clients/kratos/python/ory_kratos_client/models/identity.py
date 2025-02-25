# coding: utf-8

"""
    Ory Identities API

    This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 

    The version of the OpenAPI document: v1.3.5
    Contact: office@ory.sh
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
from ory_kratos_client.models.identity_credentials import IdentityCredentials
from ory_kratos_client.models.recovery_identity_address import RecoveryIdentityAddress
from ory_kratos_client.models.verifiable_identity_address import VerifiableIdentityAddress
from typing import Optional, Set
from typing_extensions import Self

class Identity(BaseModel):
    """
    An [identity](https://www.ory.sh/docs/kratos/concepts/identity-user-model) represents a (human) user in Ory.
    """ # noqa: E501
    created_at: Optional[datetime] = Field(default=None, description="CreatedAt is a helper struct field for gobuffalo.pop.")
    credentials: Optional[Dict[str, IdentityCredentials]] = Field(default=None, description="Credentials represents all credentials that can be used for authenticating this identity.")
    id: StrictStr = Field(description="ID is the identity's unique identifier.  The Identity ID can not be changed and can not be chosen. This ensures future compatibility and optimization for distributed stores such as CockroachDB.")
    metadata_admin: Optional[Any] = Field(default=None, description="NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-")
    metadata_public: Optional[Any] = Field(default=None, description="NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-")
    organization_id: Optional[StrictStr] = None
    recovery_addresses: Optional[List[RecoveryIdentityAddress]] = Field(default=None, description="RecoveryAddresses contains all the addresses that can be used to recover an identity.")
    schema_id: StrictStr = Field(description="SchemaID is the ID of the JSON Schema to be used for validating the identity's traits.")
    schema_url: StrictStr = Field(description="SchemaURL is the URL of the endpoint where the identity's traits schema can be fetched from.  format: url")
    state: Optional[StrictStr] = Field(default=None, description="State is the identity's state.  This value has currently no effect. active StateActive inactive StateInactive")
    state_changed_at: Optional[datetime] = None
    traits: Optional[Any] = Field(description="Traits represent an identity's traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in `schema_url`.")
    updated_at: Optional[datetime] = Field(default=None, description="UpdatedAt is a helper struct field for gobuffalo.pop.")
    verifiable_addresses: Optional[List[VerifiableIdentityAddress]] = Field(default=None, description="VerifiableAddresses contains all the addresses that can be verified by the user.")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["created_at", "credentials", "id", "metadata_admin", "metadata_public", "organization_id", "recovery_addresses", "schema_id", "schema_url", "state", "state_changed_at", "traits", "updated_at", "verifiable_addresses"]

    @field_validator('state')
    def state_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['active', 'inactive']):
            raise ValueError("must be one of enum values ('active', 'inactive')")
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
        """Create an instance of Identity from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each value in credentials (dict)
        _field_dict = {}
        if self.credentials:
            for _key_credentials in self.credentials:
                if self.credentials[_key_credentials]:
                    _field_dict[_key_credentials] = self.credentials[_key_credentials].to_dict()
            _dict['credentials'] = _field_dict
        # override the default output from pydantic by calling `to_dict()` of each item in recovery_addresses (list)
        _items = []
        if self.recovery_addresses:
            for _item_recovery_addresses in self.recovery_addresses:
                if _item_recovery_addresses:
                    _items.append(_item_recovery_addresses.to_dict())
            _dict['recovery_addresses'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in verifiable_addresses (list)
        _items = []
        if self.verifiable_addresses:
            for _item_verifiable_addresses in self.verifiable_addresses:
                if _item_verifiable_addresses:
                    _items.append(_item_verifiable_addresses.to_dict())
            _dict['verifiable_addresses'] = _items
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        # set to None if metadata_admin (nullable) is None
        # and model_fields_set contains the field
        if self.metadata_admin is None and "metadata_admin" in self.model_fields_set:
            _dict['metadata_admin'] = None

        # set to None if metadata_public (nullable) is None
        # and model_fields_set contains the field
        if self.metadata_public is None and "metadata_public" in self.model_fields_set:
            _dict['metadata_public'] = None

        # set to None if organization_id (nullable) is None
        # and model_fields_set contains the field
        if self.organization_id is None and "organization_id" in self.model_fields_set:
            _dict['organization_id'] = None

        # set to None if traits (nullable) is None
        # and model_fields_set contains the field
        if self.traits is None and "traits" in self.model_fields_set:
            _dict['traits'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Identity from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "created_at": obj.get("created_at"),
            "credentials": dict(
                (_k, IdentityCredentials.from_dict(_v))
                for _k, _v in obj["credentials"].items()
            )
            if obj.get("credentials") is not None
            else None,
            "id": obj.get("id"),
            "metadata_admin": obj.get("metadata_admin"),
            "metadata_public": obj.get("metadata_public"),
            "organization_id": obj.get("organization_id"),
            "recovery_addresses": [RecoveryIdentityAddress.from_dict(_item) for _item in obj["recovery_addresses"]] if obj.get("recovery_addresses") is not None else None,
            "schema_id": obj.get("schema_id"),
            "schema_url": obj.get("schema_url"),
            "state": obj.get("state"),
            "state_changed_at": obj.get("state_changed_at"),
            "traits": obj.get("traits"),
            "updated_at": obj.get("updated_at"),
            "verifiable_addresses": [VerifiableIdentityAddress.from_dict(_item) for _item in obj["verifiable_addresses"]] if obj.get("verifiable_addresses") is not None else None
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


