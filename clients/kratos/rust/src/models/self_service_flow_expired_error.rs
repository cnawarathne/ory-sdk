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

/// SelfServiceFlowExpiredError : Is sent when a flow is expired
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SelfServiceFlowExpiredError {
    #[serde(rename = "error", skip_serializing_if = "Option::is_none")]
    pub error: Option<Box<models::GenericError>>,
    /// When the flow has expired
    #[serde(rename = "expired_at", skip_serializing_if = "Option::is_none")]
    pub expired_at: Option<String>,
    /// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
    #[serde(rename = "since", skip_serializing_if = "Option::is_none")]
    pub since: Option<i64>,
    /// The flow ID that should be used for the new flow as it contains the correct messages.
    #[serde(rename = "use_flow_id", skip_serializing_if = "Option::is_none")]
    pub use_flow_id: Option<String>,
}

impl SelfServiceFlowExpiredError {
    /// Is sent when a flow is expired
    pub fn new() -> SelfServiceFlowExpiredError {
        SelfServiceFlowExpiredError {
            error: None,
            expired_at: None,
            since: None,
            use_flow_id: None,
        }
    }
}

