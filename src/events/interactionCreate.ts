import { Interaction } from 'discord.js'
import { BotEvent } from '../types/discord'
import checkCooldown from '../lib/cooldown'

const event: BotEvent = {
  name: 'interactionCreate',
  execute: (interaction: Interaction) => {
    if (interaction.isChatInputCommand()) {
      const command = interaction.client.slashCommands.get(
        interaction.commandName,
      )
      if (!command) return

      if (
        command.cooldown &&
        checkCooldown(
          `${interaction.commandName}-${interaction.user.username}`,
          interaction,
          command.cooldown,
        )
      )
        return

      command.execute(interaction)
    } else if (interaction.isAutocomplete()) {
      const command = interaction.client.slashCommands.get(
        interaction.commandName,
      )
      if (!command) {
        console.error(
          `No command matching ${interaction.commandName} was found.`,
        )
        return
      }
      try {
        if (!command.autocomplete) return
        command.autocomplete(interaction)
      } catch (error) {
        console.error(error)
      }
    } else if (interaction.isButton()) {
      const button = interaction.client.buttons.get(interaction.customId)
      if (!button) return

      if (
        button.cooldown &&
        checkCooldown(
          `${interaction.customId}-${interaction.user.username}`,
          interaction,
          button.cooldown,
        )
      )
        return

      button.execute(interaction)
    }
  },
}

export default event
