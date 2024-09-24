/* tslint:disable */
/* eslint-disable */
/**
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.15.4
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface AttributeFilter
 */
export interface AttributeFilter {
    /**
     * 
     * @type {string}
     * @memberof AttributeFilter
     */
    attribute?: string;
    /**
     * 
     * @type {string}
     * @memberof AttributeFilter
     */
    condition?: AttributeFilterConditionEnum;
    /**
     * 
     * @type {string}
     * @memberof AttributeFilter
     */
    value?: string;
}


/**
 * @export
 */
export const AttributeFilterConditionEnum = {
    Equals: 'equals',
    NotEquals: 'not_equals',
    Contains: 'contains',
    NotContains: 'not_contains',
    Regex: 'regex',
    NotRegex: 'not_regex',
    Set: 'set',
    NotSet: 'not_set'
} as const;
export type AttributeFilterConditionEnum = typeof AttributeFilterConditionEnum[keyof typeof AttributeFilterConditionEnum];


/**
 * Check if a given object implements the AttributeFilter interface.
 */
export function instanceOfAttributeFilter(value: object): value is AttributeFilter {
    return true;
}

export function AttributeFilterFromJSON(json: any): AttributeFilter {
    return AttributeFilterFromJSONTyped(json, false);
}

export function AttributeFilterFromJSONTyped(json: any, ignoreDiscriminator: boolean): AttributeFilter {
    if (json == null) {
        return json;
    }
    return {
        
        'attribute': json['attribute'] == null ? undefined : json['attribute'],
        'condition': json['condition'] == null ? undefined : json['condition'],
        'value': json['value'] == null ? undefined : json['value'],
    };
}

export function AttributeFilterToJSON(value?: AttributeFilter | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'attribute': value['attribute'],
        'condition': value['condition'],
        'value': value['value'],
    };
}

