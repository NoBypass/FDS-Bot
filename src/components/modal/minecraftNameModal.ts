import { ActionRowBuilder, ModalBuilder, TextInputBuilder } from "@discordjs/builders"
import { TextInputStyle } from "discord.js"

const minecraftNameModal = new ModalBuilder()
    .setTitle('Verify your Minecraft account')
    .setCustomId('minecraftNameModal')
    .addComponents(new ActionRowBuilder<TextInputBuilder>()
        .addComponents(new TextInputBuilder()
            .setCustomId('minecraftNameInput')
            .setPlaceholder('Your Minecraft username')
            .setMinLength(3)
            .setMaxLength(16)
            .setRequired(true)
            .setStyle(TextInputStyle.Short)
        )
    )

export default minecraftNameModal