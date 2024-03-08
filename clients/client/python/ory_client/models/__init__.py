# flake8: noqa

# import all models into this package
# if you have many models here with many references from one model to another this may
# raise a RecursionError
# to avoid this, import only the models that you directly need like:
# from from ory_client.model.pet import Pet
# or import this package, but before doing it, use:
# import sys
# sys.setrecursionlimit(n)

from ory_client.model.accept_o_auth2_consent_request import AcceptOAuth2ConsentRequest
from ory_client.model.accept_o_auth2_consent_request_session import AcceptOAuth2ConsentRequestSession
from ory_client.model.accept_o_auth2_login_request import AcceptOAuth2LoginRequest
from ory_client.model.active_project_in_console import ActiveProjectInConsole
from ory_client.model.attribute import Attribute
from ory_client.model.attribute_filter import AttributeFilter
from ory_client.model.attributes_count_datapoint import AttributesCountDatapoint
from ory_client.model.authenticator_assurance_level import AuthenticatorAssuranceLevel
from ory_client.model.batch_patch_identities_response import BatchPatchIdentitiesResponse
from ory_client.model.check_opl_syntax_result import CheckOplSyntaxResult
from ory_client.model.check_permission_result import CheckPermissionResult
from ory_client.model.cloud_account import CloudAccount
from ory_client.model.consistency_request_parameters import ConsistencyRequestParameters
from ory_client.model.continue_with import ContinueWith
from ory_client.model.continue_with_recovery_ui import ContinueWithRecoveryUi
from ory_client.model.continue_with_recovery_ui_flow import ContinueWithRecoveryUiFlow
from ory_client.model.continue_with_set_ory_session_token import ContinueWithSetOrySessionToken
from ory_client.model.continue_with_settings_ui import ContinueWithSettingsUi
from ory_client.model.continue_with_settings_ui_flow import ContinueWithSettingsUiFlow
from ory_client.model.continue_with_verification_ui import ContinueWithVerificationUi
from ory_client.model.continue_with_verification_ui_flow import ContinueWithVerificationUiFlow
from ory_client.model.courier_message_status import CourierMessageStatus
from ory_client.model.courier_message_type import CourierMessageType
from ory_client.model.create_custom_domain_body import CreateCustomDomainBody
from ory_client.model.create_event_stream_body import CreateEventStreamBody
from ory_client.model.create_identity_body import CreateIdentityBody
from ory_client.model.create_invite_response import CreateInviteResponse
from ory_client.model.create_json_web_key_set import CreateJsonWebKeySet
from ory_client.model.create_member_invite_response import CreateMemberInviteResponse
from ory_client.model.create_project_api_key_request import CreateProjectApiKeyRequest
from ory_client.model.create_project_body import CreateProjectBody
from ory_client.model.create_project_branding import CreateProjectBranding
from ory_client.model.create_project_member_invite_body import CreateProjectMemberInviteBody
from ory_client.model.create_project_normalized_payload import CreateProjectNormalizedPayload
from ory_client.model.create_recovery_code_for_identity_body import CreateRecoveryCodeForIdentityBody
from ory_client.model.create_recovery_link_for_identity_body import CreateRecoveryLinkForIdentityBody
from ory_client.model.create_relationship_body import CreateRelationshipBody
from ory_client.model.create_subscription_body import CreateSubscriptionBody
from ory_client.model.create_subscription_common import CreateSubscriptionCommon
from ory_client.model.create_verifiable_credential_request_body import CreateVerifiableCredentialRequestBody
from ory_client.model.create_workspace_member_invite_body import CreateWorkspaceMemberInviteBody
from ory_client.model.create_workspace_payload import CreateWorkspacePayload
from ory_client.model.create_workspace_subscription_body import CreateWorkspaceSubscriptionBody
from ory_client.model.credential_supported_draft00 import CredentialSupportedDraft00
from ory_client.model.custom_domain import CustomDomain
from ory_client.model.delete_my_sessions_count import DeleteMySessionsCount
from ory_client.model.email_template_data import EmailTemplateData
from ory_client.model.email_template_data_body import EmailTemplateDataBody
from ory_client.model.error_authenticator_assurance_level_not_satisfied import ErrorAuthenticatorAssuranceLevelNotSatisfied
from ory_client.model.error_browser_location_change_required import ErrorBrowserLocationChangeRequired
from ory_client.model.error_flow_replaced import ErrorFlowReplaced
from ory_client.model.error_generic import ErrorGeneric
from ory_client.model.error_o_auth2 import ErrorOAuth2
from ory_client.model.event_stream import EventStream
from ory_client.model.expanded_permission_tree import ExpandedPermissionTree
from ory_client.model.flow_error import FlowError
from ory_client.model.generic_error import GenericError
from ory_client.model.generic_error_content import GenericErrorContent
from ory_client.model.generic_usage import GenericUsage
from ory_client.model.get_attributes_count_response import GetAttributesCountResponse
from ory_client.model.get_managed_identity_schema_location import GetManagedIdentitySchemaLocation
from ory_client.model.get_metrics_event_attributes_response import GetMetricsEventAttributesResponse
from ory_client.model.get_metrics_event_types_response import GetMetricsEventTypesResponse
from ory_client.model.get_organization_response import GetOrganizationResponse
from ory_client.model.get_project_events_body import GetProjectEventsBody
from ory_client.model.get_project_events_response import GetProjectEventsResponse
from ory_client.model.get_project_metrics_response import GetProjectMetricsResponse
from ory_client.model.get_session_activity_response import GetSessionActivityResponse
from ory_client.model.get_version200_response import GetVersion200Response
from ory_client.model.health_not_ready_status import HealthNotReadyStatus
from ory_client.model.health_status import HealthStatus
from ory_client.model.identity import Identity
from ory_client.model.identity_credentials import IdentityCredentials
from ory_client.model.identity_credentials_code import IdentityCredentialsCode
from ory_client.model.identity_credentials_oidc import IdentityCredentialsOidc
from ory_client.model.identity_credentials_oidc_provider import IdentityCredentialsOidcProvider
from ory_client.model.identity_credentials_password import IdentityCredentialsPassword
from ory_client.model.identity_patch import IdentityPatch
from ory_client.model.identity_patch_response import IdentityPatchResponse
from ory_client.model.identity_schema_container import IdentitySchemaContainer
from ory_client.model.identity_schema_preset import IdentitySchemaPreset
from ory_client.model.identity_schema_presets import IdentitySchemaPresets
from ory_client.model.identity_schemas import IdentitySchemas
from ory_client.model.identity_with_credentials import IdentityWithCredentials
from ory_client.model.identity_with_credentials_oidc import IdentityWithCredentialsOidc
from ory_client.model.identity_with_credentials_oidc_config import IdentityWithCredentialsOidcConfig
from ory_client.model.identity_with_credentials_oidc_config_provider import IdentityWithCredentialsOidcConfigProvider
from ory_client.model.identity_with_credentials_password import IdentityWithCredentialsPassword
from ory_client.model.identity_with_credentials_password_config import IdentityWithCredentialsPasswordConfig
from ory_client.model.internal_get_project_branding_body import InternalGetProjectBrandingBody
from ory_client.model.internal_is_ax_welcome_screen_enabled_for_project_body import InternalIsAXWelcomeScreenEnabledForProjectBody
from ory_client.model.internal_is_owner_for_project_by_slug_body import InternalIsOwnerForProjectBySlugBody
from ory_client.model.internal_is_owner_for_project_by_slug_response import InternalIsOwnerForProjectBySlugResponse
from ory_client.model.introspected_o_auth2_token import IntrospectedOAuth2Token
from ory_client.model.is_owner_for_project_by_slug import IsOwnerForProjectBySlug
from ory_client.model.is_ready200_response import IsReady200Response
from ory_client.model.is_ready503_response import IsReady503Response
from ory_client.model.json_patch import JsonPatch
from ory_client.model.json_patch_document import JsonPatchDocument
from ory_client.model.json_web_key import JsonWebKey
from ory_client.model.json_web_key_set import JsonWebKeySet
from ory_client.model.keto_namespace import KetoNamespace
from ory_client.model.keto_namespaces import KetoNamespaces
from ory_client.model.list_custom_domains import ListCustomDomains
from ory_client.model.list_event_streams import ListEventStreams
from ory_client.model.list_my_workspaces_response import ListMyWorkspacesResponse
from ory_client.model.list_organizations_response import ListOrganizationsResponse
from ory_client.model.list_workspace_projects_response import ListWorkspaceProjectsResponse
from ory_client.model.login_flow import LoginFlow
from ory_client.model.login_flow_state import LoginFlowState
from ory_client.model.logout_flow import LogoutFlow
from ory_client.model.managed_identity_schema import ManagedIdentitySchema
from ory_client.model.managed_identity_schema_validation_result import ManagedIdentitySchemaValidationResult
from ory_client.model.managed_identity_schemas import ManagedIdentitySchemas
from ory_client.model.member_invite import MemberInvite
from ory_client.model.member_invites import MemberInvites
from ory_client.model.message import Message
from ory_client.model.message_dispatch import MessageDispatch
from ory_client.model.metrics_datapoint import MetricsDatapoint
from ory_client.model.migration_options import MigrationOptions
from ory_client.model.namespace import Namespace
from ory_client.model.needs_privileged_session_error import NeedsPrivilegedSessionError
from ory_client.model.normalized_project import NormalizedProject
from ory_client.model.normalized_project_revision import NormalizedProjectRevision
from ory_client.model.normalized_project_revision_courier_channel import NormalizedProjectRevisionCourierChannel
from ory_client.model.normalized_project_revision_hook import NormalizedProjectRevisionHook
from ory_client.model.normalized_project_revision_identity_schema import NormalizedProjectRevisionIdentitySchema
from ory_client.model.normalized_project_revision_identity_schemas import NormalizedProjectRevisionIdentitySchemas
from ory_client.model.normalized_project_revision_third_party_provider import NormalizedProjectRevisionThirdPartyProvider
from ory_client.model.normalized_project_revision_tokenizer_template import NormalizedProjectRevisionTokenizerTemplate
from ory_client.model.normalized_project_revision_tokenizer_templates import NormalizedProjectRevisionTokenizerTemplates
from ory_client.model.normalized_projects import NormalizedProjects
from ory_client.model.null_duration import NullDuration
from ory_client.model.o_auth2_client import OAuth2Client
from ory_client.model.o_auth2_client_token_lifespans import OAuth2ClientTokenLifespans
from ory_client.model.o_auth2_consent_request import OAuth2ConsentRequest
from ory_client.model.o_auth2_consent_request_open_id_connect_context import OAuth2ConsentRequestOpenIDConnectContext
from ory_client.model.o_auth2_consent_session import OAuth2ConsentSession
from ory_client.model.o_auth2_consent_session_expires_at import OAuth2ConsentSessionExpiresAt
from ory_client.model.o_auth2_consent_sessions import OAuth2ConsentSessions
from ory_client.model.o_auth2_login_request import OAuth2LoginRequest
from ory_client.model.o_auth2_logout_request import OAuth2LogoutRequest
from ory_client.model.o_auth2_redirect_to import OAuth2RedirectTo
from ory_client.model.o_auth2_token_exchange import OAuth2TokenExchange
from ory_client.model.oidc_configuration import OidcConfiguration
from ory_client.model.oidc_user_info import OidcUserInfo
from ory_client.model.organization import Organization
from ory_client.model.organization_body import OrganizationBody
from ory_client.model.pagination import Pagination
from ory_client.model.pagination_headers import PaginationHeaders
from ory_client.model.parse_error import ParseError
from ory_client.model.patch_identities_body import PatchIdentitiesBody
from ory_client.model.perform_native_logout_body import PerformNativeLogoutBody
from ory_client.model.permissions_on_project import PermissionsOnProject
from ory_client.model.permissions_on_workpace_response import PermissionsOnWorkpaceResponse
from ory_client.model.plan import Plan
from ory_client.model.plan_details import PlanDetails
from ory_client.model.plan_features import PlanFeatures
from ory_client.model.plans import Plans
from ory_client.model.post_check_permission_body import PostCheckPermissionBody
from ory_client.model.post_check_permission_or_error_body import PostCheckPermissionOrErrorBody
from ory_client.model.pricing import Pricing
from ory_client.model.project import Project
from ory_client.model.project_api_key import ProjectApiKey
from ory_client.model.project_api_keys import ProjectApiKeys
from ory_client.model.project_branding import ProjectBranding
from ory_client.model.project_branding_colors import ProjectBrandingColors
from ory_client.model.project_branding_theme import ProjectBrandingTheme
from ory_client.model.project_branding_themes import ProjectBrandingThemes
from ory_client.model.project_cors import ProjectCors
from ory_client.model.project_events_datapoint import ProjectEventsDatapoint
from ory_client.model.project_host import ProjectHost
from ory_client.model.project_member import ProjectMember
from ory_client.model.project_members import ProjectMembers
from ory_client.model.project_metadata import ProjectMetadata
from ory_client.model.project_metadata_list import ProjectMetadataList
from ory_client.model.project_revision_hooks import ProjectRevisionHooks
from ory_client.model.project_revision_identity_schemas import ProjectRevisionIdentitySchemas
from ory_client.model.project_revision_third_party_login_providers import ProjectRevisionThirdPartyLoginProviders
from ory_client.model.project_revisions import ProjectRevisions
from ory_client.model.project_service_identity import ProjectServiceIdentity
from ory_client.model.project_service_o_auth2 import ProjectServiceOAuth2
from ory_client.model.project_service_permission import ProjectServicePermission
from ory_client.model.project_services import ProjectServices
from ory_client.model.projects import Projects
from ory_client.model.quota_usage import QuotaUsage
from ory_client.model.rfc6749_error_json import RFC6749ErrorJson
from ory_client.model.recovery_code_for_identity import RecoveryCodeForIdentity
from ory_client.model.recovery_flow import RecoveryFlow
from ory_client.model.recovery_flow_state import RecoveryFlowState
from ory_client.model.recovery_identity_address import RecoveryIdentityAddress
from ory_client.model.recovery_link_for_identity import RecoveryLinkForIdentity
from ory_client.model.registration_flow import RegistrationFlow
from ory_client.model.registration_flow_state import RegistrationFlowState
from ory_client.model.reject_o_auth2_request import RejectOAuth2Request
from ory_client.model.relation_query import RelationQuery
from ory_client.model.relationship import Relationship
from ory_client.model.relationship_namespaces import RelationshipNamespaces
from ory_client.model.relationship_patch import RelationshipPatch
from ory_client.model.relationships import Relationships
from ory_client.model.revision_courier_channels import RevisionCourierChannels
from ory_client.model.schema_patch import SchemaPatch
from ory_client.model.self_service_flow_expired_error import SelfServiceFlowExpiredError
from ory_client.model.session import Session
from ory_client.model.session_activity_datapoint import SessionActivityDatapoint
from ory_client.model.session_authentication_method import SessionAuthenticationMethod
from ory_client.model.session_authentication_methods import SessionAuthenticationMethods
from ory_client.model.session_device import SessionDevice
from ory_client.model.set_active_project_in_console_body import SetActiveProjectInConsoleBody
from ory_client.model.set_custom_domain_body import SetCustomDomainBody
from ory_client.model.set_event_stream_body import SetEventStreamBody
from ory_client.model.set_project import SetProject
from ory_client.model.set_project_branding_theme_body import SetProjectBrandingThemeBody
from ory_client.model.settings_flow import SettingsFlow
from ory_client.model.settings_flow_state import SettingsFlowState
from ory_client.model.source_position import SourcePosition
from ory_client.model.string_slice_json_format import StringSliceJSONFormat
from ory_client.model.subject_set import SubjectSet
from ory_client.model.subscription import Subscription
from ory_client.model.successful_code_exchange_response import SuccessfulCodeExchangeResponse
from ory_client.model.successful_native_login import SuccessfulNativeLogin
from ory_client.model.successful_native_registration import SuccessfulNativeRegistration
from ory_client.model.successful_project_update import SuccessfulProjectUpdate
from ory_client.model.token_pagination import TokenPagination
from ory_client.model.token_pagination_headers import TokenPaginationHeaders
from ory_client.model.token_pagination_request_parameters import TokenPaginationRequestParameters
from ory_client.model.token_pagination_response_headers import TokenPaginationResponseHeaders
from ory_client.model.trust_o_auth2_jwt_grant_issuer import TrustOAuth2JwtGrantIssuer
from ory_client.model.trusted_o_auth2_jwt_grant_issuer import TrustedOAuth2JwtGrantIssuer
from ory_client.model.trusted_o_auth2_jwt_grant_issuers import TrustedOAuth2JwtGrantIssuers
from ory_client.model.trusted_o_auth2_jwt_grant_json_web_key import TrustedOAuth2JwtGrantJsonWebKey
from ory_client.model.ui_container import UiContainer
from ory_client.model.ui_node import UiNode
from ory_client.model.ui_node_anchor_attributes import UiNodeAnchorAttributes
from ory_client.model.ui_node_attributes import UiNodeAttributes
from ory_client.model.ui_node_image_attributes import UiNodeImageAttributes
from ory_client.model.ui_node_input_attributes import UiNodeInputAttributes
from ory_client.model.ui_node_meta import UiNodeMeta
from ory_client.model.ui_node_script_attributes import UiNodeScriptAttributes
from ory_client.model.ui_node_text_attributes import UiNodeTextAttributes
from ory_client.model.ui_nodes import UiNodes
from ory_client.model.ui_text import UiText
from ory_client.model.ui_texts import UiTexts
from ory_client.model.update_identity_body import UpdateIdentityBody
from ory_client.model.update_login_flow_body import UpdateLoginFlowBody
from ory_client.model.update_login_flow_with_code_method import UpdateLoginFlowWithCodeMethod
from ory_client.model.update_login_flow_with_lookup_secret_method import UpdateLoginFlowWithLookupSecretMethod
from ory_client.model.update_login_flow_with_oidc_method import UpdateLoginFlowWithOidcMethod
from ory_client.model.update_login_flow_with_password_method import UpdateLoginFlowWithPasswordMethod
from ory_client.model.update_login_flow_with_totp_method import UpdateLoginFlowWithTotpMethod
from ory_client.model.update_login_flow_with_web_authn_method import UpdateLoginFlowWithWebAuthnMethod
from ory_client.model.update_recovery_flow_body import UpdateRecoveryFlowBody
from ory_client.model.update_recovery_flow_with_code_method import UpdateRecoveryFlowWithCodeMethod
from ory_client.model.update_recovery_flow_with_link_method import UpdateRecoveryFlowWithLinkMethod
from ory_client.model.update_registration_flow_body import UpdateRegistrationFlowBody
from ory_client.model.update_registration_flow_with_code_method import UpdateRegistrationFlowWithCodeMethod
from ory_client.model.update_registration_flow_with_oidc_method import UpdateRegistrationFlowWithOidcMethod
from ory_client.model.update_registration_flow_with_password_method import UpdateRegistrationFlowWithPasswordMethod
from ory_client.model.update_registration_flow_with_web_authn_method import UpdateRegistrationFlowWithWebAuthnMethod
from ory_client.model.update_settings_flow_body import UpdateSettingsFlowBody
from ory_client.model.update_settings_flow_with_lookup_method import UpdateSettingsFlowWithLookupMethod
from ory_client.model.update_settings_flow_with_oidc_method import UpdateSettingsFlowWithOidcMethod
from ory_client.model.update_settings_flow_with_password_method import UpdateSettingsFlowWithPasswordMethod
from ory_client.model.update_settings_flow_with_profile_method import UpdateSettingsFlowWithProfileMethod
from ory_client.model.update_settings_flow_with_totp_method import UpdateSettingsFlowWithTotpMethod
from ory_client.model.update_settings_flow_with_web_authn_method import UpdateSettingsFlowWithWebAuthnMethod
from ory_client.model.update_subscription_body import UpdateSubscriptionBody
from ory_client.model.update_verification_flow_body import UpdateVerificationFlowBody
from ory_client.model.update_verification_flow_with_code_method import UpdateVerificationFlowWithCodeMethod
from ory_client.model.update_verification_flow_with_link_method import UpdateVerificationFlowWithLinkMethod
from ory_client.model.update_workspace_payload import UpdateWorkspacePayload
from ory_client.model.usage import Usage
from ory_client.model.verifiable_credential_priming_response import VerifiableCredentialPrimingResponse
from ory_client.model.verifiable_credential_proof import VerifiableCredentialProof
from ory_client.model.verifiable_credential_response import VerifiableCredentialResponse
from ory_client.model.verifiable_identity_address import VerifiableIdentityAddress
from ory_client.model.verification_flow import VerificationFlow
from ory_client.model.verification_flow_state import VerificationFlowState
from ory_client.model.version import Version
from ory_client.model.warning import Warning
from ory_client.model.workspace import Workspace
from ory_client.model.workspace_meta import WorkspaceMeta
from ory_client.model.workspaces import Workspaces
