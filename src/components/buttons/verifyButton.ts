import { ButtonBuilder } from '@discordjs/builders'
import { ButtonInteraction, ButtonStyle } from 'discord.js'
import verificationModel from '../modal/verificationModal'
import { Button } from '../../types/discord'

export const verifyButton: Button = {
  button: new ButtonBuilder()
    .setCustomId('verifyButton')
    .setLabel('Verify')
    .setStyle(ButtonStyle.Primary),

  execute: async (interaction: ButtonInteraction) => {
    await interaction.showModal(verificationModel.modal)
  },
  cooldown: 1,
}

export default verifyButton
