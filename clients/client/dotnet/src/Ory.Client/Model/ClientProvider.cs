/*
 * Ory APIs
 *
 * # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | - -- -- -- -- -- -- - | - -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- - | - -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- - | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 
 *
 * The version of the OpenAPI document: v1.17.2
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
    /// ClientProvider
    /// </summary>
    [DataContract(Name = "Provider")]
    public partial class ClientProvider : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientProvider" /> class.
        /// </summary>
        /// <param name="clientId">The RP&#39;s client identifier, issued by the IdP..</param>
        /// <param name="configUrl">A full path of the IdP config file..</param>
        /// <param name="domainHint">By specifying one of domain_hints values provided by the accounts endpoints, the FedCM dialog selectively shows the specified account..</param>
        /// <param name="fields">Array of strings that specifies the user information (\&quot;name\&quot;, \&quot; email\&quot;, \&quot;picture\&quot;) that RP needs IdP to share with them.  Note: Field API is supported by Chrome 132 and later..</param>
        /// <param name="loginHint">By specifying one of login_hints values provided by the accounts endpoints, the FedCM dialog selectively shows the specified account..</param>
        /// <param name="nonce">A random string to ensure the response is issued for this specific request. Prevents replay attacks..</param>
        /// <param name="parameters">Custom object that allows to specify additional key-value parameters: scope: A string value containing additional permissions that RP needs to request, for example \&quot; drive.readonly calendar.readonly\&quot; nonce: A random string to ensure the response is issued for this specific request. Prevents replay attacks.  Other custom key-value parameters.  Note: parameters is supported from Chrome 132..</param>
        public ClientProvider(string clientId = default(string), string configUrl = default(string), string domainHint = default(string), List<string> fields = default(List<string>), string loginHint = default(string), string nonce = default(string), Dictionary<string, string> parameters = default(Dictionary<string, string>))
        {
            this.ClientId = clientId;
            this.ConfigUrl = configUrl;
            this.DomainHint = domainHint;
            this.Fields = fields;
            this.LoginHint = loginHint;
            this.Nonce = nonce;
            this.Parameters = parameters;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The RP&#39;s client identifier, issued by the IdP.
        /// </summary>
        /// <value>The RP&#39;s client identifier, issued by the IdP.</value>
        [DataMember(Name = "client_id", EmitDefaultValue = false)]
        public string ClientId { get; set; }

        /// <summary>
        /// A full path of the IdP config file.
        /// </summary>
        /// <value>A full path of the IdP config file.</value>
        [DataMember(Name = "config_url", EmitDefaultValue = false)]
        public string ConfigUrl { get; set; }

        /// <summary>
        /// By specifying one of domain_hints values provided by the accounts endpoints, the FedCM dialog selectively shows the specified account.
        /// </summary>
        /// <value>By specifying one of domain_hints values provided by the accounts endpoints, the FedCM dialog selectively shows the specified account.</value>
        [DataMember(Name = "domain_hint", EmitDefaultValue = false)]
        public string DomainHint { get; set; }

        /// <summary>
        /// Array of strings that specifies the user information (\&quot;name\&quot;, \&quot; email\&quot;, \&quot;picture\&quot;) that RP needs IdP to share with them.  Note: Field API is supported by Chrome 132 and later.
        /// </summary>
        /// <value>Array of strings that specifies the user information (\&quot;name\&quot;, \&quot; email\&quot;, \&quot;picture\&quot;) that RP needs IdP to share with them.  Note: Field API is supported by Chrome 132 and later.</value>
        [DataMember(Name = "fields", EmitDefaultValue = false)]
        public List<string> Fields { get; set; }

        /// <summary>
        /// By specifying one of login_hints values provided by the accounts endpoints, the FedCM dialog selectively shows the specified account.
        /// </summary>
        /// <value>By specifying one of login_hints values provided by the accounts endpoints, the FedCM dialog selectively shows the specified account.</value>
        [DataMember(Name = "login_hint", EmitDefaultValue = false)]
        public string LoginHint { get; set; }

        /// <summary>
        /// A random string to ensure the response is issued for this specific request. Prevents replay attacks.
        /// </summary>
        /// <value>A random string to ensure the response is issued for this specific request. Prevents replay attacks.</value>
        [DataMember(Name = "nonce", EmitDefaultValue = false)]
        public string Nonce { get; set; }

        /// <summary>
        /// Custom object that allows to specify additional key-value parameters: scope: A string value containing additional permissions that RP needs to request, for example \&quot; drive.readonly calendar.readonly\&quot; nonce: A random string to ensure the response is issued for this specific request. Prevents replay attacks.  Other custom key-value parameters.  Note: parameters is supported from Chrome 132.
        /// </summary>
        /// <value>Custom object that allows to specify additional key-value parameters: scope: A string value containing additional permissions that RP needs to request, for example \&quot; drive.readonly calendar.readonly\&quot; nonce: A random string to ensure the response is issued for this specific request. Prevents replay attacks.  Other custom key-value parameters.  Note: parameters is supported from Chrome 132.</value>
        [DataMember(Name = "parameters", EmitDefaultValue = false)]
        public Dictionary<string, string> Parameters { get; set; }

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
            sb.Append("class ClientProvider {\n");
            sb.Append("  ClientId: ").Append(ClientId).Append("\n");
            sb.Append("  ConfigUrl: ").Append(ConfigUrl).Append("\n");
            sb.Append("  DomainHint: ").Append(DomainHint).Append("\n");
            sb.Append("  Fields: ").Append(Fields).Append("\n");
            sb.Append("  LoginHint: ").Append(LoginHint).Append("\n");
            sb.Append("  Nonce: ").Append(Nonce).Append("\n");
            sb.Append("  Parameters: ").Append(Parameters).Append("\n");
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
