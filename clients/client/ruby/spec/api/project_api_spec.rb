=begin
#Ory APIs

## Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

The version of the OpenAPI document: v1.17.2
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
Generator version: 7.7.0

=end

require 'spec_helper'
require 'json'

# Unit tests for OryClient::ProjectApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'ProjectApi' do
  before do
    # run before each test
    @api_instance = OryClient::ProjectApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of ProjectApi' do
    it 'should create an instance of ProjectApi' do
      expect(@api_instance).to be_instance_of(OryClient::ProjectApi)
    end
  end

  # unit tests for create_organization
  # Create an Enterprise SSO Organization
  # Creates an Enterprise SSO Organization in a project.
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @option opts [OrganizationBody] :organization_body 
  # @return [Organization]
  describe 'create_organization test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for create_project
  # Create a Project
  # Creates a new project.
  # @param [Hash] opts the optional parameters
  # @option opts [CreateProjectBody] :create_project_body 
  # @return [Project]
  describe 'create_project test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for create_project_api_key
  # Create project API key
  # Create an API key for a project.
  # @param project The Project ID or Project slug
  # @param [Hash] opts the optional parameters
  # @option opts [CreateProjectApiKeyRequest] :create_project_api_key_request 
  # @return [ProjectApiKey]
  describe 'create_project_api_key test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for delete_organization
  # Delete Enterprise SSO Organization
  # Irrecoverably deletes an Enterprise SSO Organization in a project by its ID.
  # @param project_id Project ID  The project&#39;s ID.
  # @param organization_id Organization ID  The Organization&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'delete_organization test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for delete_project_api_key
  # Delete project API key
  # Deletes an API key and immediately removes it.
  # @param project The Project ID or Project slug
  # @param token_id The Token ID
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'delete_project_api_key test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_organization
  # Get Enterprise SSO Organization by ID
  # Retrieves an Enterprise SSO Organization for a project by its ID
  # @param project_id Project ID  The project&#39;s ID.
  # @param organization_id Organization ID  The Organization&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @return [GetOrganizationResponse]
  describe 'get_organization test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_project
  # Get a Project
  # Get a projects you have access to by its ID.
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @return [Project]
  describe 'get_project test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_project_members
  # Get all members associated with this project
  # This endpoint requires the user to be a member of the project with the role &#x60;OWNER&#x60; or &#x60;DEVELOPER&#x60;.
  # @param project 
  # @param [Hash] opts the optional parameters
  # @return [Array<ProjectMember>]
  describe 'get_project_members test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for list_organizations
  # List all Enterprise SSO organizations
  # Lists all Enterprise SSO organizations in a project.
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @option opts [Integer] :page_size Items per Page  This is the number of items per page to return. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
  # @option opts [String] :page_token Next Page Token  The next page token. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
  # @option opts [String] :domain Domain  If set, only organizations with that domain will be returned.
  # @return [ListOrganizationsResponse]
  describe 'list_organizations test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for list_project_api_keys
  # List a project&#39;s API keys
  # A list of all the project&#39;s API keys.
  # @param project The Project ID or Project slug
  # @param [Hash] opts the optional parameters
  # @return [Array<ProjectApiKey>]
  describe 'list_project_api_keys test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for list_projects
  # List All Projects
  # Lists all projects you have access to.
  # @param [Hash] opts the optional parameters
  # @return [Array<ProjectMetadata>]
  describe 'list_projects test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for patch_project
  # Patch an Ory Network Project Configuration
  # Deprecated: Use the &#x60;patchProjectWithRevision&#x60; endpoint instead to specify the exact revision the patch was generated for.  This endpoints allows you to patch individual Ory Network project configuration keys for Ory&#39;s services (identity, permission, ...). The configuration format is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @option opts [Array<JsonPatch>] :json_patch 
  # @return [SuccessfulProjectUpdate]
  describe 'patch_project test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for patch_project_with_revision
  # Patch an Ory Network Project Configuration based on a revision ID
  # This endpoints allows you to patch individual Ory Network Project configuration keys for Ory&#39;s services (identity, permission, ...). The configuration format is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.
  # @param project_id Project ID  The project&#39;s ID.
  # @param revision_id Revision ID  The revision ID that this patch was generated for.
  # @param [Hash] opts the optional parameters
  # @option opts [Array<JsonPatch>] :json_patch 
  # @return [SuccessfulProjectUpdate]
  describe 'patch_project_with_revision test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for purge_project
  # Irrecoverably purge a project
  # !! Use with extreme caution !!  Using this API endpoint you can purge (completely delete) a project and its data. This action can not be undone and will delete ALL your data.  Calling this endpoint will additionally delete custom domains and other related data.  If the project is linked to a subscription, the subscription needs to be unlinked first.
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'purge_project test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for remove_project_member
  # Remove a member associated with this project
  # This also sets their invite status to &#x60;REMOVED&#x60;. This endpoint requires the user to be a member of the project with the role &#x60;OWNER&#x60;.
  # @param project 
  # @param member 
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'remove_project_member test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for set_project
  # Update an Ory Network Project Configuration
  # This endpoints allows you to update the Ory Network project configuration for individual services (identity, permission, ...). The configuration is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.  Be aware that updating any service&#39;s configuration will completely override your current configuration for that service!
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @option opts [SetProject] :set_project 
  # @return [SuccessfulProjectUpdate]
  describe 'set_project test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for update_organization
  # Update an Enterprise SSO Organization
  # Updates an Enterprise SSO Organization in a project by its ID.
  # @param project_id Project ID  The project&#39;s ID.
  # @param organization_id Organization ID  The Organization&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @option opts [OrganizationBody] :organization_body 
  # @return [Organization]
  describe 'update_organization test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

end
