/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.13.6
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
    ///  Class for testing ClientListInvoicesResponse
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the model.
    /// </remarks>
    public class ClientListInvoicesResponseTests : IDisposable
    {
        // TODO uncomment below to declare an instance variable for ClientListInvoicesResponse
        //private ClientListInvoicesResponse instance;

        public ClientListInvoicesResponseTests()
        {
            // TODO uncomment below to create an instance of ClientListInvoicesResponse
            //instance = new ClientListInvoicesResponse();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of ClientListInvoicesResponse
        /// </summary>
        [Fact]
        public void ClientListInvoicesResponseInstanceTest()
        {
            // TODO uncomment below to test "IsType" ClientListInvoicesResponse
            //Assert.IsType<ClientListInvoicesResponse>(instance);
        }

        /// <summary>
        /// Test the property 'Buckets'
        /// </summary>
        [Fact]
        public void BucketsTest()
        {
            // TODO unit test for the property 'Buckets'
        }

        /// <summary>
        /// Test the property 'HasNextPage'
        /// </summary>
        [Fact]
        public void HasNextPageTest()
        {
            // TODO unit test for the property 'HasNextPage'
        }

        /// <summary>
        /// Test the property 'NextPageToken'
        /// </summary>
        [Fact]
        public void NextPageTokenTest()
        {
            // TODO unit test for the property 'NextPageToken'
        }
    }
}
