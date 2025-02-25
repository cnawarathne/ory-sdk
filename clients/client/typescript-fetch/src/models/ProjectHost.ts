/* tslint:disable */
/* eslint-disable */
/**
 * Ory APIs
 * # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 
 *
 * The version of the OpenAPI document: v1.17.2
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ProjectHost
 */
export interface ProjectHost {
    /**
     * The project's host.
     * @type {string}
     * @memberof ProjectHost
     */
    host: string;
    /**
     * The mapping's ID.
     * @type {string}
     * @memberof ProjectHost
     */
    readonly id: string;
    /**
     * The Revision's Project ID
     * @type {string}
     * @memberof ProjectHost
     */
    project_id: string;
}

/**
 * Check if a given object implements the ProjectHost interface.
 */
export function instanceOfProjectHost(value: object): value is ProjectHost {
    if (!('host' in value) || value['host'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('project_id' in value) || value['project_id'] === undefined) return false;
    return true;
}

export function ProjectHostFromJSON(json: any): ProjectHost {
    return ProjectHostFromJSONTyped(json, false);
}

export function ProjectHostFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProjectHost {
    if (json == null) {
        return json;
    }
    return {
        
        'host': json['host'],
        'id': json['id'],
        'project_id': json['project_id'],
    };
}

export function ProjectHostToJSON(value?: Omit<ProjectHost, 'id'> | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'host': value['host'],
        'project_id': value['project_id'],
    };
}

