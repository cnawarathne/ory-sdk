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
    /// Is Account Experience Enabled For Project Request Body
    /// </summary>
    [DataContract(Name = "internalIsAXWelcomeScreenEnabledForProjectBody")]
    public partial class ClientInternalIsAXWelcomeScreenEnabledForProjectBody : IEquatable<ClientInternalIsAXWelcomeScreenEnabledForProjectBody>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientInternalIsAXWelcomeScreenEnabledForProjectBody" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientInternalIsAXWelcomeScreenEnabledForProjectBody()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientInternalIsAXWelcomeScreenEnabledForProjectBody" /> class.
        /// </summary>
        /// <param name="path">Path is the path of the request. (required).</param>
        /// <param name="projectSlug">ProjectSlug is the project&#39;s slug. (required).</param>
        public ClientInternalIsAXWelcomeScreenEnabledForProjectBody(string path = default(string), string projectSlug = default(string))
        {
            // to ensure "path" is required (not null)
            if (path == null) {
                throw new ArgumentNullException("path is a required property for ClientInternalIsAXWelcomeScreenEnabledForProjectBody and cannot be null");
            }
            this.Path = path;
            // to ensure "projectSlug" is required (not null)
            if (projectSlug == null) {
                throw new ArgumentNullException("projectSlug is a required property for ClientInternalIsAXWelcomeScreenEnabledForProjectBody and cannot be null");
            }
            this.ProjectSlug = projectSlug;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Path is the path of the request.
        /// </summary>
        /// <value>Path is the path of the request.</value>
        [DataMember(Name = "path", IsRequired = true, EmitDefaultValue = false)]
        public string Path { get; set; }

        /// <summary>
        /// ProjectSlug is the project&#39;s slug.
        /// </summary>
        /// <value>ProjectSlug is the project&#39;s slug.</value>
        [DataMember(Name = "project_slug", IsRequired = true, EmitDefaultValue = false)]
        public string ProjectSlug { get; set; }

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
            sb.Append("class ClientInternalIsAXWelcomeScreenEnabledForProjectBody {\n");
            sb.Append("  Path: ").Append(Path).Append("\n");
            sb.Append("  ProjectSlug: ").Append(ProjectSlug).Append("\n");
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
            return this.Equals(input as ClientInternalIsAXWelcomeScreenEnabledForProjectBody);
        }

        /// <summary>
        /// Returns true if ClientInternalIsAXWelcomeScreenEnabledForProjectBody instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientInternalIsAXWelcomeScreenEnabledForProjectBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientInternalIsAXWelcomeScreenEnabledForProjectBody input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Path == input.Path ||
                    (this.Path != null &&
                    this.Path.Equals(input.Path))
                ) && 
                (
                    this.ProjectSlug == input.ProjectSlug ||
                    (this.ProjectSlug != null &&
                    this.ProjectSlug.Equals(input.ProjectSlug))
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
                if (this.Path != null)
                {
                    hashCode = (hashCode * 59) + this.Path.GetHashCode();
                }
                if (this.ProjectSlug != null)
                {
                    hashCode = (hashCode * 59) + this.ProjectSlug.GetHashCode();
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
