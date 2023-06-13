import {
  ActionRowBuilder,
  ModalBuilder,
  TextInputBuilder,
} from '@discordjs/builders'
import { ModalSubmitInteraction, TextInputStyle } from 'discord.js'
import { Modal } from '../../types/discord'
import { sendMsgByInteraction } from '../../lib/common'
import { api } from '../../lib/api'
import { HypixelPlayer, MojangAccount } from '../../types/data'
import { isBooleanObject, isNumberObject } from 'util/types'

const verificationModel: Modal = {
  modal: new ModalBuilder()
    .setTitle('Verify your Minecraft account')
    .setCustomId('minecraftNameModal')
    .addComponents(
      new ActionRowBuilder<TextInputBuilder>().addComponents(
        new TextInputBuilder()
          .setCustomId('username')
          .setPlaceholder('Your Minecraft username')
          .setMinLength(1)
          .setMaxLength(16)
          .setRequired(true)
          .setStyle(TextInputStyle.Short),
      ),
    ),

  execute: async (interaction: ModalSubmitInteraction): Promise<void> => {
    let mojangAccount: Partial<MojangAccount> | null = null
    let hypixelPlayer: Partial<HypixelPlayer> | null = null
    const ign = interaction.fields.getTextInputValue('username')
    if (ign == null || isBooleanObject(ign) || isNumberObject(ign))
      throw new Error('Ign is invalid')

    try {
      mojangAccount = await api.get.mojangAccountByName(
        ign,
        'uuid, playedWith: { hypixelPlayer }',
      )
      if (mojangAccount == null) {
        mojangAccount = await api.add.mojangAccount(ign, '')
        if (!mojangAccount)
          return sendMsgByInteraction(
            interaction,
            `No Mojang account with name **${ign}** could be found.`,
          )
      }
      if (mojangAccount.playedWith == null) {
        hypixelPlayer = await api.add.hypixelPlayer('id')
        if (hypixelPlayer.id == null || mojangAccount.id == null) {
          interaction.reply({
            content: 'This error is not supposed to be possible wtf',
            ephemeral: true,
          })
          return
        }

        await api.connect.mojangAccountWithHypixelPlayer(
          {
            hypixelPlayerId: hypixelPlayer.id,
            mojangAccountId: mojangAccount.id,
          },
          '',
        )
      }

      interaction.reply({
        content: `Your Minecraft account **${ign}** has been successfully linked with your Discord account.`,
        ephemeral: true,
      })
    } catch (error) {
      interaction.reply({
        content:
          'An unexpected error occurred while talking to the API. Please try again later.',
        ephemeral: true,
      })
    }
  },
}

export default verificationModel
