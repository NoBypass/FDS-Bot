import { Message, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client, guild } from "../src/index";
import verifiedUsers from '../schemas/verified-users'

let claimedCache = []

const clearCache = () => {
    claimedCache = []
    setTimeout(clearCache, 1000 * 60 * 10)
}
clearCache()

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.permissions.has('ADMINISTRATOR')) return message.reply('You don\'t have permission to use this command, idiot!')

        var xptoadd = parseInt(args[1]) || 0

        const options = {
            upsert: true,
            new: false,
          };

        await verifiedUsers.findOneAndUpdate({ memberid: args[0] },
            {
                $inc: {
                    xp: xptoadd || 0,
                },
            },
            options,
        ).exec()

        message.reply('Added ' + xptoadd + ' xp to <@' + args[0] + '>\'s count!')

    }
}