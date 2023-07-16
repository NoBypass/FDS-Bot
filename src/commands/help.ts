import { PermissionFlagsBits, SlashCommandBuilder } from 'discord.js'
import { SlashCommand } from '../types/discord'
import { join } from 'path'
import { readdirSync } from 'fs'
import { EmbedBuilder } from '@discordjs/builders'

const HelpCommand: SlashCommand = {
  command: new SlashCommandBuilder()
    .setName('help')
    .setDescription('View all the commands and some additional info')
    .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages),

  execute: (interaction) => {
    const commands: SlashCommandBuilder[] = []
    const commandsDir = join(__dirname, '../commands')
    readdirSync(commandsDir).forEach((file) => {
      if (!file.endsWith('.ts')) return
      // eslint-disable-next-line @typescript-eslint/no-var-requires
      const command: SlashCommand = require(`${commandsDir}/${file}`).default
      if (!(command.command.name as string).startsWith('_'))
        commands.push(command.command)
    })

    const embed = new EmbedBuilder()
      .setTitle('Help Menu')
      .setDescription(
        `Here's a list of all slash commands provided by the <@${interaction.client.user.id}>`,
      )
    commands.forEach((command) => {
      embed.addFields({
        name: `/${command.name}`,
        value: command.description,
        inline: true,
      })
    })

    interaction.reply({ embeds: [embed], ephemeral: true })
  },
  cooldown: 10,
}

export default HelpCommand
