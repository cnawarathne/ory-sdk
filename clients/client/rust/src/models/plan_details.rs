/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.4.0
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PlanDetails {
    /// BaseFeeMonthly is the monthly base fee for the plan.
    #[serde(rename = "base_fee_monthly")]
    pub base_fee_monthly: i64,
    /// BaseFeeYearly is the yearly base fee for the plan.
    #[serde(rename = "base_fee_yearly")]
    pub base_fee_yearly: i64,
    /// Custom is true if the plan is custom. This means it will be hidden from the pricing page.
    #[serde(rename = "custom")]
    pub custom: bool,
    /// Description is the description of the plan.
    #[serde(rename = "description")]
    pub description: String,
    /// Features are the feature definitions included in the plan.
    #[serde(rename = "features")]
    pub features: ::std::collections::HashMap<String, crate::models::GenericUsage>,
    /// Name is the name of the plan.
    #[serde(rename = "name")]
    pub name: String,
    /// Version is the version of the plan. The combination of `name@version` must be unique.
    #[serde(rename = "version")]
    pub version: i64,
}


impl PlanDetails {
    pub fn new(base_fee_monthly: i64, base_fee_yearly: i64, custom: bool, description: String, features: ::std::collections::HashMap<String, crate::models::GenericUsage>, name: String, version: i64) -> PlanDetails {
        PlanDetails {
                base_fee_monthly,
                base_fee_yearly,
                custom,
                description,
                features,
                name,
                version,
        }
    }
}


