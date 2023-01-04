import { Message, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client, guild } from "../index";
const path = require('path');
var ownname = path.basename(__filename);
import verifiedUsers from '../schemas/verified-users'

export default {
    callback: async (message: Message, ...args: string[]) => {
        message.reply(message.member.user.tag)
    }
}