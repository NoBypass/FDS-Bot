import { PermissionFlagsBits, SlashCommandBuilder } from 'discord.js'
import { SlashCommand } from '../types/discord'
import { EmbedBuilder } from '@discordjs/builders'
import { MainModel } from '../database/schema'
import { formatNick, getMemberByInteraction } from '../lib/discord'

const DailyCommand: SlashCommand = {
  command: new SlashCommandBuilder()
    .setName('daily')
    .setDescription('Claim your daily exp through this command')
    .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages),

  execute: async (interaction) => {
    const member = await getMemberByInteraction(interaction)
    const user = await MainModel.findOne({
      memberid: interaction.user.id,
    })
    if (!user) {
      return interaction.reply({
        content:
          'You are not registered in the database, please verify your account',
        ephemeral: true,
      })
    }
    const lastClaimed = user.lastclaimed
    const isSameDay = new Date(lastClaimed).getDate() == new Date().getDate()
    if (isSameDay) {
      return interaction.reply({
        content: `You have already claimed your daily reward, please wait until tomorrow`,
        ephemeral: true,
      })
    }
    const hasLostStreak = new Date().getTime() - lastClaimed > 86400000
    const randomized = Math.random() * 500
    const xpToGive = Math.round(
      randomized + (randomized / 100) * 5 * user.streak,
    )

    const embed = new EmbedBuilder()
      .setTitle(
        hasLostStreak
          ? `${formatNick(
              member,
            )} lost their streak but gained **+${xpToGive}** xp`
          : `${formatNick(member)} ${
              xpToGive < 400 && xpToGive > 100
                ? 'claimed their daily reward'
                : `got ${xpToGive > 450 || xpToGive < 50 ? 'very' : ''} ${
                    xpToGive < 100 ? 'un' : ''
                  }lucky`
            }`,
      )
      .setDescription(
        hasLostStreak
          ? `Streak of \`\`${
              user.streak
            }\`\` was lost (streak started on ${new Date(
              new Date(new Date().getTime() - user.streak * 86400000).getDate(),
            ).toLocaleDateString('en-US', {
              weekday: 'short',
              year: 'numeric',
              month: 'short',
              day: 'numeric',
            })})`
          : `Received **+${xpToGive}** xp\nCurrent streak: \`\`${
              user.streak + 1
            }\`\``,
      )

    await interaction.reply({ embeds: [embed] })

    await MainModel.findOneAndUpdate(
      { memberid: interaction.user.id },
      {
        $inc: { xp: xpToGive },
        lastclaimed: new Date().getTime(),
        streak: hasLostStreak ? 0 : user.streak + 1,
      },
    )

    // TODO guild xp system
  },
  cooldown: 10,
}

export default DailyCommand
