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
import type { GenericError } from './GenericError';
import {
    GenericErrorFromJSON,
    GenericErrorFromJSONTyped,
    GenericErrorToJSON,
} from './GenericError';

/**
 * The standard Ory JSON API error format.
 * @export
 * @interface ErrorGeneric
 */
export interface ErrorGeneric {
    /**
     * 
     * @type {GenericError}
     * @memberof ErrorGeneric
     */
    error: GenericError;
}

/**
 * Check if a given object implements the ErrorGeneric interface.
 */
export function instanceOfErrorGeneric(value: object): value is ErrorGeneric {
    if (!('error' in value) || value['error'] === undefined) return false;
    return true;
}

export function ErrorGenericFromJSON(json: any): ErrorGeneric {
    return ErrorGenericFromJSONTyped(json, false);
}

export function ErrorGenericFromJSONTyped(json: any, ignoreDiscriminator: boolean): ErrorGeneric {
    if (json == null) {
        return json;
    }
    return {
        
        'error': GenericErrorFromJSON(json['error']),
    };
}

export function ErrorGenericToJSON(value?: ErrorGeneric | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'error': GenericErrorToJSON(value['error']),
    };
}

