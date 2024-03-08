/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.8.1
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateSubscriptionBody {
    ///  usd USD eur Euro
    #[serde(rename = "currency", skip_serializing_if = "Option::is_none")]
    pub currency: Option<CurrencyEnum>,
    ///  monthly Monthly yearly Yearly
    #[serde(rename = "interval")]
    pub interval: IntervalEnum,
    #[serde(rename = "plan")]
    pub plan: String,
    #[serde(rename = "provision_first_project")]
    pub provision_first_project: String,
    #[serde(rename = "return_to", skip_serializing_if = "Option::is_none")]
    pub return_to: Option<String>,
}


impl CreateSubscriptionBody {
    pub fn new(interval: IntervalEnum, plan: String, provision_first_project: String) -> CreateSubscriptionBody {
        CreateSubscriptionBody {
                currency: None,
                interval,
                plan,
                provision_first_project,
                return_to: None,
        }
    }
}

///  usd USD eur Euro
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum CurrencyEnum {
    #[serde(rename = "usd")]
    Usd,
    #[serde(rename = "eur")]
    Eur,
}
///  monthly Monthly yearly Yearly
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum IntervalEnum {
    #[serde(rename = "monthly")]
    Monthly,
    #[serde(rename = "yearly")]
    Yearly,
}

