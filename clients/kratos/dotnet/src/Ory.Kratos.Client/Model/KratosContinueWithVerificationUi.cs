/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.3.8
 * Contact: office@ory.sh
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
using OpenAPIDateConverter = Ory.Kratos.Client.Client.OpenAPIDateConverter;

namespace Ory.Kratos.Client.Model
{
    /// <summary>
    /// Indicates, that the UI flow could be continued by showing a verification ui
    /// </summary>
    [DataContract(Name = "continueWithVerificationUi")]
    public partial class KratosContinueWithVerificationUi : IValidatableObject
    {
        /// <summary>
        /// Action will always be &#x60;show_verification_ui&#x60; show_verification_ui ContinueWithActionShowVerificationUIString
        /// </summary>
        /// <value>Action will always be &#x60;show_verification_ui&#x60; show_verification_ui ContinueWithActionShowVerificationUIString</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum ActionEnum
        {
            /// <summary>
            /// Enum ShowVerificationUi for value: show_verification_ui
            /// </summary>
            [EnumMember(Value = "show_verification_ui")]
            ShowVerificationUi = 1
        }


        /// <summary>
        /// Action will always be &#x60;show_verification_ui&#x60; show_verification_ui ContinueWithActionShowVerificationUIString
        /// </summary>
        /// <value>Action will always be &#x60;show_verification_ui&#x60; show_verification_ui ContinueWithActionShowVerificationUIString</value>
        [DataMember(Name = "action", IsRequired = true, EmitDefaultValue = true)]
        public ActionEnum Action { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosContinueWithVerificationUi" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected KratosContinueWithVerificationUi()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosContinueWithVerificationUi" /> class.
        /// </summary>
        /// <param name="action">Action will always be &#x60;show_verification_ui&#x60; show_verification_ui ContinueWithActionShowVerificationUIString (required).</param>
        /// <param name="flow">flow (required).</param>
        public KratosContinueWithVerificationUi(ActionEnum action = default(ActionEnum), KratosContinueWithVerificationUiFlow flow = default(KratosContinueWithVerificationUiFlow))
        {
            this.Action = action;
            // to ensure "flow" is required (not null)
            if (flow == null)
            {
                throw new ArgumentNullException("flow is a required property for KratosContinueWithVerificationUi and cannot be null");
            }
            this.Flow = flow;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets Flow
        /// </summary>
        [DataMember(Name = "flow", IsRequired = true, EmitDefaultValue = true)]
        public KratosContinueWithVerificationUiFlow Flow { get; set; }

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
            sb.Append("class KratosContinueWithVerificationUi {\n");
            sb.Append("  Action: ").Append(Action).Append("\n");
            sb.Append("  Flow: ").Append(Flow).Append("\n");
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
