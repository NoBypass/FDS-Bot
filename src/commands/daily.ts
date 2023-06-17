import {PermissionFlagsBits, SlashCommandBuilder} from 'discord.js'
import {SlashCommand} from '../types/discord'
import { EmbedBuilder } from '@discordjs/builders'

const DailyCommand: SlashCommand = {
    command: new SlashCommandBuilder()
        .setName('daily')
        .setDescription('Claim your daily exp through this command')
        .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages),

    execute: interaction => {
        const xpToGive = Math.round(Math.random() * 500)
        const embed = new EmbedBuilder()
        .setTitle(`${interaction.user.username} ${xpToGive < 400 && xpToGive > 100 ? 'claimed their daily reward' : `got ${xpToGive > 450 || xpToGive < 50 ? 'very' : ''} ${xpToGive < 100 ? 'un' : ''}lucky`}`)
        .setDescription(`Received **+${xpToGive}** xp`)

        // TODO streaks
        // TODO check if level changed
        // TODO link with database for guild xp and just for giving the xp
    },
    cooldown: 10
}

export default DailyCommand