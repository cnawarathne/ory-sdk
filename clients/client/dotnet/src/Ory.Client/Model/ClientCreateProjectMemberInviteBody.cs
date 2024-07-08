/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.13.6
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
    /// Create Project MemberInvite Request Body
    /// </summary>
    [DataContract(Name = "CreateProjectMemberInviteBody")]
    public partial class ClientCreateProjectMemberInviteBody : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientCreateProjectMemberInviteBody" /> class.
        /// </summary>
        /// <param name="inviteeEmail">A email to invite.</param>
        public ClientCreateProjectMemberInviteBody(string inviteeEmail = default(string))
        {
            this.InviteeEmail = inviteeEmail;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// A email to invite
        /// </summary>
        /// <value>A email to invite</value>
        [DataMember(Name = "invitee_email", EmitDefaultValue = false)]
        public string InviteeEmail { get; set; }

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
            sb.Append("class ClientCreateProjectMemberInviteBody {\n");
            sb.Append("  InviteeEmail: ").Append(InviteeEmail).Append("\n");
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
