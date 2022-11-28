import { Message, MessageEmbed, PermissionResolvable } from "discord.js";
import descEmbed from "../counters/descEmbed";

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (message.member.roles.cache.some(role => role.name === 'Server Booster')) return descEmbed('You don\'t have permission to use this command.', message)

    }
}