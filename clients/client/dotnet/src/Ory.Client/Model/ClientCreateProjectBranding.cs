/*
 * Ory APIs
 *
 * # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | - -- -- -- -- -- -- - | - -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- - | - -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- - | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 
 *
 * The version of the OpenAPI document: v1.15.10
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
    /// Create a Project Branding
    /// </summary>
    [DataContract(Name = "createProjectBranding")]
    public partial class ClientCreateProjectBranding : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientCreateProjectBranding" /> class.
        /// </summary>
        /// <param name="faviconType">faviconType.</param>
        /// <param name="faviconUrl">faviconUrl.</param>
        /// <param name="logoType">logoType.</param>
        /// <param name="logoUrl">logoUrl.</param>
        /// <param name="name">name.</param>
        /// <param name="theme">theme.</param>
        public ClientCreateProjectBranding(string faviconType = default(string), string faviconUrl = default(string), string logoType = default(string), string logoUrl = default(string), string name = default(string), ClientProjectBrandingColors theme = default(ClientProjectBrandingColors))
        {
            this.FaviconType = faviconType;
            this.FaviconUrl = faviconUrl;
            this.LogoType = logoType;
            this.LogoUrl = logoUrl;
            this.Name = name;
            this.Theme = theme;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets FaviconType
        /// </summary>
        [DataMember(Name = "favicon_type", EmitDefaultValue = false)]
        public string FaviconType { get; set; }

        /// <summary>
        /// Gets or Sets FaviconUrl
        /// </summary>
        [DataMember(Name = "favicon_url", EmitDefaultValue = false)]
        public string FaviconUrl { get; set; }

        /// <summary>
        /// Gets or Sets LogoType
        /// </summary>
        [DataMember(Name = "logo_type", EmitDefaultValue = false)]
        public string LogoType { get; set; }

        /// <summary>
        /// Gets or Sets LogoUrl
        /// </summary>
        [DataMember(Name = "logo_url", EmitDefaultValue = false)]
        public string LogoUrl { get; set; }

        /// <summary>
        /// Gets or Sets Name
        /// </summary>
        [DataMember(Name = "name", EmitDefaultValue = false)]
        public string Name { get; set; }

        /// <summary>
        /// Gets or Sets Theme
        /// </summary>
        [DataMember(Name = "theme", EmitDefaultValue = false)]
        public ClientProjectBrandingColors Theme { get; set; }

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
            sb.Append("class ClientCreateProjectBranding {\n");
            sb.Append("  FaviconType: ").Append(FaviconType).Append("\n");
            sb.Append("  FaviconUrl: ").Append(FaviconUrl).Append("\n");
            sb.Append("  LogoType: ").Append(LogoType).Append("\n");
            sb.Append("  LogoUrl: ").Append(LogoUrl).Append("\n");
            sb.Append("  Name: ").Append(Name).Append("\n");
            sb.Append("  Theme: ").Append(Theme).Append("\n");
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
