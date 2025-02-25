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
    /// Create Identity and Import Password Credentials Configuration
    /// </summary>
    [DataContract(Name = "identityWithCredentialsPasswordConfig")]
    public partial class KratosIdentityWithCredentialsPasswordConfig : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosIdentityWithCredentialsPasswordConfig" /> class.
        /// </summary>
        /// <param name="hashedPassword">The hashed password in [PHC format](https://www.ory.sh/docs/kratos/manage-identities/import-user-accounts-identities#hashed-passwords).</param>
        /// <param name="password">The password in plain text if no hash is available..</param>
        /// <param name="usePasswordMigrationHook">If set to true, the password will be migrated using the password migration hook..</param>
        public KratosIdentityWithCredentialsPasswordConfig(string hashedPassword = default(string), string password = default(string), bool usePasswordMigrationHook = default(bool))
        {
            this.HashedPassword = hashedPassword;
            this.Password = password;
            this.UsePasswordMigrationHook = usePasswordMigrationHook;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The hashed password in [PHC format](https://www.ory.sh/docs/kratos/manage-identities/import-user-accounts-identities#hashed-passwords)
        /// </summary>
        /// <value>The hashed password in [PHC format](https://www.ory.sh/docs/kratos/manage-identities/import-user-accounts-identities#hashed-passwords)</value>
        [DataMember(Name = "hashed_password", EmitDefaultValue = false)]
        public string HashedPassword { get; set; }

        /// <summary>
        /// The password in plain text if no hash is available.
        /// </summary>
        /// <value>The password in plain text if no hash is available.</value>
        [DataMember(Name = "password", EmitDefaultValue = false)]
        public string Password { get; set; }

        /// <summary>
        /// If set to true, the password will be migrated using the password migration hook.
        /// </summary>
        /// <value>If set to true, the password will be migrated using the password migration hook.</value>
        [DataMember(Name = "use_password_migration_hook", EmitDefaultValue = true)]
        public bool UsePasswordMigrationHook { get; set; }

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
            sb.Append("class KratosIdentityWithCredentialsPasswordConfig {\n");
            sb.Append("  HashedPassword: ").Append(HashedPassword).Append("\n");
            sb.Append("  Password: ").Append(Password).Append("\n");
            sb.Append("  UsePasswordMigrationHook: ").Append(UsePasswordMigrationHook).Append("\n");
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
