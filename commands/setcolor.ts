import { ColorResolvable, Message, MessageEmbed, PermissionResolvable, RoleResolvable } from "discord.js";
import descEmbed from "../counters/descEmbed";
import { guild } from "../index";

export default {
    callback: async (message: Message, ...args: string[]) => {
        if (!message.member.roles.cache.some(role => role.name === 'Server Booster')) return descEmbed('You need to boost the server to use this command.', message)

        function isHex (hex) {
            return typeof hex === 'string'
                && hex.length === 6
                && !isNaN(Number('0x' + hex))
        }

        if (isHex(args[0] as any == false)) return descEmbed('Please format your command like this: "-setcolor <hex color (without the "#")>"', message)

        const existingRole = message.guild.roles.cache.find(x => x.name === message.member.id)

        if (existingRole == undefined) {
            try {
                var newRole = message.guild.roles.create({
                    name: message.member.id,
                    color: '#' + args[0] as ColorResolvable,
                    mentionable: false,
                    position: 20
                })
            } catch (err) {
                return descEmbed('There was an error while creating a new role: \n' + err, message)
            }

            message.member.roles.add(await newRole as RoleResolvable)
            return descEmbed('Changed your color!', message)
        } else {
            try {
                existingRole.edit({
                    color: '#' + args[0] as ColorResolvable,
                })
            } catch (err) {
                return descEmbed('There was an error while editing the role: \n' + err, message)
            }
            return descEmbed('Changed your color!', message)
        }
    }
}