import {SlashCommandBuilder, PermissionFlagsBits} from "discord.js"
import { SlashCommand } from "../types/discord"
import { EmbedBuilder } from "@discordjs/builders"

const command: SlashCommand = {
    command: new SlashCommandBuilder()
        .setName('test')
        .setDescription('Shows the bots ping and tests slash commands')
        .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages),

    execute: interaction => {
        const embed = new EmbedBuilder()
            .setColor(0x0099ff)
            .setTitle('Test successful')
            .setDescription(`Test successful, Ping: ${interaction.client.ws.ping}`)
        return interaction.channel?.send({embeds: [embed]})
    },
    cooldown: 10
}

export default command