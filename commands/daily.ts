import { Message, MessageEmbed } from "discord.js";
import { get } from "mongoose";
import { hypixel_api_key, desc, client } from "../index";
import verifiedUsers from '../schemas/verified-users'

export default {
    callback: async (message: Message, ...args: string[]) => {
        const userId = message.member.id
        const lastclaimed = (await verifiedUsers.find({ memberid: userId }).select('lastclaimed -_id') as any)[0].lastclaimed
        const randomXp = Math.floor(Math.random() * 1000)
        const currentTime = new Date().getTime()
        const day = 1000 * 60 * 60 * 24
        const yesterday = currentTime - day

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
        } else if (await currentTime - lastclaimed > day) {
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
        } else if (await currentTime - lastclaimed < day) return message.reply('You need to wait **' + convertMsToTime(currentTime - lastclaimed) + '** to claim your daily again.')
        else return message.reply('Looks like this command is still bugged!')

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
    }
}