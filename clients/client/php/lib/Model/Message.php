<?php
/**
 * Message
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
 * The version of the OpenAPI document: v1.0.2
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
 * Message Class Doc Comment
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class Message implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'message';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'body' => 'string',
        'createdAt' => '\DateTime',
        'id' => 'string',
        'recipient' => 'string',
        'sendCount' => 'int',
        'status' => '\Ory\Client\Model\CourierMessageStatus',
        'subject' => 'string',
        'templateType' => 'string',
        'type' => '\Ory\Client\Model\CourierMessageType',
        'updatedAt' => '\DateTime'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'body' => null,
        'createdAt' => 'date-time',
        'id' => 'uuid',
        'recipient' => null,
        'sendCount' => 'int64',
        'status' => null,
        'subject' => null,
        'templateType' => null,
        'type' => null,
        'updatedAt' => 'date-time'
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
        'body' => 'body',
        'createdAt' => 'created_at',
        'id' => 'id',
        'recipient' => 'recipient',
        'sendCount' => 'send_count',
        'status' => 'status',
        'subject' => 'subject',
        'templateType' => 'template_type',
        'type' => 'type',
        'updatedAt' => 'updated_at'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'body' => 'setBody',
        'createdAt' => 'setCreatedAt',
        'id' => 'setId',
        'recipient' => 'setRecipient',
        'sendCount' => 'setSendCount',
        'status' => 'setStatus',
        'subject' => 'setSubject',
        'templateType' => 'setTemplateType',
        'type' => 'setType',
        'updatedAt' => 'setUpdatedAt'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'body' => 'getBody',
        'createdAt' => 'getCreatedAt',
        'id' => 'getId',
        'recipient' => 'getRecipient',
        'sendCount' => 'getSendCount',
        'status' => 'getStatus',
        'subject' => 'getSubject',
        'templateType' => 'getTemplateType',
        'type' => 'getType',
        'updatedAt' => 'getUpdatedAt'
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

    const TEMPLATE_TYPE_RECOVERY_INVALID = 'recovery_invalid';
    const TEMPLATE_TYPE_RECOVERY_VALID = 'recovery_valid';
    const TEMPLATE_TYPE_RECOVERY_CODE_INVALID = 'recovery_code_invalid';
    const TEMPLATE_TYPE_RECOVERY_CODE_VALID = 'recovery_code_valid';
    const TEMPLATE_TYPE_VERIFICATION_INVALID = 'verification_invalid';
    const TEMPLATE_TYPE_VERIFICATION_VALID = 'verification_valid';
    const TEMPLATE_TYPE_VERIFICATION_CODE_INVALID = 'verification_code_invalid';
    const TEMPLATE_TYPE_VERIFICATION_CODE_VALID = 'verification_code_valid';
    const TEMPLATE_TYPE_OTP = 'otp';
    const TEMPLATE_TYPE_STUB = 'stub';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getTemplateTypeAllowableValues()
    {
        return [
            self::TEMPLATE_TYPE_RECOVERY_INVALID,
            self::TEMPLATE_TYPE_RECOVERY_VALID,
            self::TEMPLATE_TYPE_RECOVERY_CODE_INVALID,
            self::TEMPLATE_TYPE_RECOVERY_CODE_VALID,
            self::TEMPLATE_TYPE_VERIFICATION_INVALID,
            self::TEMPLATE_TYPE_VERIFICATION_VALID,
            self::TEMPLATE_TYPE_VERIFICATION_CODE_INVALID,
            self::TEMPLATE_TYPE_VERIFICATION_CODE_VALID,
            self::TEMPLATE_TYPE_OTP,
            self::TEMPLATE_TYPE_STUB,
        ];
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
        $this->container['body'] = $data['body'] ?? null;
        $this->container['createdAt'] = $data['createdAt'] ?? null;
        $this->container['id'] = $data['id'] ?? null;
        $this->container['recipient'] = $data['recipient'] ?? null;
        $this->container['sendCount'] = $data['sendCount'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['subject'] = $data['subject'] ?? null;
        $this->container['templateType'] = $data['templateType'] ?? null;
        $this->container['type'] = $data['type'] ?? null;
        $this->container['updatedAt'] = $data['updatedAt'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        $allowedValues = $this->getTemplateTypeAllowableValues();
        if (!is_null($this->container['templateType']) && !in_array($this->container['templateType'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'templateType', must be one of '%s'",
                $this->container['templateType'],
                implode("', '", $allowedValues)
            );
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
     * Gets body
     *
     * @return string|null
     */
    public function getBody()
    {
        return $this->container['body'];
    }

    /**
     * Sets body
     *
     * @param string|null $body body
     *
     * @return self
     */
    public function setBody($body)
    {
        $this->container['body'] = $body;

        return $this;
    }

    /**
     * Gets createdAt
     *
     * @return \DateTime|null
     */
    public function getCreatedAt()
    {
        return $this->container['createdAt'];
    }

    /**
     * Sets createdAt
     *
     * @param \DateTime|null $createdAt CreatedAt is a helper struct field for gobuffalo.pop.
     *
     * @return self
     */
    public function setCreatedAt($createdAt)
    {
        $this->container['createdAt'] = $createdAt;

        return $this;
    }

    /**
     * Gets id
     *
     * @return string|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param string|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets recipient
     *
     * @return string|null
     */
    public function getRecipient()
    {
        return $this->container['recipient'];
    }

    /**
     * Sets recipient
     *
     * @param string|null $recipient recipient
     *
     * @return self
     */
    public function setRecipient($recipient)
    {
        $this->container['recipient'] = $recipient;

        return $this;
    }

    /**
     * Gets sendCount
     *
     * @return int|null
     */
    public function getSendCount()
    {
        return $this->container['sendCount'];
    }

    /**
     * Sets sendCount
     *
     * @param int|null $sendCount sendCount
     *
     * @return self
     */
    public function setSendCount($sendCount)
    {
        $this->container['sendCount'] = $sendCount;

        return $this;
    }

    /**
     * Gets status
     *
     * @return \Ory\Client\Model\CourierMessageStatus|null
     */
    public function getStatus()
    {
        return $this->container['status'];
    }

    /**
     * Sets status
     *
     * @param \Ory\Client\Model\CourierMessageStatus|null $status status
     *
     * @return self
     */
    public function setStatus($status)
    {
        $this->container['status'] = $status;

        return $this;
    }

    /**
     * Gets subject
     *
     * @return string|null
     */
    public function getSubject()
    {
        return $this->container['subject'];
    }

    /**
     * Sets subject
     *
     * @param string|null $subject subject
     *
     * @return self
     */
    public function setSubject($subject)
    {
        $this->container['subject'] = $subject;

        return $this;
    }

    /**
     * Gets templateType
     *
     * @return string|null
     */
    public function getTemplateType()
    {
        return $this->container['templateType'];
    }

    /**
     * Sets templateType
     *
     * @param string|null $templateType templateType
     *
     * @return self
     */
    public function setTemplateType($templateType)
    {
        $allowedValues = $this->getTemplateTypeAllowableValues();
        if (!is_null($templateType) && !in_array($templateType, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'templateType', must be one of '%s'",
                    $templateType,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['templateType'] = $templateType;

        return $this;
    }

    /**
     * Gets type
     *
     * @return \Ory\Client\Model\CourierMessageType|null
     */
    public function getType()
    {
        return $this->container['type'];
    }

    /**
     * Sets type
     *
     * @param \Ory\Client\Model\CourierMessageType|null $type type
     *
     * @return self
     */
    public function setType($type)
    {
        $this->container['type'] = $type;

        return $this;
    }

    /**
     * Gets updatedAt
     *
     * @return \DateTime|null
     */
    public function getUpdatedAt()
    {
        return $this->container['updatedAt'];
    }

    /**
     * Sets updatedAt
     *
     * @param \DateTime|null $updatedAt UpdatedAt is a helper struct field for gobuffalo.pop.
     *
     * @return self
     */
    public function setUpdatedAt($updatedAt)
    {
        $this->container['updatedAt'] = $updatedAt;

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


