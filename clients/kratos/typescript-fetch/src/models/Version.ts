/* tslint:disable */
/* eslint-disable */
/**
 * Ory Identities API
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.3.5
 * Contact: office@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface Version
 */
export interface Version {
    /**
     * Version is the service's version.
     * @type {string}
     * @memberof Version
     */
    version?: string;
}

/**
 * Check if a given object implements the Version interface.
 */
export function instanceOfVersion(value: object): value is Version {
    return true;
}

export function VersionFromJSON(json: any): Version {
    return VersionFromJSONTyped(json, false);
}

export function VersionFromJSONTyped(json: any, ignoreDiscriminator: boolean): Version {
    if (json == null) {
        return json;
    }
    return {
        
        'version': json['version'] == null ? undefined : json['version'],
    };
}

export function VersionToJSON(value?: Version | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'version': value['version'],
    };
}

