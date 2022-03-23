=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v0.0.1-alpha.145
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.4.0

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryClient::UpdateCustomHostnameBody
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OryClient::UpdateCustomHostnameBody do
  let(:instance) { OryClient::UpdateCustomHostnameBody.new }

  describe 'test an instance of UpdateCustomHostnameBody' do
    it 'should create an instance of UpdateCustomHostnameBody' do
      expect(instance).to be_instance_of(OryClient::UpdateCustomHostnameBody)
    end
  end
  describe 'test attribute "cookie_domain"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "hostname"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
