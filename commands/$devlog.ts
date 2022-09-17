import { ColorResolvable, Message, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client } from "../index";

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.permissions.has('ADMINISTRATOR')) return message.reply('You don\'t have permission to use this command.')
        const msg = args.join(' ').split(' || ')

        if (msg[0] == 'fix') {
            var type = '**Fix**'
            var color = '#edd539'
        } else if (msg[0] == 'new') {
            var type = '**New**'
            var color = '#2db333'
        } else if (msg[0] == 'removal') {
            var type = '**Removal**'
            var color = '#e63420'
        } else if (msg[0] == 'update') {
            var type = '**Update**'
            var color = '#1590e8'
        } else {
            var type = '**Misc**'
            var color = '#000000'
        }
        const fieldCount = msg.length - 2

        const embed = new MessageEmbed()
            .setColor(color as ColorResolvable)
            .setDescription(type + ' - ' + msg[1] || '**error**')
            .setFooter({ text: desc });

        message.delete()
        const aEmbed = await message.channel.send({ embeds: [embed] });

        if (fieldCount > 0) {
            const fields = msg.splice(2)
            for (let i = 0; i < fieldCount; i++) {
                const argss = fields[i].split(', ')

                const newEmbed = aEmbed.embeds[0]
                newEmbed.addFields({ name: argss[0], value: argss[1], inline: false });
                
                aEmbed.edit({
                    embeds: [newEmbed]
                })
            }
        }
    }
}