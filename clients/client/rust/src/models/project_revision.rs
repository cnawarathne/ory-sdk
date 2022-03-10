/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.124
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ProjectRevision {
    /// The Project's Revision Creation Date
    #[serde(rename = "created_at", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    /// Configures the Ory Kratos Cookie SameSite Attribute  This governs the \"cookies.same_site\" setting.
    #[serde(rename = "kratos_cookies_same_site", skip_serializing_if = "Option::is_none")]
    pub kratos_cookies_same_site: Option<String>,
    /// Configures the Ory Kratos SMTP Connection URI  This governs the \"courier.smtp.connection_uri\" setting.
    #[serde(rename = "kratos_courier_smtp_connection_uri", skip_serializing_if = "Option::is_none")]
    pub kratos_courier_smtp_connection_uri: Option<String>,
    /// Configures the Ory Kratos SMTP From Address  This governs the \"courier.smtp.from_address\" setting.
    #[serde(rename = "kratos_courier_smtp_from_address", skip_serializing_if = "Option::is_none")]
    pub kratos_courier_smtp_from_address: Option<String>,
    /// Configures the Ory Kratos SMTP From Name  This governs the \"courier.smtp.from_name\" setting.
    #[serde(rename = "kratos_courier_smtp_from_name", skip_serializing_if = "Option::is_none")]
    pub kratos_courier_smtp_from_name: Option<String>,
    /// NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-
    #[serde(rename = "kratos_courier_smtp_headers", skip_serializing_if = "Option::is_none")]
    pub kratos_courier_smtp_headers: Option<serde_json::Value>,
    #[serde(rename = "kratos_identity_schemas", skip_serializing_if = "Option::is_none")]
    pub kratos_identity_schemas: Option<Vec<crate::models::ProjectRevisionIdentitySchema>>,
    #[serde(rename = "kratos_secrets_cipher", skip_serializing_if = "Option::is_none")]
    pub kratos_secrets_cipher: Option<Vec<String>>,
    #[serde(rename = "kratos_secrets_cookie", skip_serializing_if = "Option::is_none")]
    pub kratos_secrets_cookie: Option<Vec<String>>,
    #[serde(rename = "kratos_secrets_default", skip_serializing_if = "Option::is_none")]
    pub kratos_secrets_default: Option<Vec<String>>,
    #[serde(rename = "kratos_selfservice_allowed_return_urls", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_allowed_return_urls: Option<Vec<String>>,
    /// Configures the Ory Kratos Default Return URL  This governs the \"selfservice.allowed_return_urls\" setting.
    #[serde(rename = "kratos_selfservice_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Error UI URL  This governs the \"selfservice.flows.error.ui_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_error_ui_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_error_ui_url: Option<String>,
    #[serde(rename = "kratos_selfservice_flows_hooks", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_hooks: Option<Vec<crate::models::ProjectRevisionHook>>,
    /// Configures the Ory Kratos Login Default Return URL  This governs the \"selfservice.flows.login.after.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_login_after_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_login_after_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Login After OIDC Default Return URL  This governs the \"selfservice.flows.login.after.oidc.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_login_after_oidc_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_login_after_oidc_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Login After Password Default Return URL  This governs the \"selfservice.flows.login.after.password.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_login_after_password_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_login_after_password_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Login Lifespan  This governs the \"selfservice.flows.login.lifespan\" setting.
    #[serde(rename = "kratos_selfservice_flows_login_lifespan", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_login_lifespan: Option<String>,
    /// Configures the Ory Kratos Login UI URL  This governs the \"selfservice.flows.login.ui_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_login_ui_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_login_ui_url: Option<String>,
    /// Configures the Ory Kratos Logout Default Return URL  This governs the \"selfservice.flows.logout.after.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_logout_after_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_logout_after_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Recovery Default Return URL  This governs the \"selfservice.flows.recovery.after.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_recovery_after_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_recovery_after_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Recovery Enabled Setting  This governs the \"selfservice.flows.recovery.enabled\" setting.
    #[serde(rename = "kratos_selfservice_flows_recovery_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_recovery_enabled: Option<bool>,
    /// Configures the Ory Kratos Recovery Lifespan  This governs the \"selfservice.flows.recovery.lifespan\" setting.
    #[serde(rename = "kratos_selfservice_flows_recovery_lifespan", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_recovery_lifespan: Option<String>,
    /// Configures the Ory Kratos Recovery UI URL  This governs the \"selfservice.flows.recovery.ui_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_recovery_ui_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_recovery_ui_url: Option<String>,
    /// Configures the Ory Kratos Registration Default Return URL  This governs the \"selfservice.flows.registration.after.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_registration_after_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_registration_after_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Registration After OIDC Default Return URL  This governs the \"selfservice.flows.registration.after.oidc.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_registration_after_oidc_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_registration_after_oidc_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Registration After Password Default Return URL  This governs the \"selfservice.flows.registration.after.password.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_registration_after_password_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_registration_after_password_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Registration Lifespan  This governs the \"selfservice.flows.registration.lifespan\" setting.
    #[serde(rename = "kratos_selfservice_flows_registration_lifespan", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_registration_lifespan: Option<String>,
    /// Configures the Ory Kratos Registration UI URL  This governs the \"selfservice.flows.registration.ui_url\" setting.0
    #[serde(rename = "kratos_selfservice_flows_registration_ui_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_registration_ui_url: Option<String>,
    /// Configures the Ory Kratos Settings Default Return URL  This governs the \"selfservice.flows.settings.after.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_settings_after_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_settings_after_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Settings Default Return URL After Updating Passwords  This governs the \"selfservice.flows.settings.after.password.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_settings_after_password_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_settings_after_password_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Settings Default Return URL After Updating Profiles  This governs the \"selfservice.flows.settings.after.profile.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_settings_after_profile_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_settings_after_profile_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Settings Lifespan  This governs the \"selfservice.flows.settings.lifespan\" setting.
    #[serde(rename = "kratos_selfservice_flows_settings_lifespan", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_settings_lifespan: Option<String>,
    /// Configures the Ory Kratos Settings Privileged Session Max Age  This governs the \"selfservice.flows.settings.privileged_session_max_age\" setting.
    #[serde(rename = "kratos_selfservice_flows_settings_privileged_session_max_age", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_settings_privileged_session_max_age: Option<String>,
    /// Configures the Ory Kratos Settings Required AAL  This governs the \"selfservice.flows.settings.required_aal\" setting.
    #[serde(rename = "kratos_selfservice_flows_settings_required_aal", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_settings_required_aal: Option<String>,
    /// Configures the Ory Kratos Settings UI URL  This governs the \"selfservice.flows.settings.ui_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_settings_ui_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_settings_ui_url: Option<String>,
    /// Configures the Ory Kratos Verification Default Return URL  This governs the \"selfservice.flows.verification.after.default_browser_return_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_verification_after_default_browser_return_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_verification_after_default_browser_return_url: Option<String>,
    /// Configures the Ory Kratos Verification Enabled Setting  This governs the \"selfservice.flows.verification.enabled\" setting.
    #[serde(rename = "kratos_selfservice_flows_verification_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_verification_enabled: Option<bool>,
    /// Configures the Ory Kratos Verification Lifespan  This governs the \"selfservice.flows.verification.lifespan\" setting.
    #[serde(rename = "kratos_selfservice_flows_verification_lifespan", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_verification_lifespan: Option<String>,
    /// Configures the Ory Kratos Verification UI URL  This governs the \"selfservice.flows.verification.ui_url\" setting.
    #[serde(rename = "kratos_selfservice_flows_verification_ui_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_flows_verification_ui_url: Option<String>,
    /// Configures the Base URL which Recovery, Verification, and Login Links Point to  It is recommended to leave this value empty. It will be appropriately configured to the best matching domain (e.g. when using custom domains) automatically.  This governs the \"selfservice.methods.link.config.base_url\" setting.
    #[serde(rename = "kratos_selfservice_methods_link_config_base_url", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_link_config_base_url: Option<String>,
    /// Configures whether Ory Kratos Link Method is enabled  This governs the \"selfservice.methods.link.config.lifespan\" setting.
    #[serde(rename = "kratos_selfservice_methods_link_config_lifespan", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_link_config_lifespan: Option<String>,
    #[serde(rename = "kratos_selfservice_methods_link_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_link_enabled: Option<bool>,
    #[serde(rename = "kratos_selfservice_methods_lookup_secret_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_lookup_secret_enabled: Option<bool>,
    #[serde(rename = "kratos_selfservice_methods_oidc_config_providers", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_oidc_config_providers: Option<Vec<crate::models::ProjectRevisionThirdPartyLoginProvider>>,
    /// Configures whether Ory Kratos Third Party / OpenID Connect Login is enabled  This governs the \"selfservice.methods.oidc.enabled\" setting.
    #[serde(rename = "kratos_selfservice_methods_oidc_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_oidc_enabled: Option<bool>,
    #[serde(rename = "kratos_selfservice_methods_password_config_haveibeenpwned_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_password_config_haveibeenpwned_enabled: Option<bool>,
    #[serde(rename = "kratos_selfservice_methods_password_config_ignore_network_errors", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_password_config_ignore_network_errors: Option<bool>,
    /// Configures Ory Kratos Password Max Breaches Detection  This governs the \"selfservice.methods.password.enabled\" setting.
    #[serde(rename = "kratos_selfservice_methods_password_config_max_breaches", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_password_config_max_breaches: Option<i64>,
    #[serde(rename = "kratos_selfservice_methods_password_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_password_enabled: Option<bool>,
    #[serde(rename = "kratos_selfservice_methods_profile_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_profile_enabled: Option<bool>,
    /// Configures Ory Kratos TOTP Issuer  This governs the \"selfservice.methods.totp.config.issuer\" setting.
    #[serde(rename = "kratos_selfservice_methods_totp_config_issuer", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_totp_config_issuer: Option<String>,
    #[serde(rename = "kratos_selfservice_methods_totp_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_totp_enabled: Option<bool>,
    /// Configures the Ory Kratos Webauthn RP Display Name  This governs the \"selfservice.methods.webauthn.config.rp.display_name\" setting.
    #[serde(rename = "kratos_selfservice_methods_webauthn_config_rp_display_name", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_webauthn_config_rp_display_name: Option<String>,
    /// Configures the Ory Kratos Webauthn RP Icon  This governs the \"selfservice.methods.webauthn.config.rp.icon\" setting.
    #[serde(rename = "kratos_selfservice_methods_webauthn_config_rp_icon", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_webauthn_config_rp_icon: Option<String>,
    /// Configures the Ory Kratos Webauthn RP ID  This governs the \"selfservice.methods.webauthn.config.rp.id\" setting.
    #[serde(rename = "kratos_selfservice_methods_webauthn_config_rp_id", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_webauthn_config_rp_id: Option<String>,
    /// Configures the Ory Kratos Webauthn RP Origin  This governs the \"selfservice.methods.webauthn.config.rp.origin\" setting.
    #[serde(rename = "kratos_selfservice_methods_webauthn_config_rp_origin", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_webauthn_config_rp_origin: Option<String>,
    #[serde(rename = "kratos_selfservice_methods_webauthn_enabled", skip_serializing_if = "Option::is_none")]
    pub kratos_selfservice_methods_webauthn_enabled: Option<bool>,
    #[serde(rename = "kratos_session_cookie_persistent", skip_serializing_if = "Option::is_none")]
    pub kratos_session_cookie_persistent: Option<bool>,
    /// Configures the Ory Kratos Session Cookie SameSite Attribute  This governs the \"session.cookie.same_site\" setting.
    #[serde(rename = "kratos_session_cookie_same_site", skip_serializing_if = "Option::is_none")]
    pub kratos_session_cookie_same_site: Option<String>,
    /// Configures the Ory Kratos Session Lifespan  This governs the \"session.lifespan\" setting.
    #[serde(rename = "kratos_session_lifespan", skip_serializing_if = "Option::is_none")]
    pub kratos_session_lifespan: Option<String>,
    /// Configures the Ory Kratos Session Whoami AAL requirement  This governs the \"session.whoami.required_aal\" setting.
    #[serde(rename = "kratos_session_whoami_required_aal", skip_serializing_if = "Option::is_none")]
    pub kratos_session_whoami_required_aal: Option<String>,
    /// The project's name.
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "project_id", skip_serializing_if = "Option::is_none")]
    pub project_id: Option<String>,
    /// Last Time Project's Revision was Updated
    #[serde(rename = "updated_at", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<String>,
}

impl ProjectRevision {
    pub fn new(name: String) -> ProjectRevision {
        ProjectRevision {
            created_at: None,
            id: None,
            kratos_cookies_same_site: None,
            kratos_courier_smtp_connection_uri: None,
            kratos_courier_smtp_from_address: None,
            kratos_courier_smtp_from_name: None,
            kratos_courier_smtp_headers: None,
            kratos_identity_schemas: None,
            kratos_secrets_cipher: None,
            kratos_secrets_cookie: None,
            kratos_secrets_default: None,
            kratos_selfservice_allowed_return_urls: None,
            kratos_selfservice_default_browser_return_url: None,
            kratos_selfservice_flows_error_ui_url: None,
            kratos_selfservice_flows_hooks: None,
            kratos_selfservice_flows_login_after_default_browser_return_url: None,
            kratos_selfservice_flows_login_after_oidc_default_browser_return_url: None,
            kratos_selfservice_flows_login_after_password_default_browser_return_url: None,
            kratos_selfservice_flows_login_lifespan: None,
            kratos_selfservice_flows_login_ui_url: None,
            kratos_selfservice_flows_logout_after_default_browser_return_url: None,
            kratos_selfservice_flows_recovery_after_default_browser_return_url: None,
            kratos_selfservice_flows_recovery_enabled: None,
            kratos_selfservice_flows_recovery_lifespan: None,
            kratos_selfservice_flows_recovery_ui_url: None,
            kratos_selfservice_flows_registration_after_default_browser_return_url: None,
            kratos_selfservice_flows_registration_after_oidc_default_browser_return_url: None,
            kratos_selfservice_flows_registration_after_password_default_browser_return_url: None,
            kratos_selfservice_flows_registration_lifespan: None,
            kratos_selfservice_flows_registration_ui_url: None,
            kratos_selfservice_flows_settings_after_default_browser_return_url: None,
            kratos_selfservice_flows_settings_after_password_default_browser_return_url: None,
            kratos_selfservice_flows_settings_after_profile_default_browser_return_url: None,
            kratos_selfservice_flows_settings_lifespan: None,
            kratos_selfservice_flows_settings_privileged_session_max_age: None,
            kratos_selfservice_flows_settings_required_aal: None,
            kratos_selfservice_flows_settings_ui_url: None,
            kratos_selfservice_flows_verification_after_default_browser_return_url: None,
            kratos_selfservice_flows_verification_enabled: None,
            kratos_selfservice_flows_verification_lifespan: None,
            kratos_selfservice_flows_verification_ui_url: None,
            kratos_selfservice_methods_link_config_base_url: None,
            kratos_selfservice_methods_link_config_lifespan: None,
            kratos_selfservice_methods_link_enabled: None,
            kratos_selfservice_methods_lookup_secret_enabled: None,
            kratos_selfservice_methods_oidc_config_providers: None,
            kratos_selfservice_methods_oidc_enabled: None,
            kratos_selfservice_methods_password_config_haveibeenpwned_enabled: None,
            kratos_selfservice_methods_password_config_ignore_network_errors: None,
            kratos_selfservice_methods_password_config_max_breaches: None,
            kratos_selfservice_methods_password_enabled: None,
            kratos_selfservice_methods_profile_enabled: None,
            kratos_selfservice_methods_totp_config_issuer: None,
            kratos_selfservice_methods_totp_enabled: None,
            kratos_selfservice_methods_webauthn_config_rp_display_name: None,
            kratos_selfservice_methods_webauthn_config_rp_icon: None,
            kratos_selfservice_methods_webauthn_config_rp_id: None,
            kratos_selfservice_methods_webauthn_config_rp_origin: None,
            kratos_selfservice_methods_webauthn_enabled: None,
            kratos_session_cookie_persistent: None,
            kratos_session_cookie_same_site: None,
            kratos_session_lifespan: None,
            kratos_session_whoami_required_aal: None,
            name,
            project_id: None,
            updated_at: None,
        }
    }
}


