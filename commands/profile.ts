import { Message, MessageEmbed } from "discord.js";
import fetch from "node-fetch";
import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
var ownname = path.basename(__filename);
import verifiedUsers from '../schemas/verified-users'
import { getApiData } from '../counters/get-api-data'

export default {
    callback: async (message: Message, ...args: string[]) => {
        var mmb = '$' as string
        if (args.length > 1) return message.reply('You can\'t view the profiles of multiple users at once.')
        if (args.length == 0) mmb = (await verifiedUsers.findOne().where({ memberid: message.member.id }) as any).ign
        if (args.length == 1) mmb = await getDisplayName(args[0])

        async function getDisplayName(ign) {
            const url = 'https://api.ashcon.app/mojang/v2/user/' + ign

            let settings = { method: "Get" };

            const response = await fetch(url, settings);
            const result = await response.json();

            return result.username;
        }

        console.log(mmb)
        if (!await verifiedUsers.findOne().where({ ign: await mmb })) return message.reply('The member you provided is not in this server. Please use "-profile <ign>"')

        const memberM = await verifiedUsers.findOne().where({ ign: await mmb }) as any
        const memberDC = await client.users.fetch(await memberM.memberid).catch(err => console.log(err))
        const ign = await memberM.ign
        const level = await memberM.level
        const overXP = await memberM.xp

        const getNeededXP = (level) => {
            if (level < 10) {
                return Math.pow(level, 2) * 100
            } else if (level >= 10) {
                return 10000
            }
        }

        async function getXP(level) {
            if (level < 10) {
                let total = 0
                for (let i = 0; i < level; i++) {
                    total += Math.pow(i + 1, 2)
                }
                return (total - 1) * 100
            } else if (level >= 10) {
                return 28400 + (level - 9) * 10000
            }
        }

        const levelXP = getXP(await level)
        const totalXP = await levelXP + await overXP
        const neededXP = getNeededXP(await level) - await overXP
        const xpPercentage = Math.round(100 / getNeededXP(await level) * overXP)
        const trueeC = Math.round(xpPercentage / 5)
        var progressStr = ''

        for (let i = 0; i < trueeC; i++) {
            progressStr += '▓'
        }

        const till20 = 20 - progressStr.length

        for (let i = 0; i < till20; i++) {
            progressStr += '░'
        }

        const pBar = '[' + progressStr + ']'

        const embed = new MessageEmbed()
            .setColor('#2F3136')
            .setTitle(await ign + '\'s Profile [lvl: ' + await level + ']')
            .setThumbnail((memberDC as any).displayAvatarURL())
            .addFields(
                { name: 'Progress: ' + xpPercentage + '%' || '**error**', value: pBar || '**error**', inline: false },
                { name: 'Total XP: ' + totalXP || '**error**', value: '**XP needed: ' + neededXP + '**' || '**error**', inline: false },
            )
            .setFooter({ text: desc });

        // if (args[1] == '$t') {
        //     const nodeHtmlToImage = require('node-html-to-image')
// 
        //     const img = nodeHtmlToImage({
        //         quality: 100,
        //         //output: './image.png',
        //         type: 'png',
        //         html: '<html><body>Hello {{name}}!</body></html>',
        //         content: { name: 'you' },
        //         puppeteerArgs: {
        //             args: ['--no-sandbox'],
        //           },
        //         encoding: 'buffer'
        //     })
        //     message.channel.send(new AttachmentBuilder(img, { name: 'profile-image.png' }))
        //     
        // }

        return message.channel.send({ embeds: [embed] });
    }
}