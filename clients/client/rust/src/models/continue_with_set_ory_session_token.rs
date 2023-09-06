/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.0
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// ContinueWithSetOrySessionToken : Indicates that a session was issued, and the application should use this token for authenticated requests



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ContinueWithSetOrySessionToken {
    /// Action will always be `set_ory_session_token` set_ory_session_token ContinueWithActionSetOrySessionToken show_verification_ui ContinueWithActionShowVerificationUI
    #[serde(rename = "action")]
    pub action: ActionEnum,
    /// Token is the token of the session
    #[serde(rename = "ory_session_token")]
    pub ory_session_token: String,
}


impl ContinueWithSetOrySessionToken {
    /// Indicates that a session was issued, and the application should use this token for authenticated requests
    pub fn new(action: ActionEnum, ory_session_token: String) -> ContinueWithSetOrySessionToken {
        ContinueWithSetOrySessionToken {
                action,
                ory_session_token,
        }
    }
}

/// Action will always be `set_ory_session_token` set_ory_session_token ContinueWithActionSetOrySessionToken show_verification_ui ContinueWithActionShowVerificationUI
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum ActionEnum {
    #[serde(rename = "set_ory_session_token")]
    SetOrySessionToken,
    #[serde(rename = "show_verification_ui")]
    ShowVerificationUi,
}

