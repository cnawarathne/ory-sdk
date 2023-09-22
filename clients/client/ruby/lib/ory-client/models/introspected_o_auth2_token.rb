=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v1.2.9
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'date'
require 'time'

module OryClient
  # Introspection contains an access token's session data as specified by [IETF RFC 7662](https://tools.ietf.org/html/rfc7662)
  class IntrospectedOAuth2Token
    # Active is a boolean indicator of whether or not the presented token is currently active.  The specifics of a token's \"active\" state will vary depending on the implementation of the authorization server and the information it keeps about its tokens, but a \"true\" value return for the \"active\" property will generally indicate that a given token has been issued by this authorization server, has not been revoked by the resource owner, and is within its given time window of validity (e.g., after its issuance time and before its expiration time).
    attr_accessor :active

    # Audience contains a list of the token's intended audiences.
    attr_accessor :aud

    # ID is aclient identifier for the OAuth 2.0 client that requested this token.
    attr_accessor :client_id

    # Expires at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token will expire.
    attr_accessor :exp

    # Extra is arbitrary data set by the session.
    attr_accessor :ext

    # Issued at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token was originally issued.
    attr_accessor :iat

    # IssuerURL is a string representing the issuer of this token
    attr_accessor :iss

    # NotBefore is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token is not to be used before.
    attr_accessor :nbf

    # ObfuscatedSubject is set when the subject identifier algorithm was set to \"pairwise\" during authorization. It is the `sub` value of the ID Token that was issued.
    attr_accessor :obfuscated_subject

    # Scope is a JSON string containing a space-separated list of scopes associated with this token.
    attr_accessor :scope

    # Subject of the token, as defined in JWT [RFC7519]. Usually a machine-readable identifier of the resource owner who authorized this token.
    attr_accessor :sub

    # TokenType is the introspected token's type, typically `Bearer`.
    attr_accessor :token_type

    # TokenUse is the introspected token's use, for example `access_token` or `refresh_token`.
    attr_accessor :token_use

    # Username is a human-readable identifier for the resource owner who authorized this token.
    attr_accessor :username

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'active' => :'active',
        :'aud' => :'aud',
        :'client_id' => :'client_id',
        :'exp' => :'exp',
        :'ext' => :'ext',
        :'iat' => :'iat',
        :'iss' => :'iss',
        :'nbf' => :'nbf',
        :'obfuscated_subject' => :'obfuscated_subject',
        :'scope' => :'scope',
        :'sub' => :'sub',
        :'token_type' => :'token_type',
        :'token_use' => :'token_use',
        :'username' => :'username'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'active' => :'Boolean',
        :'aud' => :'Array<String>',
        :'client_id' => :'String',
        :'exp' => :'Integer',
        :'ext' => :'Hash<String, Object>',
        :'iat' => :'Integer',
        :'iss' => :'String',
        :'nbf' => :'Integer',
        :'obfuscated_subject' => :'String',
        :'scope' => :'String',
        :'sub' => :'String',
        :'token_type' => :'String',
        :'token_use' => :'String',
        :'username' => :'String'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `OryClient::IntrospectedOAuth2Token` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `OryClient::IntrospectedOAuth2Token`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'active')
        self.active = attributes[:'active']
      end

      if attributes.key?(:'aud')
        if (value = attributes[:'aud']).is_a?(Array)
          self.aud = value
        end
      end

      if attributes.key?(:'client_id')
        self.client_id = attributes[:'client_id']
      end

      if attributes.key?(:'exp')
        self.exp = attributes[:'exp']
      end

      if attributes.key?(:'ext')
        if (value = attributes[:'ext']).is_a?(Hash)
          self.ext = value
        end
      end

      if attributes.key?(:'iat')
        self.iat = attributes[:'iat']
      end

      if attributes.key?(:'iss')
        self.iss = attributes[:'iss']
      end

      if attributes.key?(:'nbf')
        self.nbf = attributes[:'nbf']
      end

      if attributes.key?(:'obfuscated_subject')
        self.obfuscated_subject = attributes[:'obfuscated_subject']
      end

      if attributes.key?(:'scope')
        self.scope = attributes[:'scope']
      end

      if attributes.key?(:'sub')
        self.sub = attributes[:'sub']
      end

      if attributes.key?(:'token_type')
        self.token_type = attributes[:'token_type']
      end

      if attributes.key?(:'token_use')
        self.token_use = attributes[:'token_use']
      end

      if attributes.key?(:'username')
        self.username = attributes[:'username']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      invalid_properties = Array.new
      if @active.nil?
        invalid_properties.push('invalid value for "active", active cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      return false if @active.nil?
      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          active == o.active &&
          aud == o.aud &&
          client_id == o.client_id &&
          exp == o.exp &&
          ext == o.ext &&
          iat == o.iat &&
          iss == o.iss &&
          nbf == o.nbf &&
          obfuscated_subject == o.obfuscated_subject &&
          scope == o.scope &&
          sub == o.sub &&
          token_type == o.token_type &&
          token_use == o.token_use &&
          username == o.username
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [active, aud, client_id, exp, ext, iat, iss, nbf, obfuscated_subject, scope, sub, token_type, token_use, username].hash
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def self.build_from_hash(attributes)
      new.build_from_hash(attributes)
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def build_from_hash(attributes)
      return nil unless attributes.is_a?(Hash)
      attributes = attributes.transform_keys(&:to_sym)
      self.class.openapi_types.each_pair do |key, type|
        if attributes[self.class.attribute_map[key]].nil? && self.class.openapi_nullable.include?(key)
          self.send("#{key}=", nil)
        elsif type =~ /\AArray<(.*)>/i
          # check to ensure the input is an array given that the attribute
          # is documented as an array but the input is not
          if attributes[self.class.attribute_map[key]].is_a?(Array)
            self.send("#{key}=", attributes[self.class.attribute_map[key]].map { |v| _deserialize($1, v) })
          end
        elsif !attributes[self.class.attribute_map[key]].nil?
          self.send("#{key}=", _deserialize(type, attributes[self.class.attribute_map[key]]))
        end
      end

      self
    end

    # Deserializes the data based on type
    # @param string type Data type
    # @param string value Value to be deserialized
    # @return [Object] Deserialized data
    def _deserialize(type, value)
      case type.to_sym
      when :Time
        Time.parse(value)
      when :Date
        Date.parse(value)
      when :String
        value.to_s
      when :Integer
        value.to_i
      when :Float
        value.to_f
      when :Boolean
        if value.to_s =~ /\A(true|t|yes|y|1)\z/i
          true
        else
          false
        end
      when :Object
        # generic object (usually a Hash), return directly
        value
      when /\AArray<(?<inner_type>.+)>\z/
        inner_type = Regexp.last_match[:inner_type]
        value.map { |v| _deserialize(inner_type, v) }
      when /\AHash<(?<k_type>.+?), (?<v_type>.+)>\z/
        k_type = Regexp.last_match[:k_type]
        v_type = Regexp.last_match[:v_type]
        {}.tap do |hash|
          value.each do |k, v|
            hash[_deserialize(k_type, k)] = _deserialize(v_type, v)
          end
        end
      else # model
        # models (e.g. Pet) or oneOf
        klass = OryClient.const_get(type)
        klass.respond_to?(:openapi_one_of) ? klass.build(value) : klass.build_from_hash(value)
      end
    end

    # Returns the string representation of the object
    # @return [String] String presentation of the object
    def to_s
      to_hash.to_s
    end

    # to_body is an alias to to_hash (backward compatibility)
    # @return [Hash] Returns the object in the form of hash
    def to_body
      to_hash
    end

    # Returns the object in the form of hash
    # @return [Hash] Returns the object in the form of hash
    def to_hash
      hash = {}
      self.class.attribute_map.each_pair do |attr, param|
        value = self.send(attr)
        if value.nil?
          is_nullable = self.class.openapi_nullable.include?(attr)
          next if !is_nullable || (is_nullable && !instance_variable_defined?(:"@#{attr}"))
        end

        hash[param] = _to_hash(value)
      end
      hash
    end

    # Outputs non-array value in the form of hash
    # For object, use to_hash. Otherwise, just return the value
    # @param [Object] value Any valid value
    # @return [Hash] Returns the value in the form of hash
    def _to_hash(value)
      if value.is_a?(Array)
        value.compact.map { |v| _to_hash(v) }
      elsif value.is_a?(Hash)
        {}.tap do |hash|
          value.each { |k, v| hash[k] = _to_hash(v) }
        end
      elsif value.respond_to? :to_hash
        value.to_hash
      else
        value
      end
    end

  end

end
