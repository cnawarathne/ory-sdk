/*
 * Ory Hydra API
 * Documentation for all of Ory Hydra's APIs. 
 *
 * The version of the OpenAPI document: v2.4.0-alpha.1
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.hydra.model;

import java.util.Objects;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.openapitools.jackson.nullable.JsonNullable;
import sh.ory.hydra.model.AcceptOAuth2ConsentRequestSession;

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

import sh.ory.hydra.JSON;

/**
 * AcceptOAuth2ConsentRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2025-02-09T15:15:46.432729038Z[Etc/UTC]", comments = "Generator version: 7.7.0")
public class AcceptOAuth2ConsentRequest {
  public static final String SERIALIZED_NAME_CONTEXT = "context";
  @SerializedName(SERIALIZED_NAME_CONTEXT)
  private Object context = null;

  public static final String SERIALIZED_NAME_GRANT_ACCESS_TOKEN_AUDIENCE = "grant_access_token_audience";
  @SerializedName(SERIALIZED_NAME_GRANT_ACCESS_TOKEN_AUDIENCE)
  private List<String> grantAccessTokenAudience = new ArrayList<>();

  public static final String SERIALIZED_NAME_GRANT_SCOPE = "grant_scope";
  @SerializedName(SERIALIZED_NAME_GRANT_SCOPE)
  private List<String> grantScope = new ArrayList<>();

  public static final String SERIALIZED_NAME_HANDLED_AT = "handled_at";
  @SerializedName(SERIALIZED_NAME_HANDLED_AT)
  private OffsetDateTime handledAt;

  public static final String SERIALIZED_NAME_REMEMBER = "remember";
  @SerializedName(SERIALIZED_NAME_REMEMBER)
  private Boolean remember;

  public static final String SERIALIZED_NAME_REMEMBER_FOR = "remember_for";
  @SerializedName(SERIALIZED_NAME_REMEMBER_FOR)
  private Long rememberFor;

  public static final String SERIALIZED_NAME_SESSION = "session";
  @SerializedName(SERIALIZED_NAME_SESSION)
  private AcceptOAuth2ConsentRequestSession session;

  public AcceptOAuth2ConsentRequest() {
  }

  public AcceptOAuth2ConsentRequest context(Object context) {
    this.context = context;
    return this;
  }

  /**
   * Get context
   * @return context
   */
  @javax.annotation.Nullable
  public Object getContext() {
    return context;
  }

  public void setContext(Object context) {
    this.context = context;
  }


  public AcceptOAuth2ConsentRequest grantAccessTokenAudience(List<String> grantAccessTokenAudience) {
    this.grantAccessTokenAudience = grantAccessTokenAudience;
    return this;
  }

  public AcceptOAuth2ConsentRequest addGrantAccessTokenAudienceItem(String grantAccessTokenAudienceItem) {
    if (this.grantAccessTokenAudience == null) {
      this.grantAccessTokenAudience = new ArrayList<>();
    }
    this.grantAccessTokenAudience.add(grantAccessTokenAudienceItem);
    return this;
  }

  /**
   * Get grantAccessTokenAudience
   * @return grantAccessTokenAudience
   */
  @javax.annotation.Nullable
  public List<String> getGrantAccessTokenAudience() {
    return grantAccessTokenAudience;
  }

  public void setGrantAccessTokenAudience(List<String> grantAccessTokenAudience) {
    this.grantAccessTokenAudience = grantAccessTokenAudience;
  }


  public AcceptOAuth2ConsentRequest grantScope(List<String> grantScope) {
    this.grantScope = grantScope;
    return this;
  }

  public AcceptOAuth2ConsentRequest addGrantScopeItem(String grantScopeItem) {
    if (this.grantScope == null) {
      this.grantScope = new ArrayList<>();
    }
    this.grantScope.add(grantScopeItem);
    return this;
  }

  /**
   * Get grantScope
   * @return grantScope
   */
  @javax.annotation.Nullable
  public List<String> getGrantScope() {
    return grantScope;
  }

  public void setGrantScope(List<String> grantScope) {
    this.grantScope = grantScope;
  }


  public AcceptOAuth2ConsentRequest handledAt(OffsetDateTime handledAt) {
    this.handledAt = handledAt;
    return this;
  }

  /**
   * Get handledAt
   * @return handledAt
   */
  @javax.annotation.Nullable
  public OffsetDateTime getHandledAt() {
    return handledAt;
  }

  public void setHandledAt(OffsetDateTime handledAt) {
    this.handledAt = handledAt;
  }


  public AcceptOAuth2ConsentRequest remember(Boolean remember) {
    this.remember = remember;
    return this;
  }

  /**
   * Remember, if set to true, tells ORY Hydra to remember this consent authorization and reuse it if the same client asks the same user for the same, or a subset of, scope.
   * @return remember
   */
  @javax.annotation.Nullable
  public Boolean getRemember() {
    return remember;
  }

  public void setRemember(Boolean remember) {
    this.remember = remember;
  }


  public AcceptOAuth2ConsentRequest rememberFor(Long rememberFor) {
    this.rememberFor = rememberFor;
    return this;
  }

  /**
   * RememberFor sets how long the consent authorization should be remembered for in seconds. If set to &#x60;0&#x60;, the authorization will be remembered indefinitely.
   * @return rememberFor
   */
  @javax.annotation.Nullable
  public Long getRememberFor() {
    return rememberFor;
  }

  public void setRememberFor(Long rememberFor) {
    this.rememberFor = rememberFor;
  }


  public AcceptOAuth2ConsentRequest session(AcceptOAuth2ConsentRequestSession session) {
    this.session = session;
    return this;
  }

  /**
   * Get session
   * @return session
   */
  @javax.annotation.Nullable
  public AcceptOAuth2ConsentRequestSession getSession() {
    return session;
  }

  public void setSession(AcceptOAuth2ConsentRequestSession session) {
    this.session = session;
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
   * @return the AcceptOAuth2ConsentRequest instance itself
   */
  public AcceptOAuth2ConsentRequest putAdditionalProperty(String key, Object value) {
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
    AcceptOAuth2ConsentRequest acceptOAuth2ConsentRequest = (AcceptOAuth2ConsentRequest) o;
    return Objects.equals(this.context, acceptOAuth2ConsentRequest.context) &&
        Objects.equals(this.grantAccessTokenAudience, acceptOAuth2ConsentRequest.grantAccessTokenAudience) &&
        Objects.equals(this.grantScope, acceptOAuth2ConsentRequest.grantScope) &&
        Objects.equals(this.handledAt, acceptOAuth2ConsentRequest.handledAt) &&
        Objects.equals(this.remember, acceptOAuth2ConsentRequest.remember) &&
        Objects.equals(this.rememberFor, acceptOAuth2ConsentRequest.rememberFor) &&
        Objects.equals(this.session, acceptOAuth2ConsentRequest.session)&&
        Objects.equals(this.additionalProperties, acceptOAuth2ConsentRequest.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(context, grantAccessTokenAudience, grantScope, handledAt, remember, rememberFor, session, additionalProperties);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AcceptOAuth2ConsentRequest {\n");
    sb.append("    context: ").append(toIndentedString(context)).append("\n");
    sb.append("    grantAccessTokenAudience: ").append(toIndentedString(grantAccessTokenAudience)).append("\n");
    sb.append("    grantScope: ").append(toIndentedString(grantScope)).append("\n");
    sb.append("    handledAt: ").append(toIndentedString(handledAt)).append("\n");
    sb.append("    remember: ").append(toIndentedString(remember)).append("\n");
    sb.append("    rememberFor: ").append(toIndentedString(rememberFor)).append("\n");
    sb.append("    session: ").append(toIndentedString(session)).append("\n");
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
    openapiFields.add("context");
    openapiFields.add("grant_access_token_audience");
    openapiFields.add("grant_scope");
    openapiFields.add("handled_at");
    openapiFields.add("remember");
    openapiFields.add("remember_for");
    openapiFields.add("session");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

  /**
   * Validates the JSON Element and throws an exception if issues found
   *
   * @param jsonElement JSON Element
   * @throws IOException if the JSON Element is invalid with respect to AcceptOAuth2ConsentRequest
   */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!AcceptOAuth2ConsentRequest.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in AcceptOAuth2ConsentRequest is not found in the empty JSON string", AcceptOAuth2ConsentRequest.openapiRequiredFields.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // ensure the optional json data is an array if present
      if (jsonObj.get("grant_access_token_audience") != null && !jsonObj.get("grant_access_token_audience").isJsonNull() && !jsonObj.get("grant_access_token_audience").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `grant_access_token_audience` to be an array in the JSON string but got `%s`", jsonObj.get("grant_access_token_audience").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("grant_scope") != null && !jsonObj.get("grant_scope").isJsonNull() && !jsonObj.get("grant_scope").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `grant_scope` to be an array in the JSON string but got `%s`", jsonObj.get("grant_scope").toString()));
      }
      // validate the optional field `session`
      if (jsonObj.get("session") != null && !jsonObj.get("session").isJsonNull()) {
        AcceptOAuth2ConsentRequestSession.validateJsonElement(jsonObj.get("session"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!AcceptOAuth2ConsentRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'AcceptOAuth2ConsentRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<AcceptOAuth2ConsentRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(AcceptOAuth2ConsentRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<AcceptOAuth2ConsentRequest>() {
           @Override
           public void write(JsonWriter out, AcceptOAuth2ConsentRequest value) throws IOException {
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
           public AcceptOAuth2ConsentRequest read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             JsonObject jsonObj = jsonElement.getAsJsonObject();
             // store additional fields in the deserialized instance
             AcceptOAuth2ConsentRequest instance = thisAdapter.fromJsonTree(jsonObj);
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
   * Create an instance of AcceptOAuth2ConsentRequest given an JSON string
   *
   * @param jsonString JSON string
   * @return An instance of AcceptOAuth2ConsentRequest
   * @throws IOException if the JSON string is invalid with respect to AcceptOAuth2ConsentRequest
   */
  public static AcceptOAuth2ConsentRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, AcceptOAuth2ConsentRequest.class);
  }

  /**
   * Convert an instance of AcceptOAuth2ConsentRequest to an JSON string
   *
   * @return JSON string
   */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

