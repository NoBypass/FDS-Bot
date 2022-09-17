import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
var ownname = path.basename(__filename);

module.exports = client => {
    try {
        const channelID = '960119399099031574'

        const updateMembers = guild => {
            const channel = guild.channels.cache.get(channelID)
            channel.setName('Total Members: ' + guild.memberCount.toLocaleString())
        }

        client.on('guildMemberAdd', member => updateMembers(member.guild))
        client.on('guildMemberRemove', member => updateMembers(member.guild))

        const guild = client.guilds.cache.get('897642715368534027')
        updateMembers(guild)
    } catch (e) {
        const channel = client.channels.cache.find(channel => channel.id === '973262334539747348') as any
        channel.send('**Error in** ' + ownname + ' : ' + e)
    }
}