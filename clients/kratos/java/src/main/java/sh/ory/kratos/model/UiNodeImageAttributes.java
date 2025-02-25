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
 * UiNodeImageAttributes
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2025-02-25T15:51:40.899187259Z[Etc/UTC]", comments = "Generator version: 7.7.0")
public class UiNodeImageAttributes {
  public static final String SERIALIZED_NAME_HEIGHT = "height";
  @SerializedName(SERIALIZED_NAME_HEIGHT)
  private Long height;

  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private String id;

  /**
   * NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \&quot;img\&quot;. text Text input Input img Image a Anchor script Script
   */
  @JsonAdapter(NodeTypeEnum.Adapter.class)
  public enum NodeTypeEnum {
    TEXT("text"),
    
    INPUT("input"),
    
    IMG("img"),
    
    A("a"),
    
    SCRIPT("script");

    private String value;

    NodeTypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static NodeTypeEnum fromValue(String value) {
      for (NodeTypeEnum b : NodeTypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<NodeTypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final NodeTypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public NodeTypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return NodeTypeEnum.fromValue(value);
      }
    }

    public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      String value = jsonElement.getAsString();
      NodeTypeEnum.fromValue(value);
    }
  }

  public static final String SERIALIZED_NAME_NODE_TYPE = "node_type";
  @SerializedName(SERIALIZED_NAME_NODE_TYPE)
  private NodeTypeEnum nodeType;

  public static final String SERIALIZED_NAME_SRC = "src";
  @SerializedName(SERIALIZED_NAME_SRC)
  private String src;

  public static final String SERIALIZED_NAME_WIDTH = "width";
  @SerializedName(SERIALIZED_NAME_WIDTH)
  private Long width;

  public UiNodeImageAttributes() {
  }

  public UiNodeImageAttributes height(Long height) {
    this.height = height;
    return this;
  }

  /**
   * Height of the image
   * @return height
   */
  @javax.annotation.Nonnull
  public Long getHeight() {
    return height;
  }

  public void setHeight(Long height) {
    this.height = height;
  }


  public UiNodeImageAttributes id(String id) {
    this.id = id;
    return this;
  }

  /**
   * A unique identifier
   * @return id
   */
  @javax.annotation.Nonnull
  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public UiNodeImageAttributes nodeType(NodeTypeEnum nodeType) {
    this.nodeType = nodeType;
    return this;
  }

  /**
   * NodeType represents this node&#39;s types. It is a mirror of &#x60;node.type&#x60; and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \&quot;img\&quot;. text Text input Input img Image a Anchor script Script
   * @return nodeType
   */
  @javax.annotation.Nonnull
  public NodeTypeEnum getNodeType() {
    return nodeType;
  }

  public void setNodeType(NodeTypeEnum nodeType) {
    this.nodeType = nodeType;
  }


  public UiNodeImageAttributes src(String src) {
    this.src = src;
    return this;
  }

  /**
   * The image&#39;s source URL.  format: uri
   * @return src
   */
  @javax.annotation.Nonnull
  public String getSrc() {
    return src;
  }

  public void setSrc(String src) {
    this.src = src;
  }


  public UiNodeImageAttributes width(Long width) {
    this.width = width;
    return this;
  }

  /**
   * Width of the image
   * @return width
   */
  @javax.annotation.Nonnull
  public Long getWidth() {
    return width;
  }

  public void setWidth(Long width) {
    this.width = width;
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
   * @return the UiNodeImageAttributes instance itself
   */
  public UiNodeImageAttributes putAdditionalProperty(String key, Object value) {
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
    UiNodeImageAttributes uiNodeImageAttributes = (UiNodeImageAttributes) o;
    return Objects.equals(this.height, uiNodeImageAttributes.height) &&
        Objects.equals(this.id, uiNodeImageAttributes.id) &&
        Objects.equals(this.nodeType, uiNodeImageAttributes.nodeType) &&
        Objects.equals(this.src, uiNodeImageAttributes.src) &&
        Objects.equals(this.width, uiNodeImageAttributes.width)&&
        Objects.equals(this.additionalProperties, uiNodeImageAttributes.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(height, id, nodeType, src, width, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UiNodeImageAttributes {\n");
    sb.append("    height: ").append(toIndentedString(height)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    nodeType: ").append(toIndentedString(nodeType)).append("\n");
    sb.append("    src: ").append(toIndentedString(src)).append("\n");
    sb.append("    width: ").append(toIndentedString(width)).append("\n");
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
    openapiFields.add("height");
    openapiFields.add("id");
    openapiFields.add("node_type");
    openapiFields.add("src");
    openapiFields.add("width");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("height");
    openapiRequiredFields.add("id");
    openapiRequiredFields.add("node_type");
    openapiRequiredFields.add("src");
    openapiRequiredFields.add("width");
  }

  /**
   * Validates the JSON Element and throws an exception if issues found
   *
   * @param jsonElement JSON Element
   * @throws IOException if the JSON Element is invalid with respect to UiNodeImageAttributes
   */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!UiNodeImageAttributes.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in UiNodeImageAttributes is not found in the empty JSON string", UiNodeImageAttributes.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : UiNodeImageAttributes.openapiRequiredFields) {
        if (jsonElement.getAsJsonObject().get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if (!jsonObj.get("id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("id").toString()));
      }
      if (!jsonObj.get("node_type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `node_type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("node_type").toString()));
      }
      // validate the required field `node_type`
      NodeTypeEnum.validateJsonElement(jsonObj.get("node_type"));
      if (!jsonObj.get("src").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `src` to be a primitive type in the JSON string but got `%s`", jsonObj.get("src").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UiNodeImageAttributes.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UiNodeImageAttributes' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UiNodeImageAttributes> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UiNodeImageAttributes.class));

       return (TypeAdapter<T>) new TypeAdapter<UiNodeImageAttributes>() {
           @Override
           public void write(JsonWriter out, UiNodeImageAttributes value) throws IOException {
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
           public UiNodeImageAttributes read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             JsonObject jsonObj = jsonElement.getAsJsonObject();
             // store additional fields in the deserialized instance
             UiNodeImageAttributes instance = thisAdapter.fromJsonTree(jsonObj);
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
   * Create an instance of UiNodeImageAttributes given an JSON string
   *
   * @param jsonString JSON string
   * @return An instance of UiNodeImageAttributes
   * @throws IOException if the JSON string is invalid with respect to UiNodeImageAttributes
   */
  public static UiNodeImageAttributes fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UiNodeImageAttributes.class);
  }

  /**
   * Convert an instance of UiNodeImageAttributes to an JSON string
   *
   * @return JSON string
   */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

