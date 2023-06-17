import {PermissionFlagsBits, SlashCommandBuilder} from 'discord.js'
import {SlashCommand} from '../types/discord'
import htmlToImage from 'html-to-image'
import { EmbedBuilder } from '@discordjs/builders'

const ProfileCommand: SlashCommand = {
    command: new SlashCommandBuilder()
        .setName('profile')
        .setDescription('View your server stats')
        .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages),

    execute: interaction => {
        const embed = new EmbedBuilder()
        const measurements = { width: 800, height: 300 }
        const { totalXp, level } = { totalXp: 5, level: 5 } // TODO link with database for guild xp and just for giving the xp
        const { levelXp, xpToGet } = { levelXp: 5, xpToGet: 5 } // TODO
        const html: any = `
            <style>
                .main {
                    width: ${measurements.width};
                    height: ${measurements.height};
                }
                .bar {
                    width: 100%;
                    height: 12px;
                    position: absolute;
                    border-radius: 1000px;
                    background-color: greenyellow;
                }
                .progress {
                    color: white;
                }
            </style>
            <div class="main">
                <div class="bar">
                    <p class="progress"></p>
                </div>
            </div>
        `

        htmlToImage.toPng(html, measurements)
            .then((dataUrl: string) => {
                embed.setImage(dataUrl)
            })
            .catch((error: Error) => console.error(error))
        embed.setTitle(`${interaction.user.username} is level ${level}`)
    },
    cooldown: 10
}

export default ProfileCommand