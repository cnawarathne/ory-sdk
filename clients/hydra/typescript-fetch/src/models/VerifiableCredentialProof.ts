/* tslint:disable */
/* eslint-disable */
/**
 * Ory Hydra API
 * Documentation for all of Ory Hydra\'s APIs. 
 *
 * The version of the OpenAPI document: v2.4.0-alpha.1
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface VerifiableCredentialProof
 */
export interface VerifiableCredentialProof {
    /**
     * 
     * @type {string}
     * @memberof VerifiableCredentialProof
     */
    jwt?: string;
    /**
     * 
     * @type {string}
     * @memberof VerifiableCredentialProof
     */
    proof_type?: string;
}

/**
 * Check if a given object implements the VerifiableCredentialProof interface.
 */
export function instanceOfVerifiableCredentialProof(value: object): value is VerifiableCredentialProof {
    return true;
}

export function VerifiableCredentialProofFromJSON(json: any): VerifiableCredentialProof {
    return VerifiableCredentialProofFromJSONTyped(json, false);
}

export function VerifiableCredentialProofFromJSONTyped(json: any, ignoreDiscriminator: boolean): VerifiableCredentialProof {
    if (json == null) {
        return json;
    }
    return {
        
        'jwt': json['jwt'] == null ? undefined : json['jwt'],
        'proof_type': json['proof_type'] == null ? undefined : json['proof_type'],
    };
}

export function VerifiableCredentialProofToJSON(value?: VerifiableCredentialProof | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'jwt': value['jwt'],
        'proof_type': value['proof_type'],
    };
}

