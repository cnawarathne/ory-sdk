<?php
/**
 * IntrospectedOAuth2Token
 *
 * PHP version 7.3
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.
 *
 * The version of the OpenAPI document: v1.2.0
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.4.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace Ory\Client\Model;

use \ArrayAccess;
use \Ory\Client\ObjectSerializer;

/**
 * IntrospectedOAuth2Token Class Doc Comment
 *
 * @category Class
 * @description Introspection contains an access token&#39;s session data as specified by [IETF RFC 7662](https://tools.ietf.org/html/rfc7662)
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class IntrospectedOAuth2Token implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'introspectedOAuth2Token';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'active' => 'bool',
        'aud' => 'string[]',
        'clientId' => 'string',
        'exp' => 'int',
        'ext' => 'array<string,mixed>',
        'iat' => 'int',
        'iss' => 'string',
        'nbf' => 'int',
        'obfuscatedSubject' => 'string',
        'scope' => 'string',
        'sub' => 'string',
        'tokenType' => 'string',
        'tokenUse' => 'string',
        'username' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'active' => null,
        'aud' => null,
        'clientId' => null,
        'exp' => 'int64',
        'ext' => null,
        'iat' => 'int64',
        'iss' => null,
        'nbf' => 'int64',
        'obfuscatedSubject' => null,
        'scope' => null,
        'sub' => null,
        'tokenType' => null,
        'tokenUse' => null,
        'username' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'active' => 'active',
        'aud' => 'aud',
        'clientId' => 'client_id',
        'exp' => 'exp',
        'ext' => 'ext',
        'iat' => 'iat',
        'iss' => 'iss',
        'nbf' => 'nbf',
        'obfuscatedSubject' => 'obfuscated_subject',
        'scope' => 'scope',
        'sub' => 'sub',
        'tokenType' => 'token_type',
        'tokenUse' => 'token_use',
        'username' => 'username'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'active' => 'setActive',
        'aud' => 'setAud',
        'clientId' => 'setClientId',
        'exp' => 'setExp',
        'ext' => 'setExt',
        'iat' => 'setIat',
        'iss' => 'setIss',
        'nbf' => 'setNbf',
        'obfuscatedSubject' => 'setObfuscatedSubject',
        'scope' => 'setScope',
        'sub' => 'setSub',
        'tokenType' => 'setTokenType',
        'tokenUse' => 'setTokenUse',
        'username' => 'setUsername'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'active' => 'getActive',
        'aud' => 'getAud',
        'clientId' => 'getClientId',
        'exp' => 'getExp',
        'ext' => 'getExt',
        'iat' => 'getIat',
        'iss' => 'getIss',
        'nbf' => 'getNbf',
        'obfuscatedSubject' => 'getObfuscatedSubject',
        'scope' => 'getScope',
        'sub' => 'getSub',
        'tokenType' => 'getTokenType',
        'tokenUse' => 'getTokenUse',
        'username' => 'getUsername'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['active'] = $data['active'] ?? null;
        $this->container['aud'] = $data['aud'] ?? null;
        $this->container['clientId'] = $data['clientId'] ?? null;
        $this->container['exp'] = $data['exp'] ?? null;
        $this->container['ext'] = $data['ext'] ?? null;
        $this->container['iat'] = $data['iat'] ?? null;
        $this->container['iss'] = $data['iss'] ?? null;
        $this->container['nbf'] = $data['nbf'] ?? null;
        $this->container['obfuscatedSubject'] = $data['obfuscatedSubject'] ?? null;
        $this->container['scope'] = $data['scope'] ?? null;
        $this->container['sub'] = $data['sub'] ?? null;
        $this->container['tokenType'] = $data['tokenType'] ?? null;
        $this->container['tokenUse'] = $data['tokenUse'] ?? null;
        $this->container['username'] = $data['username'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['active'] === null) {
            $invalidProperties[] = "'active' can't be null";
        }
        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets active
     *
     * @return bool
     */
    public function getActive()
    {
        return $this->container['active'];
    }

    /**
     * Sets active
     *
     * @param bool $active Active is a boolean indicator of whether or not the presented token is currently active.  The specifics of a token's \"active\" state will vary depending on the implementation of the authorization server and the information it keeps about its tokens, but a \"true\" value return for the \"active\" property will generally indicate that a given token has been issued by this authorization server, has not been revoked by the resource owner, and is within its given time window of validity (e.g., after its issuance time and before its expiration time).
     *
     * @return self
     */
    public function setActive($active)
    {
        $this->container['active'] = $active;

        return $this;
    }

    /**
     * Gets aud
     *
     * @return string[]|null
     */
    public function getAud()
    {
        return $this->container['aud'];
    }

    /**
     * Sets aud
     *
     * @param string[]|null $aud Audience contains a list of the token's intended audiences.
     *
     * @return self
     */
    public function setAud($aud)
    {
        $this->container['aud'] = $aud;

        return $this;
    }

    /**
     * Gets clientId
     *
     * @return string|null
     */
    public function getClientId()
    {
        return $this->container['clientId'];
    }

    /**
     * Sets clientId
     *
     * @param string|null $clientId ID is aclient identifier for the OAuth 2.0 client that requested this token.
     *
     * @return self
     */
    public function setClientId($clientId)
    {
        $this->container['clientId'] = $clientId;

        return $this;
    }

    /**
     * Gets exp
     *
     * @return int|null
     */
    public function getExp()
    {
        return $this->container['exp'];
    }

    /**
     * Sets exp
     *
     * @param int|null $exp Expires at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token will expire.
     *
     * @return self
     */
    public function setExp($exp)
    {
        $this->container['exp'] = $exp;

        return $this;
    }

    /**
     * Gets ext
     *
     * @return array<string,mixed>|null
     */
    public function getExt()
    {
        return $this->container['ext'];
    }

    /**
     * Sets ext
     *
     * @param array<string,mixed>|null $ext Extra is arbitrary data set by the session.
     *
     * @return self
     */
    public function setExt($ext)
    {
        $this->container['ext'] = $ext;

        return $this;
    }

    /**
     * Gets iat
     *
     * @return int|null
     */
    public function getIat()
    {
        return $this->container['iat'];
    }

    /**
     * Sets iat
     *
     * @param int|null $iat Issued at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token was originally issued.
     *
     * @return self
     */
    public function setIat($iat)
    {
        $this->container['iat'] = $iat;

        return $this;
    }

    /**
     * Gets iss
     *
     * @return string|null
     */
    public function getIss()
    {
        return $this->container['iss'];
    }

    /**
     * Sets iss
     *
     * @param string|null $iss IssuerURL is a string representing the issuer of this token
     *
     * @return self
     */
    public function setIss($iss)
    {
        $this->container['iss'] = $iss;

        return $this;
    }

    /**
     * Gets nbf
     *
     * @return int|null
     */
    public function getNbf()
    {
        return $this->container['nbf'];
    }

    /**
     * Sets nbf
     *
     * @param int|null $nbf NotBefore is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token is not to be used before.
     *
     * @return self
     */
    public function setNbf($nbf)
    {
        $this->container['nbf'] = $nbf;

        return $this;
    }

    /**
     * Gets obfuscatedSubject
     *
     * @return string|null
     */
    public function getObfuscatedSubject()
    {
        return $this->container['obfuscatedSubject'];
    }

    /**
     * Sets obfuscatedSubject
     *
     * @param string|null $obfuscatedSubject ObfuscatedSubject is set when the subject identifier algorithm was set to \"pairwise\" during authorization. It is the `sub` value of the ID Token that was issued.
     *
     * @return self
     */
    public function setObfuscatedSubject($obfuscatedSubject)
    {
        $this->container['obfuscatedSubject'] = $obfuscatedSubject;

        return $this;
    }

    /**
     * Gets scope
     *
     * @return string|null
     */
    public function getScope()
    {
        return $this->container['scope'];
    }

    /**
     * Sets scope
     *
     * @param string|null $scope Scope is a JSON string containing a space-separated list of scopes associated with this token.
     *
     * @return self
     */
    public function setScope($scope)
    {
        $this->container['scope'] = $scope;

        return $this;
    }

    /**
     * Gets sub
     *
     * @return string|null
     */
    public function getSub()
    {
        return $this->container['sub'];
    }

    /**
     * Sets sub
     *
     * @param string|null $sub Subject of the token, as defined in JWT [RFC7519]. Usually a machine-readable identifier of the resource owner who authorized this token.
     *
     * @return self
     */
    public function setSub($sub)
    {
        $this->container['sub'] = $sub;

        return $this;
    }

    /**
     * Gets tokenType
     *
     * @return string|null
     */
    public function getTokenType()
    {
        return $this->container['tokenType'];
    }

    /**
     * Sets tokenType
     *
     * @param string|null $tokenType TokenType is the introspected token's type, typically `Bearer`.
     *
     * @return self
     */
    public function setTokenType($tokenType)
    {
        $this->container['tokenType'] = $tokenType;

        return $this;
    }

    /**
     * Gets tokenUse
     *
     * @return string|null
     */
    public function getTokenUse()
    {
        return $this->container['tokenUse'];
    }

    /**
     * Sets tokenUse
     *
     * @param string|null $tokenUse TokenUse is the introspected token's use, for example `access_token` or `refresh_token`.
     *
     * @return self
     */
    public function setTokenUse($tokenUse)
    {
        $this->container['tokenUse'] = $tokenUse;

        return $this;
    }

    /**
     * Gets username
     *
     * @return string|null
     */
    public function getUsername()
    {
        return $this->container['username'];
    }

    /**
     * Sets username
     *
     * @param string|null $username Username is a human-readable identifier for the resource owner who authorized this token.
     *
     * @return self
     */
    public function setUsername($username)
    {
        $this->container['username'] = $username;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


