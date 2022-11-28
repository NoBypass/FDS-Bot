import { Message, MessageEmbed, PermissionResolvable } from "discord.js";
import descEmbed from "../counters/descEmbed";

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.permissions.has('Server Booster' as PermissionResolvable)) return descEmbed('You don\'t have permission to use this command.', message)

    }
}