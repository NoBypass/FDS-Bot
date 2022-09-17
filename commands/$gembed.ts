import { Message, MessageEmbed, MessageActionRow, MessageButton, MessageAttachment, } from "discord.js";
import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
var ownname = path.basename(__filename);

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.permissions.has('ADMINISTRATOR')) return message.reply('You don\'t have permission to use this command.')
        //const attachment = new MessageAttachment('../assets/fds_emp.png')
        const firstEmbed = new MessageEmbed()
            .setColor('#000000')
            .setImage('https://imgur.com/YVkel12.png')
        const secondEmbed = new MessageEmbed()
            .setColor('#000000')
            .setTitle('Requirements:')
            .setDescription('For the requirements you can choose one of the two types of guild members: "Grinder" and "Sweat".')
            .addFields(
                { name: 'Grinder', value: ' - 150k GEXP/week', inline: false },
                { name: 'Sweat', value: ' - No GEXP requirement', inline: false },
                { name: 'You need to at least meet one of those for "Sweat":',
                value: ' - **BedWars:** 200 Stars and 5 FKDR or 50 Index\r - **Duels:** 7\'500 Wins and 3 WLR or 100 Index\r - **Bridge:** 5\'000 Wins and 5 WLR or 80 Index\r - **SkyWars:** 12 Stars and 3 KDR\r - **WoolGames:** 50 Stars and 3 WLR',
                inline: false },
            )
            .setFooter({ text: desc });
        const buttonRow = new MessageActionRow().addComponents(
                new MessageButton()
                    .setLabel('Apply for "Sweat"')
                    .setStyle('SUCCESS')
                    .setCustomId('swe'),
                new MessageButton()
                    .setLabel('Apply for "Grinder"')
                    .setStyle('SUCCESS')
                    .setCustomId('gri'),
                new MessageButton()
                    .setLabel('Create Misc Ticket')
                    .setStyle('DANGER')
                    .setCustomId('abs'),
                new MessageButton()
                    .setLabel('Forum Post')
                    .setStyle('LINK')
                    .setURL('https://hypixel.net/threads/fds-employees-fds-❂-lvl-6-❂-pvp-guild.5082995/'),
            )

        message.channel.send({ embeds: [firstEmbed] })
        message.channel.send({ embeds: [secondEmbed], components: [buttonRow] })
        message.delete()
    }
}
