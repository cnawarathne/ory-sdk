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
import type { IdentityWithCredentialsPasswordConfig } from './IdentityWithCredentialsPasswordConfig';
import {
    IdentityWithCredentialsPasswordConfigFromJSON,
    IdentityWithCredentialsPasswordConfigFromJSONTyped,
    IdentityWithCredentialsPasswordConfigToJSON,
} from './IdentityWithCredentialsPasswordConfig';

/**
 * Create Identity and Import Password Credentials
 * @export
 * @interface IdentityWithCredentialsPassword
 */
export interface IdentityWithCredentialsPassword {
    /**
     * 
     * @type {IdentityWithCredentialsPasswordConfig}
     * @memberof IdentityWithCredentialsPassword
     */
    config?: IdentityWithCredentialsPasswordConfig;
}

/**
 * Check if a given object implements the IdentityWithCredentialsPassword interface.
 */
export function instanceOfIdentityWithCredentialsPassword(value: object): value is IdentityWithCredentialsPassword {
    return true;
}

export function IdentityWithCredentialsPasswordFromJSON(json: any): IdentityWithCredentialsPassword {
    return IdentityWithCredentialsPasswordFromJSONTyped(json, false);
}

export function IdentityWithCredentialsPasswordFromJSONTyped(json: any, ignoreDiscriminator: boolean): IdentityWithCredentialsPassword {
    if (json == null) {
        return json;
    }
    return {
        
        'config': json['config'] == null ? undefined : IdentityWithCredentialsPasswordConfigFromJSON(json['config']),
    };
}

export function IdentityWithCredentialsPasswordToJSON(value?: IdentityWithCredentialsPassword | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'config': IdentityWithCredentialsPasswordConfigToJSON(value['config']),
    };
}

