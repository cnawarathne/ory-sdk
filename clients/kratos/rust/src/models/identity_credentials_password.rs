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

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct IdentityCredentialsPassword {
    /// HashedPassword is a hash-representation of the password.
    #[serde(rename = "hashed_password", skip_serializing_if = "Option::is_none")]
    pub hashed_password: Option<String>,
    /// UsePasswordMigrationHook is set to true if the password should be migrated using the password migration hook. If set, and the HashedPassword is empty, a webhook will be called during login to migrate the password.
    #[serde(rename = "use_password_migration_hook", skip_serializing_if = "Option::is_none")]
    pub use_password_migration_hook: Option<bool>,
}

impl IdentityCredentialsPassword {
    pub fn new() -> IdentityCredentialsPassword {
        IdentityCredentialsPassword {
            hashed_password: None,
            use_password_migration_hook: None,
        }
    }
}

