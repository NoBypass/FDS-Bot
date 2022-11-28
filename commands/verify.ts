import { Message, MessageEmbed } from "discord.js";
import fetch from "node-fetch";
import { hypixel_api_key, desc, client, guildID } from "../index";
import verifiedUsers from '../schemas/verified-users'
import { getApiData } from '../counters/get-api-data'
import { getIndex } from '../counters/index-count'
import { getBedWarsStar } from '../counters/index-count'
import descEmbed from '../counters/descEmbed'

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.roles.cache.some(role => role.name === 'Guest')) return message.reply('You\'re already verified.')

        fetch('https://api.mojang.com/users/profiles/minecraft/' + args)
            .then(response => response.json())
            .then(async mdata => {
                var uuid = (mdata as any).id

                if (uuid === undefined) return descEmbed('Player either does not exist or the API couldn\'t respond in time. If thats the case please try again in a minute.', message)

                /*
                const gData = 'https://api.hypixel.net/guild?key=' + hypixel_api_key + '&id=' + guildID
                const gMemberArr = (await gData as any).guild.members
                if (await gMemberArr.some(e => e.uuid === uuid.replace('-', ''))) var gmBo = true
                else gmBo = false

                const GMrole = message.guild.roles.cache.find(r => r.id === "1001868724955009055")
                if (gmBo == true) message.member.roles.add(GMrole)
                */

                const data = await getApiData(uuid) as any
                const bwlevel = getBedWarsStar((await data)?.player?.stats?.Bedwars?.Experience)
                const indexes = getIndex(await data) as any
                const index = (await indexes).index
                const indexarr = (await indexes).indexarr

                var facepng = 'https://crafatar.com/avatars/' + uuid + '?size=256&default=MHF_Alex&overlay'

                if (index < 100) {
                    var rolestr = 'Bronze'
                    var role = message.guild.roles.cache.find(r => r.id === "964653007117627453")
                    message.member.roles.add(role)
                }
                else if (index >= 100 && index < 200) {
                    var rolestr = 'Silver'
                    var role = message.guild.roles.cache.find(r => r.id === "964652871113142282")
                    message.member.roles.remove("964653007117627453");
                    message.member.roles.add(role)
                }
                else if (index >= 200 && index < 300) {
                    var rolestr = 'Platinum'
                    var role = message.guild.roles.cache.find(r => r.id === "964652446154633256")
                    message.member.roles.remove("964653007117627453");
                    message.member.roles.remove("964652871113142282");
                    message.member.roles.add(role)
                }
                else if (index >= 400) {
                    var rolestr = 'Titanium'
                    var role = message.guild.roles.cache.find(r => r.id === "964651864052334683")
                    message.member.roles.remove("964653007117627453");
                    message.member.roles.remove("964652871113142282");
                    message.member.roles.remove("964652446154633256");
                    message.member.roles.add(role)
                }

                const channel = client.channels.cache.find(channel => channel.id === '973634176169414686') as any
                channel.send((await data).player.displayname + ' **verified at** ' + new Date().toLocaleString())

                if (message.member.roles.cache.some(role => role.name === "Guest")) {
                    message.member.roles.remove("964653096611491870");
                }
                if (!message.member.permissions.has('ADMINISTRATOR')) message.member.setNickname((await data)?.player?.displayname).catch(err => { console.log(err) })

                var userID = message.member.id

                const currentTime = new Date().getTime()
                const yesterday = currentTime - 1000 * 60 * 60 * 24

                async function verify() {
                    await new verifiedUsers({
                        ign: (await data)?.player?.displayname,
                        uuid: (await data)?.player?.uuid,
                        memberid: userID,
                        stats: {
                            duelswins: (await data)?.player?.stats?.Duels?.wins || 0,
                            duelsdeaths: (await data)?.player?.stats?.Duels?.deaths + (await data)?.player?.stats?.Duels?.bridge_deaths || 0,
                            duelskills: (await data)?.player?.stats?.Duels?.kills || 0,
                            bridgewins: (await data)?.player?.achievements?.duels_bridge_wins || 0,
                            bedwarsfinals: (await data)?.player?.stats?.Bedwars?.final_kills_bedwars || 0,
                            bedwarsstars: await bwlevel || 0,
                        },
                        customstats: {
                            index: await index,
                            //gm: gmBo,
                        },
                        xp: 0,
                        level: 1,
                        lastclaimed: yesterday,
                    }).save()
                }
                verify()

                const embed = new MessageEmbed()
                .setColor('#2F3136')
                .setTitle('Successfully linked ' + message.member.user.tag + ' with ' + (await data)?.player?.displayname)
                .setDescription('Your index is **' + await index + '** thus your division is <@&' + role + '>')
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
    }
}