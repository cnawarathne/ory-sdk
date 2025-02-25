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
import type { Session } from './Session';
import {
    SessionFromJSON,
    SessionFromJSONTyped,
    SessionToJSON,
} from './Session';

/**
 * The Response for Registration Flows via API
 * @export
 * @interface SuccessfulCodeExchangeResponse
 */
export interface SuccessfulCodeExchangeResponse {
    /**
     * 
     * @type {Session}
     * @memberof SuccessfulCodeExchangeResponse
     */
    session: Session;
    /**
     * The Session Token
     * 
     * A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization
     * Header:
     * 
     * Authorization: bearer ${session-token}
     * 
     * The session token is only issued for API flows, not for Browser flows!
     * @type {string}
     * @memberof SuccessfulCodeExchangeResponse
     */
    session_token?: string;
}

/**
 * Check if a given object implements the SuccessfulCodeExchangeResponse interface.
 */
export function instanceOfSuccessfulCodeExchangeResponse(value: object): value is SuccessfulCodeExchangeResponse {
    if (!('session' in value) || value['session'] === undefined) return false;
    return true;
}

export function SuccessfulCodeExchangeResponseFromJSON(json: any): SuccessfulCodeExchangeResponse {
    return SuccessfulCodeExchangeResponseFromJSONTyped(json, false);
}

export function SuccessfulCodeExchangeResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SuccessfulCodeExchangeResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'session': SessionFromJSON(json['session']),
        'session_token': json['session_token'] == null ? undefined : json['session_token'],
    };
}

export function SuccessfulCodeExchangeResponseToJSON(value?: SuccessfulCodeExchangeResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'session': SessionToJSON(value['session']),
        'session_token': value['session_token'],
    };
}

