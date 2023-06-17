import { ButtonBuilder, ModalBuilder } from '@discordjs/builders'
import {
  CommandInteraction,
  Collection,
  AutocompleteInteraction,
  SlashCommandBuilder,
  ButtonInteraction,
  ModalSubmitInteraction,
} from 'discord.js'
import { JwtResponse } from './api'

export interface SlashCommand {
  command: SlashCommandBuilder | any
  execute: (interaction: CommandInteraction) => void
  autocomplete?: (interaction: AutocompleteInteraction) => void
  cooldown?: number
}

export interface Button {
  button: ButtonBuilder
  execute: (interaction: ButtonInteraction) => void
  cooldown?: number
}

export interface Modal {
  modal: ModalBuilder
  execute: (interaction: ModalSubmitInteraction) => void
}

export interface BotEvent {
  name: string
  once?: boolean | false
  execute: (...args: any[]) => void
}

export interface Attachment {
  files: {
    attachment: string
    name: string
  }[]
}

declare module 'discord.js' {
  export interface Client {
    slashCommands: Collection<string, SlashCommand>
    cooldowns: Collection<string, number>
    buttons: Collection<string, Button>
    auth: JwtResponse
  }
}

declare global {
  namespace NodeJS {
    interface ProcessEnv {
      TOKEN: string
      API_URI: string
      API_VERSION: string
      CLIENT_ID: string
      ERROR_CHANNEL_ID: string
    }
  }
}
