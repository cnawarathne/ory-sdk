/* tslint:disable */
/* eslint-disable */
/**
 * Ory APIs
 * # Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 
 *
 * The version of the OpenAPI document: v1.15.10
 * Contact: support@ory.sh
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

