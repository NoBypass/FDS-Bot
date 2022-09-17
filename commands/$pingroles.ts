import { Message, MessageActionRow, MessageButton, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
var ownname = path.basename(__filename);

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.permissions.has('ADMINISTRATOR')) return message.reply('You don\'t have permission to use this command.')
        const exampleEmbed = new MessageEmbed()
            .setColor('#000000')
            .setTitle('Pingroles:')
            .setDescription('Click on the buttons to get notified on specific events. Click again to remove the role.')
            .addFields(
                { name: '📢', value: 'Announcements notifications', inline: true },
                { name: '🤖', value: 'Bot update notifications', inline: true },
                { name: '🏆', value: 'Tournament notifications', inline: true },
                { name: '⚔️', value: 'Scrims notifications', inline: true },
                { name: '⭐', value: 'Bridge queue notifications', inline: true },
                { name: '🌟', value: 'BedWars queue notifications', inline: true },
                { name: '✨', value: 'WoolGames queue notifications', inline: true },
            )
            .setFooter({ text: desc });

        const buttonRow1 = new MessageActionRow().addComponents(
            new MessageButton()
                .setEmoji('📢')
                .setStyle('SECONDARY')
                .setCustomId('ann'),
            new MessageButton()
                .setEmoji('🤖')
                .setStyle('SECONDARY')
                .setCustomId('bot'),
            new MessageButton()
                .setEmoji('🏆')
                .setStyle('SECONDARY')
                .setCustomId('tou'),
            new MessageButton()
                .setEmoji('⚔️')
                .setStyle('SECONDARY')
                .setCustomId('scr'),
            new MessageButton()
                .setEmoji('⭐')
                .setStyle('SECONDARY')
                .setCustomId('bri'),
        )
        const buttonRow2 = new MessageActionRow().addComponents(
            new MessageButton()
                .setEmoji('🌟')
                .setStyle('SECONDARY')
                .setCustomId('bed'),
            new MessageButton()
                .setEmoji('✨')
                .setStyle('SECONDARY')
                .setCustomId('woo')
        )

        message.channel.send({ embeds: [exampleEmbed], components: [buttonRow1, buttonRow2] });
    }
}