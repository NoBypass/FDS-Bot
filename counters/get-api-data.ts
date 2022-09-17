import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
import fetch from "node-fetch";

export async function getApiData(uuid) {
    const url = 'https://api.hypixel.net/player?key=' + hypixel_api_key + '&uuid=' + uuid

    let settings = { method: "Get" };

    const response = await fetch(url, settings);
    const result = await response.json();

    return result;
}
