import { Client } from "discord.js";
import getFiles from './get-files'
import verifiedUsers from "../schemas/verified-users";

let prefix = '$'

export default (client: Client) => {
    const commands = {} as {
        [key: string]: any
    }
    console.log("Prefix: ", prefix)
    
    const suffix = '.ts'
    
    const commandFiles = getFiles('./commands', suffix)
    console.log(commandFiles)

    for (const command of commandFiles) {
        let commandFile = require(command)
        if (commandFile.default) commandFile = commandFile.default

        const split = command.replace(/\\/g, '/').split('/')
        const commandName = split[split.length - 1].replace(suffix, '')

        commands[commandName.toLowerCase()] = commandFile
    }
    console.log(commands)
    client.on('messageCreate', (message) => {
        verifiedUsers.updateOne(
            { memberid: message.member.id },
            {
                customstats: {
                    day: {
                        $inc: {
                            messagesSent: 1,
                        },
                    }
                }
            }
        )
        if (message.author.bot || !message.content.startsWith(prefix)) {
            return
        }
        const args = message.content.slice(1).split(/ +/)
        const commandName = args.shift()!.toLowerCase()

        if (!commands[commandName]) {
            return
        }

        verifiedUsers.updateOne(
            { memberid: message.member.id },
            {
                customstats: {
                    day: {
                        $inc: {
                            commandsExecuted: 1,
                        },
                    }
                }
            }
        )

        console.log(message.author.tag + ' used "' + commandName + '" at ' + new Date().toLocaleString())

        try {
            commands[commandName].callback(message, ...args)
        } catch (error) {
            console.error(error)
        }
    })
}