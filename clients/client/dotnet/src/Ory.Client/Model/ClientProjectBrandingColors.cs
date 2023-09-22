/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.9
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
    /// ClientProjectBrandingColors
    /// </summary>
    [DataContract(Name = "projectBrandingColors")]
    public partial class ClientProjectBrandingColors : IEquatable<ClientProjectBrandingColors>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientProjectBrandingColors" /> class.
        /// </summary>
        /// <param name="accentDefaultColor">AccentDefaultColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="accentDisabledColor">AccentDisabledColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="accentEmphasisColor">AccentEmphasisColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="accentMutedColor">AccentMutedColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="accentSubtleColor">AccentSubtleColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="backgroundCanvasColor">BackgroundCanvasColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="backgroundSubtleColor">BackgroundSubtleColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="backgroundSurfaceColor">BackgroundSurfaceColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="borderDefaultColor">BorderDefaultColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="errorDefaultColor">ErrorDefaultColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="errorEmphasisColor">ErrorEmphasisColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="errorMutedColor">ErrorMutedColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="errorSubtleColor">ErrorSubtleColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="foregroundDefaultColor">ForegroundDefaultColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="foregroundDisabledColor">ForegroundDisabledColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="foregroundMutedColor">ForegroundMutedColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="foregroundOnAccentColor">ForegroundOnAccentColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="foregroundOnDarkColor">ForegroundOnDarkColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="foregroundOnDisabledColor">ForegroundOnDisabledColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="foregroundSubtleColor">ForegroundSubtleColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="inputBackgroundColor">InputBackgroundColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="inputDisabledColor">InputDisabledColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="inputPlaceholderColor">InputPlaceholderColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="inputTextColor">InputTextColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="primaryColor">Primary color is an hsla color value used to derive the other colors from for the Ory Account Experience theme..</param>
        /// <param name="secondaryColor">Secondary color is a hsla color code used to derive the other colors from for the Ory Account Experience theme..</param>
        /// <param name="successEmphasisColor">SuccessEmphasisColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="textDefaultColor">TextDefaultColor is a hex color code used by the Ory Account Experience theme..</param>
        /// <param name="textDisabledColor">TextDisabledColor is a hex color code used by the Ory Account Experience theme..</param>
        public ClientProjectBrandingColors(string accentDefaultColor = default(string), string accentDisabledColor = default(string), string accentEmphasisColor = default(string), string accentMutedColor = default(string), string accentSubtleColor = default(string), string backgroundCanvasColor = default(string), string backgroundSubtleColor = default(string), string backgroundSurfaceColor = default(string), string borderDefaultColor = default(string), string errorDefaultColor = default(string), string errorEmphasisColor = default(string), string errorMutedColor = default(string), string errorSubtleColor = default(string), string foregroundDefaultColor = default(string), string foregroundDisabledColor = default(string), string foregroundMutedColor = default(string), string foregroundOnAccentColor = default(string), string foregroundOnDarkColor = default(string), string foregroundOnDisabledColor = default(string), string foregroundSubtleColor = default(string), string inputBackgroundColor = default(string), string inputDisabledColor = default(string), string inputPlaceholderColor = default(string), string inputTextColor = default(string), string primaryColor = default(string), string secondaryColor = default(string), string successEmphasisColor = default(string), string textDefaultColor = default(string), string textDisabledColor = default(string))
        {
            this.AccentDefaultColor = accentDefaultColor;
            this.AccentDisabledColor = accentDisabledColor;
            this.AccentEmphasisColor = accentEmphasisColor;
            this.AccentMutedColor = accentMutedColor;
            this.AccentSubtleColor = accentSubtleColor;
            this.BackgroundCanvasColor = backgroundCanvasColor;
            this.BackgroundSubtleColor = backgroundSubtleColor;
            this.BackgroundSurfaceColor = backgroundSurfaceColor;
            this.BorderDefaultColor = borderDefaultColor;
            this.ErrorDefaultColor = errorDefaultColor;
            this.ErrorEmphasisColor = errorEmphasisColor;
            this.ErrorMutedColor = errorMutedColor;
            this.ErrorSubtleColor = errorSubtleColor;
            this.ForegroundDefaultColor = foregroundDefaultColor;
            this.ForegroundDisabledColor = foregroundDisabledColor;
            this.ForegroundMutedColor = foregroundMutedColor;
            this.ForegroundOnAccentColor = foregroundOnAccentColor;
            this.ForegroundOnDarkColor = foregroundOnDarkColor;
            this.ForegroundOnDisabledColor = foregroundOnDisabledColor;
            this.ForegroundSubtleColor = foregroundSubtleColor;
            this.InputBackgroundColor = inputBackgroundColor;
            this.InputDisabledColor = inputDisabledColor;
            this.InputPlaceholderColor = inputPlaceholderColor;
            this.InputTextColor = inputTextColor;
            this.PrimaryColor = primaryColor;
            this.SecondaryColor = secondaryColor;
            this.SuccessEmphasisColor = successEmphasisColor;
            this.TextDefaultColor = textDefaultColor;
            this.TextDisabledColor = textDisabledColor;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// AccentDefaultColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>AccentDefaultColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "accent_default_color", EmitDefaultValue = false)]
        public string AccentDefaultColor { get; set; }

        /// <summary>
        /// AccentDisabledColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>AccentDisabledColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "accent_disabled_color", EmitDefaultValue = false)]
        public string AccentDisabledColor { get; set; }

        /// <summary>
        /// AccentEmphasisColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>AccentEmphasisColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "accent_emphasis_color", EmitDefaultValue = false)]
        public string AccentEmphasisColor { get; set; }

        /// <summary>
        /// AccentMutedColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>AccentMutedColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "accent_muted_color", EmitDefaultValue = false)]
        public string AccentMutedColor { get; set; }

        /// <summary>
        /// AccentSubtleColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>AccentSubtleColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "accent_subtle_color", EmitDefaultValue = false)]
        public string AccentSubtleColor { get; set; }

        /// <summary>
        /// BackgroundCanvasColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>BackgroundCanvasColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "background_canvas_color", EmitDefaultValue = false)]
        public string BackgroundCanvasColor { get; set; }

        /// <summary>
        /// BackgroundSubtleColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>BackgroundSubtleColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "background_subtle_color", EmitDefaultValue = false)]
        public string BackgroundSubtleColor { get; set; }

        /// <summary>
        /// BackgroundSurfaceColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>BackgroundSurfaceColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "background_surface_color", EmitDefaultValue = false)]
        public string BackgroundSurfaceColor { get; set; }

        /// <summary>
        /// BorderDefaultColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>BorderDefaultColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "border_default_color", EmitDefaultValue = false)]
        public string BorderDefaultColor { get; set; }

        /// <summary>
        /// ErrorDefaultColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ErrorDefaultColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "error_default_color", EmitDefaultValue = false)]
        public string ErrorDefaultColor { get; set; }

        /// <summary>
        /// ErrorEmphasisColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ErrorEmphasisColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "error_emphasis_color", EmitDefaultValue = false)]
        public string ErrorEmphasisColor { get; set; }

        /// <summary>
        /// ErrorMutedColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ErrorMutedColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "error_muted_color", EmitDefaultValue = false)]
        public string ErrorMutedColor { get; set; }

        /// <summary>
        /// ErrorSubtleColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ErrorSubtleColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "error_subtle_color", EmitDefaultValue = false)]
        public string ErrorSubtleColor { get; set; }

        /// <summary>
        /// ForegroundDefaultColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ForegroundDefaultColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "foreground_default_color", EmitDefaultValue = false)]
        public string ForegroundDefaultColor { get; set; }

        /// <summary>
        /// ForegroundDisabledColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ForegroundDisabledColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "foreground_disabled_color", EmitDefaultValue = false)]
        public string ForegroundDisabledColor { get; set; }

        /// <summary>
        /// ForegroundMutedColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ForegroundMutedColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "foreground_muted_color", EmitDefaultValue = false)]
        public string ForegroundMutedColor { get; set; }

        /// <summary>
        /// ForegroundOnAccentColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ForegroundOnAccentColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "foreground_on_accent_color", EmitDefaultValue = false)]
        public string ForegroundOnAccentColor { get; set; }

        /// <summary>
        /// ForegroundOnDarkColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ForegroundOnDarkColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "foreground_on_dark_color", EmitDefaultValue = false)]
        public string ForegroundOnDarkColor { get; set; }

        /// <summary>
        /// ForegroundOnDisabledColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ForegroundOnDisabledColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "foreground_on_disabled_color", EmitDefaultValue = false)]
        public string ForegroundOnDisabledColor { get; set; }

        /// <summary>
        /// ForegroundSubtleColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>ForegroundSubtleColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "foreground_subtle_color", EmitDefaultValue = false)]
        public string ForegroundSubtleColor { get; set; }

        /// <summary>
        /// InputBackgroundColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>InputBackgroundColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "input_background_color", EmitDefaultValue = false)]
        public string InputBackgroundColor { get; set; }

        /// <summary>
        /// InputDisabledColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>InputDisabledColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "input_disabled_color", EmitDefaultValue = false)]
        public string InputDisabledColor { get; set; }

        /// <summary>
        /// InputPlaceholderColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>InputPlaceholderColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "input_placeholder_color", EmitDefaultValue = false)]
        public string InputPlaceholderColor { get; set; }

        /// <summary>
        /// InputTextColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>InputTextColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "input_text_color", EmitDefaultValue = false)]
        public string InputTextColor { get; set; }

        /// <summary>
        /// Primary color is an hsla color value used to derive the other colors from for the Ory Account Experience theme.
        /// </summary>
        /// <value>Primary color is an hsla color value used to derive the other colors from for the Ory Account Experience theme.</value>
        [DataMember(Name = "primary_color", EmitDefaultValue = false)]
        public string PrimaryColor { get; set; }

        /// <summary>
        /// Secondary color is a hsla color code used to derive the other colors from for the Ory Account Experience theme.
        /// </summary>
        /// <value>Secondary color is a hsla color code used to derive the other colors from for the Ory Account Experience theme.</value>
        [DataMember(Name = "secondary_color", EmitDefaultValue = false)]
        public string SecondaryColor { get; set; }

        /// <summary>
        /// SuccessEmphasisColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>SuccessEmphasisColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "success_emphasis_color", EmitDefaultValue = false)]
        public string SuccessEmphasisColor { get; set; }

        /// <summary>
        /// TextDefaultColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>TextDefaultColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "text_default_color", EmitDefaultValue = false)]
        public string TextDefaultColor { get; set; }

        /// <summary>
        /// TextDisabledColor is a hex color code used by the Ory Account Experience theme.
        /// </summary>
        /// <value>TextDisabledColor is a hex color code used by the Ory Account Experience theme.</value>
        [DataMember(Name = "text_disabled_color", EmitDefaultValue = false)]
        public string TextDisabledColor { get; set; }

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
            sb.Append("class ClientProjectBrandingColors {\n");
            sb.Append("  AccentDefaultColor: ").Append(AccentDefaultColor).Append("\n");
            sb.Append("  AccentDisabledColor: ").Append(AccentDisabledColor).Append("\n");
            sb.Append("  AccentEmphasisColor: ").Append(AccentEmphasisColor).Append("\n");
            sb.Append("  AccentMutedColor: ").Append(AccentMutedColor).Append("\n");
            sb.Append("  AccentSubtleColor: ").Append(AccentSubtleColor).Append("\n");
            sb.Append("  BackgroundCanvasColor: ").Append(BackgroundCanvasColor).Append("\n");
            sb.Append("  BackgroundSubtleColor: ").Append(BackgroundSubtleColor).Append("\n");
            sb.Append("  BackgroundSurfaceColor: ").Append(BackgroundSurfaceColor).Append("\n");
            sb.Append("  BorderDefaultColor: ").Append(BorderDefaultColor).Append("\n");
            sb.Append("  ErrorDefaultColor: ").Append(ErrorDefaultColor).Append("\n");
            sb.Append("  ErrorEmphasisColor: ").Append(ErrorEmphasisColor).Append("\n");
            sb.Append("  ErrorMutedColor: ").Append(ErrorMutedColor).Append("\n");
            sb.Append("  ErrorSubtleColor: ").Append(ErrorSubtleColor).Append("\n");
            sb.Append("  ForegroundDefaultColor: ").Append(ForegroundDefaultColor).Append("\n");
            sb.Append("  ForegroundDisabledColor: ").Append(ForegroundDisabledColor).Append("\n");
            sb.Append("  ForegroundMutedColor: ").Append(ForegroundMutedColor).Append("\n");
            sb.Append("  ForegroundOnAccentColor: ").Append(ForegroundOnAccentColor).Append("\n");
            sb.Append("  ForegroundOnDarkColor: ").Append(ForegroundOnDarkColor).Append("\n");
            sb.Append("  ForegroundOnDisabledColor: ").Append(ForegroundOnDisabledColor).Append("\n");
            sb.Append("  ForegroundSubtleColor: ").Append(ForegroundSubtleColor).Append("\n");
            sb.Append("  InputBackgroundColor: ").Append(InputBackgroundColor).Append("\n");
            sb.Append("  InputDisabledColor: ").Append(InputDisabledColor).Append("\n");
            sb.Append("  InputPlaceholderColor: ").Append(InputPlaceholderColor).Append("\n");
            sb.Append("  InputTextColor: ").Append(InputTextColor).Append("\n");
            sb.Append("  PrimaryColor: ").Append(PrimaryColor).Append("\n");
            sb.Append("  SecondaryColor: ").Append(SecondaryColor).Append("\n");
            sb.Append("  SuccessEmphasisColor: ").Append(SuccessEmphasisColor).Append("\n");
            sb.Append("  TextDefaultColor: ").Append(TextDefaultColor).Append("\n");
            sb.Append("  TextDisabledColor: ").Append(TextDisabledColor).Append("\n");
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
            return this.Equals(input as ClientProjectBrandingColors);
        }

        /// <summary>
        /// Returns true if ClientProjectBrandingColors instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientProjectBrandingColors to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientProjectBrandingColors input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.AccentDefaultColor == input.AccentDefaultColor ||
                    (this.AccentDefaultColor != null &&
                    this.AccentDefaultColor.Equals(input.AccentDefaultColor))
                ) && 
                (
                    this.AccentDisabledColor == input.AccentDisabledColor ||
                    (this.AccentDisabledColor != null &&
                    this.AccentDisabledColor.Equals(input.AccentDisabledColor))
                ) && 
                (
                    this.AccentEmphasisColor == input.AccentEmphasisColor ||
                    (this.AccentEmphasisColor != null &&
                    this.AccentEmphasisColor.Equals(input.AccentEmphasisColor))
                ) && 
                (
                    this.AccentMutedColor == input.AccentMutedColor ||
                    (this.AccentMutedColor != null &&
                    this.AccentMutedColor.Equals(input.AccentMutedColor))
                ) && 
                (
                    this.AccentSubtleColor == input.AccentSubtleColor ||
                    (this.AccentSubtleColor != null &&
                    this.AccentSubtleColor.Equals(input.AccentSubtleColor))
                ) && 
                (
                    this.BackgroundCanvasColor == input.BackgroundCanvasColor ||
                    (this.BackgroundCanvasColor != null &&
                    this.BackgroundCanvasColor.Equals(input.BackgroundCanvasColor))
                ) && 
                (
                    this.BackgroundSubtleColor == input.BackgroundSubtleColor ||
                    (this.BackgroundSubtleColor != null &&
                    this.BackgroundSubtleColor.Equals(input.BackgroundSubtleColor))
                ) && 
                (
                    this.BackgroundSurfaceColor == input.BackgroundSurfaceColor ||
                    (this.BackgroundSurfaceColor != null &&
                    this.BackgroundSurfaceColor.Equals(input.BackgroundSurfaceColor))
                ) && 
                (
                    this.BorderDefaultColor == input.BorderDefaultColor ||
                    (this.BorderDefaultColor != null &&
                    this.BorderDefaultColor.Equals(input.BorderDefaultColor))
                ) && 
                (
                    this.ErrorDefaultColor == input.ErrorDefaultColor ||
                    (this.ErrorDefaultColor != null &&
                    this.ErrorDefaultColor.Equals(input.ErrorDefaultColor))
                ) && 
                (
                    this.ErrorEmphasisColor == input.ErrorEmphasisColor ||
                    (this.ErrorEmphasisColor != null &&
                    this.ErrorEmphasisColor.Equals(input.ErrorEmphasisColor))
                ) && 
                (
                    this.ErrorMutedColor == input.ErrorMutedColor ||
                    (this.ErrorMutedColor != null &&
                    this.ErrorMutedColor.Equals(input.ErrorMutedColor))
                ) && 
                (
                    this.ErrorSubtleColor == input.ErrorSubtleColor ||
                    (this.ErrorSubtleColor != null &&
                    this.ErrorSubtleColor.Equals(input.ErrorSubtleColor))
                ) && 
                (
                    this.ForegroundDefaultColor == input.ForegroundDefaultColor ||
                    (this.ForegroundDefaultColor != null &&
                    this.ForegroundDefaultColor.Equals(input.ForegroundDefaultColor))
                ) && 
                (
                    this.ForegroundDisabledColor == input.ForegroundDisabledColor ||
                    (this.ForegroundDisabledColor != null &&
                    this.ForegroundDisabledColor.Equals(input.ForegroundDisabledColor))
                ) && 
                (
                    this.ForegroundMutedColor == input.ForegroundMutedColor ||
                    (this.ForegroundMutedColor != null &&
                    this.ForegroundMutedColor.Equals(input.ForegroundMutedColor))
                ) && 
                (
                    this.ForegroundOnAccentColor == input.ForegroundOnAccentColor ||
                    (this.ForegroundOnAccentColor != null &&
                    this.ForegroundOnAccentColor.Equals(input.ForegroundOnAccentColor))
                ) && 
                (
                    this.ForegroundOnDarkColor == input.ForegroundOnDarkColor ||
                    (this.ForegroundOnDarkColor != null &&
                    this.ForegroundOnDarkColor.Equals(input.ForegroundOnDarkColor))
                ) && 
                (
                    this.ForegroundOnDisabledColor == input.ForegroundOnDisabledColor ||
                    (this.ForegroundOnDisabledColor != null &&
                    this.ForegroundOnDisabledColor.Equals(input.ForegroundOnDisabledColor))
                ) && 
                (
                    this.ForegroundSubtleColor == input.ForegroundSubtleColor ||
                    (this.ForegroundSubtleColor != null &&
                    this.ForegroundSubtleColor.Equals(input.ForegroundSubtleColor))
                ) && 
                (
                    this.InputBackgroundColor == input.InputBackgroundColor ||
                    (this.InputBackgroundColor != null &&
                    this.InputBackgroundColor.Equals(input.InputBackgroundColor))
                ) && 
                (
                    this.InputDisabledColor == input.InputDisabledColor ||
                    (this.InputDisabledColor != null &&
                    this.InputDisabledColor.Equals(input.InputDisabledColor))
                ) && 
                (
                    this.InputPlaceholderColor == input.InputPlaceholderColor ||
                    (this.InputPlaceholderColor != null &&
                    this.InputPlaceholderColor.Equals(input.InputPlaceholderColor))
                ) && 
                (
                    this.InputTextColor == input.InputTextColor ||
                    (this.InputTextColor != null &&
                    this.InputTextColor.Equals(input.InputTextColor))
                ) && 
                (
                    this.PrimaryColor == input.PrimaryColor ||
                    (this.PrimaryColor != null &&
                    this.PrimaryColor.Equals(input.PrimaryColor))
                ) && 
                (
                    this.SecondaryColor == input.SecondaryColor ||
                    (this.SecondaryColor != null &&
                    this.SecondaryColor.Equals(input.SecondaryColor))
                ) && 
                (
                    this.SuccessEmphasisColor == input.SuccessEmphasisColor ||
                    (this.SuccessEmphasisColor != null &&
                    this.SuccessEmphasisColor.Equals(input.SuccessEmphasisColor))
                ) && 
                (
                    this.TextDefaultColor == input.TextDefaultColor ||
                    (this.TextDefaultColor != null &&
                    this.TextDefaultColor.Equals(input.TextDefaultColor))
                ) && 
                (
                    this.TextDisabledColor == input.TextDisabledColor ||
                    (this.TextDisabledColor != null &&
                    this.TextDisabledColor.Equals(input.TextDisabledColor))
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
                if (this.AccentDefaultColor != null)
                {
                    hashCode = (hashCode * 59) + this.AccentDefaultColor.GetHashCode();
                }
                if (this.AccentDisabledColor != null)
                {
                    hashCode = (hashCode * 59) + this.AccentDisabledColor.GetHashCode();
                }
                if (this.AccentEmphasisColor != null)
                {
                    hashCode = (hashCode * 59) + this.AccentEmphasisColor.GetHashCode();
                }
                if (this.AccentMutedColor != null)
                {
                    hashCode = (hashCode * 59) + this.AccentMutedColor.GetHashCode();
                }
                if (this.AccentSubtleColor != null)
                {
                    hashCode = (hashCode * 59) + this.AccentSubtleColor.GetHashCode();
                }
                if (this.BackgroundCanvasColor != null)
                {
                    hashCode = (hashCode * 59) + this.BackgroundCanvasColor.GetHashCode();
                }
                if (this.BackgroundSubtleColor != null)
                {
                    hashCode = (hashCode * 59) + this.BackgroundSubtleColor.GetHashCode();
                }
                if (this.BackgroundSurfaceColor != null)
                {
                    hashCode = (hashCode * 59) + this.BackgroundSurfaceColor.GetHashCode();
                }
                if (this.BorderDefaultColor != null)
                {
                    hashCode = (hashCode * 59) + this.BorderDefaultColor.GetHashCode();
                }
                if (this.ErrorDefaultColor != null)
                {
                    hashCode = (hashCode * 59) + this.ErrorDefaultColor.GetHashCode();
                }
                if (this.ErrorEmphasisColor != null)
                {
                    hashCode = (hashCode * 59) + this.ErrorEmphasisColor.GetHashCode();
                }
                if (this.ErrorMutedColor != null)
                {
                    hashCode = (hashCode * 59) + this.ErrorMutedColor.GetHashCode();
                }
                if (this.ErrorSubtleColor != null)
                {
                    hashCode = (hashCode * 59) + this.ErrorSubtleColor.GetHashCode();
                }
                if (this.ForegroundDefaultColor != null)
                {
                    hashCode = (hashCode * 59) + this.ForegroundDefaultColor.GetHashCode();
                }
                if (this.ForegroundDisabledColor != null)
                {
                    hashCode = (hashCode * 59) + this.ForegroundDisabledColor.GetHashCode();
                }
                if (this.ForegroundMutedColor != null)
                {
                    hashCode = (hashCode * 59) + this.ForegroundMutedColor.GetHashCode();
                }
                if (this.ForegroundOnAccentColor != null)
                {
                    hashCode = (hashCode * 59) + this.ForegroundOnAccentColor.GetHashCode();
                }
                if (this.ForegroundOnDarkColor != null)
                {
                    hashCode = (hashCode * 59) + this.ForegroundOnDarkColor.GetHashCode();
                }
                if (this.ForegroundOnDisabledColor != null)
                {
                    hashCode = (hashCode * 59) + this.ForegroundOnDisabledColor.GetHashCode();
                }
                if (this.ForegroundSubtleColor != null)
                {
                    hashCode = (hashCode * 59) + this.ForegroundSubtleColor.GetHashCode();
                }
                if (this.InputBackgroundColor != null)
                {
                    hashCode = (hashCode * 59) + this.InputBackgroundColor.GetHashCode();
                }
                if (this.InputDisabledColor != null)
                {
                    hashCode = (hashCode * 59) + this.InputDisabledColor.GetHashCode();
                }
                if (this.InputPlaceholderColor != null)
                {
                    hashCode = (hashCode * 59) + this.InputPlaceholderColor.GetHashCode();
                }
                if (this.InputTextColor != null)
                {
                    hashCode = (hashCode * 59) + this.InputTextColor.GetHashCode();
                }
                if (this.PrimaryColor != null)
                {
                    hashCode = (hashCode * 59) + this.PrimaryColor.GetHashCode();
                }
                if (this.SecondaryColor != null)
                {
                    hashCode = (hashCode * 59) + this.SecondaryColor.GetHashCode();
                }
                if (this.SuccessEmphasisColor != null)
                {
                    hashCode = (hashCode * 59) + this.SuccessEmphasisColor.GetHashCode();
                }
                if (this.TextDefaultColor != null)
                {
                    hashCode = (hashCode * 59) + this.TextDefaultColor.GetHashCode();
                }
                if (this.TextDisabledColor != null)
                {
                    hashCode = (hashCode * 59) + this.TextDisabledColor.GetHashCode();
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
