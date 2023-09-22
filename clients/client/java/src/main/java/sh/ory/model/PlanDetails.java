/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.9
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
import java.util.HashMap;
import java.util.Map;
import sh.ory.model.GenericUsage;

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
 * PlanDetails
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-09-22T10:41:25.890682106Z[Etc/UTC]")
public class PlanDetails {
  public static final String SERIALIZED_NAME_BASE_FEE_MONTHLY = "base_fee_monthly";
  @SerializedName(SERIALIZED_NAME_BASE_FEE_MONTHLY)
  private Long baseFeeMonthly;

  public static final String SERIALIZED_NAME_BASE_FEE_YEARLY = "base_fee_yearly";
  @SerializedName(SERIALIZED_NAME_BASE_FEE_YEARLY)
  private Long baseFeeYearly;

  public static final String SERIALIZED_NAME_CUSTOM = "custom";
  @SerializedName(SERIALIZED_NAME_CUSTOM)
  private Boolean custom;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_FEATURES = "features";
  @SerializedName(SERIALIZED_NAME_FEATURES)
  private Map<String, GenericUsage> features = new HashMap<>();

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_VERSION = "version";
  @SerializedName(SERIALIZED_NAME_VERSION)
  private Long version;

  public PlanDetails() {
  }

  public PlanDetails baseFeeMonthly(Long baseFeeMonthly) {
    
    this.baseFeeMonthly = baseFeeMonthly;
    return this;
  }

   /**
   * BaseFeeMonthly is the monthly base fee for the plan.
   * @return baseFeeMonthly
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "BaseFeeMonthly is the monthly base fee for the plan.")

  public Long getBaseFeeMonthly() {
    return baseFeeMonthly;
  }


  public void setBaseFeeMonthly(Long baseFeeMonthly) {
    this.baseFeeMonthly = baseFeeMonthly;
  }


  public PlanDetails baseFeeYearly(Long baseFeeYearly) {
    
    this.baseFeeYearly = baseFeeYearly;
    return this;
  }

   /**
   * BaseFeeYearly is the yearly base fee for the plan.
   * @return baseFeeYearly
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "BaseFeeYearly is the yearly base fee for the plan.")

  public Long getBaseFeeYearly() {
    return baseFeeYearly;
  }


  public void setBaseFeeYearly(Long baseFeeYearly) {
    this.baseFeeYearly = baseFeeYearly;
  }


  public PlanDetails custom(Boolean custom) {
    
    this.custom = custom;
    return this;
  }

   /**
   * Custom is true if the plan is custom. This means it will be hidden from the pricing page.
   * @return custom
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Custom is true if the plan is custom. This means it will be hidden from the pricing page.")

  public Boolean getCustom() {
    return custom;
  }


  public void setCustom(Boolean custom) {
    this.custom = custom;
  }


  public PlanDetails description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Description is the description of the plan.
   * @return description
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Description is the description of the plan.")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public PlanDetails features(Map<String, GenericUsage> features) {
    
    this.features = features;
    return this;
  }

  public PlanDetails putFeaturesItem(String key, GenericUsage featuresItem) {
    this.features.put(key, featuresItem);
    return this;
  }

   /**
   * Features are the feature definitions included in the plan.
   * @return features
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Features are the feature definitions included in the plan.")

  public Map<String, GenericUsage> getFeatures() {
    return features;
  }


  public void setFeatures(Map<String, GenericUsage> features) {
    this.features = features;
  }


  public PlanDetails name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name is the name of the plan.
   * @return name
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Name is the name of the plan.")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public PlanDetails version(Long version) {
    
    this.version = version;
    return this;
  }

   /**
   * Version is the version of the plan. The combination of &#x60;name@version&#x60; must be unique.
   * @return version
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Version is the version of the plan. The combination of `name@version` must be unique.")

  public Long getVersion() {
    return version;
  }


  public void setVersion(Long version) {
    this.version = version;
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
   * @return the PlanDetails instance itself
   */
  public PlanDetails putAdditionalProperty(String key, Object value) {
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
    PlanDetails planDetails = (PlanDetails) o;
    return Objects.equals(this.baseFeeMonthly, planDetails.baseFeeMonthly) &&
        Objects.equals(this.baseFeeYearly, planDetails.baseFeeYearly) &&
        Objects.equals(this.custom, planDetails.custom) &&
        Objects.equals(this.description, planDetails.description) &&
        Objects.equals(this.features, planDetails.features) &&
        Objects.equals(this.name, planDetails.name) &&
        Objects.equals(this.version, planDetails.version)&&
        Objects.equals(this.additionalProperties, planDetails.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(baseFeeMonthly, baseFeeYearly, custom, description, features, name, version, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PlanDetails {\n");
    sb.append("    baseFeeMonthly: ").append(toIndentedString(baseFeeMonthly)).append("\n");
    sb.append("    baseFeeYearly: ").append(toIndentedString(baseFeeYearly)).append("\n");
    sb.append("    custom: ").append(toIndentedString(custom)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    features: ").append(toIndentedString(features)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
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
    openapiFields.add("base_fee_monthly");
    openapiFields.add("base_fee_yearly");
    openapiFields.add("custom");
    openapiFields.add("description");
    openapiFields.add("features");
    openapiFields.add("name");
    openapiFields.add("version");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("base_fee_monthly");
    openapiRequiredFields.add("base_fee_yearly");
    openapiRequiredFields.add("custom");
    openapiRequiredFields.add("description");
    openapiRequiredFields.add("features");
    openapiRequiredFields.add("name");
    openapiRequiredFields.add("version");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to PlanDetails
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!PlanDetails.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in PlanDetails is not found in the empty JSON string", PlanDetails.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : PlanDetails.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("description").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `description` to be a primitive type in the JSON string but got `%s`", jsonObj.get("description").toString()));
      }
      if (!jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!PlanDetails.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'PlanDetails' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<PlanDetails> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(PlanDetails.class));

       return (TypeAdapter<T>) new TypeAdapter<PlanDetails>() {
           @Override
           public void write(JsonWriter out, PlanDetails value) throws IOException {
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
           public PlanDetails read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             PlanDetails instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of PlanDetails given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of PlanDetails
  * @throws IOException if the JSON string is invalid with respect to PlanDetails
  */
  public static PlanDetails fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, PlanDetails.class);
  }

 /**
  * Convert an instance of PlanDetails to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

