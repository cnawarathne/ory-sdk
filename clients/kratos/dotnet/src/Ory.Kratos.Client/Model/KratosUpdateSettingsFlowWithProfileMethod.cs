// <auto-generated>
/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.4.0-alpha.0
 * Contact: office@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */

#nullable enable

using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Text;
using System.Text.RegularExpressions;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.ComponentModel.DataAnnotations;
using Ory.Kratos.Client.Client;

namespace Ory.Kratos.Client.Model
{
    /// <summary>
    /// Update Settings Flow with Profile Method
    /// </summary>
    public partial class KratosUpdateSettingsFlowWithProfileMethod : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="KratosUpdateSettingsFlowWithProfileMethod" /> class.
        /// </summary>
        /// <param name="method">Method  Should be set to profile when trying to update a profile.</param>
        /// <param name="traits">Traits  The identity&#39;s traits.</param>
        /// <param name="csrfToken">The Anti-CSRF Token  This token is only required when performing browser flows.</param>
        /// <param name="transientPayload">Transient data to pass along to any webhooks</param>
        [JsonConstructor]
        public KratosUpdateSettingsFlowWithProfileMethod(string method, Object traits, Option<string?> csrfToken = default, Option<Object?> transientPayload = default)
        {
            Method = method;
            Traits = traits;
            CsrfTokenOption = csrfToken;
            TransientPayloadOption = transientPayload;
            OnCreated();
        }

        partial void OnCreated();

        /// <summary>
        /// Method  Should be set to profile when trying to update a profile.
        /// </summary>
        /// <value>Method  Should be set to profile when trying to update a profile.</value>
        [JsonPropertyName("method")]
        public string Method { get; set; }

        /// <summary>
        /// Traits  The identity&#39;s traits.
        /// </summary>
        /// <value>Traits  The identity&#39;s traits.</value>
        [JsonPropertyName("traits")]
        public Object Traits { get; set; }

        /// <summary>
        /// Used to track the state of CsrfToken
        /// </summary>
        [JsonIgnore]
        [global::System.ComponentModel.EditorBrowsable(global::System.ComponentModel.EditorBrowsableState.Never)]
        public Option<string?> CsrfTokenOption { get; private set; }

        /// <summary>
        /// The Anti-CSRF Token  This token is only required when performing browser flows.
        /// </summary>
        /// <value>The Anti-CSRF Token  This token is only required when performing browser flows.</value>
        [JsonPropertyName("csrf_token")]
        public string? CsrfToken { get { return this.CsrfTokenOption; } set { this.CsrfTokenOption = new(value); } }

        /// <summary>
        /// Used to track the state of TransientPayload
        /// </summary>
        [JsonIgnore]
        [global::System.ComponentModel.EditorBrowsable(global::System.ComponentModel.EditorBrowsableState.Never)]
        public Option<Object?> TransientPayloadOption { get; private set; }

        /// <summary>
        /// Transient data to pass along to any webhooks
        /// </summary>
        /// <value>Transient data to pass along to any webhooks</value>
        [JsonPropertyName("transient_payload")]
        public Object? TransientPayload { get { return this.TransientPayloadOption; } set { this.TransientPayloadOption = new(value); } }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public Dictionary<string, JsonElement> AdditionalProperties { get; } = new Dictionary<string, JsonElement>();

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class KratosUpdateSettingsFlowWithProfileMethod {\n");
            sb.Append("  Method: ").Append(Method).Append("\n");
            sb.Append("  Traits: ").Append(Traits).Append("\n");
            sb.Append("  CsrfToken: ").Append(CsrfToken).Append("\n");
            sb.Append("  TransientPayload: ").Append(TransientPayload).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
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

    /// <summary>
    /// A Json converter for type <see cref="KratosUpdateSettingsFlowWithProfileMethod" />
    /// </summary>
    public class KratosUpdateSettingsFlowWithProfileMethodJsonConverter : JsonConverter<KratosUpdateSettingsFlowWithProfileMethod>
    {
        /// <summary>
        /// Deserializes json to <see cref="KratosUpdateSettingsFlowWithProfileMethod" />
        /// </summary>
        /// <param name="utf8JsonReader"></param>
        /// <param name="typeToConvert"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <returns></returns>
        /// <exception cref="JsonException"></exception>
        public override KratosUpdateSettingsFlowWithProfileMethod Read(ref Utf8JsonReader utf8JsonReader, Type typeToConvert, JsonSerializerOptions jsonSerializerOptions)
        {
            int currentDepth = utf8JsonReader.CurrentDepth;

            if (utf8JsonReader.TokenType != JsonTokenType.StartObject && utf8JsonReader.TokenType != JsonTokenType.StartArray)
                throw new JsonException();

            JsonTokenType startingTokenType = utf8JsonReader.TokenType;

            Option<string?> method = default;
            Option<Object?> traits = default;
            Option<string?> csrfToken = default;
            Option<Object?> transientPayload = default;

            while (utf8JsonReader.Read())
            {
                if (startingTokenType == JsonTokenType.StartObject && utf8JsonReader.TokenType == JsonTokenType.EndObject && currentDepth == utf8JsonReader.CurrentDepth)
                    break;

                if (startingTokenType == JsonTokenType.StartArray && utf8JsonReader.TokenType == JsonTokenType.EndArray && currentDepth == utf8JsonReader.CurrentDepth)
                    break;

                if (utf8JsonReader.TokenType == JsonTokenType.PropertyName && currentDepth == utf8JsonReader.CurrentDepth - 1)
                {
                    string? localVarJsonPropertyName = utf8JsonReader.GetString();
                    utf8JsonReader.Read();

                    switch (localVarJsonPropertyName)
                    {
                        case "method":
                            method = new Option<string?>(utf8JsonReader.GetString()!);
                            break;
                        case "traits":
                            if (utf8JsonReader.TokenType != JsonTokenType.Null)
                                traits = new Option<Object?>(JsonSerializer.Deserialize<Object>(ref utf8JsonReader, jsonSerializerOptions)!);
                            break;
                        case "csrf_token":
                            csrfToken = new Option<string?>(utf8JsonReader.GetString()!);
                            break;
                        case "transient_payload":
                            if (utf8JsonReader.TokenType != JsonTokenType.Null)
                                transientPayload = new Option<Object?>(JsonSerializer.Deserialize<Object>(ref utf8JsonReader, jsonSerializerOptions)!);
                            break;
                        default:
                            break;
                    }
                }
            }

            if (!method.IsSet)
                throw new ArgumentException("Property is required for class KratosUpdateSettingsFlowWithProfileMethod.", nameof(method));

            if (!traits.IsSet)
                throw new ArgumentException("Property is required for class KratosUpdateSettingsFlowWithProfileMethod.", nameof(traits));

            if (method.IsSet && method.Value == null)
                throw new ArgumentNullException(nameof(method), "Property is not nullable for class KratosUpdateSettingsFlowWithProfileMethod.");

            if (traits.IsSet && traits.Value == null)
                throw new ArgumentNullException(nameof(traits), "Property is not nullable for class KratosUpdateSettingsFlowWithProfileMethod.");

            if (csrfToken.IsSet && csrfToken.Value == null)
                throw new ArgumentNullException(nameof(csrfToken), "Property is not nullable for class KratosUpdateSettingsFlowWithProfileMethod.");

            if (transientPayload.IsSet && transientPayload.Value == null)
                throw new ArgumentNullException(nameof(transientPayload), "Property is not nullable for class KratosUpdateSettingsFlowWithProfileMethod.");

            return new KratosUpdateSettingsFlowWithProfileMethod(method.Value!, traits.Value!, csrfToken, transientPayload);
        }

        /// <summary>
        /// Serializes a <see cref="KratosUpdateSettingsFlowWithProfileMethod" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="kratosUpdateSettingsFlowWithProfileMethod"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public override void Write(Utf8JsonWriter writer, KratosUpdateSettingsFlowWithProfileMethod kratosUpdateSettingsFlowWithProfileMethod, JsonSerializerOptions jsonSerializerOptions)
        {
            writer.WriteStartObject();

            WriteProperties(writer, kratosUpdateSettingsFlowWithProfileMethod, jsonSerializerOptions);
            writer.WriteEndObject();
        }

        /// <summary>
        /// Serializes the properties of <see cref="KratosUpdateSettingsFlowWithProfileMethod" />
        /// </summary>
        /// <param name="writer"></param>
        /// <param name="kratosUpdateSettingsFlowWithProfileMethod"></param>
        /// <param name="jsonSerializerOptions"></param>
        /// <exception cref="NotImplementedException"></exception>
        public void WriteProperties(Utf8JsonWriter writer, KratosUpdateSettingsFlowWithProfileMethod kratosUpdateSettingsFlowWithProfileMethod, JsonSerializerOptions jsonSerializerOptions)
        {
            if (kratosUpdateSettingsFlowWithProfileMethod.Method == null)
                throw new ArgumentNullException(nameof(kratosUpdateSettingsFlowWithProfileMethod.Method), "Property is required for class KratosUpdateSettingsFlowWithProfileMethod.");

            if (kratosUpdateSettingsFlowWithProfileMethod.Traits == null)
                throw new ArgumentNullException(nameof(kratosUpdateSettingsFlowWithProfileMethod.Traits), "Property is required for class KratosUpdateSettingsFlowWithProfileMethod.");

            if (kratosUpdateSettingsFlowWithProfileMethod.CsrfTokenOption.IsSet && kratosUpdateSettingsFlowWithProfileMethod.CsrfToken == null)
                throw new ArgumentNullException(nameof(kratosUpdateSettingsFlowWithProfileMethod.CsrfToken), "Property is required for class KratosUpdateSettingsFlowWithProfileMethod.");

            if (kratosUpdateSettingsFlowWithProfileMethod.TransientPayloadOption.IsSet && kratosUpdateSettingsFlowWithProfileMethod.TransientPayload == null)
                throw new ArgumentNullException(nameof(kratosUpdateSettingsFlowWithProfileMethod.TransientPayload), "Property is required for class KratosUpdateSettingsFlowWithProfileMethod.");

            writer.WriteString("method", kratosUpdateSettingsFlowWithProfileMethod.Method);

            writer.WritePropertyName("traits");
            JsonSerializer.Serialize(writer, kratosUpdateSettingsFlowWithProfileMethod.Traits, jsonSerializerOptions);
            if (kratosUpdateSettingsFlowWithProfileMethod.CsrfTokenOption.IsSet)
                writer.WriteString("csrf_token", kratosUpdateSettingsFlowWithProfileMethod.CsrfToken);

            if (kratosUpdateSettingsFlowWithProfileMethod.TransientPayloadOption.IsSet)
            {
                writer.WritePropertyName("transient_payload");
                JsonSerializer.Serialize(writer, kratosUpdateSettingsFlowWithProfileMethod.TransientPayload, jsonSerializerOptions);
            }
        }
    }
}
