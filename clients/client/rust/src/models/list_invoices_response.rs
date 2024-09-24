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

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ListInvoicesResponse {
    #[serde(rename = "buckets")]
    pub buckets: Vec<models::BillingPeriodBucket>,
    #[serde(rename = "has_next_page")]
    pub has_next_page: bool,
    #[serde(rename = "next_page_token")]
    pub next_page_token: String,
}

impl ListInvoicesResponse {
    pub fn new(buckets: Vec<models::BillingPeriodBucket>, has_next_page: bool, next_page_token: String) -> ListInvoicesResponse {
        ListInvoicesResponse {
            buckets,
            has_next_page,
            next_page_token,
        }
    }
}

