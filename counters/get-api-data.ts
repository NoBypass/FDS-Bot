import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
import fetch from "node-fetch";
var rateLimit = require('function-rate-limit');

export const getApiData = rateLimit(110, 1000*60, async function (uuid: string) {
    const url = 'https://api.hypixel.net/player?key=' + hypixel_api_key + '&uuid=' + uuid

    let settings = { method: "Get" };

    const response = await fetch(url, settings);
    const result = await response.json();

    return result;
});

//export async function getApiData(uuid) {
//    const url = 'https://api.hypixel.net/player?key=' + hypixel_api_key + '&uuid=' + uuid
//
//    let settings = { method: "Get" };
//
//    const response = await fetch(url, settings);
//    const result = await response.json();
//
//    return result;
//}
