import { Message, MessageEmbed } from "discord.js";
import descEmbed from "../counters/descEmbed";
import { hypixel_api_key, desc, client } from "../index";

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.permissions.has('ADMINISTRATOR')) return descEmbed('You don\'t have permission to use this command.', message)
        const roles = new MessageEmbed()
            .setColor('#2F3136')
            .setTitle('Roles:')
            .setDescription('All the Roles and how to obtain them. \rThe formula for the index is: \r```√(BedWars Stars * 12 + FKDR^2.5 * 3.5) + (Duels Wins / 3.5 + Duels WLR^1.5) + (SkyWars Stars^2.3 + KDR^5) + (Bridge Wins + Bridge WLR^1.7) + (WoolGames Stars * 15 + (WoolWars Kills + WoolWars Assits / 2)^3.5 * 3.5) + (Achievements * 2.2 + Nwetworklevel)) / 100\' ^1.85```')
            .addFields(
                { name: 'The Trio', value: '<@&' + message.guild.roles.cache.find(role => role.name == "The Trio") + '> Should be called "The Duo" since Kake is lame and inactive.', inline: true },
                { name: 'Trusted', value: '<@&' + message.guild.roles.cache.find(role => role.name == "Trusted") + '> People we know fairly well', inline: true },
                { name: 'Bot', value: '<@&' + message.guild.roles.cache.find(role => role.name == "Bot") + '> Bots.', inline: true },
                { name: 'Guest', value: '<@&' + message.guild.roles.cache.find(role => role.name == "Guest") + '> Default role', inline: true },
                { name: 'Bronze', value: '<@&' + message.guild.roles.cache.find(role => role.name == "Bronze") + '> Verified people and people who have less than 200 index points', inline: true },
                { name: 'Silver', value: '<@&' + message.guild.roles.cache.find(role => role.name == "Silver") + '> Verified people and people who have between 200 and 300 index points', inline: true },
                { name: 'Platinum', value: '<@&' + message.guild.roles.cache.find(role => role.name == "Platinum") + '> Verified people and people who have between 300 and 400 index points', inline: true },
                { name: 'Titanium', value: '<@&' + message.guild.roles.cache.find(role => role.name == "Titanium") + '> Verified people and people who have over 400 index points', inline: true }
            )
            .setFooter({ text: desc });

        message.channel.send({ embeds: [roles] })
        message.delete()
    }
}