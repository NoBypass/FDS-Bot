import { CommandInteraction, GuildMember } from 'discord.js'
import chalk from 'chalk'

export const getMemberByInteraction = async (
  interaction: CommandInteraction,
) => {
  const member = await interaction.guild?.members.fetch(interaction.user.id)
  if (!member) {
    throw new Error(
      chalk.red('Member not found at interaction: ' + interaction),
    )
  }
  return member
}

export const formatNick = (member: GuildMember) => {
  return member.nickname?.split('[')[0] || member.user.username
}
