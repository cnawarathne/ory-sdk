/* tslint:disable */
/* eslint-disable */
/**
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.13.6
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import type { ContinueWithRecoveryUi } from './ContinueWithRecoveryUi';
import {
    instanceOfContinueWithRecoveryUi,
    ContinueWithRecoveryUiFromJSON,
    ContinueWithRecoveryUiFromJSONTyped,
    ContinueWithRecoveryUiToJSON,
} from './ContinueWithRecoveryUi';
import type { ContinueWithSetOrySessionToken } from './ContinueWithSetOrySessionToken';
import {
    instanceOfContinueWithSetOrySessionToken,
    ContinueWithSetOrySessionTokenFromJSON,
    ContinueWithSetOrySessionTokenFromJSONTyped,
    ContinueWithSetOrySessionTokenToJSON,
} from './ContinueWithSetOrySessionToken';
import type { ContinueWithSettingsUi } from './ContinueWithSettingsUi';
import {
    instanceOfContinueWithSettingsUi,
    ContinueWithSettingsUiFromJSON,
    ContinueWithSettingsUiFromJSONTyped,
    ContinueWithSettingsUiToJSON,
} from './ContinueWithSettingsUi';
import type { ContinueWithVerificationUi } from './ContinueWithVerificationUi';
import {
    instanceOfContinueWithVerificationUi,
    ContinueWithVerificationUiFromJSON,
    ContinueWithVerificationUiFromJSONTyped,
    ContinueWithVerificationUiToJSON,
} from './ContinueWithVerificationUi';

/**
 * @type ContinueWith
 * 
 * @export
 */
export type ContinueWith = { action: 'set_ory_session_token' } & ContinueWithSetOrySessionToken | { action: 'show_recovery_ui' } & ContinueWithRecoveryUi | { action: 'show_settings_ui' } & ContinueWithSettingsUi | { action: 'show_verification_ui' } & ContinueWithVerificationUi;

export function ContinueWithFromJSON(json: any): ContinueWith {
    return ContinueWithFromJSONTyped(json, false);
}

export function ContinueWithFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContinueWith {
    if (json == null) {
        return json;
    }
    switch (json['action']) {
        case 'set_ory_session_token':
            return Object.assign({}, ContinueWithSetOrySessionTokenFromJSONTyped(json, true), { action: 'set_ory_session_token' } as const);
        case 'show_recovery_ui':
            return Object.assign({}, ContinueWithRecoveryUiFromJSONTyped(json, true), { action: 'show_recovery_ui' } as const);
        case 'show_settings_ui':
            return Object.assign({}, ContinueWithSettingsUiFromJSONTyped(json, true), { action: 'show_settings_ui' } as const);
        case 'show_verification_ui':
            return Object.assign({}, ContinueWithVerificationUiFromJSONTyped(json, true), { action: 'show_verification_ui' } as const);
        default:
            throw new Error(`No variant of ContinueWith exists with 'action=${json['action']}'`);
    }
}

export function ContinueWithToJSON(value?: ContinueWith | null): any {
    if (value == null) {
        return value;
    }
    switch (value['action']) {
        case 'set_ory_session_token':
            return ContinueWithSetOrySessionTokenToJSON(value);
        case 'show_recovery_ui':
            return ContinueWithRecoveryUiToJSON(value);
        case 'show_settings_ui':
            return ContinueWithSettingsUiToJSON(value);
        case 'show_verification_ui':
            return ContinueWithVerificationUiToJSON(value);
        default:
            throw new Error(`No variant of ContinueWith exists with 'action=${value['action']}'`);
    }

}

