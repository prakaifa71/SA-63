/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPatientroomEdges,
    EntPatientroomEdgesFromJSON,
    EntPatientroomEdgesFromJSONTyped,
    EntPatientroomEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientroom
 */
export interface EntPatientroom {
    /**
     * Typeroom holds the value of the "Typeroom" field.
     * @type {string}
     * @memberof EntPatientroom
     */
    typeroom?: string;
    /**
     * 
     * @type {EntPatientroomEdges}
     * @memberof EntPatientroom
     */
    edges?: EntPatientroomEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatientroom
     */
    id?: number;
}

export function EntPatientroomFromJSON(json: any): EntPatientroom {
    return EntPatientroomFromJSONTyped(json, false);
}

export function EntPatientroomFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientroom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'typeroom': !exists(json, 'Typeroom') ? undefined : json['Typeroom'],
        'edges': !exists(json, 'edges') ? undefined : EntPatientroomEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPatientroomToJSON(value?: EntPatientroom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Typeroom': value.typeroom,
        'edges': EntPatientroomEdgesToJSON(value.edges),
        'id': value.id,
    };
}

