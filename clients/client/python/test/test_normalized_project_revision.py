# coding: utf-8

"""
    Ory APIs

    # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

    The version of the OpenAPI document: v1.17.1
    Contact: support@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from ory_client.models.normalized_project_revision import NormalizedProjectRevision

class TestNormalizedProjectRevision(unittest.TestCase):
    """NormalizedProjectRevision unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> NormalizedProjectRevision:
        """Test NormalizedProjectRevision
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `NormalizedProjectRevision`
        """
        model = NormalizedProjectRevision()
        if include_optional:
            return NormalizedProjectRevision(
                account_experience_favicon_dark = '',
                account_experience_favicon_light = '',
                account_experience_logo_dark = '',
                account_experience_logo_light = '',
                account_experience_theme_variables_dark = '',
                account_experience_theme_variables_light = '',
                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                disable_account_experience_welcome_screen = True,
                enable_ax_v2 = True,
                hydra_oauth2_allowed_top_level_claims = [
                    ''
                    ],
                hydra_oauth2_client_credentials_default_grant_allowed_scope = True,
                hydra_oauth2_exclude_not_before_claim = True,
                hydra_oauth2_grant_jwt_iat_optional = True,
                hydra_oauth2_grant_jwt_jti_optional = True,
                hydra_oauth2_grant_jwt_max_ttl = '720h',
                hydra_oauth2_grant_refresh_token_rotation_grace_period = '',
                hydra_oauth2_mirror_top_level_claims = True,
                hydra_oauth2_pkce_enforced = True,
                hydra_oauth2_pkce_enforced_for_public_clients = True,
                hydra_oauth2_refresh_token_hook = '',
                hydra_oauth2_token_hook = '',
                hydra_oidc_dynamic_client_registration_default_scope = [
                    ''
                    ],
                hydra_oidc_dynamic_client_registration_enabled = True,
                hydra_oidc_subject_identifiers_pairwise_salt = '',
                hydra_oidc_subject_identifiers_supported_types = [
                    ''
                    ],
                hydra_secrets_cookie = [
                    ''
                    ],
                hydra_secrets_system = [
                    ''
                    ],
                hydra_serve_cookies_same_site_legacy_workaround = True,
                hydra_serve_cookies_same_site_mode = '',
                hydra_strategies_access_token = 'opaque',
                hydra_strategies_jwt_scope_claim = 'list',
                hydra_strategies_scope = 'wildcard',
                hydra_ttl_access_token = '30m',
                hydra_ttl_auth_code = '720h',
                hydra_ttl_id_token = '30m',
                hydra_ttl_login_consent_request = '30m',
                hydra_ttl_refresh_token = '720h',
                hydra_urls_consent = '',
                hydra_urls_error = '',
                hydra_urls_login = '',
                hydra_urls_logout = '',
                hydra_urls_post_logout_redirect = '',
                hydra_urls_registration = '',
                hydra_urls_self_issuer = '',
                hydra_webfinger_jwks_broadcast_keys = [
                    ''
                    ],
                hydra_webfinger_oidc_discovery_auth_url = '',
                hydra_webfinger_oidc_discovery_client_registration_url = '',
                hydra_webfinger_oidc_discovery_jwks_url = '',
                hydra_webfinger_oidc_discovery_supported_claims = [
                    ''
                    ],
                hydra_webfinger_oidc_discovery_supported_scope = [
                    ''
                    ],
                hydra_webfinger_oidc_discovery_token_url = '',
                hydra_webfinger_oidc_discovery_userinfo_url = '',
                id = '',
                keto_namespace_configuration = '',
                keto_namespaces = [
                    ory_client.models.keto_namespace.KetoNamespace(
                        id = 56, 
                        name = '', )
                    ],
                kratos_cookies_same_site = '',
                kratos_courier_channels = [
                    ory_client.models.normalized_project_revision_courier_channel.NormalizedProjectRevisionCourierChannel(
                        channel_id = '', 
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        request_config_auth_config_api_key_in = 'header', 
                        request_config_auth_config_api_key_name = '', 
                        request_config_auth_config_api_key_value = '', 
                        request_config_auth_config_basic_auth_password = '', 
                        request_config_auth_config_basic_auth_user = '', 
                        request_config_auth_type = 'basic_auth', 
                        request_config_body = '', 
                        request_config_headers = ory_client.models.null_json_raw_message.nullJsonRawMessage(), 
                        request_config_method = 'POST', 
                        request_config_url = '', 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                    ],
                kratos_courier_delivery_strategy = 'smtp',
                kratos_courier_http_request_config_auth_api_key_in = '',
                kratos_courier_http_request_config_auth_api_key_name = '',
                kratos_courier_http_request_config_auth_api_key_value = '',
                kratos_courier_http_request_config_auth_basic_auth_password = '',
                kratos_courier_http_request_config_auth_basic_auth_user = '',
                kratos_courier_http_request_config_auth_type = 'empty (no authentication)',
                kratos_courier_http_request_config_body = '',
                kratos_courier_http_request_config_headers = ory_client.models.null_json_raw_message.nullJsonRawMessage(),
                kratos_courier_http_request_config_method = 'POST',
                kratos_courier_http_request_config_url = '',
                kratos_courier_smtp_connection_uri = '',
                kratos_courier_smtp_from_address = '',
                kratos_courier_smtp_from_name = '',
                kratos_courier_smtp_headers = ory_client.models.null_json_raw_message.nullJsonRawMessage(),
                kratos_courier_smtp_local_name = '',
                kratos_courier_templates_login_code_valid_email_body_html = '',
                kratos_courier_templates_login_code_valid_email_body_plaintext = '',
                kratos_courier_templates_login_code_valid_email_subject = '',
                kratos_courier_templates_login_code_valid_sms_body_plaintext = '',
                kratos_courier_templates_recovery_code_invalid_email_body_html = '',
                kratos_courier_templates_recovery_code_invalid_email_body_plaintext = '',
                kratos_courier_templates_recovery_code_invalid_email_subject = '',
                kratos_courier_templates_recovery_code_valid_email_body_html = '',
                kratos_courier_templates_recovery_code_valid_email_body_plaintext = '',
                kratos_courier_templates_recovery_code_valid_email_subject = '',
                kratos_courier_templates_recovery_invalid_email_body_html = '',
                kratos_courier_templates_recovery_invalid_email_body_plaintext = '',
                kratos_courier_templates_recovery_invalid_email_subject = '',
                kratos_courier_templates_recovery_valid_email_body_html = '',
                kratos_courier_templates_recovery_valid_email_body_plaintext = '',
                kratos_courier_templates_recovery_valid_email_subject = '',
                kratos_courier_templates_registration_code_valid_email_body_html = '',
                kratos_courier_templates_registration_code_valid_email_body_plaintext = '',
                kratos_courier_templates_registration_code_valid_email_subject = '',
                kratos_courier_templates_registration_code_valid_sms_body_plaintext = '',
                kratos_courier_templates_verification_code_invalid_email_body_html = '',
                kratos_courier_templates_verification_code_invalid_email_body_plaintext = '',
                kratos_courier_templates_verification_code_invalid_email_subject = '',
                kratos_courier_templates_verification_code_valid_email_body_html = '',
                kratos_courier_templates_verification_code_valid_email_body_plaintext = '',
                kratos_courier_templates_verification_code_valid_email_subject = '',
                kratos_courier_templates_verification_code_valid_sms_body_plaintext = '',
                kratos_courier_templates_verification_invalid_email_body_html = '',
                kratos_courier_templates_verification_invalid_email_body_plaintext = '',
                kratos_courier_templates_verification_invalid_email_subject = '',
                kratos_courier_templates_verification_valid_email_body_html = '',
                kratos_courier_templates_verification_valid_email_body_plaintext = '',
                kratos_courier_templates_verification_valid_email_subject = '',
                kratos_feature_flags_cacheable_sessions = True,
                kratos_feature_flags_cacheable_sessions_max_age = '',
                kratos_feature_flags_faster_session_extend = True,
                kratos_feature_flags_use_continue_with_transitions = True,
                kratos_identity_schemas = [
                    ory_client.models.normalized_project_revision_identity_schema.normalizedProjectRevisionIdentitySchema(
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        id = '', 
                        identity_schema = ory_client.models.schema_represents_an_ory_kratos_identity_schema.Schema represents an Ory Kratos Identity Schema(
                            blob_name = '', 
                            blob_url = '', 
                            content_hash = '', 
                            created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            id = '', 
                            name = 'CustomerIdentity', 
                            updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ), 
                        identity_schema_id = '', 
                        import_id = '', 
                        import_url = 'base64://ey...', 
                        is_default = True, 
                        preset = '', 
                        project_revision_id = '', 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                    ],
                kratos_oauth2_provider_headers = ory_client.models.null_json_raw_message.nullJsonRawMessage(),
                kratos_oauth2_provider_override_return_to = True,
                kratos_oauth2_provider_url = '',
                kratos_preview_default_read_consistency_level = '',
                kratos_secrets_cipher = [
                    ''
                    ],
                kratos_secrets_cookie = [
                    ''
                    ],
                kratos_secrets_default = [
                    ''
                    ],
                kratos_selfservice_allowed_return_urls = [
                    ''
                    ],
                kratos_selfservice_default_browser_return_url = '',
                kratos_selfservice_flows_error_ui_url = '',
                kratos_selfservice_flows_login_after_code_default_browser_return_url = '',
                kratos_selfservice_flows_login_after_default_browser_return_url = '',
                kratos_selfservice_flows_login_after_lookup_secret_default_browser_return_url = '',
                kratos_selfservice_flows_login_after_oidc_default_browser_return_url = '',
                kratos_selfservice_flows_login_after_passkey_default_browser_return_url = '',
                kratos_selfservice_flows_login_after_password_default_browser_return_url = '',
                kratos_selfservice_flows_login_after_totp_default_browser_return_url = '',
                kratos_selfservice_flows_login_after_webauthn_default_browser_return_url = '',
                kratos_selfservice_flows_login_lifespan = '',
                kratos_selfservice_flows_login_ui_url = '',
                kratos_selfservice_flows_logout_after_default_browser_return_url = '',
                kratos_selfservice_flows_recovery_after_default_browser_return_url = '',
                kratos_selfservice_flows_recovery_enabled = True,
                kratos_selfservice_flows_recovery_lifespan = '',
                kratos_selfservice_flows_recovery_notify_unknown_recipients = True,
                kratos_selfservice_flows_recovery_ui_url = '',
                kratos_selfservice_flows_recovery_use = 'link',
                kratos_selfservice_flows_registration_after_code_default_browser_return_url = '',
                kratos_selfservice_flows_registration_after_default_browser_return_url = '',
                kratos_selfservice_flows_registration_after_oidc_default_browser_return_url = '',
                kratos_selfservice_flows_registration_after_passkey_default_browser_return_url = '',
                kratos_selfservice_flows_registration_after_password_default_browser_return_url = '',
                kratos_selfservice_flows_registration_after_webauthn_default_browser_return_url = '',
                kratos_selfservice_flows_registration_enable_legacy_one_step = True,
                kratos_selfservice_flows_registration_enabled = True,
                kratos_selfservice_flows_registration_lifespan = '',
                kratos_selfservice_flows_registration_login_hints = True,
                kratos_selfservice_flows_registration_ui_url = '',
                kratos_selfservice_flows_settings_after_default_browser_return_url = '',
                kratos_selfservice_flows_settings_after_lookup_secret_default_browser_return_url = '',
                kratos_selfservice_flows_settings_after_oidc_default_browser_return_url = '',
                kratos_selfservice_flows_settings_after_passkey_default_browser_return_url = '',
                kratos_selfservice_flows_settings_after_password_default_browser_return_url = '',
                kratos_selfservice_flows_settings_after_profile_default_browser_return_url = '',
                kratos_selfservice_flows_settings_after_totp_default_browser_return_url = '',
                kratos_selfservice_flows_settings_after_webauthn_default_browser_return_url = '',
                kratos_selfservice_flows_settings_lifespan = '',
                kratos_selfservice_flows_settings_privileged_session_max_age = '',
                kratos_selfservice_flows_settings_required_aal = '',
                kratos_selfservice_flows_settings_ui_url = '',
                kratos_selfservice_flows_verification_after_default_browser_return_url = '',
                kratos_selfservice_flows_verification_enabled = True,
                kratos_selfservice_flows_verification_lifespan = '',
                kratos_selfservice_flows_verification_notify_unknown_recipients = True,
                kratos_selfservice_flows_verification_ui_url = '',
                kratos_selfservice_flows_verification_use = 'link',
                kratos_selfservice_methods_captcha_config_cf_turnstile_sitekey = '',
                kratos_selfservice_methods_captcha_enabled = True,
                kratos_selfservice_methods_code_config_lifespan = '',
                kratos_selfservice_methods_code_config_missing_credential_fallback_enabled = True,
                kratos_selfservice_methods_code_enabled = True,
                kratos_selfservice_methods_code_mfa_enabled = True,
                kratos_selfservice_methods_code_passwordless_enabled = True,
                kratos_selfservice_methods_code_passwordless_login_fallback_enabled = True,
                kratos_selfservice_methods_link_config_base_url = '',
                kratos_selfservice_methods_link_config_lifespan = '',
                kratos_selfservice_methods_link_enabled = True,
                kratos_selfservice_methods_lookup_secret_enabled = True,
                kratos_selfservice_methods_oidc_config_base_redirect_uri = '',
                kratos_selfservice_methods_oidc_config_providers = [
                    ory_client.models.normalized_project_revision_third_party_provider.normalizedProjectRevisionThirdPartyProvider(
                        additional_id_token_audiences = [
                            ''
                            ], 
                        apple_private_key = '', 
                        apple_private_key_id = 'UX56C66723', 
                        apple_team_id = 'KP76DQS54M', 
                        auth_url = 'https://www.googleapis.com/oauth2/v2/auth', 
                        azure_tenant = 'contoso.onmicrosoft.com', 
                        claims_source = '', 
                        client_id = '', 
                        client_secret = , 
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        fedcm_config_url = , 
                        id = '', 
                        issuer_url = 'https://accounts.google.com', 
                        label = '', 
                        mapper_url = '', 
                        net_id_token_origin_header = , 
                        organization_id = '', 
                        pkce = 'auto', 
                        project_revision_id = '', 
                        provider = 'google', 
                        provider_id = '', 
                        requested_claims = ory_client.models.json_raw_message_represents_a_json/raw_message_that_works_well_with_json,_sql,_and_swagger/.JSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger.(), 
                        scope = [
                            ''
                            ], 
                        state = 'enabled', 
                        subject_source = , 
                        token_url = 'https://www.googleapis.com/oauth2/v4/token', 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                    ],
                kratos_selfservice_methods_oidc_enabled = True,
                kratos_selfservice_methods_passkey_config_rp_display_name = '',
                kratos_selfservice_methods_passkey_config_rp_id = '',
                kratos_selfservice_methods_passkey_config_rp_origins = [
                    ''
                    ],
                kratos_selfservice_methods_passkey_enabled = True,
                kratos_selfservice_methods_password_config_haveibeenpwned_enabled = True,
                kratos_selfservice_methods_password_config_identifier_similarity_check_enabled = True,
                kratos_selfservice_methods_password_config_ignore_network_errors = True,
                kratos_selfservice_methods_password_config_max_breaches = 56,
                kratos_selfservice_methods_password_config_min_password_length = 56,
                kratos_selfservice_methods_password_enabled = True,
                kratos_selfservice_methods_profile_enabled = True,
                kratos_selfservice_methods_saml_config_providers = [
                    ory_client.models.normalized_project_revision_saml_provider.normalizedProjectRevisionSAMLProvider(
                        client_id = '', 
                        client_secret = '', 
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        id = '', 
                        label = '', 
                        mapper_url = '', 
                        organization_id = '', 
                        project_revision_id = '', 
                        provider_id = '', 
                        raw_idp_metadata_xml = '', 
                        state = 'enabled', 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                    ],
                kratos_selfservice_methods_saml_enabled = True,
                kratos_selfservice_methods_totp_config_issuer = '',
                kratos_selfservice_methods_totp_enabled = True,
                kratos_selfservice_methods_webauthn_config_passwordless = True,
                kratos_selfservice_methods_webauthn_config_rp_display_name = '',
                kratos_selfservice_methods_webauthn_config_rp_icon = '',
                kratos_selfservice_methods_webauthn_config_rp_id = '',
                kratos_selfservice_methods_webauthn_config_rp_origins = [
                    ''
                    ],
                kratos_selfservice_methods_webauthn_enabled = True,
                kratos_session_cookie_persistent = True,
                kratos_session_cookie_same_site = '',
                kratos_session_lifespan = '',
                kratos_session_whoami_required_aal = '',
                kratos_session_whoami_tokenizer_templates = [
                    ory_client.models.normalized_project_revision_tokenizer_template.normalizedProjectRevisionTokenizerTemplate(
                        claims_mapper_url = '', 
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        id = '', 
                        jwks_url = '', 
                        key = '', 
                        project_revision_id = '', 
                        ttl = '1m', 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                    ],
                name = '',
                project_id = '',
                project_revision_hooks = [
                    ory_client.models.normalized_project_revision_hook.normalizedProjectRevisionHook(
                        config_key = '', 
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        hook = '', 
                        id = '', 
                        project_revision_id = '', 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        web_hook_config_auth_api_key_in = 'header', 
                        web_hook_config_auth_api_key_name = 'X-API-Key', 
                        web_hook_config_auth_api_key_value = 'eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ', 
                        web_hook_config_auth_basic_auth_password = '', 
                        web_hook_config_auth_basic_auth_user = '', 
                        web_hook_config_auth_type = '', 
                        web_hook_config_body = 'base64://ZnVuY3Rpb24oY3R4KSB7CiAgaWRlbnRpdHlfaWQ6IGlmIGN0eFsiaWRlbnRpdHkiXSAhPSBudWxsIHRoZW4gY3R4LmlkZW50aXR5LmlkLAp9=', 
                        web_hook_config_can_interrupt = True, 
                        web_hook_config_method = 'POST', 
                        web_hook_config_response_ignore = True, 
                        web_hook_config_response_parse = True, 
                        web_hook_config_url = 'https://www.example.org/web-hook-listener', )
                    ],
                serve_admin_cors_allowed_origins = [
                    ''
                    ],
                serve_admin_cors_enabled = True,
                serve_public_cors_allowed_origins = [
                    ''
                    ],
                serve_public_cors_enabled = True,
                strict_security = True,
                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else:
            return NormalizedProjectRevision(
                name = '',
        )
        """

    def testNormalizedProjectRevision(self):
        """Test NormalizedProjectRevision"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
