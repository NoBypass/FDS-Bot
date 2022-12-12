import { Message, MessageEmbed } from "discord.js";
import fetch from "node-fetch";
import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
var ownname = path.basename(__filename);
import verifiedUsers from '../schemas/verified-users'
import { getApiData } from '../counters/get-api-data'
import { getIndex } from '../counters/index-count'

export default {
    callback: async (message: Message, ...args: string[]) => {
        const ign = (await verifiedUsers.findOne().where({memberid: message.member.id}).select('ign -_id') as any).ign
        if (args.length == 0) args = [ign]
        if (args.length !== 1) return message.reply('You need to provide one player. (-getindex <playername>)')
        try {
            fetch('https://api.mojang.com/users/profiles/minecraft/' + args)
                .then(response => response.json())
                .then(async mdata => {
                    var uuid = (mdata as any).id
                    const data = await getApiData(uuid) as any
                    const indexes = getIndex(await data) as any
                    const index = (await indexes).index
                    const indexarr = (await indexes).indexarr

                    if ((await data).player.stats == undefined) return message.reply('Player either does not exist or never joined hypixel.')

                    var facepng = 'https://crafatar.com/avatars/' + uuid + '?size=256&default=MHF_Alex&overlay'

                    const embed = new MessageEmbed()
                        .setColor('#2F3136')
                        .setTitle((await data).player.displayname + '\'s index is: ' + (await index).toLocaleString())
                        .setThumbnail(facepng)
                        .addFields(
                            { name: 'BedWars', value: 'Index: ' + (await indexarr[0]), inline: true },
                            { name: 'Duels', value: 'Index: ' + (await indexarr[5]), inline: true },
                            { name: 'SkyWars', value: 'Index: ' + (await indexarr[4]), inline: true },
                            { name: 'Bridge', value: 'Index: ' + (await indexarr[1]), inline: true },
                            { name: 'WoolGames', value: 'Index: ' + (await indexarr[3]), inline: true },
                            { name: 'Misc', value: 'Index: ' + (await indexarr[2]), inline: true },
                        )
                        .setFooter({ text: desc });
                            
                    return message.channel.send({ embeds: [embed] });
                })
            } catch (e) {
                message.reply('**Sorry, there was an error fetching the index or api data, this error likely means that the API couln\'t respond in time and you can try again in a couple oif minutes.**' + e)
        }
    }
}