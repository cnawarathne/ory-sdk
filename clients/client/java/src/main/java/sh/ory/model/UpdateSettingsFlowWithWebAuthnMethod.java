/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.4.0
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

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

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import sh.ory.JSON;

/**
 * Update Settings Flow with WebAuthn Method
 */
@ApiModel(description = "Update Settings Flow with WebAuthn Method")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-11-15T13:59:44.429413179Z[Etc/UTC]")
public class UpdateSettingsFlowWithWebAuthnMethod {
  public static final String SERIALIZED_NAME_CSRF_TOKEN = "csrf_token";
  @SerializedName(SERIALIZED_NAME_CSRF_TOKEN)
  private String csrfToken;

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private String method;

  public static final String SERIALIZED_NAME_WEBAUTHN_REGISTER = "webauthn_register";
  @SerializedName(SERIALIZED_NAME_WEBAUTHN_REGISTER)
  private String webauthnRegister;

  public static final String SERIALIZED_NAME_WEBAUTHN_REGISTER_DISPLAYNAME = "webauthn_register_displayname";
  @SerializedName(SERIALIZED_NAME_WEBAUTHN_REGISTER_DISPLAYNAME)
  private String webauthnRegisterDisplayname;

  public static final String SERIALIZED_NAME_WEBAUTHN_REMOVE = "webauthn_remove";
  @SerializedName(SERIALIZED_NAME_WEBAUTHN_REMOVE)
  private String webauthnRemove;

  public UpdateSettingsFlowWithWebAuthnMethod() {
  }

  public UpdateSettingsFlowWithWebAuthnMethod csrfToken(String csrfToken) {
    
    this.csrfToken = csrfToken;
    return this;
  }

   /**
   * CSRFToken is the anti-CSRF token
   * @return csrfToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "CSRFToken is the anti-CSRF token")

  public String getCsrfToken() {
    return csrfToken;
  }


  public void setCsrfToken(String csrfToken) {
    this.csrfToken = csrfToken;
  }


  public UpdateSettingsFlowWithWebAuthnMethod method(String method) {
    
    this.method = method;
    return this;
  }

   /**
   * Method  Should be set to \&quot;webauthn\&quot; when trying to add, update, or remove a webAuthn pairing.
   * @return method
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Method  Should be set to \"webauthn\" when trying to add, update, or remove a webAuthn pairing.")

  public String getMethod() {
    return method;
  }


  public void setMethod(String method) {
    this.method = method;
  }


  public UpdateSettingsFlowWithWebAuthnMethod webauthnRegister(String webauthnRegister) {
    
    this.webauthnRegister = webauthnRegister;
    return this;
  }

   /**
   * Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here.
   * @return webauthnRegister
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here.")

  public String getWebauthnRegister() {
    return webauthnRegister;
  }


  public void setWebauthnRegister(String webauthnRegister) {
    this.webauthnRegister = webauthnRegister;
  }


  public UpdateSettingsFlowWithWebAuthnMethod webauthnRegisterDisplayname(String webauthnRegisterDisplayname) {
    
    this.webauthnRegisterDisplayname = webauthnRegisterDisplayname;
    return this;
  }

   /**
   * Name of the WebAuthn Security Key to be Added  A human-readable name for the security key which will be added.
   * @return webauthnRegisterDisplayname
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Name of the WebAuthn Security Key to be Added  A human-readable name for the security key which will be added.")

  public String getWebauthnRegisterDisplayname() {
    return webauthnRegisterDisplayname;
  }


  public void setWebauthnRegisterDisplayname(String webauthnRegisterDisplayname) {
    this.webauthnRegisterDisplayname = webauthnRegisterDisplayname;
  }


  public UpdateSettingsFlowWithWebAuthnMethod webauthnRemove(String webauthnRemove) {
    
    this.webauthnRemove = webauthnRemove;
    return this;
  }

   /**
   * Remove a WebAuthn Security Key  This must contain the ID of the WebAuthN connection.
   * @return webauthnRemove
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Remove a WebAuthn Security Key  This must contain the ID of the WebAuthN connection.")

  public String getWebauthnRemove() {
    return webauthnRemove;
  }


  public void setWebauthnRemove(String webauthnRemove) {
    this.webauthnRemove = webauthnRemove;
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
   * @return the UpdateSettingsFlowWithWebAuthnMethod instance itself
   */
  public UpdateSettingsFlowWithWebAuthnMethod putAdditionalProperty(String key, Object value) {
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
    UpdateSettingsFlowWithWebAuthnMethod updateSettingsFlowWithWebAuthnMethod = (UpdateSettingsFlowWithWebAuthnMethod) o;
    return Objects.equals(this.csrfToken, updateSettingsFlowWithWebAuthnMethod.csrfToken) &&
        Objects.equals(this.method, updateSettingsFlowWithWebAuthnMethod.method) &&
        Objects.equals(this.webauthnRegister, updateSettingsFlowWithWebAuthnMethod.webauthnRegister) &&
        Objects.equals(this.webauthnRegisterDisplayname, updateSettingsFlowWithWebAuthnMethod.webauthnRegisterDisplayname) &&
        Objects.equals(this.webauthnRemove, updateSettingsFlowWithWebAuthnMethod.webauthnRemove)&&
        Objects.equals(this.additionalProperties, updateSettingsFlowWithWebAuthnMethod.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(csrfToken, method, webauthnRegister, webauthnRegisterDisplayname, webauthnRemove, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdateSettingsFlowWithWebAuthnMethod {\n");
    sb.append("    csrfToken: ").append(toIndentedString(csrfToken)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    webauthnRegister: ").append(toIndentedString(webauthnRegister)).append("\n");
    sb.append("    webauthnRegisterDisplayname: ").append(toIndentedString(webauthnRegisterDisplayname)).append("\n");
    sb.append("    webauthnRemove: ").append(toIndentedString(webauthnRemove)).append("\n");
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
    openapiFields.add("webauthn_register");
    openapiFields.add("webauthn_register_displayname");
    openapiFields.add("webauthn_remove");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("method");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to UpdateSettingsFlowWithWebAuthnMethod
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!UpdateSettingsFlowWithWebAuthnMethod.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in UpdateSettingsFlowWithWebAuthnMethod is not found in the empty JSON string", UpdateSettingsFlowWithWebAuthnMethod.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : UpdateSettingsFlowWithWebAuthnMethod.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if ((jsonObj.get("csrf_token") != null && !jsonObj.get("csrf_token").isJsonNull()) && !jsonObj.get("csrf_token").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `csrf_token` to be a primitive type in the JSON string but got `%s`", jsonObj.get("csrf_token").toString()));
      }
      if (!jsonObj.get("method").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `method` to be a primitive type in the JSON string but got `%s`", jsonObj.get("method").toString()));
      }
      if ((jsonObj.get("webauthn_register") != null && !jsonObj.get("webauthn_register").isJsonNull()) && !jsonObj.get("webauthn_register").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `webauthn_register` to be a primitive type in the JSON string but got `%s`", jsonObj.get("webauthn_register").toString()));
      }
      if ((jsonObj.get("webauthn_register_displayname") != null && !jsonObj.get("webauthn_register_displayname").isJsonNull()) && !jsonObj.get("webauthn_register_displayname").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `webauthn_register_displayname` to be a primitive type in the JSON string but got `%s`", jsonObj.get("webauthn_register_displayname").toString()));
      }
      if ((jsonObj.get("webauthn_remove") != null && !jsonObj.get("webauthn_remove").isJsonNull()) && !jsonObj.get("webauthn_remove").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `webauthn_remove` to be a primitive type in the JSON string but got `%s`", jsonObj.get("webauthn_remove").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UpdateSettingsFlowWithWebAuthnMethod.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UpdateSettingsFlowWithWebAuthnMethod' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UpdateSettingsFlowWithWebAuthnMethod> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UpdateSettingsFlowWithWebAuthnMethod.class));

       return (TypeAdapter<T>) new TypeAdapter<UpdateSettingsFlowWithWebAuthnMethod>() {
           @Override
           public void write(JsonWriter out, UpdateSettingsFlowWithWebAuthnMethod value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             obj.remove("additionalProperties");
             // serialize additonal properties
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
                   obj.add(entry.getKey(), gson.toJsonTree(entry.getValue()).getAsJsonObject());
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public UpdateSettingsFlowWithWebAuthnMethod read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             UpdateSettingsFlowWithWebAuthnMethod instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of UpdateSettingsFlowWithWebAuthnMethod given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of UpdateSettingsFlowWithWebAuthnMethod
  * @throws IOException if the JSON string is invalid with respect to UpdateSettingsFlowWithWebAuthnMethod
  */
  public static UpdateSettingsFlowWithWebAuthnMethod fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UpdateSettingsFlowWithWebAuthnMethod.class);
  }

 /**
  * Convert an instance of UpdateSettingsFlowWithWebAuthnMethod to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

