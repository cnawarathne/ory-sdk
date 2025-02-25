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
import type { UiText } from './UiText';
import {
    UiTextFromJSON,
    UiTextFromJSONTyped,
    UiTextToJSON,
} from './UiText';

/**
 * 
 * @export
 * @interface UiNodeAnchorAttributes
 */
export interface UiNodeAnchorAttributes {
    /**
     * The link's href (destination) URL.
     * 
     * format: uri
     * @type {string}
     * @memberof UiNodeAnchorAttributes
     */
    href: string;
    /**
     * A unique identifier
     * @type {string}
     * @memberof UiNodeAnchorAttributes
     */
    id: string;
    /**
     * NodeType represents this node's types. It is a mirror of `node.type` and
     * is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is "a".
     * text Text
     * input Input
     * img Image
     * a Anchor
     * script Script
     * @type {string}
     * @memberof UiNodeAnchorAttributes
     */
    node_type: UiNodeAnchorAttributesNodeTypeEnum;
    /**
     * 
     * @type {UiText}
     * @memberof UiNodeAnchorAttributes
     */
    title: UiText;
}


/**
 * @export
 */
export const UiNodeAnchorAttributesNodeTypeEnum = {
    Text: 'text',
    Input: 'input',
    Img: 'img',
    A: 'a',
    Script: 'script'
} as const;
export type UiNodeAnchorAttributesNodeTypeEnum = typeof UiNodeAnchorAttributesNodeTypeEnum[keyof typeof UiNodeAnchorAttributesNodeTypeEnum];


/**
 * Check if a given object implements the UiNodeAnchorAttributes interface.
 */
export function instanceOfUiNodeAnchorAttributes(value: object): value is UiNodeAnchorAttributes {
    if (!('href' in value) || value['href'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('node_type' in value) || value['node_type'] === undefined) return false;
    if (!('title' in value) || value['title'] === undefined) return false;
    return true;
}

export function UiNodeAnchorAttributesFromJSON(json: any): UiNodeAnchorAttributes {
    return UiNodeAnchorAttributesFromJSONTyped(json, false);
}

export function UiNodeAnchorAttributesFromJSONTyped(json: any, ignoreDiscriminator: boolean): UiNodeAnchorAttributes {
    if (json == null) {
        return json;
    }
    return {
        
        'href': json['href'],
        'id': json['id'],
        'node_type': json['node_type'],
        'title': UiTextFromJSON(json['title']),
    };
}

export function UiNodeAnchorAttributesToJSON(value?: UiNodeAnchorAttributes | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'href': value['href'],
        'id': value['id'],
        'node_type': value['node_type'],
        'title': UiTextToJSON(value['title']),
    };
}

