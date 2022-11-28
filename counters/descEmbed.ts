import { Message, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client } from "../index";

async function descEmbed (description: string, message: Message) {
    const roles = new MessageEmbed()
        .setColor('#2F3136')
        .setDescription(description)
        .setFooter({ text: desc });

    message.channel.send({ embeds: [roles] });
}

export default descEmbed;