/* tslint:disable */
/* eslint-disable */
/**
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.15.4
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
 * @interface IdentityCredentialsPassword
 */
export interface IdentityCredentialsPassword {
    /**
     * HashedPassword is a hash-representation of the password.
     * @type {string}
     * @memberof IdentityCredentialsPassword
     */
    hashed_password?: string;
    /**
     * UsePasswordMigrationHook is set to true if the password should be migrated
     * using the password migration hook. If set, and the HashedPassword is empty, a
     * webhook will be called during login to migrate the password.
     * @type {boolean}
     * @memberof IdentityCredentialsPassword
     */
    use_password_migration_hook?: boolean;
}

/**
 * Check if a given object implements the IdentityCredentialsPassword interface.
 */
export function instanceOfIdentityCredentialsPassword(value: object): value is IdentityCredentialsPassword {
    return true;
}

export function IdentityCredentialsPasswordFromJSON(json: any): IdentityCredentialsPassword {
    return IdentityCredentialsPasswordFromJSONTyped(json, false);
}

export function IdentityCredentialsPasswordFromJSONTyped(json: any, ignoreDiscriminator: boolean): IdentityCredentialsPassword {
    if (json == null) {
        return json;
    }
    return {
        
        'hashed_password': json['hashed_password'] == null ? undefined : json['hashed_password'],
        'use_password_migration_hook': json['use_password_migration_hook'] == null ? undefined : json['use_password_migration_hook'],
    };
}

export function IdentityCredentialsPasswordToJSON(value?: IdentityCredentialsPassword | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'hashed_password': value['hashed_password'],
        'use_password_migration_hook': value['use_password_migration_hook'],
    };
}

