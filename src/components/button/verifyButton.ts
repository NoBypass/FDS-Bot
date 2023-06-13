import { ButtonBuilder } from "@discordjs/builders"
import { ButtonInteraction, ButtonStyle } from "discord.js"

export const verifyButton = {
    button: new ButtonBuilder()
        .setCustomId('verifyButton')
        .setLabel('Verify')
        .setStyle(ButtonStyle.Primary),
    execute: async (interaction: ButtonInteraction) => {

    }
}

export default verifyButton