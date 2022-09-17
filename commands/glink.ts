import { Message, MessageEmbed } from "discord.js";
import fetch from "node-fetch";
import { hypixel_api_key, desc, client, guildID } from "../index";
import verifiedUsers from '../schemas/verified-users'

export default {
    callback: async (message: Message, ...args: string[]) => {
        const ign = (await verifiedUsers.findOne().where({ memberid: message.member.id }).select('ign -_id') as any).ign
        const linked = (await verifiedUsers.findOne().where({ memberid: message.member.id }).select('customstats -_id') as any).customstats.gm
        console.log(linked)
        if (linked == true) return message.reply('You\'re already linked.')

        fetch('https://api.ashcon.app/mojang/v2/user/' + ign)
            .then(response => response.json())
            .then(async mdata => {
                const uuid = (mdata as any).uuid.replaceAll('-', '')

                if (uuid === undefined) return message.reply('Unexpected error: 0')

                const gURL = 'https://api.hypixel.net/guild?key=' + hypixel_api_key + '&id=' + guildID
                fetch(gURL)
                    .then(response => response.json())
                    .then(async gData => {
                        const gMemberArr = gData.guild.members
                        console.log(gMemberArr)
                        console.log(uuid)
                        var gmBo = false;
                        for (var i = 0; i < gMemberArr.length; i++) {
                            if (gMemberArr[i].uuid == uuid) {
                                gmBo = true;
                                break;
                            }
                        }

                        const userID = message.member.id

                        await verifiedUsers.findOneAndUpdate(
                            {
                                memberid: userID
                            },
                            {
                                customstats: {
                                    gm: gmBo,
                                },
                            }).catch(err => {
                                console.log(err)
                                message.reply('Unexpected error: c1')
                            })

                        const GMrole = message.guild.roles.cache.find(r => r.id === "1001868724955009055")
                        if (gmBo == true) {
                            message.member.roles.add(GMrole)
                            message.reply('Successfully verified as a guild member. You can now access all the guild-related channels.')
                        } else {
                            message.reply('You are not a member of "FDS Employees", please join the guild by applying in #〔👑〕join-guild and then use the command again.')
                        }

                        if (!message.member.permissions.has('ADMINISTRATOR')) message.member.setNickname('❂ ' + message.member.displayName).catch(err => { console.log(err) })
                    })
            })
    }

}