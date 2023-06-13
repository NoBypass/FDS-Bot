import { SlashCommandBuilder } from '@discordjs/builders'
import { SlashCommand } from '../types/discord'
import { PermissionFlagsBits } from 'discord.js'
import { sendMsgByInteraction } from '../lib/common'
import { api } from '../lib/api'
import { HypixelPlayer, MojangAccount } from '../types/data'
import { isBooleanObject, isNumberObject } from 'util/types'

const command: SlashCommand = {
    command: new SlashCommandBuilder()
    .setName('verifyMenu')
    .setDescription('Verify and link your account with Hypixel')
    .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages)
    .addStringOption(option => option
        .setName('username')
        .setDescription('Your Minecraft username')
        .setRequired(true)
    ),
    
    execute: async interaction => {
        let mojangAccount: Partial<MojangAccount> | null = null
        let hypixelPlayer: Partial<HypixelPlayer> | null = null
        const ign = interaction.options.get('username')?.value
        if (ign == null || isBooleanObject(ign) || isNumberObject(ign)) throw new Error('Ign is invalid')

        try {
            mojangAccount = await api.get.mojangAccountByName(ign, 'uuid, playedWith: { hypixelPlayer }')
            if (mojangAccount == null) {
                mojangAccount = await api.add.mojangAccount(ign, '')
                if (!mojangAccount) return sendMsgByInteraction(interaction, `No Mojang account with name **${ign}** could be found.`)
            }
            if (mojangAccount.playedWith == null) {
                hypixelPlayer = await api.add.hypixelPlayer('id')
                if (hypixelPlayer.id == null || mojangAccount.id == null) return interaction.reply({
                        content: 'This error is not supposed to be possible wtf',
                        ephemeral: true
                })

                await api.connect.mojangAccountWithHypixelPlayer({
                    hypixelPlayerId: hypixelPlayer.id,
                    mojangAccountId: mojangAccount.id
                }, '')
            }
            
        } catch (error) {
            interaction.reply({
                content: 'An unexpected error occurred while talking to the API. Please try again later.',
                ephemeral: true
            })
        }


    },

    cooldown: 2
}