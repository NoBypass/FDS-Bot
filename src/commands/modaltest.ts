import {
  SlashCommandBuilder,
  PermissionFlagsBits,
  TextInputStyle,
} from 'discord.js'
import { SlashCommand } from '../types/discord'
import {
  ActionRowBuilder,
  ModalBuilder,
  TextInputBuilder,
} from '@discordjs/builders'

const TestCommand: SlashCommand = {
  command: new SlashCommandBuilder()
    .setName('modaltest')
    .setDescription('Shows a test modal')
    .setDefaultMemberPermissions(PermissionFlagsBits.Administrator),

  execute: (interaction) => {
    const modal = new ModalBuilder()
      .setTitle('Test modal')
      .setCustomId('test_modal')

    const input = new TextInputBuilder()
      .setCustomId('test_input')
      .setPlaceholder('Test input')
      .setRequired(true)
      .setStyle(TextInputStyle.Short)

    modal.addComponents(
      new ActionRowBuilder<TextInputBuilder>().addComponents(input),
    )

    interaction.showModal(modal)

    return interaction.reply({ content: 'Modal shown' })
  },
  cooldown: 10,
}

export default TestCommand
