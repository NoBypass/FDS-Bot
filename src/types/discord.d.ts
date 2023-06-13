import {
  CommandInteraction,
  Collection,
  PermissionResolvable,
  Message,
  AutocompleteInteraction,
  SlashCommandBuilder,
  User,
  Channel,
} from 'discord.js'

export interface SlashCommand {
  command: SlashCommandBuilder | any
  execute: (interaction: CommandInteraction) => void
  autocomplete?: (interaction: AutocompleteInteraction) => void
  cooldown?: number
}

export interface Command {
  name: string
  execute: (message: Message, args: Array<string>) => void
  permissions: Array<PermissionResolvable>
  aliases: Array<string>
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
    commands: Collection<string, Command>
    cooldowns: Collection<string, number>
  }
}