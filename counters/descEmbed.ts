import { Message, MessageEmbed } from "discord.js";

async function descEmbed (desc: string, message: Message) {
    const roles = new MessageEmbed()
        .setColor('#2F3136')
        .setTitle('Booster Perks:')
        .setDescription(desc)
        .setFooter({ text: desc });

    message.channel.send({ embeds: [roles] });
}

export default descEmbed;