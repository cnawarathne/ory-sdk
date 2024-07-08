/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.13.6
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
 * OpenID Connect Userinfo
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-07-08T10:36:18.530884942Z[Etc/UTC]", comments = "Generator version: 7.7.0")
public class OidcUserInfo {
  public static final String SERIALIZED_NAME_BIRTHDATE = "birthdate";
  @SerializedName(SERIALIZED_NAME_BIRTHDATE)
  private String birthdate;

  public static final String SERIALIZED_NAME_EMAIL = "email";
  @SerializedName(SERIALIZED_NAME_EMAIL)
  private String email;

  public static final String SERIALIZED_NAME_EMAIL_VERIFIED = "email_verified";
  @SerializedName(SERIALIZED_NAME_EMAIL_VERIFIED)
  private Boolean emailVerified;

  public static final String SERIALIZED_NAME_FAMILY_NAME = "family_name";
  @SerializedName(SERIALIZED_NAME_FAMILY_NAME)
  private String familyName;

  public static final String SERIALIZED_NAME_GENDER = "gender";
  @SerializedName(SERIALIZED_NAME_GENDER)
  private String gender;

  public static final String SERIALIZED_NAME_GIVEN_NAME = "given_name";
  @SerializedName(SERIALIZED_NAME_GIVEN_NAME)
  private String givenName;

  public static final String SERIALIZED_NAME_LOCALE = "locale";
  @SerializedName(SERIALIZED_NAME_LOCALE)
  private String locale;

  public static final String SERIALIZED_NAME_MIDDLE_NAME = "middle_name";
  @SerializedName(SERIALIZED_NAME_MIDDLE_NAME)
  private String middleName;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_NICKNAME = "nickname";
  @SerializedName(SERIALIZED_NAME_NICKNAME)
  private String nickname;

  public static final String SERIALIZED_NAME_PHONE_NUMBER = "phone_number";
  @SerializedName(SERIALIZED_NAME_PHONE_NUMBER)
  private String phoneNumber;

  public static final String SERIALIZED_NAME_PHONE_NUMBER_VERIFIED = "phone_number_verified";
  @SerializedName(SERIALIZED_NAME_PHONE_NUMBER_VERIFIED)
  private Boolean phoneNumberVerified;

  public static final String SERIALIZED_NAME_PICTURE = "picture";
  @SerializedName(SERIALIZED_NAME_PICTURE)
  private String picture;

  public static final String SERIALIZED_NAME_PREFERRED_USERNAME = "preferred_username";
  @SerializedName(SERIALIZED_NAME_PREFERRED_USERNAME)
  private String preferredUsername;

  public static final String SERIALIZED_NAME_PROFILE = "profile";
  @SerializedName(SERIALIZED_NAME_PROFILE)
  private String profile;

  public static final String SERIALIZED_NAME_SUB = "sub";
  @SerializedName(SERIALIZED_NAME_SUB)
  private String sub;

  public static final String SERIALIZED_NAME_UPDATED_AT = "updated_at";
  @SerializedName(SERIALIZED_NAME_UPDATED_AT)
  private Long updatedAt;

  public static final String SERIALIZED_NAME_WEBSITE = "website";
  @SerializedName(SERIALIZED_NAME_WEBSITE)
  private String website;

  public static final String SERIALIZED_NAME_ZONEINFO = "zoneinfo";
  @SerializedName(SERIALIZED_NAME_ZONEINFO)
  private String zoneinfo;

  public OidcUserInfo() {
  }

  public OidcUserInfo birthdate(String birthdate) {
    this.birthdate = birthdate;
    return this;
  }

  /**
   * End-User&#39;s birthday, represented as an ISO 8601:2004 [ISO8601‑2004] YYYY-MM-DD format. The year MAY be 0000, indicating that it is omitted. To represent only the year, YYYY format is allowed. Note that depending on the underlying platform&#39;s date related function, providing just year can result in varying month and day, so the implementers need to take this factor into account to correctly process the dates.
   * @return birthdate
   */
  @javax.annotation.Nullable
  public String getBirthdate() {
    return birthdate;
  }

  public void setBirthdate(String birthdate) {
    this.birthdate = birthdate;
  }


  public OidcUserInfo email(String email) {
    this.email = email;
    return this;
  }

  /**
   * End-User&#39;s preferred e-mail address. Its value MUST conform to the RFC 5322 [RFC5322] addr-spec syntax. The RP MUST NOT rely upon this value being unique, as discussed in Section 5.7.
   * @return email
   */
  @javax.annotation.Nullable
  public String getEmail() {
    return email;
  }

  public void setEmail(String email) {
    this.email = email;
  }


  public OidcUserInfo emailVerified(Boolean emailVerified) {
    this.emailVerified = emailVerified;
    return this;
  }

  /**
   * True if the End-User&#39;s e-mail address has been verified; otherwise false. When this Claim Value is true, this means that the OP took affirmative steps to ensure that this e-mail address was controlled by the End-User at the time the verification was performed. The means by which an e-mail address is verified is context-specific, and dependent upon the trust framework or contractual agreements within which the parties are operating.
   * @return emailVerified
   */
  @javax.annotation.Nullable
  public Boolean getEmailVerified() {
    return emailVerified;
  }

  public void setEmailVerified(Boolean emailVerified) {
    this.emailVerified = emailVerified;
  }


  public OidcUserInfo familyName(String familyName) {
    this.familyName = familyName;
    return this;
  }

  /**
   * Surname(s) or last name(s) of the End-User. Note that in some cultures, people can have multiple family names or no family name; all can be present, with the names being separated by space characters.
   * @return familyName
   */
  @javax.annotation.Nullable
  public String getFamilyName() {
    return familyName;
  }

  public void setFamilyName(String familyName) {
    this.familyName = familyName;
  }


  public OidcUserInfo gender(String gender) {
    this.gender = gender;
    return this;
  }

  /**
   * End-User&#39;s gender. Values defined by this specification are female and male. Other values MAY be used when neither of the defined values are applicable.
   * @return gender
   */
  @javax.annotation.Nullable
  public String getGender() {
    return gender;
  }

  public void setGender(String gender) {
    this.gender = gender;
  }


  public OidcUserInfo givenName(String givenName) {
    this.givenName = givenName;
    return this;
  }

  /**
   * Given name(s) or first name(s) of the End-User. Note that in some cultures, people can have multiple given names; all can be present, with the names being separated by space characters.
   * @return givenName
   */
  @javax.annotation.Nullable
  public String getGivenName() {
    return givenName;
  }

  public void setGivenName(String givenName) {
    this.givenName = givenName;
  }


  public OidcUserInfo locale(String locale) {
    this.locale = locale;
    return this;
  }

  /**
   * End-User&#39;s locale, represented as a BCP47 [RFC5646] language tag. This is typically an ISO 639-1 Alpha-2 [ISO639‑1] language code in lowercase and an ISO 3166-1 Alpha-2 [ISO3166‑1] country code in uppercase, separated by a dash. For example, en-US or fr-CA. As a compatibility note, some implementations have used an underscore as the separator rather than a dash, for example, en_US; Relying Parties MAY choose to accept this locale syntax as well.
   * @return locale
   */
  @javax.annotation.Nullable
  public String getLocale() {
    return locale;
  }

  public void setLocale(String locale) {
    this.locale = locale;
  }


  public OidcUserInfo middleName(String middleName) {
    this.middleName = middleName;
    return this;
  }

  /**
   * Middle name(s) of the End-User. Note that in some cultures, people can have multiple middle names; all can be present, with the names being separated by space characters. Also note that in some cultures, middle names are not used.
   * @return middleName
   */
  @javax.annotation.Nullable
  public String getMiddleName() {
    return middleName;
  }

  public void setMiddleName(String middleName) {
    this.middleName = middleName;
  }


  public OidcUserInfo name(String name) {
    this.name = name;
    return this;
  }

  /**
   * End-User&#39;s full name in displayable form including all name parts, possibly including titles and suffixes, ordered according to the End-User&#39;s locale and preferences.
   * @return name
   */
  @javax.annotation.Nullable
  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }


  public OidcUserInfo nickname(String nickname) {
    this.nickname = nickname;
    return this;
  }

  /**
   * Casual name of the End-User that may or may not be the same as the given_name. For instance, a nickname value of Mike might be returned alongside a given_name value of Michael.
   * @return nickname
   */
  @javax.annotation.Nullable
  public String getNickname() {
    return nickname;
  }

  public void setNickname(String nickname) {
    this.nickname = nickname;
  }


  public OidcUserInfo phoneNumber(String phoneNumber) {
    this.phoneNumber = phoneNumber;
    return this;
  }

  /**
   * End-User&#39;s preferred telephone number. E.164 [E.164] is RECOMMENDED as the format of this Claim, for example, +1 (425) 555-1212 or +56 (2) 687 2400. If the phone number contains an extension, it is RECOMMENDED that the extension be represented using the RFC 3966 [RFC3966] extension syntax, for example, +1 (604) 555-1234;ext&#x3D;5678.
   * @return phoneNumber
   */
  @javax.annotation.Nullable
  public String getPhoneNumber() {
    return phoneNumber;
  }

  public void setPhoneNumber(String phoneNumber) {
    this.phoneNumber = phoneNumber;
  }


  public OidcUserInfo phoneNumberVerified(Boolean phoneNumberVerified) {
    this.phoneNumberVerified = phoneNumberVerified;
    return this;
  }

  /**
   * True if the End-User&#39;s phone number has been verified; otherwise false. When this Claim Value is true, this means that the OP took affirmative steps to ensure that this phone number was controlled by the End-User at the time the verification was performed. The means by which a phone number is verified is context-specific, and dependent upon the trust framework or contractual agreements within which the parties are operating. When true, the phone_number Claim MUST be in E.164 format and any extensions MUST be represented in RFC 3966 format.
   * @return phoneNumberVerified
   */
  @javax.annotation.Nullable
  public Boolean getPhoneNumberVerified() {
    return phoneNumberVerified;
  }

  public void setPhoneNumberVerified(Boolean phoneNumberVerified) {
    this.phoneNumberVerified = phoneNumberVerified;
  }


  public OidcUserInfo picture(String picture) {
    this.picture = picture;
    return this;
  }

  /**
   * URL of the End-User&#39;s profile picture. This URL MUST refer to an image file (for example, a PNG, JPEG, or GIF image file), rather than to a Web page containing an image. Note that this URL SHOULD specifically reference a profile photo of the End-User suitable for displaying when describing the End-User, rather than an arbitrary photo taken by the End-User.
   * @return picture
   */
  @javax.annotation.Nullable
  public String getPicture() {
    return picture;
  }

  public void setPicture(String picture) {
    this.picture = picture;
  }


  public OidcUserInfo preferredUsername(String preferredUsername) {
    this.preferredUsername = preferredUsername;
    return this;
  }

  /**
   * Non-unique shorthand name by which the End-User wishes to be referred to at the RP, such as janedoe or j.doe. This value MAY be any valid JSON string including special characters such as @, /, or whitespace.
   * @return preferredUsername
   */
  @javax.annotation.Nullable
  public String getPreferredUsername() {
    return preferredUsername;
  }

  public void setPreferredUsername(String preferredUsername) {
    this.preferredUsername = preferredUsername;
  }


  public OidcUserInfo profile(String profile) {
    this.profile = profile;
    return this;
  }

  /**
   * URL of the End-User&#39;s profile page. The contents of this Web page SHOULD be about the End-User.
   * @return profile
   */
  @javax.annotation.Nullable
  public String getProfile() {
    return profile;
  }

  public void setProfile(String profile) {
    this.profile = profile;
  }


  public OidcUserInfo sub(String sub) {
    this.sub = sub;
    return this;
  }

  /**
   * Subject - Identifier for the End-User at the IssuerURL.
   * @return sub
   */
  @javax.annotation.Nullable
  public String getSub() {
    return sub;
  }

  public void setSub(String sub) {
    this.sub = sub;
  }


  public OidcUserInfo updatedAt(Long updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

  /**
   * Time the End-User&#39;s information was last updated. Its value is a JSON number representing the number of seconds from 1970-01-01T0:0:0Z as measured in UTC until the date/time.
   * @return updatedAt
   */
  @javax.annotation.Nullable
  public Long getUpdatedAt() {
    return updatedAt;
  }

  public void setUpdatedAt(Long updatedAt) {
    this.updatedAt = updatedAt;
  }


  public OidcUserInfo website(String website) {
    this.website = website;
    return this;
  }

  /**
   * URL of the End-User&#39;s Web page or blog. This Web page SHOULD contain information published by the End-User or an organization that the End-User is affiliated with.
   * @return website
   */
  @javax.annotation.Nullable
  public String getWebsite() {
    return website;
  }

  public void setWebsite(String website) {
    this.website = website;
  }


  public OidcUserInfo zoneinfo(String zoneinfo) {
    this.zoneinfo = zoneinfo;
    return this;
  }

  /**
   * String from zoneinfo [zoneinfo] time zone database representing the End-User&#39;s time zone. For example, Europe/Paris or America/Los_Angeles.
   * @return zoneinfo
   */
  @javax.annotation.Nullable
  public String getZoneinfo() {
    return zoneinfo;
  }

  public void setZoneinfo(String zoneinfo) {
    this.zoneinfo = zoneinfo;
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
   * @return the OidcUserInfo instance itself
   */
  public OidcUserInfo putAdditionalProperty(String key, Object value) {
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
    OidcUserInfo oidcUserInfo = (OidcUserInfo) o;
    return Objects.equals(this.birthdate, oidcUserInfo.birthdate) &&
        Objects.equals(this.email, oidcUserInfo.email) &&
        Objects.equals(this.emailVerified, oidcUserInfo.emailVerified) &&
        Objects.equals(this.familyName, oidcUserInfo.familyName) &&
        Objects.equals(this.gender, oidcUserInfo.gender) &&
        Objects.equals(this.givenName, oidcUserInfo.givenName) &&
        Objects.equals(this.locale, oidcUserInfo.locale) &&
        Objects.equals(this.middleName, oidcUserInfo.middleName) &&
        Objects.equals(this.name, oidcUserInfo.name) &&
        Objects.equals(this.nickname, oidcUserInfo.nickname) &&
        Objects.equals(this.phoneNumber, oidcUserInfo.phoneNumber) &&
        Objects.equals(this.phoneNumberVerified, oidcUserInfo.phoneNumberVerified) &&
        Objects.equals(this.picture, oidcUserInfo.picture) &&
        Objects.equals(this.preferredUsername, oidcUserInfo.preferredUsername) &&
        Objects.equals(this.profile, oidcUserInfo.profile) &&
        Objects.equals(this.sub, oidcUserInfo.sub) &&
        Objects.equals(this.updatedAt, oidcUserInfo.updatedAt) &&
        Objects.equals(this.website, oidcUserInfo.website) &&
        Objects.equals(this.zoneinfo, oidcUserInfo.zoneinfo)&&
        Objects.equals(this.additionalProperties, oidcUserInfo.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(birthdate, email, emailVerified, familyName, gender, givenName, locale, middleName, name, nickname, phoneNumber, phoneNumberVerified, picture, preferredUsername, profile, sub, updatedAt, website, zoneinfo, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class OidcUserInfo {\n");
    sb.append("    birthdate: ").append(toIndentedString(birthdate)).append("\n");
    sb.append("    email: ").append(toIndentedString(email)).append("\n");
    sb.append("    emailVerified: ").append(toIndentedString(emailVerified)).append("\n");
    sb.append("    familyName: ").append(toIndentedString(familyName)).append("\n");
    sb.append("    gender: ").append(toIndentedString(gender)).append("\n");
    sb.append("    givenName: ").append(toIndentedString(givenName)).append("\n");
    sb.append("    locale: ").append(toIndentedString(locale)).append("\n");
    sb.append("    middleName: ").append(toIndentedString(middleName)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    nickname: ").append(toIndentedString(nickname)).append("\n");
    sb.append("    phoneNumber: ").append(toIndentedString(phoneNumber)).append("\n");
    sb.append("    phoneNumberVerified: ").append(toIndentedString(phoneNumberVerified)).append("\n");
    sb.append("    picture: ").append(toIndentedString(picture)).append("\n");
    sb.append("    preferredUsername: ").append(toIndentedString(preferredUsername)).append("\n");
    sb.append("    profile: ").append(toIndentedString(profile)).append("\n");
    sb.append("    sub: ").append(toIndentedString(sub)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
    sb.append("    website: ").append(toIndentedString(website)).append("\n");
    sb.append("    zoneinfo: ").append(toIndentedString(zoneinfo)).append("\n");
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
    openapiFields.add("birthdate");
    openapiFields.add("email");
    openapiFields.add("email_verified");
    openapiFields.add("family_name");
    openapiFields.add("gender");
    openapiFields.add("given_name");
    openapiFields.add("locale");
    openapiFields.add("middle_name");
    openapiFields.add("name");
    openapiFields.add("nickname");
    openapiFields.add("phone_number");
    openapiFields.add("phone_number_verified");
    openapiFields.add("picture");
    openapiFields.add("preferred_username");
    openapiFields.add("profile");
    openapiFields.add("sub");
    openapiFields.add("updated_at");
    openapiFields.add("website");
    openapiFields.add("zoneinfo");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

  /**
   * Validates the JSON Element and throws an exception if issues found
   *
   * @param jsonElement JSON Element
   * @throws IOException if the JSON Element is invalid with respect to OidcUserInfo
   */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!OidcUserInfo.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in OidcUserInfo is not found in the empty JSON string", OidcUserInfo.openapiRequiredFields.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("birthdate") != null && !jsonObj.get("birthdate").isJsonNull()) && !jsonObj.get("birthdate").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `birthdate` to be a primitive type in the JSON string but got `%s`", jsonObj.get("birthdate").toString()));
      }
      if ((jsonObj.get("email") != null && !jsonObj.get("email").isJsonNull()) && !jsonObj.get("email").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `email` to be a primitive type in the JSON string but got `%s`", jsonObj.get("email").toString()));
      }
      if ((jsonObj.get("family_name") != null && !jsonObj.get("family_name").isJsonNull()) && !jsonObj.get("family_name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `family_name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("family_name").toString()));
      }
      if ((jsonObj.get("gender") != null && !jsonObj.get("gender").isJsonNull()) && !jsonObj.get("gender").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `gender` to be a primitive type in the JSON string but got `%s`", jsonObj.get("gender").toString()));
      }
      if ((jsonObj.get("given_name") != null && !jsonObj.get("given_name").isJsonNull()) && !jsonObj.get("given_name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `given_name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("given_name").toString()));
      }
      if ((jsonObj.get("locale") != null && !jsonObj.get("locale").isJsonNull()) && !jsonObj.get("locale").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `locale` to be a primitive type in the JSON string but got `%s`", jsonObj.get("locale").toString()));
      }
      if ((jsonObj.get("middle_name") != null && !jsonObj.get("middle_name").isJsonNull()) && !jsonObj.get("middle_name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `middle_name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("middle_name").toString()));
      }
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("nickname") != null && !jsonObj.get("nickname").isJsonNull()) && !jsonObj.get("nickname").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `nickname` to be a primitive type in the JSON string but got `%s`", jsonObj.get("nickname").toString()));
      }
      if ((jsonObj.get("phone_number") != null && !jsonObj.get("phone_number").isJsonNull()) && !jsonObj.get("phone_number").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `phone_number` to be a primitive type in the JSON string but got `%s`", jsonObj.get("phone_number").toString()));
      }
      if ((jsonObj.get("picture") != null && !jsonObj.get("picture").isJsonNull()) && !jsonObj.get("picture").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `picture` to be a primitive type in the JSON string but got `%s`", jsonObj.get("picture").toString()));
      }
      if ((jsonObj.get("preferred_username") != null && !jsonObj.get("preferred_username").isJsonNull()) && !jsonObj.get("preferred_username").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `preferred_username` to be a primitive type in the JSON string but got `%s`", jsonObj.get("preferred_username").toString()));
      }
      if ((jsonObj.get("profile") != null && !jsonObj.get("profile").isJsonNull()) && !jsonObj.get("profile").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `profile` to be a primitive type in the JSON string but got `%s`", jsonObj.get("profile").toString()));
      }
      if ((jsonObj.get("sub") != null && !jsonObj.get("sub").isJsonNull()) && !jsonObj.get("sub").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `sub` to be a primitive type in the JSON string but got `%s`", jsonObj.get("sub").toString()));
      }
      if ((jsonObj.get("website") != null && !jsonObj.get("website").isJsonNull()) && !jsonObj.get("website").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `website` to be a primitive type in the JSON string but got `%s`", jsonObj.get("website").toString()));
      }
      if ((jsonObj.get("zoneinfo") != null && !jsonObj.get("zoneinfo").isJsonNull()) && !jsonObj.get("zoneinfo").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `zoneinfo` to be a primitive type in the JSON string but got `%s`", jsonObj.get("zoneinfo").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!OidcUserInfo.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'OidcUserInfo' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<OidcUserInfo> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(OidcUserInfo.class));

       return (TypeAdapter<T>) new TypeAdapter<OidcUserInfo>() {
           @Override
           public void write(JsonWriter out, OidcUserInfo value) throws IOException {
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
           public OidcUserInfo read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             JsonObject jsonObj = jsonElement.getAsJsonObject();
             // store additional fields in the deserialized instance
             OidcUserInfo instance = thisAdapter.fromJsonTree(jsonObj);
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
   * Create an instance of OidcUserInfo given an JSON string
   *
   * @param jsonString JSON string
   * @return An instance of OidcUserInfo
   * @throws IOException if the JSON string is invalid with respect to OidcUserInfo
   */
  public static OidcUserInfo fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, OidcUserInfo.class);
  }

  /**
   * Convert an instance of OidcUserInfo to an JSON string
   *
   * @return JSON string
   */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

