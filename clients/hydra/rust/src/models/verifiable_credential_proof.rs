/*
 * Ory Hydra API
 *
 * Documentation for all of Ory Hydra's APIs. 
 *
 * The version of the OpenAPI document: v2.4.0-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct VerifiableCredentialProof {
    #[serde(rename = "jwt", skip_serializing_if = "Option::is_none")]
    pub jwt: Option<String>,
    #[serde(rename = "proof_type", skip_serializing_if = "Option::is_none")]
    pub proof_type: Option<String>,
}

impl VerifiableCredentialProof {
    pub fn new() -> VerifiableCredentialProof {
        VerifiableCredentialProof {
            jwt: None,
            proof_type: None,
        }
    }
}

