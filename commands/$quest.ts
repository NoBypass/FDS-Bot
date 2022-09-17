import { Message, MessageEmbed, MessageActionRow, MessageButton, MessageAttachment, } from "discord.js";
import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
var ownname = path.basename(__filename);

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.permissions.has('ADMINISTRATOR')) return message.reply('You don\'t have permission to use this command.')
        const quests = [
            'Get the guild on one of the bridge guild leaderboards (either doubles or threes).',
            'Get on any weekly or monthly leaderboard.',
            'Get 200k weekly guild xp.',
            'Complete all the tnt games quests.',
            'Level up at least once.',
            'Invite at least 2 new guild members (no alts).',
            'Level up at least twice on discord.'
        ]
        const rewards = [
            '10k',
            '7.5k',
            '12.5k',
            '12.5k',
            '7.5k',
            '10k',
            '7.5k'
        ]
        let i = Math.floor(Math.random() * quests.length)
        const embed = new MessageEmbed()
            .setColor('#000000')
            .setTitle('Current guild quest:')
            .setDescription(quests[i] + ' ')
            .addFields(
                { name: 'Rewards:', value: rewards[i] + ' server xp', inline: false },
                { name: 'How to submit:', value: 'Open a "Misc" ticket in #〔👑〕join-guild and send a screenshot with proof. For the quests that require a party, you can mention your party members in the ticket. After submission you will recieve a ping and the xp.', inline: false }
            )
            .setFooter({ text: desc })

        message.channel.send({ embeds: [embed] })
        message.delete()
    }
}