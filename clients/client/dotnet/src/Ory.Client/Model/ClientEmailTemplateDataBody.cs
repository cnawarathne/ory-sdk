/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.0
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
    /// ClientEmailTemplateDataBody
    /// </summary>
    [DataContract(Name = "emailTemplateDataBody")]
    public partial class ClientEmailTemplateDataBody : IEquatable<ClientEmailTemplateDataBody>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientEmailTemplateDataBody" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientEmailTemplateDataBody()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientEmailTemplateDataBody" /> class.
        /// </summary>
        /// <param name="html">html (required).</param>
        /// <param name="plaintext">plaintext (required).</param>
        public ClientEmailTemplateDataBody(string html = default(string), string plaintext = default(string))
        {
            // to ensure "html" is required (not null)
            if (html == null) {
                throw new ArgumentNullException("html is a required property for ClientEmailTemplateDataBody and cannot be null");
            }
            this.Html = html;
            // to ensure "plaintext" is required (not null)
            if (plaintext == null) {
                throw new ArgumentNullException("plaintext is a required property for ClientEmailTemplateDataBody and cannot be null");
            }
            this.Plaintext = plaintext;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets Html
        /// </summary>
        [DataMember(Name = "html", IsRequired = true, EmitDefaultValue = false)]
        public string Html { get; set; }

        /// <summary>
        /// Gets or Sets Plaintext
        /// </summary>
        [DataMember(Name = "plaintext", IsRequired = true, EmitDefaultValue = false)]
        public string Plaintext { get; set; }

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
            sb.Append("class ClientEmailTemplateDataBody {\n");
            sb.Append("  Html: ").Append(Html).Append("\n");
            sb.Append("  Plaintext: ").Append(Plaintext).Append("\n");
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
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as ClientEmailTemplateDataBody);
        }

        /// <summary>
        /// Returns true if ClientEmailTemplateDataBody instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientEmailTemplateDataBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientEmailTemplateDataBody input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Html == input.Html ||
                    (this.Html != null &&
                    this.Html.Equals(input.Html))
                ) && 
                (
                    this.Plaintext == input.Plaintext ||
                    (this.Plaintext != null &&
                    this.Plaintext.Equals(input.Plaintext))
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.Html != null)
                {
                    hashCode = (hashCode * 59) + this.Html.GetHashCode();
                }
                if (this.Plaintext != null)
                {
                    hashCode = (hashCode * 59) + this.Plaintext.GetHashCode();
                }
                if (this.AdditionalProperties != null)
                {
                    hashCode = (hashCode * 59) + this.AdditionalProperties.GetHashCode();
                }
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
