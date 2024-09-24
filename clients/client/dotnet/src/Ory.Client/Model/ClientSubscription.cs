/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.15.4
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Client.Client.OpenAPIDateConverter;

namespace Ory.Client.Model
{
    /// <summary>
    /// ClientSubscription
    /// </summary>
    [DataContract(Name = "subscription")]
    public partial class ClientSubscription : IValidatableObject
    {
        /// <summary>
        /// The currency of the subscription. To change this, a new subscription must be created. usd USD eur Euro
        /// </summary>
        /// <value>The currency of the subscription. To change this, a new subscription must be created. usd USD eur Euro</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum CurrencyEnum
        {
            /// <summary>
            /// Enum Usd for value: usd
            /// </summary>
            [EnumMember(Value = "usd")]
            Usd = 1,

            /// <summary>
            /// Enum Eur for value: eur
            /// </summary>
            [EnumMember(Value = "eur")]
            Eur = 2
        }


        /// <summary>
        /// The currency of the subscription. To change this, a new subscription must be created. usd USD eur Euro
        /// </summary>
        /// <value>The currency of the subscription. To change this, a new subscription must be created. usd USD eur Euro</value>
        [DataMember(Name = "currency", IsRequired = true, EmitDefaultValue = true)]
        public CurrencyEnum Currency { get; set; }

        /// <summary>
        /// Returns false as Currency should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeCurrency()
        {
            return false;
        }
        /// <summary>
        /// The currently active interval of the subscription monthly Monthly yearly Yearly
        /// </summary>
        /// <value>The currently active interval of the subscription monthly Monthly yearly Yearly</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum CurrentIntervalEnum
        {
            /// <summary>
            /// Enum Monthly for value: monthly
            /// </summary>
            [EnumMember(Value = "monthly")]
            Monthly = 1,

            /// <summary>
            /// Enum Yearly for value: yearly
            /// </summary>
            [EnumMember(Value = "yearly")]
            Yearly = 2
        }


        /// <summary>
        /// The currently active interval of the subscription monthly Monthly yearly Yearly
        /// </summary>
        /// <value>The currently active interval of the subscription monthly Monthly yearly Yearly</value>
        [DataMember(Name = "current_interval", IsRequired = true, EmitDefaultValue = true)]
        public CurrentIntervalEnum CurrentInterval { get; set; }

        /// <summary>
        /// Returns false as CurrentInterval should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeCurrentInterval()
        {
            return false;
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientSubscription" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientSubscription()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientSubscription" /> class.
        /// </summary>
        /// <param name="currentPlanDetails">currentPlanDetails.</param>
        /// <param name="intervalChangesTo">intervalChangesTo (required).</param>
        /// <param name="ongoingStripeCheckoutId">ongoingStripeCheckoutId.</param>
        /// <param name="planChangesAt">planChangesAt.</param>
        /// <param name="planChangesTo">planChangesTo (required).</param>
        /// <param name="status">For &#x60;collection_method&#x3D;charge_automatically&#x60; a subscription moves into &#x60;incomplete&#x60; if the initial payment attempt fails. A subscription in this status can only have metadata and default_source updated. Once the first invoice is paid, the subscription moves into an &#x60;active&#x60; status. If the first invoice is not paid within 23 hours, the subscription transitions to &#x60;incomplete_expired&#x60;. This is a terminal status, the open invoice will be voided and no further invoices will be generated.  A subscription that is currently in a trial period is &#x60;trialing&#x60; and moves to &#x60;active&#x60; when the trial period is over.  A subscription can only enter a &#x60;paused&#x60; status [when a trial ends without a payment method](https://stripe.com/billing/subscriptions/trials#create-free-trials-without-payment). A &#x60;paused&#x60; subscription doesn&#39;t generate invoices and can be resumed after your customer adds their payment method. The &#x60;paused&#x60; status is different from [pausing collection](https://stripe.com/billing/subscriptions/pause-payment), which still generates invoices and leaves the subscription&#39;s status unchanged.  If subscription &#x60;collection_method&#x3D;charge_automatically&#x60;, it becomes &#x60;past_due&#x60; when payment is required but cannot be paid (due to failed payment or awaiting additional user actions). Once Stripe has exhausted all payment retry attempts, the subscription will become &#x60;canceled&#x60; or &#x60;unpaid&#x60; (depending on your subscriptions settings).  If subscription &#x60;collection_method&#x3D;send_invoice&#x60; it becomes &#x60;past_due&#x60; when its invoice is not paid by the due date, and &#x60;canceled&#x60; or &#x60;unpaid&#x60; if it is still not paid by an additional deadline after that. Note that when a subscription has a status of &#x60;unpaid&#x60;, no subsequent invoices will be attempted (invoices will be created, but then immediately automatically closed). After receiving updated payment information from a customer, you may choose to reopen and pay their closed invoices. (required).</param>
        /// <param name="stripeCheckoutExpiresAt">stripeCheckoutExpiresAt.</param>
        public ClientSubscription(ClientPlanDetails currentPlanDetails = default(ClientPlanDetails), string intervalChangesTo = default(string), string ongoingStripeCheckoutId = default(string), DateTime planChangesAt = default(DateTime), string planChangesTo = default(string), string status = default(string), DateTime stripeCheckoutExpiresAt = default(DateTime))
        {
            // to ensure "intervalChangesTo" is required (not null)
            if (intervalChangesTo == null)
            {
                throw new ArgumentNullException("intervalChangesTo is a required property for ClientSubscription and cannot be null");
            }
            this.IntervalChangesTo = intervalChangesTo;
            // to ensure "planChangesTo" is required (not null)
            if (planChangesTo == null)
            {
                throw new ArgumentNullException("planChangesTo is a required property for ClientSubscription and cannot be null");
            }
            this.PlanChangesTo = planChangesTo;
            // to ensure "status" is required (not null)
            if (status == null)
            {
                throw new ArgumentNullException("status is a required property for ClientSubscription and cannot be null");
            }
            this.Status = status;
            this.CurrentPlanDetails = currentPlanDetails;
            this.OngoingStripeCheckoutId = ongoingStripeCheckoutId;
            this.PlanChangesAt = planChangesAt;
            this.StripeCheckoutExpiresAt = stripeCheckoutExpiresAt;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets CreatedAt
        /// </summary>
        [DataMember(Name = "created_at", IsRequired = true, EmitDefaultValue = true)]
        public DateTime CreatedAt { get; private set; }

        /// <summary>
        /// Returns false as CreatedAt should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeCreatedAt()
        {
            return false;
        }
        /// <summary>
        /// The currently active plan of the subscription
        /// </summary>
        /// <value>The currently active plan of the subscription</value>
        [DataMember(Name = "current_plan", IsRequired = true, EmitDefaultValue = true)]
        public string CurrentPlan { get; private set; }

        /// <summary>
        /// Returns false as CurrentPlan should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeCurrentPlan()
        {
            return false;
        }
        /// <summary>
        /// Gets or Sets CurrentPlanDetails
        /// </summary>
        [DataMember(Name = "current_plan_details", EmitDefaultValue = false)]
        public ClientPlanDetails CurrentPlanDetails { get; set; }

        /// <summary>
        /// The ID of the stripe customer
        /// </summary>
        /// <value>The ID of the stripe customer</value>
        [DataMember(Name = "customer_id", IsRequired = true, EmitDefaultValue = true)]
        public string CustomerId { get; private set; }

        /// <summary>
        /// Returns false as CustomerId should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeCustomerId()
        {
            return false;
        }
        /// <summary>
        /// The ID of the subscription
        /// </summary>
        /// <value>The ID of the subscription</value>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = true)]
        public string Id { get; private set; }

        /// <summary>
        /// Returns false as Id should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeId()
        {
            return false;
        }
        /// <summary>
        /// Gets or Sets IntervalChangesTo
        /// </summary>
        [DataMember(Name = "interval_changes_to", IsRequired = true, EmitDefaultValue = true)]
        public string IntervalChangesTo { get; set; }

        /// <summary>
        /// Gets or Sets OngoingStripeCheckoutId
        /// </summary>
        [DataMember(Name = "ongoing_stripe_checkout_id", EmitDefaultValue = true)]
        public string OngoingStripeCheckoutId { get; set; }

        /// <summary>
        /// Until when the subscription is payed
        /// </summary>
        /// <value>Until when the subscription is payed</value>
        [DataMember(Name = "payed_until", IsRequired = true, EmitDefaultValue = true)]
        public DateTime PayedUntil { get; private set; }

        /// <summary>
        /// Returns false as PayedUntil should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializePayedUntil()
        {
            return false;
        }
        /// <summary>
        /// Gets or Sets PlanChangesAt
        /// </summary>
        [DataMember(Name = "plan_changes_at", EmitDefaultValue = false)]
        public DateTime PlanChangesAt { get; set; }

        /// <summary>
        /// Gets or Sets PlanChangesTo
        /// </summary>
        [DataMember(Name = "plan_changes_to", IsRequired = true, EmitDefaultValue = true)]
        public string PlanChangesTo { get; set; }

        /// <summary>
        /// For &#x60;collection_method&#x3D;charge_automatically&#x60; a subscription moves into &#x60;incomplete&#x60; if the initial payment attempt fails. A subscription in this status can only have metadata and default_source updated. Once the first invoice is paid, the subscription moves into an &#x60;active&#x60; status. If the first invoice is not paid within 23 hours, the subscription transitions to &#x60;incomplete_expired&#x60;. This is a terminal status, the open invoice will be voided and no further invoices will be generated.  A subscription that is currently in a trial period is &#x60;trialing&#x60; and moves to &#x60;active&#x60; when the trial period is over.  A subscription can only enter a &#x60;paused&#x60; status [when a trial ends without a payment method](https://stripe.com/billing/subscriptions/trials#create-free-trials-without-payment). A &#x60;paused&#x60; subscription doesn&#39;t generate invoices and can be resumed after your customer adds their payment method. The &#x60;paused&#x60; status is different from [pausing collection](https://stripe.com/billing/subscriptions/pause-payment), which still generates invoices and leaves the subscription&#39;s status unchanged.  If subscription &#x60;collection_method&#x3D;charge_automatically&#x60;, it becomes &#x60;past_due&#x60; when payment is required but cannot be paid (due to failed payment or awaiting additional user actions). Once Stripe has exhausted all payment retry attempts, the subscription will become &#x60;canceled&#x60; or &#x60;unpaid&#x60; (depending on your subscriptions settings).  If subscription &#x60;collection_method&#x3D;send_invoice&#x60; it becomes &#x60;past_due&#x60; when its invoice is not paid by the due date, and &#x60;canceled&#x60; or &#x60;unpaid&#x60; if it is still not paid by an additional deadline after that. Note that when a subscription has a status of &#x60;unpaid&#x60;, no subsequent invoices will be attempted (invoices will be created, but then immediately automatically closed). After receiving updated payment information from a customer, you may choose to reopen and pay their closed invoices.
        /// </summary>
        /// <value>For &#x60;collection_method&#x3D;charge_automatically&#x60; a subscription moves into &#x60;incomplete&#x60; if the initial payment attempt fails. A subscription in this status can only have metadata and default_source updated. Once the first invoice is paid, the subscription moves into an &#x60;active&#x60; status. If the first invoice is not paid within 23 hours, the subscription transitions to &#x60;incomplete_expired&#x60;. This is a terminal status, the open invoice will be voided and no further invoices will be generated.  A subscription that is currently in a trial period is &#x60;trialing&#x60; and moves to &#x60;active&#x60; when the trial period is over.  A subscription can only enter a &#x60;paused&#x60; status [when a trial ends without a payment method](https://stripe.com/billing/subscriptions/trials#create-free-trials-without-payment). A &#x60;paused&#x60; subscription doesn&#39;t generate invoices and can be resumed after your customer adds their payment method. The &#x60;paused&#x60; status is different from [pausing collection](https://stripe.com/billing/subscriptions/pause-payment), which still generates invoices and leaves the subscription&#39;s status unchanged.  If subscription &#x60;collection_method&#x3D;charge_automatically&#x60;, it becomes &#x60;past_due&#x60; when payment is required but cannot be paid (due to failed payment or awaiting additional user actions). Once Stripe has exhausted all payment retry attempts, the subscription will become &#x60;canceled&#x60; or &#x60;unpaid&#x60; (depending on your subscriptions settings).  If subscription &#x60;collection_method&#x3D;send_invoice&#x60; it becomes &#x60;past_due&#x60; when its invoice is not paid by the due date, and &#x60;canceled&#x60; or &#x60;unpaid&#x60; if it is still not paid by an additional deadline after that. Note that when a subscription has a status of &#x60;unpaid&#x60;, no subsequent invoices will be attempted (invoices will be created, but then immediately automatically closed). After receiving updated payment information from a customer, you may choose to reopen and pay their closed invoices.</value>
        [DataMember(Name = "status", IsRequired = true, EmitDefaultValue = true)]
        public string Status { get; set; }

        /// <summary>
        /// Gets or Sets StripeCheckoutExpiresAt
        /// </summary>
        [DataMember(Name = "stripe_checkout_expires_at", EmitDefaultValue = false)]
        public DateTime StripeCheckoutExpiresAt { get; set; }

        /// <summary>
        /// Gets or Sets UpdatedAt
        /// </summary>
        [DataMember(Name = "updated_at", IsRequired = true, EmitDefaultValue = true)]
        public DateTime UpdatedAt { get; private set; }

        /// <summary>
        /// Returns false as UpdatedAt should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeUpdatedAt()
        {
            return false;
        }
        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class ClientSubscription {\n");
            sb.Append("  CreatedAt: ").Append(CreatedAt).Append("\n");
            sb.Append("  Currency: ").Append(Currency).Append("\n");
            sb.Append("  CurrentInterval: ").Append(CurrentInterval).Append("\n");
            sb.Append("  CurrentPlan: ").Append(CurrentPlan).Append("\n");
            sb.Append("  CurrentPlanDetails: ").Append(CurrentPlanDetails).Append("\n");
            sb.Append("  CustomerId: ").Append(CustomerId).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  IntervalChangesTo: ").Append(IntervalChangesTo).Append("\n");
            sb.Append("  OngoingStripeCheckoutId: ").Append(OngoingStripeCheckoutId).Append("\n");
            sb.Append("  PayedUntil: ").Append(PayedUntil).Append("\n");
            sb.Append("  PlanChangesAt: ").Append(PlanChangesAt).Append("\n");
            sb.Append("  PlanChangesTo: ").Append(PlanChangesTo).Append("\n");
            sb.Append("  Status: ").Append(Status).Append("\n");
            sb.Append("  StripeCheckoutExpiresAt: ").Append(StripeCheckoutExpiresAt).Append("\n");
            sb.Append("  UpdatedAt: ").Append(UpdatedAt).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
