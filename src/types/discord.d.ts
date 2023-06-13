import { ButtonBuilder } from '@discordjs/builders'
import {
  CommandInteraction,
  Collection,
  PermissionResolvable,
  Message,
  AutocompleteInteraction,
  SlashCommandBuilder,
  User,
  Channel,
  ButtonInteraction,
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

export interface BotEvent {
  name: string
  once?: boolean | false
  execute: (...args: any[]) => void
}

declare module 'discord.js' {
  export interface Client {
    slashCommands: Collection<string, SlashCommand>
    cooldowns: Collection<string, number>
    buttons: Collection<string, Button>
    auth: JwtResponse
  }
}