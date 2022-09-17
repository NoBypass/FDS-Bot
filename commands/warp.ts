import { Message } from "discord.js";
import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
var ownname = path.basename(__filename);

export default {
    callback: async (message: Message, ...args: string[]) => {
        const channel2pullFrom = message.guild.channels.cache.get('964675235314020442')
        if (channel2pullFrom.type != 'GUILD_VOICE') {
            return
        } else {
            //if (!message.member.permissions.has('MOVE_MEMBERS')) return
            if (!message.member.voice.channel) return message.reply("Error: Executor of command is not in a Voice Channel.")

            const sendersChannel = message.member.voice.channel.id
            const totalCount = channel2pullFrom.members.size
       
            channel2pullFrom.members.forEach((member) => {
                member.voice.setChannel(sendersChannel)
            })

            message.reply(`Moved ${totalCount} members.`)
        }
    }
}