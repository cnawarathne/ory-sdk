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
    /// KratosUiNodeScriptAttributes
    /// </summary>
    [DataContract(Name = "uiNodeScriptAttributes")]
    public partial class KratosUiNodeScriptAttributes : IValidatableObject
    {
        /// <summary>
        /// NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \&quot;script\&quot;. text Text input Input img Image a Anchor script Script
        /// </summary>
        /// <value>NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \&quot;script\&quot;. text Text input Input img Image a Anchor script Script</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum NodeTypeEnum
        {
            /// <summary>
            /// Enum Text for value: text
            /// </summary>
            [EnumMember(Value = "text")]
            Text = 1,

            /// <summary>
            /// Enum Input for value: input
            /// </summary>
            [EnumMember(Value = "input")]
            Input = 2,

            /// <summary>
            /// Enum Img for value: img
            /// </summary>
            [EnumMember(Value = "img")]
            Img = 3,

            /// <summary>
            /// Enum A for value: a
            /// </summary>
            [EnumMember(Value = "a")]
            A = 4,

            /// <summary>
            /// Enum Script for value: script
            /// </summary>
            [EnumMember(Value = "script")]
            Script = 5
        }


        /// <summary>
        /// NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \&quot;script\&quot;. text Text input Input img Image a Anchor script Script
        /// </summary>
        /// <value>NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \&quot;script\&quot;. text Text input Input img Image a Anchor script Script</value>
        [DataMember(Name = "node_type", IsRequired = true, EmitDefaultValue = true)]
        public NodeTypeEnum NodeType { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosUiNodeScriptAttributes" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected KratosUiNodeScriptAttributes()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosUiNodeScriptAttributes" /> class.
        /// </summary>
        /// <param name="async">The script async type (required).</param>
        /// <param name="crossorigin">The script cross origin policy (required).</param>
        /// <param name="id">A unique identifier (required).</param>
        /// <param name="integrity">The script&#39;s integrity hash (required).</param>
        /// <param name="nodeType">NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \&quot;script\&quot;. text Text input Input img Image a Anchor script Script (required).</param>
        /// <param name="nonce">Nonce for CSP  A nonce you may want to use to improve your Content Security Policy. You do not have to use this value but if you want to improve your CSP policies you may use it. You can also choose to use your own nonce value! (required).</param>
        /// <param name="referrerpolicy">The script referrer policy (required).</param>
        /// <param name="src">The script source (required).</param>
        /// <param name="type">The script MIME type (required).</param>
        public KratosUiNodeScriptAttributes(bool async = default(bool), string crossorigin = default(string), string id = default(string), string integrity = default(string), NodeTypeEnum nodeType = default(NodeTypeEnum), string nonce = default(string), string referrerpolicy = default(string), string src = default(string), string type = default(string))
        {
            this.Async = async;
            // to ensure "crossorigin" is required (not null)
            if (crossorigin == null)
            {
                throw new ArgumentNullException("crossorigin is a required property for KratosUiNodeScriptAttributes and cannot be null");
            }
            this.Crossorigin = crossorigin;
            // to ensure "id" is required (not null)
            if (id == null)
            {
                throw new ArgumentNullException("id is a required property for KratosUiNodeScriptAttributes and cannot be null");
            }
            this.Id = id;
            // to ensure "integrity" is required (not null)
            if (integrity == null)
            {
                throw new ArgumentNullException("integrity is a required property for KratosUiNodeScriptAttributes and cannot be null");
            }
            this.Integrity = integrity;
            this.NodeType = nodeType;
            // to ensure "nonce" is required (not null)
            if (nonce == null)
            {
                throw new ArgumentNullException("nonce is a required property for KratosUiNodeScriptAttributes and cannot be null");
            }
            this.Nonce = nonce;
            // to ensure "referrerpolicy" is required (not null)
            if (referrerpolicy == null)
            {
                throw new ArgumentNullException("referrerpolicy is a required property for KratosUiNodeScriptAttributes and cannot be null");
            }
            this.Referrerpolicy = referrerpolicy;
            // to ensure "src" is required (not null)
            if (src == null)
            {
                throw new ArgumentNullException("src is a required property for KratosUiNodeScriptAttributes and cannot be null");
            }
            this.Src = src;
            // to ensure "type" is required (not null)
            if (type == null)
            {
                throw new ArgumentNullException("type is a required property for KratosUiNodeScriptAttributes and cannot be null");
            }
            this.Type = type;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The script async type
        /// </summary>
        /// <value>The script async type</value>
        [DataMember(Name = "async", IsRequired = true, EmitDefaultValue = true)]
        public bool Async { get; set; }

        /// <summary>
        /// The script cross origin policy
        /// </summary>
        /// <value>The script cross origin policy</value>
        [DataMember(Name = "crossorigin", IsRequired = true, EmitDefaultValue = true)]
        public string Crossorigin { get; set; }

        /// <summary>
        /// A unique identifier
        /// </summary>
        /// <value>A unique identifier</value>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = true)]
        public string Id { get; set; }

        /// <summary>
        /// The script&#39;s integrity hash
        /// </summary>
        /// <value>The script&#39;s integrity hash</value>
        [DataMember(Name = "integrity", IsRequired = true, EmitDefaultValue = true)]
        public string Integrity { get; set; }

        /// <summary>
        /// Nonce for CSP  A nonce you may want to use to improve your Content Security Policy. You do not have to use this value but if you want to improve your CSP policies you may use it. You can also choose to use your own nonce value!
        /// </summary>
        /// <value>Nonce for CSP  A nonce you may want to use to improve your Content Security Policy. You do not have to use this value but if you want to improve your CSP policies you may use it. You can also choose to use your own nonce value!</value>
        [DataMember(Name = "nonce", IsRequired = true, EmitDefaultValue = true)]
        public string Nonce { get; set; }

        /// <summary>
        /// The script referrer policy
        /// </summary>
        /// <value>The script referrer policy</value>
        [DataMember(Name = "referrerpolicy", IsRequired = true, EmitDefaultValue = true)]
        public string Referrerpolicy { get; set; }

        /// <summary>
        /// The script source
        /// </summary>
        /// <value>The script source</value>
        [DataMember(Name = "src", IsRequired = true, EmitDefaultValue = true)]
        public string Src { get; set; }

        /// <summary>
        /// The script MIME type
        /// </summary>
        /// <value>The script MIME type</value>
        [DataMember(Name = "type", IsRequired = true, EmitDefaultValue = true)]
        public string Type { get; set; }

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
            sb.Append("class KratosUiNodeScriptAttributes {\n");
            sb.Append("  Async: ").Append(Async).Append("\n");
            sb.Append("  Crossorigin: ").Append(Crossorigin).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Integrity: ").Append(Integrity).Append("\n");
            sb.Append("  NodeType: ").Append(NodeType).Append("\n");
            sb.Append("  Nonce: ").Append(Nonce).Append("\n");
            sb.Append("  Referrerpolicy: ").Append(Referrerpolicy).Append("\n");
            sb.Append("  Src: ").Append(Src).Append("\n");
            sb.Append("  Type: ").Append(Type).Append("\n");
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
