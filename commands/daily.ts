import { Message, MessageEmbed } from "discord.js";
import { get } from "mongoose";
import { hypixel_api_key, desc, client } from "../index";
import verifiedUsers from '../schemas/verified-users'
import fetch from "node-fetch";

export default {
    callback: async (message: Message, ...args: string[]) => {
        const userId = message.member.id
        const lastclaimed = (await verifiedUsers.find({ memberid: userId }).select('lastclaimed -_id') as any)[0].lastclaimed
        const randomXp = Math.floor(Math.random() * 1000)
        const currentTime = (new Date).getTime()
        const dayBeginning = Date.parse(new Date().toString().slice(0, 13))
        const day = 1000 * 60 * 60 * 24

        if (await lastclaimed == undefined) {
            await verifiedUsers.findOneAndUpdate(
                {
                    memberid: userId,
                },
                {
                    $inc: {
                        xp: randomXp,
                    },
                    lastclaimed: currentTime,
                },
            )
            return message.reply('You recieved **' + randomXp + '** as a daily reward!')
        } else if (await lastclaimed < dayBeginning) {
            await verifiedUsers.findOneAndUpdate(
                {
                    memberid: userId,
                },
                {
                    $inc: {
                        xp: randomXp,
                    },
                    lastclaimed: currentTime,
                },
            )
            return reward()
        } else if (await lastclaimed > dayBeginning) return denied()
        else return message.reply('Unexpected erooor: 3')

        function padTo2Digits(num) {
            return num.toString().padStart(2, '0');
        }

        function convertMsToTime(milliseconds) {
            milliseconds = day - milliseconds
            let seconds = Math.floor(milliseconds / 1000);
            let minutes = Math.floor(seconds / 60);
            let hours = Math.floor(minutes / 60);

            seconds = seconds % 60;
            minutes = minutes % 60;

            hours = hours % 24;

            return `${padTo2Digits(hours)}h ${padTo2Digits(minutes)}m ${padTo2Digits(
                seconds,
            )}s`;
        }

        async function denied() {
            const roles = new MessageEmbed()
                .setColor('#e31010')
                .setTitle('Cannot claim daily right now')
                .setDescription('You need to wait **' + convertMsToTime(dayBeginning + day - currentTime) + '**  to claim your daily again.')
                .setFooter({ text: desc });

            message.channel.send({ embeds: [roles] });
        }

        async function reward() {
            const roles = new MessageEmbed()
            .setColor('#1f8f17')
            .setTitle('Claimed daily reward')
            .setDescription('You recieved **' + randomXp + '** ' + await getGuildExp() + 'as a daily reward!')
            .setFooter({ text: desc });

        message.channel.send({ embeds: [roles] });
        }

        async function getGuildExp() {
            const member = await verifiedUsers.findOne({ memberid: message.member.id })
            const uuid = member.uuid
            const url = 'https://api.hypixel.net/guild?key=' + hypixel_api_key + '&player=' + uuid

            let settings = { method: "Get" };
            const response = await fetch(url, settings);
            const data = await response.json();
        
            if (data.guild._id != '62e15cc48ea8c9296133317f') return ''
            const index = data.guild.members.findIndex(object => {
                return object.uuid === uuid;
            });
            const bonus = Math.round(data.guild.members[index].expHistory[formatDate(new Date())] / 50)

            await verifiedUsers.findOneAndUpdate(
                {
                    memberid: userId,
                },
                {
                    $inc: {
                        xp: bonus,
                    },
                    lastclaimed: currentTime,
                },
            )

            return '**+ ' + bonus + '** *(Guild Member Bonus)* '
        }

        function padTo2Digits2(num) {
            return num.toString().padStart(2, '0');
        }
          
        function formatDate(date) {
            return [
                date.getFullYear(),
                padTo2Digits2(date.getMonth() + 1),
                padTo2Digits2(date.getDate() - 1),
            ].join('-');
        }
    }
}