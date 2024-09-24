/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.15.4
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

import java.util.Objects;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.util.Arrays;
import sh.ory.model.ProjectBrandingColors;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import com.google.gson.TypeAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

import sh.ory.JSON;

/**
 * SetProjectBrandingThemeBody
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-09-24T13:15:05.321914806Z[Etc/UTC]", comments = "Generator version: 7.7.0")
public class SetProjectBrandingThemeBody {
  public static final String SERIALIZED_NAME_FAVICON_TYPE = "favicon_type";
  @SerializedName(SERIALIZED_NAME_FAVICON_TYPE)
  private String faviconType;

  public static final String SERIALIZED_NAME_FAVICON_URL = "favicon_url";
  @SerializedName(SERIALIZED_NAME_FAVICON_URL)
  private String faviconUrl;

  public static final String SERIALIZED_NAME_LOGO_TYPE = "logo_type";
  @SerializedName(SERIALIZED_NAME_LOGO_TYPE)
  private String logoType;

  public static final String SERIALIZED_NAME_LOGO_URL = "logo_url";
  @SerializedName(SERIALIZED_NAME_LOGO_URL)
  private String logoUrl;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_THEME = "theme";
  @SerializedName(SERIALIZED_NAME_THEME)
  private ProjectBrandingColors theme;

  public SetProjectBrandingThemeBody() {
  }

  public SetProjectBrandingThemeBody faviconType(String faviconType) {
    this.faviconType = faviconType;
    return this;
  }

  /**
   * Favicon Type
   * @return faviconType
   */
  @javax.annotation.Nullable
  public String getFaviconType() {
    return faviconType;
  }

  public void setFaviconType(String faviconType) {
    this.faviconType = faviconType;
  }


  public SetProjectBrandingThemeBody faviconUrl(String faviconUrl) {
    this.faviconUrl = faviconUrl;
    return this;
  }

  /**
   * Favicon URL
   * @return faviconUrl
   */
  @javax.annotation.Nullable
  public String getFaviconUrl() {
    return faviconUrl;
  }

  public void setFaviconUrl(String faviconUrl) {
    this.faviconUrl = faviconUrl;
  }


  public SetProjectBrandingThemeBody logoType(String logoType) {
    this.logoType = logoType;
    return this;
  }

  /**
   * Logo type
   * @return logoType
   */
  @javax.annotation.Nullable
  public String getLogoType() {
    return logoType;
  }

  public void setLogoType(String logoType) {
    this.logoType = logoType;
  }


  public SetProjectBrandingThemeBody logoUrl(String logoUrl) {
    this.logoUrl = logoUrl;
    return this;
  }

  /**
   * Logo URL
   * @return logoUrl
   */
  @javax.annotation.Nullable
  public String getLogoUrl() {
    return logoUrl;
  }

  public void setLogoUrl(String logoUrl) {
    this.logoUrl = logoUrl;
  }


  public SetProjectBrandingThemeBody name(String name) {
    this.name = name;
    return this;
  }

  /**
   * Branding name
   * @return name
   */
  @javax.annotation.Nullable
  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }


  public SetProjectBrandingThemeBody theme(ProjectBrandingColors theme) {
    this.theme = theme;
    return this;
  }

  /**
   * Get theme
   * @return theme
   */
  @javax.annotation.Nullable
  public ProjectBrandingColors getTheme() {
    return theme;
  }

  public void setTheme(ProjectBrandingColors theme) {
    this.theme = theme;
  }

  /**
   * A container for additional, undeclared properties.
   * This is a holder for any undeclared properties as specified with
   * the 'additionalProperties' keyword in the OAS document.
   */
  private Map<String, Object> additionalProperties;

  /**
   * Set the additional (undeclared) property with the specified name and value.
   * If the property does not already exist, create it otherwise replace it.
   *
   * @param key name of the property
   * @param value value of the property
   * @return the SetProjectBrandingThemeBody instance itself
   */
  public SetProjectBrandingThemeBody putAdditionalProperty(String key, Object value) {
    if (this.additionalProperties == null) {
        this.additionalProperties = new HashMap<String, Object>();
    }
    this.additionalProperties.put(key, value);
    return this;
  }

  /**
   * Return the additional (undeclared) property.
   *
   * @return a map of objects
   */
  public Map<String, Object> getAdditionalProperties() {
    return additionalProperties;
  }

  /**
   * Return the additional (undeclared) property with the specified name.
   *
   * @param key name of the property
   * @return an object
   */
  public Object getAdditionalProperty(String key) {
    if (this.additionalProperties == null) {
        return null;
    }
    return this.additionalProperties.get(key);
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SetProjectBrandingThemeBody setProjectBrandingThemeBody = (SetProjectBrandingThemeBody) o;
    return Objects.equals(this.faviconType, setProjectBrandingThemeBody.faviconType) &&
        Objects.equals(this.faviconUrl, setProjectBrandingThemeBody.faviconUrl) &&
        Objects.equals(this.logoType, setProjectBrandingThemeBody.logoType) &&
        Objects.equals(this.logoUrl, setProjectBrandingThemeBody.logoUrl) &&
        Objects.equals(this.name, setProjectBrandingThemeBody.name) &&
        Objects.equals(this.theme, setProjectBrandingThemeBody.theme)&&
        Objects.equals(this.additionalProperties, setProjectBrandingThemeBody.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(faviconType, faviconUrl, logoType, logoUrl, name, theme, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SetProjectBrandingThemeBody {\n");
    sb.append("    faviconType: ").append(toIndentedString(faviconType)).append("\n");
    sb.append("    faviconUrl: ").append(toIndentedString(faviconUrl)).append("\n");
    sb.append("    logoType: ").append(toIndentedString(logoType)).append("\n");
    sb.append("    logoUrl: ").append(toIndentedString(logoUrl)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    theme: ").append(toIndentedString(theme)).append("\n");
    sb.append("    additionalProperties: ").append(toIndentedString(additionalProperties)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("favicon_type");
    openapiFields.add("favicon_url");
    openapiFields.add("logo_type");
    openapiFields.add("logo_url");
    openapiFields.add("name");
    openapiFields.add("theme");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

  /**
   * Validates the JSON Element and throws an exception if issues found
   *
   * @param jsonElement JSON Element
   * @throws IOException if the JSON Element is invalid with respect to SetProjectBrandingThemeBody
   */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!SetProjectBrandingThemeBody.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in SetProjectBrandingThemeBody is not found in the empty JSON string", SetProjectBrandingThemeBody.openapiRequiredFields.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("favicon_type") != null && !jsonObj.get("favicon_type").isJsonNull()) && !jsonObj.get("favicon_type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `favicon_type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("favicon_type").toString()));
      }
      if ((jsonObj.get("favicon_url") != null && !jsonObj.get("favicon_url").isJsonNull()) && !jsonObj.get("favicon_url").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `favicon_url` to be a primitive type in the JSON string but got `%s`", jsonObj.get("favicon_url").toString()));
      }
      if ((jsonObj.get("logo_type") != null && !jsonObj.get("logo_type").isJsonNull()) && !jsonObj.get("logo_type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `logo_type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("logo_type").toString()));
      }
      if ((jsonObj.get("logo_url") != null && !jsonObj.get("logo_url").isJsonNull()) && !jsonObj.get("logo_url").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `logo_url` to be a primitive type in the JSON string but got `%s`", jsonObj.get("logo_url").toString()));
      }
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      // validate the optional field `theme`
      if (jsonObj.get("theme") != null && !jsonObj.get("theme").isJsonNull()) {
        ProjectBrandingColors.validateJsonElement(jsonObj.get("theme"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!SetProjectBrandingThemeBody.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'SetProjectBrandingThemeBody' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<SetProjectBrandingThemeBody> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(SetProjectBrandingThemeBody.class));

       return (TypeAdapter<T>) new TypeAdapter<SetProjectBrandingThemeBody>() {
           @Override
           public void write(JsonWriter out, SetProjectBrandingThemeBody value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             obj.remove("additionalProperties");
             // serialize additional properties
             if (value.getAdditionalProperties() != null) {
               for (Map.Entry<String, Object> entry : value.getAdditionalProperties().entrySet()) {
                 if (entry.getValue() instanceof String)
                   obj.addProperty(entry.getKey(), (String) entry.getValue());
                 else if (entry.getValue() instanceof Number)
                   obj.addProperty(entry.getKey(), (Number) entry.getValue());
                 else if (entry.getValue() instanceof Boolean)
                   obj.addProperty(entry.getKey(), (Boolean) entry.getValue());
                 else if (entry.getValue() instanceof Character)
                   obj.addProperty(entry.getKey(), (Character) entry.getValue());
                 else {
                   JsonElement jsonElement = gson.toJsonTree(entry.getValue());
                   if (jsonElement.isJsonArray()) {
                     obj.add(entry.getKey(), jsonElement.getAsJsonArray());
                   } else {
                     obj.add(entry.getKey(), jsonElement.getAsJsonObject());
                   }
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public SetProjectBrandingThemeBody read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             JsonObject jsonObj = jsonElement.getAsJsonObject();
             // store additional fields in the deserialized instance
             SetProjectBrandingThemeBody instance = thisAdapter.fromJsonTree(jsonObj);
             for (Map.Entry<String, JsonElement> entry : jsonObj.entrySet()) {
               if (!openapiFields.contains(entry.getKey())) {
                 if (entry.getValue().isJsonPrimitive()) { // primitive type
                   if (entry.getValue().getAsJsonPrimitive().isString())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsString());
                   else if (entry.getValue().getAsJsonPrimitive().isNumber())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsNumber());
                   else if (entry.getValue().getAsJsonPrimitive().isBoolean())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsBoolean());
                   else
                     throw new IllegalArgumentException(String.format("The field `%s` has unknown primitive type. Value: %s", entry.getKey(), entry.getValue().toString()));
                 } else if (entry.getValue().isJsonArray()) {
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), List.class));
                 } else { // JSON object
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), HashMap.class));
                 }
               }
             }
             return instance;
           }

       }.nullSafe();
    }
  }

  /**
   * Create an instance of SetProjectBrandingThemeBody given an JSON string
   *
   * @param jsonString JSON string
   * @return An instance of SetProjectBrandingThemeBody
   * @throws IOException if the JSON string is invalid with respect to SetProjectBrandingThemeBody
   */
  public static SetProjectBrandingThemeBody fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, SetProjectBrandingThemeBody.class);
  }

  /**
   * Convert an instance of SetProjectBrandingThemeBody to an JSON string
   *
   * @return JSON string
   */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

