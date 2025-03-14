/* tslint:disable */
/* eslint-disable */
/**
 * Ory Identities API
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.4.0-alpha.0
 * Contact: office@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * Update Recovery Flow with Link Method
 * @export
 * @interface UpdateRecoveryFlowWithLinkMethod
 */
export interface UpdateRecoveryFlowWithLinkMethod {
    /**
     * Sending the anti-csrf token is only required for browser login flows.
     * @type {string}
     * @memberof UpdateRecoveryFlowWithLinkMethod
     */
    csrf_token?: string;
    /**
     * Email to Recover
     * 
     * Needs to be set when initiating the flow. If the email is a registered
     * recovery email, a recovery link will be sent. If the email is not known,
     * a email with details on what happened will be sent instead.
     * 
     * format: email
     * @type {string}
     * @memberof UpdateRecoveryFlowWithLinkMethod
     */
    email: string;
    /**
     * Method is the method that should be used for this recovery flow
     * 
     * Allowed values are `link` and `code`
     * link RecoveryStrategyLink
     * code RecoveryStrategyCode
     * @type {string}
     * @memberof UpdateRecoveryFlowWithLinkMethod
     */
    method: UpdateRecoveryFlowWithLinkMethodMethodEnum;
    /**
     * Transient data to pass along to any webhooks
     * @type {object}
     * @memberof UpdateRecoveryFlowWithLinkMethod
     */
    transient_payload?: object;
}


/**
 * @export
 */
export const UpdateRecoveryFlowWithLinkMethodMethodEnum = {
    Link: 'link',
    Code: 'code',
    UnknownDefaultOpenApi: '11184809'
} as const;
export type UpdateRecoveryFlowWithLinkMethodMethodEnum = typeof UpdateRecoveryFlowWithLinkMethodMethodEnum[keyof typeof UpdateRecoveryFlowWithLinkMethodMethodEnum];


/**
 * Check if a given object implements the UpdateRecoveryFlowWithLinkMethod interface.
 */
export function instanceOfUpdateRecoveryFlowWithLinkMethod(value: object): value is UpdateRecoveryFlowWithLinkMethod {
    if (!('email' in value) || value['email'] === undefined) return false;
    if (!('method' in value) || value['method'] === undefined) return false;
    return true;
}

export function UpdateRecoveryFlowWithLinkMethodFromJSON(json: any): UpdateRecoveryFlowWithLinkMethod {
    return UpdateRecoveryFlowWithLinkMethodFromJSONTyped(json, false);
}

export function UpdateRecoveryFlowWithLinkMethodFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateRecoveryFlowWithLinkMethod {
    if (json == null) {
        return json;
    }
    return {
        
        'csrf_token': json['csrf_token'] == null ? undefined : json['csrf_token'],
        'email': json['email'],
        'method': json['method'],
        'transient_payload': json['transient_payload'] == null ? undefined : json['transient_payload'],
    };
}

export function UpdateRecoveryFlowWithLinkMethodToJSON(json: any): UpdateRecoveryFlowWithLinkMethod {
    return UpdateRecoveryFlowWithLinkMethodToJSONTyped(json, false);
}

export function UpdateRecoveryFlowWithLinkMethodToJSONTyped(value?: UpdateRecoveryFlowWithLinkMethod | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'csrf_token': value['csrf_token'],
        'email': value['email'],
        'method': value['method'],
        'transient_payload': value['transient_payload'],
    };
}

