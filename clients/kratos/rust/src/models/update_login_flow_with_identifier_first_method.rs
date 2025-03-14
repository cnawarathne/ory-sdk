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

/// UpdateLoginFlowWithIdentifierFirstMethod : Update Login Flow with Multi-Step Method
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UpdateLoginFlowWithIdentifierFirstMethod {
    /// Sending the anti-csrf token is only required for browser login flows.
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Identifier is the email or username of the user trying to log in.
    #[serde(rename = "identifier")]
    pub identifier: String,
    /// Method should be set to \"password\" when logging in using the identifier and password strategy.
    #[serde(rename = "method")]
    pub method: String,
    /// Transient data to pass along to any webhooks
    #[serde(rename = "transient_payload", skip_serializing_if = "Option::is_none")]
    pub transient_payload: Option<serde_json::Value>,
}

impl UpdateLoginFlowWithIdentifierFirstMethod {
    /// Update Login Flow with Multi-Step Method
    pub fn new(identifier: String, method: String) -> UpdateLoginFlowWithIdentifierFirstMethod {
        UpdateLoginFlowWithIdentifierFirstMethod {
            csrf_token: None,
            identifier,
            method,
            transient_payload: None,
        }
    }
}

