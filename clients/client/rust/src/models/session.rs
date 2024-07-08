/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.13.6
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

/// Session : A Session
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Session {
    /// Active state. If false the session is no longer active.
    #[serde(rename = "active", skip_serializing_if = "Option::is_none")]
    pub active: Option<bool>,
    /// The Session Authentication Timestamp  When this session was authenticated at. If multi-factor authentication was used this is the time when the last factor was authenticated (e.g. the TOTP code challenge was completed).
    #[serde(rename = "authenticated_at", skip_serializing_if = "Option::is_none")]
    pub authenticated_at: Option<String>,
    /// A list of authenticators which were used to authenticate the session.
    #[serde(rename = "authentication_methods", skip_serializing_if = "Option::is_none")]
    pub authentication_methods: Option<Vec<models::SessionAuthenticationMethod>>,
    #[serde(rename = "authenticator_assurance_level", skip_serializing_if = "Option::is_none")]
    pub authenticator_assurance_level: Option<models::AuthenticatorAssuranceLevel>,
    /// Devices has history of all endpoints where the session was used
    #[serde(rename = "devices", skip_serializing_if = "Option::is_none")]
    pub devices: Option<Vec<models::SessionDevice>>,
    /// The Session Expiry  When this session expires at.
    #[serde(rename = "expires_at", skip_serializing_if = "Option::is_none")]
    pub expires_at: Option<String>,
    /// Session ID
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "identity", skip_serializing_if = "Option::is_none")]
    pub identity: Option<Box<models::Identity>>,
    /// The Session Issuance Timestamp  When this session was issued at. Usually equal or close to `authenticated_at`.
    #[serde(rename = "issued_at", skip_serializing_if = "Option::is_none")]
    pub issued_at: Option<String>,
    /// Tokenized is the tokenized (e.g. JWT) version of the session.  It is only set when the `tokenize` query parameter was set to a valid tokenize template during calls to `/session/whoami`.
    #[serde(rename = "tokenized", skip_serializing_if = "Option::is_none")]
    pub tokenized: Option<String>,
}

impl Session {
    /// A Session
    pub fn new(id: String) -> Session {
        Session {
            active: None,
            authenticated_at: None,
            authentication_methods: None,
            authenticator_assurance_level: None,
            devices: None,
            expires_at: None,
            id,
            identity: None,
            issued_at: None,
            tokenized: None,
        }
    }
}

