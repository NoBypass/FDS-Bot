import { Message, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client, guild } from "../index";
const path = require('path');
var ownname = path.basename(__filename);
import verifiedUsers from '../schemas/verified-users'
import fetch from "node-fetch";
import descEmbed from "../counters/descEmbed";
import {getApiData} from "../counters/get-api-data";

export default {
    callback: async (message: Message, ...args: string[]) => {
        message.reply(message.member.user.tag)
        fetch('https://api.mojang.com/users/profiles/minecraft/' + args)
            .then(response => response.json())
            .catch(err => console.log(err))
            .then(async mdata => {
                var uuid = (mdata as any).id
                const data = await getApiData(uuid) as any
                message.reply(data?.socialMedia?.links?.DISCORD)
            })
    }
}