/* tslint:disable */
/* eslint-disable */
/**
 * Ory APIs
 * # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 
 *
 * The version of the OpenAPI document: v1.17.1
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * Used when an administrator creates a recovery link for an identity.
 * @export
 * @interface RecoveryLinkForIdentity
 */
export interface RecoveryLinkForIdentity {
    /**
     * Recovery Link Expires At
     * 
     * The timestamp when the recovery link expires.
     * @type {Date}
     * @memberof RecoveryLinkForIdentity
     */
    expires_at?: Date;
    /**
     * Recovery Link
     * 
     * This link can be used to recover the account.
     * @type {string}
     * @memberof RecoveryLinkForIdentity
     */
    recovery_link: string;
}

/**
 * Check if a given object implements the RecoveryLinkForIdentity interface.
 */
export function instanceOfRecoveryLinkForIdentity(value: object): value is RecoveryLinkForIdentity {
    if (!('recovery_link' in value) || value['recovery_link'] === undefined) return false;
    return true;
}

export function RecoveryLinkForIdentityFromJSON(json: any): RecoveryLinkForIdentity {
    return RecoveryLinkForIdentityFromJSONTyped(json, false);
}

export function RecoveryLinkForIdentityFromJSONTyped(json: any, ignoreDiscriminator: boolean): RecoveryLinkForIdentity {
    if (json == null) {
        return json;
    }
    return {
        
        'expires_at': json['expires_at'] == null ? undefined : (new Date(json['expires_at'])),
        'recovery_link': json['recovery_link'],
    };
}

export function RecoveryLinkForIdentityToJSON(value?: RecoveryLinkForIdentity | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'expires_at': value['expires_at'] == null ? undefined : ((value['expires_at']).toISOString()),
        'recovery_link': value['recovery_link'],
    };
}

