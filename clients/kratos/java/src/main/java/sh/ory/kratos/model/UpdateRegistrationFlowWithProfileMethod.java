/*
 * Ory Identities API
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.3.8
 * Contact: office@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.kratos.model;

import java.util.Objects;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.util.Arrays;

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

import sh.ory.kratos.JSON;

/**
 * Update Registration Flow with Profile Method
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2025-02-25T15:51:40.899187259Z[Etc/UTC]", comments = "Generator version: 7.7.0")
public class UpdateRegistrationFlowWithProfileMethod {
  public static final String SERIALIZED_NAME_CSRF_TOKEN = "csrf_token";
  @SerializedName(SERIALIZED_NAME_CSRF_TOKEN)
  private String csrfToken;

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private String method;

  /**
   * Screen requests navigation to a previous screen.  This must be set to credential-selection to go back to the credential selection screen. credential-selection RegistrationScreenCredentialSelection nolint:gosec // not a credential previous RegistrationScreenPrevious
   */
  @JsonAdapter(ScreenEnum.Adapter.class)
  public enum ScreenEnum {
    CREDENTIAL_SELECTION("credential-selection"),
    
    PREVIOUS("previous");

    private String value;

    ScreenEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static ScreenEnum fromValue(String value) {
      for (ScreenEnum b : ScreenEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<ScreenEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final ScreenEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public ScreenEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return ScreenEnum.fromValue(value);
      }
    }

    public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      String value = jsonElement.getAsString();
      ScreenEnum.fromValue(value);
    }
  }

  public static final String SERIALIZED_NAME_SCREEN = "screen";
  @SerializedName(SERIALIZED_NAME_SCREEN)
  private ScreenEnum screen;

  public static final String SERIALIZED_NAME_TRAITS = "traits";
  @SerializedName(SERIALIZED_NAME_TRAITS)
  private Object traits;

  public static final String SERIALIZED_NAME_TRANSIENT_PAYLOAD = "transient_payload";
  @SerializedName(SERIALIZED_NAME_TRANSIENT_PAYLOAD)
  private Object transientPayload;

  public UpdateRegistrationFlowWithProfileMethod() {
  }

  public UpdateRegistrationFlowWithProfileMethod csrfToken(String csrfToken) {
    this.csrfToken = csrfToken;
    return this;
  }

  /**
   * The Anti-CSRF Token  This token is only required when performing browser flows.
   * @return csrfToken
   */
  @javax.annotation.Nullable
  public String getCsrfToken() {
    return csrfToken;
  }

  public void setCsrfToken(String csrfToken) {
    this.csrfToken = csrfToken;
  }


  public UpdateRegistrationFlowWithProfileMethod method(String method) {
    this.method = method;
    return this;
  }

  /**
   * Method  Should be set to profile when trying to update a profile.
   * @return method
   */
  @javax.annotation.Nonnull
  public String getMethod() {
    return method;
  }

  public void setMethod(String method) {
    this.method = method;
  }


  public UpdateRegistrationFlowWithProfileMethod screen(ScreenEnum screen) {
    this.screen = screen;
    return this;
  }

  /**
   * Screen requests navigation to a previous screen.  This must be set to credential-selection to go back to the credential selection screen. credential-selection RegistrationScreenCredentialSelection nolint:gosec // not a credential previous RegistrationScreenPrevious
   * @return screen
   */
  @javax.annotation.Nullable
  public ScreenEnum getScreen() {
    return screen;
  }

  public void setScreen(ScreenEnum screen) {
    this.screen = screen;
  }


  public UpdateRegistrationFlowWithProfileMethod traits(Object traits) {
    this.traits = traits;
    return this;
  }

  /**
   * Traits  The identity&#39;s traits.
   * @return traits
   */
  @javax.annotation.Nonnull
  public Object getTraits() {
    return traits;
  }

  public void setTraits(Object traits) {
    this.traits = traits;
  }


  public UpdateRegistrationFlowWithProfileMethod transientPayload(Object transientPayload) {
    this.transientPayload = transientPayload;
    return this;
  }

  /**
   * Transient data to pass along to any webhooks
   * @return transientPayload
   */
  @javax.annotation.Nullable
  public Object getTransientPayload() {
    return transientPayload;
  }

  public void setTransientPayload(Object transientPayload) {
    this.transientPayload = transientPayload;
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
   * @return the UpdateRegistrationFlowWithProfileMethod instance itself
   */
  public UpdateRegistrationFlowWithProfileMethod putAdditionalProperty(String key, Object value) {
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
    UpdateRegistrationFlowWithProfileMethod updateRegistrationFlowWithProfileMethod = (UpdateRegistrationFlowWithProfileMethod) o;
    return Objects.equals(this.csrfToken, updateRegistrationFlowWithProfileMethod.csrfToken) &&
        Objects.equals(this.method, updateRegistrationFlowWithProfileMethod.method) &&
        Objects.equals(this.screen, updateRegistrationFlowWithProfileMethod.screen) &&
        Objects.equals(this.traits, updateRegistrationFlowWithProfileMethod.traits) &&
        Objects.equals(this.transientPayload, updateRegistrationFlowWithProfileMethod.transientPayload)&&
        Objects.equals(this.additionalProperties, updateRegistrationFlowWithProfileMethod.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(csrfToken, method, screen, traits, transientPayload, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdateRegistrationFlowWithProfileMethod {\n");
    sb.append("    csrfToken: ").append(toIndentedString(csrfToken)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    screen: ").append(toIndentedString(screen)).append("\n");
    sb.append("    traits: ").append(toIndentedString(traits)).append("\n");
    sb.append("    transientPayload: ").append(toIndentedString(transientPayload)).append("\n");
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
    openapiFields.add("csrf_token");
    openapiFields.add("method");
    openapiFields.add("screen");
    openapiFields.add("traits");
    openapiFields.add("transient_payload");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("method");
    openapiRequiredFields.add("traits");
  }

  /**
   * Validates the JSON Element and throws an exception if issues found
   *
   * @param jsonElement JSON Element
   * @throws IOException if the JSON Element is invalid with respect to UpdateRegistrationFlowWithProfileMethod
   */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!UpdateRegistrationFlowWithProfileMethod.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in UpdateRegistrationFlowWithProfileMethod is not found in the empty JSON string", UpdateRegistrationFlowWithProfileMethod.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : UpdateRegistrationFlowWithProfileMethod.openapiRequiredFields) {
        if (jsonElement.getAsJsonObject().get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("csrf_token") != null && !jsonObj.get("csrf_token").isJsonNull()) && !jsonObj.get("csrf_token").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `csrf_token` to be a primitive type in the JSON string but got `%s`", jsonObj.get("csrf_token").toString()));
      }
      if (!jsonObj.get("method").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `method` to be a primitive type in the JSON string but got `%s`", jsonObj.get("method").toString()));
      }
      if ((jsonObj.get("screen") != null && !jsonObj.get("screen").isJsonNull()) && !jsonObj.get("screen").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `screen` to be a primitive type in the JSON string but got `%s`", jsonObj.get("screen").toString()));
      }
      // validate the optional field `screen`
      if (jsonObj.get("screen") != null && !jsonObj.get("screen").isJsonNull()) {
        ScreenEnum.validateJsonElement(jsonObj.get("screen"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UpdateRegistrationFlowWithProfileMethod.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UpdateRegistrationFlowWithProfileMethod' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UpdateRegistrationFlowWithProfileMethod> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UpdateRegistrationFlowWithProfileMethod.class));

       return (TypeAdapter<T>) new TypeAdapter<UpdateRegistrationFlowWithProfileMethod>() {
           @Override
           public void write(JsonWriter out, UpdateRegistrationFlowWithProfileMethod value) throws IOException {
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
           public UpdateRegistrationFlowWithProfileMethod read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             JsonObject jsonObj = jsonElement.getAsJsonObject();
             // store additional fields in the deserialized instance
             UpdateRegistrationFlowWithProfileMethod instance = thisAdapter.fromJsonTree(jsonObj);
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
   * Create an instance of UpdateRegistrationFlowWithProfileMethod given an JSON string
   *
   * @param jsonString JSON string
   * @return An instance of UpdateRegistrationFlowWithProfileMethod
   * @throws IOException if the JSON string is invalid with respect to UpdateRegistrationFlowWithProfileMethod
   */
  public static UpdateRegistrationFlowWithProfileMethod fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UpdateRegistrationFlowWithProfileMethod.class);
  }

  /**
   * Convert an instance of UpdateRegistrationFlowWithProfileMethod to an JSON string
   *
   * @return JSON string
   */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

