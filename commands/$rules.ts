import { Message, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client } from "../index";
import descEmbed from "../counters/descEmbed";

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.permissions.has('ADMINISTRATOR')) return descEmbed('You don\'t have permission to use this command.', message)
        const exampleEmbed = new MessageEmbed()
            .setColor('#2F3136')
            .setTitle('Rules:')
            .setDescription('Ignoring these rules can lead to you getting muted or even banned.')
            .addFields(
                { name: 'No Discrimination or Herassment', value: 'Slurs are only allowed to a certain extent. Only exception is the rule "No not bullying Shoe"', inline: true },
                { name: 'No Doxxing', value: 'Sharing & publishing of personal information is not allowed.', inline: true },
                { name: 'No NSFW', value: 'Don\'t even talk about it... it\'s cringe.', inline: true },
                { name: 'No Spamming', value: 'This includes mass pinging and ghost pinging.', inline: true },
                { name: 'No Advertising', value: 'At least without permission.', inline: true },
                { name: 'No Rickrolling', value: 'No explanation required.', inline: true },
                { name: 'No not bullying Shoe', value: 'You always have to call "Shoetimee" "Shoe" and also do your best to bully him.', inline: true },
                { name: 'No Being an Idiot', value: 'Just imagine being an idiot...', inline: true },
                { name: 'No Hypixel Rule Breaking', value: 'Ignoring the rules set by hypixel is disallowed. Even on alts.', inline: true },
            )
            .setFooter({ text: desc });

        message.channel.send({ embeds: [exampleEmbed]});
    }
}