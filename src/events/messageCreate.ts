import { BotEvent } from '../types/discord'
import { Message } from 'discord.js'
import { MainModel } from '../database/schema'
import { getNeededXp, isLevelUp } from '../lib/leveling'

const event: BotEvent = {
  name: 'messageCreate',
  once: true,
  execute: async (message: Message) => {
    const xpPerWord = 15
    const user = await MainModel.findOneAndUpdate(
      { memberid: message.interaction?.user.id },
      { $inc: { xp: message.content.split(' ').length * xpPerWord } },
      { new: true },
    )

    if (!user)
      return console.log('User not found at messageCreate event: ', message)

    if (isLevelUp(user.xp, user.level)) {
      const updatedLevel = user.level + 1
      const updatedXp = user.xp - getNeededXp(user.level)

      await MainModel.findOneAndUpdate(
        { memberid: message.interaction?.user.id },
        { $set: { level: updatedLevel, xp: updatedXp } },
      )

      await message.reply(`You leveled up to level ${updatedLevel}`)
    }
  },
}

export default event
