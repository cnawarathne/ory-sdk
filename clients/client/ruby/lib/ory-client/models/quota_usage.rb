=begin
#Ory APIs

## Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

The version of the OpenAPI document: v1.15.10
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
Generator version: 7.7.0

=end

require 'date'
require 'time'

module OryClient
  class QuotaUsage
    # The additional price per unit in cents.
    attr_accessor :additional_price

    attr_accessor :can_use_more

    #  production_projects ProductionProjects staging_projects StagingProjects development_projects DevelopmentProjects daily_active_users DailyActiveUsers custom_domains CustomDomains event_streams EventStreams event_stream_events EventStreamEvents sla SLA collaborator_seats CollaboratorSeats edge_cache EdgeCache branding_themes BrandingThemes zendesk_support ZendeskSupport project_metrics ProjectMetrics project_metrics_time_window ProjectMetricsTimeWindow project_metrics_events_history ProjectMetricsEventsHistory organizations Organizations rop_grant ResourceOwnerPasswordGrant concierge_onboarding ConciergeOnboarding credit Credit data_location_global DataLocationGlobal data_location_us DataLocationUS data_location_asiane DataLocationAsiaNorthEast m2m_token_issuance M2MTokenIssuance permission_checks PermissionChecks captcha Captcha data_location_regional DataLocationRegional  Required Features rate_limit_tier RateLimitTier session_rate_limit_tier RateLimitTierSessions identities_list_rate_limit_tier RateLimitTierIdentitiesList permission_checks_rate_limit_tier RateLimitTierPermissionChecks oauth2_introspect_rate_limit_tier RateLimitTierOAuth2Introspect
    attr_accessor :feature

    attr_accessor :feature_available

    attr_accessor :included

    attr_accessor :is_unlimited

    attr_accessor :used

    class EnumAttributeValidator
      attr_reader :datatype
      attr_reader :allowable_values

      def initialize(datatype, allowable_values)
        @allowable_values = allowable_values.map do |value|
          case datatype.to_s
          when /Integer/i
            value.to_i
          when /Float/i
            value.to_f
          else
            value
          end
        end
      end

      def valid?(value)
        !value || allowable_values.include?(value)
      end
    end

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'additional_price' => :'additional_price',
        :'can_use_more' => :'can_use_more',
        :'feature' => :'feature',
        :'feature_available' => :'feature_available',
        :'included' => :'included',
        :'is_unlimited' => :'is_unlimited',
        :'used' => :'used'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'additional_price' => :'String',
        :'can_use_more' => :'Boolean',
        :'feature' => :'String',
        :'feature_available' => :'Boolean',
        :'included' => :'Integer',
        :'is_unlimited' => :'Boolean',
        :'used' => :'Integer'
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
        fail ArgumentError, "The input argument (attributes) must be a hash in `OryClient::QuotaUsage` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `OryClient::QuotaUsage`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'additional_price')
        self.additional_price = attributes[:'additional_price']
      else
        self.additional_price = nil
      end

      if attributes.key?(:'can_use_more')
        self.can_use_more = attributes[:'can_use_more']
      else
        self.can_use_more = nil
      end

      if attributes.key?(:'feature')
        self.feature = attributes[:'feature']
      else
        self.feature = nil
      end

      if attributes.key?(:'feature_available')
        self.feature_available = attributes[:'feature_available']
      else
        self.feature_available = nil
      end

      if attributes.key?(:'included')
        self.included = attributes[:'included']
      else
        self.included = nil
      end

      if attributes.key?(:'is_unlimited')
        self.is_unlimited = attributes[:'is_unlimited']
      else
        self.is_unlimited = nil
      end

      if attributes.key?(:'used')
        self.used = attributes[:'used']
      else
        self.used = nil
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = Array.new
      if @additional_price.nil?
        invalid_properties.push('invalid value for "additional_price", additional_price cannot be nil.')
      end

      if @can_use_more.nil?
        invalid_properties.push('invalid value for "can_use_more", can_use_more cannot be nil.')
      end

      if @feature.nil?
        invalid_properties.push('invalid value for "feature", feature cannot be nil.')
      end

      if @feature_available.nil?
        invalid_properties.push('invalid value for "feature_available", feature_available cannot be nil.')
      end

      if @included.nil?
        invalid_properties.push('invalid value for "included", included cannot be nil.')
      end

      if @is_unlimited.nil?
        invalid_properties.push('invalid value for "is_unlimited", is_unlimited cannot be nil.')
      end

      if @used.nil?
        invalid_properties.push('invalid value for "used", used cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if @additional_price.nil?
      return false if @can_use_more.nil?
      return false if @feature.nil?
      feature_validator = EnumAttributeValidator.new('String', ["production_projects", "staging_projects", "development_projects", "daily_active_users", "custom_domains", "event_streams", "event_stream_events", "sla", "collaborator_seats", "edge_cache", "branding_themes", "zendesk_support", "project_metrics", "project_metrics_time_window", "project_metrics_events_history", "organizations", "rop_grant", "concierge_onboarding", "credit", "data_location_global", "data_location_us", "data_location_asiane", "m2m_token_issuance", "permission_checks", "captcha", "data_location_regional", "rate_limit_tier", "session_rate_limit_tier", "identities_list_rate_limit_tier", "permission_checks_rate_limit_tier", "oauth2_introspect_rate_limit_tier"])
      return false unless feature_validator.valid?(@feature)
      return false if @feature_available.nil?
      return false if @included.nil?
      return false if @is_unlimited.nil?
      return false if @used.nil?
      true
    end

    # Custom attribute writer method checking allowed values (enum).
    # @param [Object] feature Object to be assigned
    def feature=(feature)
      validator = EnumAttributeValidator.new('String', ["production_projects", "staging_projects", "development_projects", "daily_active_users", "custom_domains", "event_streams", "event_stream_events", "sla", "collaborator_seats", "edge_cache", "branding_themes", "zendesk_support", "project_metrics", "project_metrics_time_window", "project_metrics_events_history", "organizations", "rop_grant", "concierge_onboarding", "credit", "data_location_global", "data_location_us", "data_location_asiane", "m2m_token_issuance", "permission_checks", "captcha", "data_location_regional", "rate_limit_tier", "session_rate_limit_tier", "identities_list_rate_limit_tier", "permission_checks_rate_limit_tier", "oauth2_introspect_rate_limit_tier"])
      unless validator.valid?(feature)
        fail ArgumentError, "invalid value for \"feature\", must be one of #{validator.allowable_values}."
      end
      @feature = feature
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          additional_price == o.additional_price &&
          can_use_more == o.can_use_more &&
          feature == o.feature &&
          feature_available == o.feature_available &&
          included == o.included &&
          is_unlimited == o.is_unlimited &&
          used == o.used
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [additional_price, can_use_more, feature, feature_available, included, is_unlimited, used].hash
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def self.build_from_hash(attributes)
      return nil unless attributes.is_a?(Hash)
      attributes = attributes.transform_keys(&:to_sym)
      transformed_hash = {}
      openapi_types.each_pair do |key, type|
        if attributes.key?(attribute_map[key]) && attributes[attribute_map[key]].nil?
          transformed_hash["#{key}"] = nil
        elsif type =~ /\AArray<(.*)>/i
          # check to ensure the input is an array given that the attribute
          # is documented as an array but the input is not
          if attributes[attribute_map[key]].is_a?(Array)
            transformed_hash["#{key}"] = attributes[attribute_map[key]].map { |v| _deserialize($1, v) }
          end
        elsif !attributes[attribute_map[key]].nil?
          transformed_hash["#{key}"] = _deserialize(type, attributes[attribute_map[key]])
        end
      end
      new(transformed_hash)
    end

    # Deserializes the data based on type
    # @param string type Data type
    # @param string value Value to be deserialized
    # @return [Object] Deserialized data
    def self._deserialize(type, value)
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
        klass.respond_to?(:openapi_any_of) || klass.respond_to?(:openapi_one_of) ? klass.build(value) : klass.build_from_hash(value)
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
