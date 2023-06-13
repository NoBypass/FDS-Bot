import { ButtonInteraction, CommandInteraction } from 'discord.js'

const checkCooldown = (
  key: string,
  interaction: CommandInteraction | ButtonInteraction,
  itemCooldown: number,
): boolean => {
  const cooldown = interaction.client.cooldowns.get(key)
  if (itemCooldown && cooldown) {
    if (Date.now() < cooldown) {
      interaction.reply({
        content: `You have to wait ${Math.floor(
          Math.abs(Date.now() - cooldown) / 1000,
        )} second(s) to use this feature again.`,
        ephemeral: true,
      })
      setTimeout(() => interaction.deleteReply(), 5000)
      return false
    }

    interaction.client.cooldowns.set(key, Date.now() + itemCooldown * 1000)
    setTimeout(() => {
      interaction.client.cooldowns.delete(key)
    }, itemCooldown * 1000)
  } else if (itemCooldown && !cooldown) {
    interaction.client.cooldowns.set(key, Date.now() + itemCooldown * 1000)
  }

  return true
}

export default checkCooldown
