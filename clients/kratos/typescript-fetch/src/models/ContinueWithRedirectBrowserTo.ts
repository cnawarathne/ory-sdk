/* tslint:disable */
/* eslint-disable */
/**
 * Ory Identities API
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.3.8
 * Contact: office@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * Indicates, that the UI flow could be continued by showing a recovery ui
 * @export
 * @interface ContinueWithRedirectBrowserTo
 */
export interface ContinueWithRedirectBrowserTo {
    /**
     * Action will always be `redirect_browser_to`
     * redirect_browser_to ContinueWithActionRedirectBrowserToString
     * @type {string}
     * @memberof ContinueWithRedirectBrowserTo
     */
    action: ContinueWithRedirectBrowserToActionEnum;
    /**
     * The URL to redirect the browser to
     * @type {string}
     * @memberof ContinueWithRedirectBrowserTo
     */
    redirect_browser_to: string;
}


/**
 * @export
 */
export const ContinueWithRedirectBrowserToActionEnum = {
    RedirectBrowserTo: 'redirect_browser_to'
} as const;
export type ContinueWithRedirectBrowserToActionEnum = typeof ContinueWithRedirectBrowserToActionEnum[keyof typeof ContinueWithRedirectBrowserToActionEnum];


/**
 * Check if a given object implements the ContinueWithRedirectBrowserTo interface.
 */
export function instanceOfContinueWithRedirectBrowserTo(value: object): value is ContinueWithRedirectBrowserTo {
    if (!('action' in value) || value['action'] === undefined) return false;
    if (!('redirect_browser_to' in value) || value['redirect_browser_to'] === undefined) return false;
    return true;
}

export function ContinueWithRedirectBrowserToFromJSON(json: any): ContinueWithRedirectBrowserTo {
    return ContinueWithRedirectBrowserToFromJSONTyped(json, false);
}

export function ContinueWithRedirectBrowserToFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContinueWithRedirectBrowserTo {
    if (json == null) {
        return json;
    }
    return {
        
        'action': json['action'],
        'redirect_browser_to': json['redirect_browser_to'],
    };
}

export function ContinueWithRedirectBrowserToToJSON(value?: ContinueWithRedirectBrowserTo | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'action': value['action'],
        'redirect_browser_to': value['redirect_browser_to'],
    };
}

