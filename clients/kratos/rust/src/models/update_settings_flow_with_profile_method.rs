/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.4.0-alpha.0
 * Contact: office@ory.sh
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

/// UpdateSettingsFlowWithProfileMethod : Update Settings Flow with Profile Method
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UpdateSettingsFlowWithProfileMethod {
    /// The Anti-CSRF Token  This token is only required when performing browser flows.
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method  Should be set to profile when trying to update a profile.
    #[serde(rename = "method")]
    pub method: String,
    /// Traits  The identity's traits.
    #[serde(rename = "traits")]
    pub traits: serde_json::Value,
    /// Transient data to pass along to any webhooks
    #[serde(rename = "transient_payload", skip_serializing_if = "Option::is_none")]
    pub transient_payload: Option<serde_json::Value>,
}

impl UpdateSettingsFlowWithProfileMethod {
    /// Update Settings Flow with Profile Method
    pub fn new(method: String, traits: serde_json::Value) -> UpdateSettingsFlowWithProfileMethod {
        UpdateSettingsFlowWithProfileMethod {
            csrf_token: None,
            method,
            traits,
            transient_payload: None,
        }
    }
}

