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
 * @interface OAuth2ConsentSessionExpiresAt
 */
export interface OAuth2ConsentSessionExpiresAt {
    /**
     * 
     * @type {Date}
     * @memberof OAuth2ConsentSessionExpiresAt
     */
    access_token?: Date;
    /**
     * 
     * @type {Date}
     * @memberof OAuth2ConsentSessionExpiresAt
     */
    authorize_code?: Date;
    /**
     * 
     * @type {Date}
     * @memberof OAuth2ConsentSessionExpiresAt
     */
    id_token?: Date;
    /**
     * 
     * @type {Date}
     * @memberof OAuth2ConsentSessionExpiresAt
     */
    par_context?: Date;
    /**
     * 
     * @type {Date}
     * @memberof OAuth2ConsentSessionExpiresAt
     */
    refresh_token?: Date;
}

/**
 * Check if a given object implements the OAuth2ConsentSessionExpiresAt interface.
 */
export function instanceOfOAuth2ConsentSessionExpiresAt(value: object): value is OAuth2ConsentSessionExpiresAt {
    return true;
}

export function OAuth2ConsentSessionExpiresAtFromJSON(json: any): OAuth2ConsentSessionExpiresAt {
    return OAuth2ConsentSessionExpiresAtFromJSONTyped(json, false);
}

export function OAuth2ConsentSessionExpiresAtFromJSONTyped(json: any, ignoreDiscriminator: boolean): OAuth2ConsentSessionExpiresAt {
    if (json == null) {
        return json;
    }
    return {
        
        'access_token': json['access_token'] == null ? undefined : (new Date(json['access_token'])),
        'authorize_code': json['authorize_code'] == null ? undefined : (new Date(json['authorize_code'])),
        'id_token': json['id_token'] == null ? undefined : (new Date(json['id_token'])),
        'par_context': json['par_context'] == null ? undefined : (new Date(json['par_context'])),
        'refresh_token': json['refresh_token'] == null ? undefined : (new Date(json['refresh_token'])),
    };
}

export function OAuth2ConsentSessionExpiresAtToJSON(value?: OAuth2ConsentSessionExpiresAt | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'access_token': value['access_token'] == null ? undefined : ((value['access_token']).toISOString()),
        'authorize_code': value['authorize_code'] == null ? undefined : ((value['authorize_code']).toISOString()),
        'id_token': value['id_token'] == null ? undefined : ((value['id_token']).toISOString()),
        'par_context': value['par_context'] == null ? undefined : ((value['par_context']).toISOString()),
        'refresh_token': value['refresh_token'] == null ? undefined : ((value['refresh_token']).toISOString()),
    };
}

