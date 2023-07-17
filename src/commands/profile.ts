import { PermissionFlagsBits, SlashCommandBuilder } from 'discord.js'
import { SlashCommand } from '../types/discord'
import generateImageFromHtml from '../lib/image'
import { MainModel } from '../database/schema'
import { getNeededXp, getTotalXp } from '../lib/leveling'
import { formatNick, getMemberByInteraction } from '../lib/discord'

const ProfileCommand: SlashCommand = {
  command: new SlashCommandBuilder()
    .setName('profile')
    .setDescription('View your server stats')
    .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages)
    .addUserOption((option) =>
      option
        .setName('user')
        .setDescription('The user to view the profile of')
        .setRequired(false),
    ),
  execute: async (interaction) => {
    const member = await getMemberByInteraction(interaction)
    if (!member) {
      return interaction.reply({
        content: `User not found`,
        ephemeral: true,
      })
    }

    const measurements = { width: 800, height: 174 }
    const dbEntry = await MainModel.findOne({ memberid: member.id })
    if (!dbEntry) {
      return interaction.reply({
        content:
          'You are not registered in the database, please verify your account',
        ephemeral: true,
      })
    }
    const { xp, level } = dbEntry
    const neededXp = getNeededXp(level + 1)
    const totalXp = Math.round(getTotalXp(level, xp))
    const avatar = member.displayAvatarURL({ forceStatic: true })

    const html = `
      <main>
        <h1>${formatNick(member)} is level <b>${level}</b></h1>
        <div class="bar" >
            <div class="completed" />
        </div>
        <p class="progress">${Math.round(xp)}/${neededXp}</p>
        <p class="total">Total xp: <b>${totalXp}</b></p>
      </main>
    `
    const css = `
      <style>
        body {
          padding: 0;
          margin: 0;
          overflow: hidden;
        }
        main {
          position: relative;
          text-align: center;
          width: ${measurements.width};
          height: ${measurements.height};
          background-color: #2C2F33;
          margin-left: -8px;
          margin-top: -8px;
          padding-top: 1px;
          z-index: 1;
        }
        main::before {
          content: '';
          position: absolute;
          top: 0;
          left: 0;
          width: 100%;
          height: 100%;
          background-image: url(${avatar});
          background-position: center center;
          background-repeat: no-repeat;
          background-size: cover;
          filter: brightness(0.3) saturate(1.2) blur(4px);
        }
        .bar {
          width: 90%;
          height: 28px;
          position: absolute;
          border-radius: 1000px;
          margin-left: 5%;
          background-color: #2C2F33FF;
          z-index: 1;
        }
        .completed {
          width: ${Math.round((xp / neededXp) * 100)}%;
          height: 28px;
          position: absolute;
          border-radius: 1000px;
          background-color: #7289DAFF;
          z-index: 2;
        }
        .progress {
          margin-top: -1px;
          font-size: 24px;
          font-weight: bold;
          position: relative;
          z-index: 3;
        }
        .total {
          font-size: 24px;
          z-index: 1;
        }
        h1 {
          z-index: 2;
          position: relative;
        }
        h1, p, b {
          background-color: transparent;
          color: white;
        }
        p {
          font-size: 20px;
          z-index: 1;
        }
      </style>
    `

    await interaction.reply({
      files: [await generateImageFromHtml(html, css, measurements)],
    })
  },
  cooldown: 10,
}

export default ProfileCommand
