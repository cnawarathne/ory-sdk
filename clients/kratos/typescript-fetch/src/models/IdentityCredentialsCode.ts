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
import type { IdentityCredentialsCodeAddress } from './IdentityCredentialsCodeAddress';
import {
    IdentityCredentialsCodeAddressFromJSON,
    IdentityCredentialsCodeAddressFromJSONTyped,
    IdentityCredentialsCodeAddressToJSON,
} from './IdentityCredentialsCodeAddress';

/**
 * CredentialsCode represents a one time login/registration code
 * @export
 * @interface IdentityCredentialsCode
 */
export interface IdentityCredentialsCode {
    /**
     * 
     * @type {Array<IdentityCredentialsCodeAddress>}
     * @memberof IdentityCredentialsCode
     */
    addresses?: Array<IdentityCredentialsCodeAddress>;
}

/**
 * Check if a given object implements the IdentityCredentialsCode interface.
 */
export function instanceOfIdentityCredentialsCode(value: object): value is IdentityCredentialsCode {
    return true;
}

export function IdentityCredentialsCodeFromJSON(json: any): IdentityCredentialsCode {
    return IdentityCredentialsCodeFromJSONTyped(json, false);
}

export function IdentityCredentialsCodeFromJSONTyped(json: any, ignoreDiscriminator: boolean): IdentityCredentialsCode {
    if (json == null) {
        return json;
    }
    return {
        
        'addresses': json['addresses'] == null ? undefined : ((json['addresses'] as Array<any>).map(IdentityCredentialsCodeAddressFromJSON)),
    };
}

export function IdentityCredentialsCodeToJSON(value?: IdentityCredentialsCode | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'addresses': value['addresses'] == null ? undefined : ((value['addresses'] as Array<any>).map(IdentityCredentialsCodeAddressToJSON)),
    };
}

