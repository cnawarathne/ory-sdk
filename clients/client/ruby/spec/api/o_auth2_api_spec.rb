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

# Unit tests for OryClient::OAuth2Api
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'OAuth2Api' do
  before do
    # run before each test
    @api_instance = OryClient::OAuth2Api.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of OAuth2Api' do
    it 'should create an instance of OAuth2Api' do
      expect(@api_instance).to be_instance_of(OryClient::OAuth2Api)
    end
  end

  # unit tests for accept_o_auth2_consent_request
  # Accept OAuth 2.0 Consent Request
  # When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell Ory now about it. If the subject authenticated, he/she must now be asked if the OAuth 2.0 Client which initiated the flow should be allowed to access the resources on the subject&#39;s behalf.  The consent challenge is appended to the consent provider&#39;s URL to which the subject&#39;s user-agent (browser) is redirected to. The consent provider uses that challenge to fetch information on the OAuth2 request and then tells Ory if the subject accepted or rejected the request.  This endpoint tells Ory that the subject has authorized the OAuth 2.0 client to access resources on his/her behalf. The consent provider includes additional information, such as session data for access and ID tokens, and if the consent request should be used as basis for future requests.  The response contains a redirect URL which the consent provider should redirect the user-agent to.  The default consent provider is available via the Ory Managed Account Experience. To customize the consent provider, please head over to the OAuth 2.0 documentation.
  # @param consent_challenge OAuth 2.0 Consent Request Challenge
  # @param [Hash] opts the optional parameters
  # @option opts [AcceptOAuth2ConsentRequest] :accept_o_auth2_consent_request 
  # @return [OAuth2RedirectTo]
  describe 'accept_o_auth2_consent_request test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for accept_o_auth2_login_request
  # Accept OAuth 2.0 Login Request
  # When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell the Ory OAuth2 Service about it.  The authentication challenge is appended to the login provider URL to which the subject&#39;s user-agent (browser) is redirected to. The login provider uses that challenge to fetch information on the OAuth2 request and then accept or reject the requested authentication process.  This endpoint tells Ory that the subject has successfully authenticated and includes additional information such as the subject&#39;s ID and if Ory should remember the subject&#39;s subject agent for future authentication attempts by setting a cookie.  The response contains a redirect URL which the login provider should redirect the user-agent to.
  # @param login_challenge OAuth 2.0 Login Request Challenge
  # @param [Hash] opts the optional parameters
  # @option opts [AcceptOAuth2LoginRequest] :accept_o_auth2_login_request 
  # @return [OAuth2RedirectTo]
  describe 'accept_o_auth2_login_request test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for accept_o_auth2_logout_request
  # Accept OAuth 2.0 Session Logout Request
  # When a user or an application requests Ory OAuth 2.0 to remove the session state of a subject, this endpoint is used to confirm that logout request.  The response contains a redirect URL which the consent provider should redirect the user-agent to.
  # @param logout_challenge OAuth 2.0 Logout Request Challenge
  # @param [Hash] opts the optional parameters
  # @return [OAuth2RedirectTo]
  describe 'accept_o_auth2_logout_request test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for create_o_auth2_client
  # Create OAuth 2.0 Client
  # Create a new OAuth 2.0 client. If you pass &#x60;client_secret&#x60; the secret is used, otherwise a random secret is generated. The secret is echoed in the response. It is not possible to retrieve it later on.
  # @param o_auth2_client OAuth 2.0 Client Request Body
  # @param [Hash] opts the optional parameters
  # @return [OAuth2Client]
  describe 'create_o_auth2_client test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for delete_o_auth2_client
  # Delete OAuth 2.0 Client
  # Delete an existing OAuth 2.0 Client by its ID.  OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.  Make sure that this endpoint is well protected and only callable by first-party components.
  # @param id The id of the OAuth 2.0 Client.
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'delete_o_auth2_client test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for delete_o_auth2_token
  # Delete OAuth 2.0 Access Tokens from specific OAuth 2.0 Client
  # This endpoint deletes OAuth2 access tokens issued to an OAuth 2.0 Client from the database.
  # @param client_id OAuth 2.0 Client ID
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'delete_o_auth2_token test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for delete_trusted_o_auth2_jwt_grant_issuer
  # Delete Trusted OAuth2 JWT Bearer Grant Type Issuer
  # Use this endpoint to delete trusted JWT Bearer Grant Type Issuer. The ID is the one returned when you created the trust relationship.  Once deleted, the associated issuer will no longer be able to perform the JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grant.
  # @param id The id of the desired grant
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'delete_trusted_o_auth2_jwt_grant_issuer test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_o_auth2_client
  # Get an OAuth 2.0 Client
  # Get an OAuth 2.0 client by its ID. This endpoint never returns the client secret.  OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
  # @param id The id of the OAuth 2.0 Client.
  # @param [Hash] opts the optional parameters
  # @return [OAuth2Client]
  describe 'get_o_auth2_client test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_o_auth2_consent_request
  # Get OAuth 2.0 Consent Request
  # When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell Ory now about it. If the subject authenticated, he/she must now be asked if the OAuth 2.0 Client which initiated the flow should be allowed to access the resources on the subject&#39;s behalf.  The consent challenge is appended to the consent provider&#39;s URL to which the subject&#39;s user-agent (browser) is redirected to. The consent provider uses that challenge to fetch information on the OAuth2 request and then tells Ory if the subject accepted or rejected the request.  The default consent provider is available via the Ory Managed Account Experience. To customize the consent provider, please head over to the OAuth 2.0 documentation.
  # @param consent_challenge OAuth 2.0 Consent Request Challenge
  # @param [Hash] opts the optional parameters
  # @return [OAuth2ConsentRequest]
  describe 'get_o_auth2_consent_request test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_o_auth2_login_request
  # Get OAuth 2.0 Login Request
  # When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell the Ory OAuth2 Service about it.  Per default, the login provider is Ory itself. You may use a different login provider which needs to be a web-app you write and host, and it must be able to authenticate (\&quot;show the subject a login screen\&quot;) a subject (in OAuth2 the proper name for subject is \&quot;resource owner\&quot;).  The authentication challenge is appended to the login provider URL to which the subject&#39;s user-agent (browser) is redirected to. The login provider uses that challenge to fetch information on the OAuth2 request and then accept or reject the requested authentication process.
  # @param login_challenge OAuth 2.0 Login Request Challenge
  # @param [Hash] opts the optional parameters
  # @return [OAuth2LoginRequest]
  describe 'get_o_auth2_login_request test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_o_auth2_logout_request
  # Get OAuth 2.0 Session Logout Request
  # Use this endpoint to fetch an Ory OAuth 2.0 logout request.
  # @param logout_challenge 
  # @param [Hash] opts the optional parameters
  # @return [OAuth2LogoutRequest]
  describe 'get_o_auth2_logout_request test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_trusted_o_auth2_jwt_grant_issuer
  # Get Trusted OAuth2 JWT Bearer Grant Type Issuer
  # Use this endpoint to get a trusted JWT Bearer Grant Type Issuer. The ID is the one returned when you created the trust relationship.
  # @param id The id of the desired grant
  # @param [Hash] opts the optional parameters
  # @return [TrustedOAuth2JwtGrantIssuer]
  describe 'get_trusted_o_auth2_jwt_grant_issuer test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for introspect_o_auth2_token
  # Introspect OAuth2 Access and Refresh Tokens
  # The introspection endpoint allows to check if a token (both refresh and access) is active or not. An active token is neither expired nor revoked. If a token is active, additional information on the token will be included. You can set additional data for a token by setting &#x60;session.access_token&#x60; during the consent flow.
  # @param token The string value of the token. For access tokens, this is the \\\&quot;access_token\\\&quot; value returned from the token endpoint defined in OAuth 2.0. For refresh tokens, this is the \\\&quot;refresh_token\\\&quot; value returned.
  # @param [Hash] opts the optional parameters
  # @option opts [String] :scope An optional, space separated list of required scopes. If the access token was not granted one of the scopes, the result of active will be false.
  # @return [IntrospectedOAuth2Token]
  describe 'introspect_o_auth2_token test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for list_o_auth2_clients
  # List OAuth 2.0 Clients
  # This endpoint lists all clients in the database, and never returns client secrets. As a default it lists the first 100 clients.
  # @param [Hash] opts the optional parameters
  # @option opts [Integer] :page_size Items per Page  This is the number of items per page to return. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
  # @option opts [String] :page_token Next Page Token  The next page token. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
  # @option opts [String] :client_name The name of the clients to filter by.
  # @option opts [String] :owner The owner of the clients to filter by.
  # @return [Array<OAuth2Client>]
  describe 'list_o_auth2_clients test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for list_o_auth2_consent_sessions
  # List OAuth 2.0 Consent Sessions of a Subject
  # This endpoint lists all subject&#39;s granted consent sessions, including client and granted scope. If the subject is unknown or has not granted any consent sessions yet, the endpoint returns an empty JSON array with status code 200 OK.
  # @param subject The subject to list the consent sessions for.
  # @param [Hash] opts the optional parameters
  # @option opts [Integer] :page_size Items per Page  This is the number of items per page to return. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
  # @option opts [String] :page_token Next Page Token  The next page token. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
  # @option opts [String] :login_session_id The login session id to list the consent sessions for.
  # @return [Array<OAuth2ConsentSession>]
  describe 'list_o_auth2_consent_sessions test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for list_trusted_o_auth2_jwt_grant_issuers
  # List Trusted OAuth2 JWT Bearer Grant Type Issuers
  # Use this endpoint to list all trusted JWT Bearer Grant Type Issuers.
  # @param [Hash] opts the optional parameters
  # @option opts [Integer] :max_items 
  # @option opts [Integer] :default_items 
  # @option opts [String] :issuer If optional \&quot;issuer\&quot; is supplied, only jwt-bearer grants with this issuer will be returned.
  # @return [Array<TrustedOAuth2JwtGrantIssuer>]
  describe 'list_trusted_o_auth2_jwt_grant_issuers test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for o_auth2_authorize
  # OAuth 2.0 Authorize Endpoint
  # Use open source libraries to perform OAuth 2.0 and OpenID Connect available for any programming language. You can find a list of libraries at https://oauth.net/code/  This endpoint should not be used via the Ory SDK and is only included for technical reasons. Instead, use one of the libraries linked above.
  # @param [Hash] opts the optional parameters
  # @return [ErrorOAuth2]
  describe 'o_auth2_authorize test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for oauth2_token_exchange
  # The OAuth 2.0 Token Endpoint
  # Use open source libraries to perform OAuth 2.0 and OpenID Connect available for any programming language. You can find a list of libraries here https://oauth.net/code/  This endpoint should not be used via the Ory SDK and is only included for technical reasons. Instead, use one of the libraries linked above.
  # @param grant_type 
  # @param [Hash] opts the optional parameters
  # @option opts [String] :client_id 
  # @option opts [String] :code 
  # @option opts [String] :redirect_uri 
  # @option opts [String] :refresh_token 
  # @return [OAuth2TokenExchange]
  describe 'oauth2_token_exchange test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for patch_o_auth2_client
  # Patch OAuth 2.0 Client
  # Patch an existing OAuth 2.0 Client using JSON Patch. If you pass &#x60;client_secret&#x60; the secret will be updated and returned via the API. This is the only time you will be able to retrieve the client secret, so write it down and keep it safe.  OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
  # @param id The id of the OAuth 2.0 Client.
  # @param json_patch OAuth 2.0 Client JSON Patch Body
  # @param [Hash] opts the optional parameters
  # @return [OAuth2Client]
  describe 'patch_o_auth2_client test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for reject_o_auth2_consent_request
  # Reject OAuth 2.0 Consent Request
  # When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell Ory now about it. If the subject authenticated, he/she must now be asked if the OAuth 2.0 Client which initiated the flow should be allowed to access the resources on the subject&#39;s behalf.  The consent challenge is appended to the consent provider&#39;s URL to which the subject&#39;s user-agent (browser) is redirected to. The consent provider uses that challenge to fetch information on the OAuth2 request and then tells Ory if the subject accepted or rejected the request.  This endpoint tells Ory that the subject has not authorized the OAuth 2.0 client to access resources on his/her behalf. The consent provider must include a reason why the consent was not granted.  The response contains a redirect URL which the consent provider should redirect the user-agent to.  The default consent provider is available via the Ory Managed Account Experience. To customize the consent provider, please head over to the OAuth 2.0 documentation.
  # @param consent_challenge OAuth 2.0 Consent Request Challenge
  # @param [Hash] opts the optional parameters
  # @option opts [RejectOAuth2Request] :reject_o_auth2_request 
  # @return [OAuth2RedirectTo]
  describe 'reject_o_auth2_consent_request test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for reject_o_auth2_login_request
  # Reject OAuth 2.0 Login Request
  # When an authorization code, hybrid, or implicit OAuth 2.0 Flow is initiated, Ory asks the login provider to authenticate the subject and then tell the Ory OAuth2 Service about it.  The authentication challenge is appended to the login provider URL to which the subject&#39;s user-agent (browser) is redirected to. The login provider uses that challenge to fetch information on the OAuth2 request and then accept or reject the requested authentication process.  This endpoint tells Ory that the subject has not authenticated and includes a reason why the authentication was denied.  The response contains a redirect URL which the login provider should redirect the user-agent to.
  # @param login_challenge OAuth 2.0 Login Request Challenge
  # @param [Hash] opts the optional parameters
  # @option opts [RejectOAuth2Request] :reject_o_auth2_request 
  # @return [OAuth2RedirectTo]
  describe 'reject_o_auth2_login_request test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for reject_o_auth2_logout_request
  # Reject OAuth 2.0 Session Logout Request
  # When a user or an application requests Ory OAuth 2.0 to remove the session state of a subject, this endpoint is used to deny that logout request. No HTTP request body is required.  The response is empty as the logout provider has to chose what action to perform next.
  # @param logout_challenge 
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'reject_o_auth2_logout_request test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for revoke_o_auth2_consent_sessions
  # Revoke OAuth 2.0 Consent Sessions of a Subject
  # This endpoint revokes a subject&#39;s granted consent sessions and invalidates all associated OAuth 2.0 Access Tokens. You may also only revoke sessions for a specific OAuth 2.0 Client ID.
  # @param [Hash] opts the optional parameters
  # @option opts [String] :subject OAuth 2.0 Consent Subject  The subject whose consent sessions should be deleted.
  # @option opts [String] :client OAuth 2.0 Client ID  If set, deletes only those consent sessions that have been granted to the specified OAuth 2.0 Client ID.
  # @option opts [String] :consent_request_id Consent Request ID  If set, revoke all token chains derived from this particular consent request ID.
  # @option opts [Boolean] :all Revoke All Consent Sessions  If set to &#x60;true&#x60; deletes all consent sessions by the Subject that have been granted.
  # @return [nil]
  describe 'revoke_o_auth2_consent_sessions test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for revoke_o_auth2_login_sessions
  # Revokes OAuth 2.0 Login Sessions by either a Subject or a SessionID
  # This endpoint invalidates authentication sessions. After revoking the authentication session(s), the subject has to re-authenticate at the Ory OAuth2 Provider. This endpoint does not invalidate any tokens.  If you send the subject in a query param, all authentication sessions that belong to that subject are revoked. No OpenID Connect Front- or Back-channel logout is performed in this case.  Alternatively, you can send a SessionID via &#x60;sid&#x60; query param, in which case, only the session that is connected to that SessionID is revoked. OpenID Connect Back-channel logout is performed in this case.  When using Ory for the identity provider, the login provider will also invalidate the session cookie.
  # @param [Hash] opts the optional parameters
  # @option opts [String] :subject OAuth 2.0 Subject  The subject to revoke authentication sessions for.
  # @option opts [String] :sid Login Session ID  The login session to revoke.
  # @return [nil]
  describe 'revoke_o_auth2_login_sessions test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for revoke_o_auth2_token
  # Revoke OAuth 2.0 Access or Refresh Token
  # Revoking a token (both access and refresh) means that the tokens will be invalid. A revoked access token can no longer be used to make access requests, and a revoked refresh token can no longer be used to refresh an access token. Revoking a refresh token also invalidates the access token that was created with it. A token may only be revoked by the client the token was generated for.
  # @param token 
  # @param [Hash] opts the optional parameters
  # @option opts [String] :client_id 
  # @option opts [String] :client_secret 
  # @return [nil]
  describe 'revoke_o_auth2_token test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for set_o_auth2_client
  # Set OAuth 2.0 Client
  # Replaces an existing OAuth 2.0 Client with the payload you send. If you pass &#x60;client_secret&#x60; the secret is used, otherwise the existing secret is used.  If set, the secret is echoed in the response. It is not possible to retrieve it later on.  OAuth 2.0 Clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
  # @param id OAuth 2.0 Client ID
  # @param o_auth2_client OAuth 2.0 Client Request Body
  # @param [Hash] opts the optional parameters
  # @return [OAuth2Client]
  describe 'set_o_auth2_client test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for set_o_auth2_client_lifespans
  # Set OAuth2 Client Token Lifespans
  # Set lifespans of different token types issued for this OAuth 2.0 client. Does not modify other fields.
  # @param id OAuth 2.0 Client ID
  # @param [Hash] opts the optional parameters
  # @option opts [OAuth2ClientTokenLifespans] :o_auth2_client_token_lifespans 
  # @return [OAuth2Client]
  describe 'set_o_auth2_client_lifespans test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for trust_o_auth2_jwt_grant_issuer
  # Trust OAuth2 JWT Bearer Grant Type Issuer
  # Use this endpoint to establish a trust relationship for a JWT issuer to perform JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants [RFC7523](https://datatracker.ietf.org/doc/html/rfc7523).
  # @param [Hash] opts the optional parameters
  # @option opts [TrustOAuth2JwtGrantIssuer] :trust_o_auth2_jwt_grant_issuer 
  # @return [TrustedOAuth2JwtGrantIssuer]
  describe 'trust_o_auth2_jwt_grant_issuer test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

end
