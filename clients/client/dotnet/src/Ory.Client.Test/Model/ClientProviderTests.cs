/*
 * Ory APIs
 *
 * # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | - -- -- -- -- -- -- - | - -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- - | - -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- - | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 
 *
 * The version of the OpenAPI document: v1.17.1
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using Xunit;

using System;
using System.Linq;
using System.IO;
using System.Collections.Generic;
using Ory.Client.Model;
using Ory.Client.Client;
using System.Reflection;
using Newtonsoft.Json;

namespace Ory.Client.Test.Model
{
    /// <summary>
    ///  Class for testing ClientProvider
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class ClientProviderTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for ClientProvider
        //private ClientProvider instance;

        public ClientProviderTests()
        {
            // TODO uncomment below to create an instance of ClientProvider
            //instance = new ClientProvider();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of ClientProvider
        /// </summary>
        [Fact]
        public void ClientProviderInstanceTest()
        {
            // TODO uncomment below to test "IsType" ClientProvider
            //Assert.IsType<ClientProvider>(instance);
        }

        /// <summary>
        /// Test the property 'ClientId'
        /// </summary>
        [Fact]
        public void ClientIdTest()
        {
            // TODO unit test for the property 'ClientId'
        }

        /// <summary>
        /// Test the property 'ConfigUrl'
        /// </summary>
        [Fact]
        public void ConfigUrlTest()
        {
            // TODO unit test for the property 'ConfigUrl'
        }

        /// <summary>
        /// Test the property 'DomainHint'
        /// </summary>
        [Fact]
        public void DomainHintTest()
        {
            // TODO unit test for the property 'DomainHint'
        }

        /// <summary>
        /// Test the property 'Fields'
        /// </summary>
        [Fact]
        public void FieldsTest()
        {
            // TODO unit test for the property 'Fields'
        }

        /// <summary>
        /// Test the property 'LoginHint'
        /// </summary>
        [Fact]
        public void LoginHintTest()
        {
            // TODO unit test for the property 'LoginHint'
        }

        /// <summary>
        /// Test the property 'Nonce'
        /// </summary>
        [Fact]
        public void NonceTest()
        {
            // TODO unit test for the property 'Nonce'
        }

        /// <summary>
        /// Test the property 'Parameters'
        /// </summary>
        [Fact]
        public void ParametersTest()
        {
            // TODO unit test for the property 'Parameters'
        }
    }
}
