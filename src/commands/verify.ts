import {
    ActionRowBuilder,
    ButtonBuilder,
    EmbedBuilder,
    SlashCommandBuilder,
  } from '@discordjs/builders'
  import { SlashCommand } from '../types/discord'
  import { PermissionFlagsBits } from 'discord.js'
  import verifyButton from '../components/button/verifyButton'
  
  const VerifyCommand: SlashCommand = {
    command: new SlashCommandBuilder()
      .setName('verifymenu')
      .setDescription('Verify and link your account with Hypixel')
      .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages)
      .addStringOption((option) =>
        option
          .setName('username')
          .setDescription('Your Minecraft username')
          .setRequired(true),
      ),
  
    execute: async (interaction) => {
      interaction.reply({
        embeds: [
          new EmbedBuilder()
            .setTitle('Verify your Minecraft account')
            .setDescription('Click the button below to verify your account'),
        ],
        components: [
          new ActionRowBuilder<ButtonBuilder>().addComponents(
            verifyButton.button,
          ),
        ],
      })
    },
  
    cooldown: 2,
  }
  export default VerifyCommand