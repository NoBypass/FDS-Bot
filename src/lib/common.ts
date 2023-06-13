import { CommandInteraction, MessagePayload } from "discord.js"

export const graphqlSafeString = (str: string): string => {
    const regex = /^[a-zA-Z0-9_{}:,\[\]\(\)\!\$]*$/
    return regex.test(str) ? str : ''
}

export const errLog = (err: Error | unknown, interaction: CommandInteraction, during?: string): void => {
    sendMsgByInteraction(interaction, `Error ${during == null ? '' : 'while'} ${during} by **${interaction.user.username}**:\n ${err}`)
}

export const sendMsgByInteraction = (interaction: CommandInteraction, msg: string | MessagePayload): void => {
    const channel = interaction.client.channels.cache.get(process.env.ERROR_CHANNEL_ID)
    if (channel?.isTextBased()) channel.send(msg)
}