/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.4.0
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
    /// Create B2B SSO Organization Request Body
    /// </summary>
    [DataContract(Name = "OrganizationBody")]
    public partial class ClientOrganizationBody : IEquatable<ClientOrganizationBody>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientOrganizationBody" /> class.
        /// </summary>
        /// <param name="domains">Domains contains the list of organization&#39;s domains..</param>
        /// <param name="label">Label contains the organization&#39;s label..</param>
        public ClientOrganizationBody(List<string> domains = default(List<string>), string label = default(string))
        {
            this.Domains = domains;
            this.Label = label;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Domains contains the list of organization&#39;s domains.
        /// </summary>
        /// <value>Domains contains the list of organization&#39;s domains.</value>
        [DataMember(Name = "domains", EmitDefaultValue = false)]
        public List<string> Domains { get; set; }

        /// <summary>
        /// Label contains the organization&#39;s label.
        /// </summary>
        /// <value>Label contains the organization&#39;s label.</value>
        [DataMember(Name = "label", EmitDefaultValue = false)]
        public string Label { get; set; }

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
            sb.Append("class ClientOrganizationBody {\n");
            sb.Append("  Domains: ").Append(Domains).Append("\n");
            sb.Append("  Label: ").Append(Label).Append("\n");
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
            return this.Equals(input as ClientOrganizationBody);
        }

        /// <summary>
        /// Returns true if ClientOrganizationBody instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientOrganizationBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientOrganizationBody input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Domains == input.Domains ||
                    this.Domains != null &&
                    input.Domains != null &&
                    this.Domains.SequenceEqual(input.Domains)
                ) && 
                (
                    this.Label == input.Label ||
                    (this.Label != null &&
                    this.Label.Equals(input.Label))
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
                if (this.Domains != null)
                {
                    hashCode = (hashCode * 59) + this.Domains.GetHashCode();
                }
                if (this.Label != null)
                {
                    hashCode = (hashCode * 59) + this.Label.GetHashCode();
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
