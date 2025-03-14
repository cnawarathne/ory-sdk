/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.4.0-alpha.0
 * Contact: office@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */

#nullable enable

using System;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Ory.Kratos.Client.Client;

namespace Ory.Kratos.Client.Extensions
{
    /// <summary>
    /// Extension methods for IHostBuilder
    /// </summary>
    public static class IHostBuilderExtensions
    {
        /// <summary>
        /// Add the api to your host builder.
        /// </summary>
        /// <param name="builder"></param>
        /// <param name="options"></param>
        public static IHostBuilder ConfigureApi(this IHostBuilder builder, Action<HostBuilderContext, IServiceCollection, HostConfiguration> options)
        {
            builder.ConfigureServices((context, services) => 
            {
                HostConfiguration config = new HostConfiguration(services);

                options(context, services, config);

                IServiceCollectionExtensions.AddApi(services, config);
            });

            return builder;
        }
    }
}
