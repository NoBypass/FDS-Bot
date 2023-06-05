import { Message, MessageEmbed, PermissionResolvable } from "discord.js";
import descEmbed from "../counters/descEmbed";
import verifiedUsers from '../schemas/verified-users'

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.roles.cache.some(role => role.name === 'Server Booster')) return descEmbed('You need to boost the server to use this command.', message)

        let length = 0
        const obj = await verifiedUsers.where('memberid').equals(message.member.id).select('ign level -_id')
        length += obj[0].ign.length + 6
        if (message.member.roles.cache.some(role => role.name === 'Guild Member')) {
            var prefix = '❂ '
            length += 2;
        }
        else prefix = ''

        if (args.length > 1) return descEmbed('Please format your message like this: "-settag <your tag>"', message)
        if (args[0].length >= (16 - length)) return descEmbed('Please format your message like this: "-settag <your tag with lass than ' + (16 - length) + ' characters>"', message)

        async function pushTag(tag: string) {
            await verifiedUsers.findOneAndUpdate(
                {
                    memberid: message.member.id,
                },
                {
                    customtag: args[0]
                },
            )
        }
        pushTag(args[0])

        if (!message.member.permissions.has('ADMINISTRATOR')) message.member.setNickname(prefix + obj[0].ign + ' [' + obj[0].level + '] [' + args[0] + ']').catch(err => {console.log(err)})
    }
}