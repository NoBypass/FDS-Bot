import { Message, MessageEmbed } from "discord.js";

async function descEmbed (desc: string, message: Message) {
    const roles = new MessageEmbed()
        .setColor('#2F3136')
        .setTitle('Booster Perks:')
        .setDescription('Everything you can unlock by boosting the server. Can be changed with ``-setcolor`` and ``-settag``')
        .addFields(
            { name: 'Custom Tag', value: 'Just like the levels you can get an extra tag next to your name as long as it\'s appropriate', inline: true },
            { name: 'Custom Color', value: 'You can set your own custom role color.', inline: true },
        )
        .setFooter({ text: desc });

    message.channel.send({ embeds: [roles] });
}

export default descEmbed;