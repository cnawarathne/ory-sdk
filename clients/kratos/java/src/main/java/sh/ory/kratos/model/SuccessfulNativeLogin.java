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
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import sh.ory.kratos.model.ContinueWith;
import sh.ory.kratos.model.Session;

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
 * The Response for Login Flows via API
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2025-02-25T15:51:40.899187259Z[Etc/UTC]", comments = "Generator version: 7.7.0")
public class SuccessfulNativeLogin {
  public static final String SERIALIZED_NAME_CONTINUE_WITH = "continue_with";
  @SerializedName(SERIALIZED_NAME_CONTINUE_WITH)
  private List<ContinueWith> continueWith = new ArrayList<>();

  public static final String SERIALIZED_NAME_SESSION = "session";
  @SerializedName(SERIALIZED_NAME_SESSION)
  private Session session;

  public static final String SERIALIZED_NAME_SESSION_TOKEN = "session_token";
  @SerializedName(SERIALIZED_NAME_SESSION_TOKEN)
  private String sessionToken;

  public SuccessfulNativeLogin() {
  }

  public SuccessfulNativeLogin continueWith(List<ContinueWith> continueWith) {
    this.continueWith = continueWith;
    return this;
  }

  public SuccessfulNativeLogin addContinueWithItem(ContinueWith continueWithItem) {
    if (this.continueWith == null) {
      this.continueWith = new ArrayList<>();
    }
    this.continueWith.add(continueWithItem);
    return this;
  }

  /**
   * Contains a list of actions, that could follow this flow  It can, for example, this will contain a reference to the verification flow, created as part of the user&#39;s registration or the token of the session.
   * @return continueWith
   */
  @javax.annotation.Nullable
  public List<ContinueWith> getContinueWith() {
    return continueWith;
  }

  public void setContinueWith(List<ContinueWith> continueWith) {
    this.continueWith = continueWith;
  }


  public SuccessfulNativeLogin session(Session session) {
    this.session = session;
    return this;
  }

  /**
   * Get session
   * @return session
   */
  @javax.annotation.Nonnull
  public Session getSession() {
    return session;
  }

  public void setSession(Session session) {
    this.session = session;
  }


  public SuccessfulNativeLogin sessionToken(String sessionToken) {
    this.sessionToken = sessionToken;
    return this;
  }

  /**
   * The Session Token  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization Header:  Authorization: bearer ${session-token}  The session token is only issued for API flows, not for Browser flows!
   * @return sessionToken
   */
  @javax.annotation.Nullable
  public String getSessionToken() {
    return sessionToken;
  }

  public void setSessionToken(String sessionToken) {
    this.sessionToken = sessionToken;
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
   * @return the SuccessfulNativeLogin instance itself
   */
  public SuccessfulNativeLogin putAdditionalProperty(String key, Object value) {
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
    SuccessfulNativeLogin successfulNativeLogin = (SuccessfulNativeLogin) o;
    return Objects.equals(this.continueWith, successfulNativeLogin.continueWith) &&
        Objects.equals(this.session, successfulNativeLogin.session) &&
        Objects.equals(this.sessionToken, successfulNativeLogin.sessionToken)&&
        Objects.equals(this.additionalProperties, successfulNativeLogin.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(continueWith, session, sessionToken, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SuccessfulNativeLogin {\n");
    sb.append("    continueWith: ").append(toIndentedString(continueWith)).append("\n");
    sb.append("    session: ").append(toIndentedString(session)).append("\n");
    sb.append("    sessionToken: ").append(toIndentedString(sessionToken)).append("\n");
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
    openapiFields.add("continue_with");
    openapiFields.add("session");
    openapiFields.add("session_token");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("session");
  }

  /**
   * Validates the JSON Element and throws an exception if issues found
   *
   * @param jsonElement JSON Element
   * @throws IOException if the JSON Element is invalid with respect to SuccessfulNativeLogin
   */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!SuccessfulNativeLogin.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in SuccessfulNativeLogin is not found in the empty JSON string", SuccessfulNativeLogin.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : SuccessfulNativeLogin.openapiRequiredFields) {
        if (jsonElement.getAsJsonObject().get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if (jsonObj.get("continue_with") != null && !jsonObj.get("continue_with").isJsonNull()) {
        JsonArray jsonArraycontinueWith = jsonObj.getAsJsonArray("continue_with");
        if (jsonArraycontinueWith != null) {
          // ensure the json data is an array
          if (!jsonObj.get("continue_with").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `continue_with` to be an array in the JSON string but got `%s`", jsonObj.get("continue_with").toString()));
          }

          // validate the optional field `continue_with` (array)
          for (int i = 0; i < jsonArraycontinueWith.size(); i++) {
            ContinueWith.validateJsonElement(jsonArraycontinueWith.get(i));
          };
        }
      }
      // validate the required field `session`
      Session.validateJsonElement(jsonObj.get("session"));
      if ((jsonObj.get("session_token") != null && !jsonObj.get("session_token").isJsonNull()) && !jsonObj.get("session_token").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `session_token` to be a primitive type in the JSON string but got `%s`", jsonObj.get("session_token").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!SuccessfulNativeLogin.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'SuccessfulNativeLogin' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<SuccessfulNativeLogin> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(SuccessfulNativeLogin.class));

       return (TypeAdapter<T>) new TypeAdapter<SuccessfulNativeLogin>() {
           @Override
           public void write(JsonWriter out, SuccessfulNativeLogin value) throws IOException {
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
           public SuccessfulNativeLogin read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             JsonObject jsonObj = jsonElement.getAsJsonObject();
             // store additional fields in the deserialized instance
             SuccessfulNativeLogin instance = thisAdapter.fromJsonTree(jsonObj);
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
   * Create an instance of SuccessfulNativeLogin given an JSON string
   *
   * @param jsonString JSON string
   * @return An instance of SuccessfulNativeLogin
   * @throws IOException if the JSON string is invalid with respect to SuccessfulNativeLogin
   */
  public static SuccessfulNativeLogin fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, SuccessfulNativeLogin.class);
  }

  /**
   * Convert an instance of SuccessfulNativeLogin to an JSON string
   *
   * @return JSON string
   */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

