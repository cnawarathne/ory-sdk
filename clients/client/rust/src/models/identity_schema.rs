/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.1.0-alpha.12
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// IdentitySchema : Together the name and identity uuid are a unique index constraint. This prevents a user from having schemas with the same name. This also allows schemas to have the same name across the system.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct IdentitySchema {
    /// The gcs file name  This is a randomly generated name which is used to uniquely identify the file on the blob storage
    #[serde(rename = "blob_name")]
    pub blob_name: String,
    /// The publicly accessible url of the schema
    #[serde(rename = "blob_url")]
    pub blob_url: String,
    /// The Content Hash  Contains a hash of the schema's content.
    #[serde(rename = "content_hash", skip_serializing_if = "Option::is_none")]
    pub content_hash: Option<String>,
    /// The Schema's Creation Date
    #[serde(rename = "created_at")]
    pub created_at: String,
    #[serde(rename = "id")]
    pub id: String,
    /// The schema name  This is set by the user and is for them to easily recognise their schema
    #[serde(rename = "name")]
    pub name: String,
    /// The actual Identity JSON Schema
    #[serde(rename = "schema", skip_serializing_if = "Option::is_none")]
    pub schema: Option<serde_json::Value>,
    /// Last Time Schema was Updated
    #[serde(rename = "updated_at")]
    pub updated_at: String,
}

impl IdentitySchema {
    /// Together the name and identity uuid are a unique index constraint. This prevents a user from having schemas with the same name. This also allows schemas to have the same name across the system.
    pub fn new(blob_name: String, blob_url: String, created_at: String, id: String, name: String, updated_at: String) -> IdentitySchema {
        IdentitySchema {
            blob_name,
            blob_url,
            content_hash: None,
            created_at,
            id,
            name,
            schema: None,
            updated_at,
        }
    }
}


