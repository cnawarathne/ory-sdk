/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.3.8
 * Contact: office@ory.sh
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

/// SessionAuthenticationMethod : A singular authenticator used during authentication / login.
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SessionAuthenticationMethod {
    #[serde(rename = "aal", skip_serializing_if = "Option::is_none")]
    pub aal: Option<models::AuthenticatorAssuranceLevel>,
    /// When the authentication challenge was completed.
    #[serde(rename = "completed_at", skip_serializing_if = "Option::is_none")]
    pub completed_at: Option<String>,
    #[serde(rename = "method", skip_serializing_if = "Option::is_none")]
    pub method: Option<MethodEnum>,
    /// The Organization id used for authentication
    #[serde(rename = "organization", skip_serializing_if = "Option::is_none")]
    pub organization: Option<String>,
    /// OIDC or SAML provider id used for authentication
    #[serde(rename = "provider", skip_serializing_if = "Option::is_none")]
    pub provider: Option<String>,
}

impl SessionAuthenticationMethod {
    /// A singular authenticator used during authentication / login.
    pub fn new() -> SessionAuthenticationMethod {
        SessionAuthenticationMethod {
            aal: None,
            completed_at: None,
            method: None,
            organization: None,
            provider: None,
        }
    }
}
/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum MethodEnum {
    #[serde(rename = "link_recovery")]
    LinkRecovery,
    #[serde(rename = "code_recovery")]
    CodeRecovery,
    #[serde(rename = "password")]
    Password,
    #[serde(rename = "code")]
    Code,
    #[serde(rename = "totp")]
    Totp,
    #[serde(rename = "oidc")]
    Oidc,
    #[serde(rename = "webauthn")]
    Webauthn,
    #[serde(rename = "lookup_secret")]
    LookupSecret,
    #[serde(rename = "v0.6_legacy_session")]
    V0Period6LegacySession,
}

impl Default for MethodEnum {
    fn default() -> MethodEnum {
        Self::LinkRecovery
    }
}

