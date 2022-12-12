import { Message, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client, guild } from "../index";
const path = require('path');
var ownname = path.basename(__filename);
import verifiedUsers from '../schemas/verified-users'

export default {
    callback: async (message: Message, ...args: string[]) => {
        const TAGHOLDERS = await verifiedUsers.find({ customtag:{$exists:true} }).select('memberid -_id')
        const MEMBERS = await message.guild.members.fetch()

        for (let member of MEMBERS) {
            if (TAGHOLDERS.includes(message.guild.members.cache.find(m => m.id == member as any))) continue

            const memberid = await TAGHOLDERS.indexOf(member)
            const obj = await verifiedUsers.where('memberid').equals(memberid).select('ign level -_id')
            if (message.member.roles.cache.some(role => role.name === 'Guild Member')) var prefix = '❂ '
            else prefix = ''
            if (!message.member.permissions.has('ADMINISTRATOR')) message.member.setNickname(prefix + obj[0].ign + ' [' + obj[0].level + '] ').catch(err => {console.log(err)})
        }
    }
}