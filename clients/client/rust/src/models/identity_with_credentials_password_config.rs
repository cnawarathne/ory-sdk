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

/// IdentityWithCredentialsPasswordConfig : Create Identity and Import Password Credentials Configuration
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct IdentityWithCredentialsPasswordConfig {
    /// The hashed password in [PHC format](https://www.ory.sh/docs/kratos/manage-identities/import-user-accounts-identities#hashed-passwords)
    #[serde(rename = "hashed_password", skip_serializing_if = "Option::is_none")]
    pub hashed_password: Option<String>,
    /// The password in plain text if no hash is available.
    #[serde(rename = "password", skip_serializing_if = "Option::is_none")]
    pub password: Option<String>,
    /// If set to true, the password will be migrated using the password migration hook.
    #[serde(rename = "use_password_migration_hook", skip_serializing_if = "Option::is_none")]
    pub use_password_migration_hook: Option<bool>,
}

impl IdentityWithCredentialsPasswordConfig {
    /// Create Identity and Import Password Credentials Configuration
    pub fn new() -> IdentityWithCredentialsPasswordConfig {
        IdentityWithCredentialsPasswordConfig {
            hashed_password: None,
            password: None,
            use_password_migration_hook: None,
        }
    }
}

