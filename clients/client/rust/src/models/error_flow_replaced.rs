/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.15.4
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

/// ErrorFlowReplaced : Is sent when a flow is replaced by a different flow of the same class
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ErrorFlowReplaced {
    #[serde(rename = "error", skip_serializing_if = "Option::is_none")]
    pub error: Option<Box<models::GenericError>>,
    /// The flow ID that should be used for the new flow as it contains the correct messages.
    #[serde(rename = "use_flow_id", skip_serializing_if = "Option::is_none")]
    pub use_flow_id: Option<String>,
}

impl ErrorFlowReplaced {
    /// Is sent when a flow is replaced by a different flow of the same class
    pub fn new() -> ErrorFlowReplaced {
        ErrorFlowReplaced {
            error: None,
            use_flow_id: None,
        }
    }
}

