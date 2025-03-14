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

import type { UpdateSettingsFlowWithLookupMethod } from './UpdateSettingsFlowWithLookupMethod';
import {
    instanceOfUpdateSettingsFlowWithLookupMethod,
    UpdateSettingsFlowWithLookupMethodFromJSON,
    UpdateSettingsFlowWithLookupMethodFromJSONTyped,
    UpdateSettingsFlowWithLookupMethodToJSON,
} from './UpdateSettingsFlowWithLookupMethod';
import type { UpdateSettingsFlowWithOidcMethod } from './UpdateSettingsFlowWithOidcMethod';
import {
    instanceOfUpdateSettingsFlowWithOidcMethod,
    UpdateSettingsFlowWithOidcMethodFromJSON,
    UpdateSettingsFlowWithOidcMethodFromJSONTyped,
    UpdateSettingsFlowWithOidcMethodToJSON,
} from './UpdateSettingsFlowWithOidcMethod';
import type { UpdateSettingsFlowWithPasskeyMethod } from './UpdateSettingsFlowWithPasskeyMethod';
import {
    instanceOfUpdateSettingsFlowWithPasskeyMethod,
    UpdateSettingsFlowWithPasskeyMethodFromJSON,
    UpdateSettingsFlowWithPasskeyMethodFromJSONTyped,
    UpdateSettingsFlowWithPasskeyMethodToJSON,
} from './UpdateSettingsFlowWithPasskeyMethod';
import type { UpdateSettingsFlowWithPasswordMethod } from './UpdateSettingsFlowWithPasswordMethod';
import {
    instanceOfUpdateSettingsFlowWithPasswordMethod,
    UpdateSettingsFlowWithPasswordMethodFromJSON,
    UpdateSettingsFlowWithPasswordMethodFromJSONTyped,
    UpdateSettingsFlowWithPasswordMethodToJSON,
} from './UpdateSettingsFlowWithPasswordMethod';
import type { UpdateSettingsFlowWithProfileMethod } from './UpdateSettingsFlowWithProfileMethod';
import {
    instanceOfUpdateSettingsFlowWithProfileMethod,
    UpdateSettingsFlowWithProfileMethodFromJSON,
    UpdateSettingsFlowWithProfileMethodFromJSONTyped,
    UpdateSettingsFlowWithProfileMethodToJSON,
} from './UpdateSettingsFlowWithProfileMethod';
import type { UpdateSettingsFlowWithTotpMethod } from './UpdateSettingsFlowWithTotpMethod';
import {
    instanceOfUpdateSettingsFlowWithTotpMethod,
    UpdateSettingsFlowWithTotpMethodFromJSON,
    UpdateSettingsFlowWithTotpMethodFromJSONTyped,
    UpdateSettingsFlowWithTotpMethodToJSON,
} from './UpdateSettingsFlowWithTotpMethod';
import type { UpdateSettingsFlowWithWebAuthnMethod } from './UpdateSettingsFlowWithWebAuthnMethod';
import {
    instanceOfUpdateSettingsFlowWithWebAuthnMethod,
    UpdateSettingsFlowWithWebAuthnMethodFromJSON,
    UpdateSettingsFlowWithWebAuthnMethodFromJSONTyped,
    UpdateSettingsFlowWithWebAuthnMethodToJSON,
} from './UpdateSettingsFlowWithWebAuthnMethod';

/**
 * @type UpdateSettingsFlowBody
 * Update Settings Flow Request Body
 * @export
 */
export type UpdateSettingsFlowBody = { method: 'lookup_secret' } & UpdateSettingsFlowWithLookupMethod | { method: 'oidc' } & UpdateSettingsFlowWithOidcMethod | { method: 'passkey' } & UpdateSettingsFlowWithPasskeyMethod | { method: 'password' } & UpdateSettingsFlowWithPasswordMethod | { method: 'profile' } & UpdateSettingsFlowWithProfileMethod | { method: 'totp' } & UpdateSettingsFlowWithTotpMethod | { method: 'webauthn' } & UpdateSettingsFlowWithWebAuthnMethod;

export function UpdateSettingsFlowBodyFromJSON(json: any): UpdateSettingsFlowBody {
    return UpdateSettingsFlowBodyFromJSONTyped(json, false);
}

export function UpdateSettingsFlowBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateSettingsFlowBody {
    if (json == null) {
        return json;
    }
    switch (json['method']) {
        case 'lookup_secret':
            return Object.assign({}, UpdateSettingsFlowWithLookupMethodFromJSONTyped(json, true), { method: 'lookup_secret' } as const);
        case 'oidc':
            return Object.assign({}, UpdateSettingsFlowWithOidcMethodFromJSONTyped(json, true), { method: 'oidc' } as const);
        case 'passkey':
            return Object.assign({}, UpdateSettingsFlowWithPasskeyMethodFromJSONTyped(json, true), { method: 'passkey' } as const);
        case 'password':
            return Object.assign({}, UpdateSettingsFlowWithPasswordMethodFromJSONTyped(json, true), { method: 'password' } as const);
        case 'profile':
            return Object.assign({}, UpdateSettingsFlowWithProfileMethodFromJSONTyped(json, true), { method: 'profile' } as const);
        case 'totp':
            return Object.assign({}, UpdateSettingsFlowWithTotpMethodFromJSONTyped(json, true), { method: 'totp' } as const);
        case 'webauthn':
            return Object.assign({}, UpdateSettingsFlowWithWebAuthnMethodFromJSONTyped(json, true), { method: 'webauthn' } as const);
        default:
            throw new Error(`No variant of UpdateSettingsFlowBody exists with 'method=${json['method']}'`);
    }
}

export function UpdateSettingsFlowBodyToJSON(json: any): any {
    return UpdateSettingsFlowBodyToJSONTyped(json, false);
}

export function UpdateSettingsFlowBodyToJSONTyped(value?: UpdateSettingsFlowBody | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }
    switch (value['method']) {
        case 'lookup_secret':
            return Object.assign({}, UpdateSettingsFlowWithLookupMethodToJSON(value), { method: 'lookup_secret' } as const);
        case 'oidc':
            return Object.assign({}, UpdateSettingsFlowWithOidcMethodToJSON(value), { method: 'oidc' } as const);
        case 'passkey':
            return Object.assign({}, UpdateSettingsFlowWithPasskeyMethodToJSON(value), { method: 'passkey' } as const);
        case 'password':
            return Object.assign({}, UpdateSettingsFlowWithPasswordMethodToJSON(value), { method: 'password' } as const);
        case 'profile':
            return Object.assign({}, UpdateSettingsFlowWithProfileMethodToJSON(value), { method: 'profile' } as const);
        case 'totp':
            return Object.assign({}, UpdateSettingsFlowWithTotpMethodToJSON(value), { method: 'totp' } as const);
        case 'webauthn':
            return Object.assign({}, UpdateSettingsFlowWithWebAuthnMethodToJSON(value), { method: 'webauthn' } as const);
        default:
            throw new Error(`No variant of UpdateSettingsFlowBody exists with 'method=${value['method']}'`);
    }

}

