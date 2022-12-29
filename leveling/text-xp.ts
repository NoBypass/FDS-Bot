import fetch from "node-fetch";
import verifiedUsers from '../schemas/verified-users'
import { hypixel_api_key, desc, client, guildID } from "../index";

module.exports = (client) => {
    client.on('messageCreate', async (message) => {
        if (await verifiedUsers.exists({ memberid: message.member.id }) !== null) {
            if (message.channel.parent.id === '1000781349248057354') return;
            const args = message.content.slice(1).split(/ +/)
            const { member } = message
            if (message.channel.parent.id == '1000781575681757254') var xpPerMessage = args.length * 6
            else var xpPerMessage = args.length * 3

            addXP(member.id, xpPerMessage, message)
        }
    })
}

const getNeededXP = (level) => {
    if (level < 10) {
        return Math.pow(level, 2) * 100
    } else if (level >= 10) {
        return 10000
    }
}

const addXP = async (memberid, xpToAdd, message) => {
    const result = await verifiedUsers.findOneAndUpdate(
        {
            memberid,
        },
        {
            memberid,
            $inc: {
                xp: xpToAdd,
            },
        },
    )

    verifiedUsers.updateOne(
        { memberid: memberid },
        {
            customstats: {
                day: {
                    $inc: {
                        xpFromText: xpToAdd
                    },
                }
            }
        }
    )

    let { xp, level } = result
    const needed = getNeededXP(level)

    if (xp >= needed) {
        ++level
        xp -= needed
        message.reply(
            'You are now level ' + level + '!'
        )
        await verifiedUsers.updateOne(
            {
                memberid,
            },
            {
                level,
                xp,
            }
        )
        const obj = await verifiedUsers.where('memberid').equals(memberid).select('ign level -_id')
        if (obj[0].customtag == undefined) var tag = ''
        else tag = '[' + obj[0].customtag+ ']'
        if (message.member.roles.cache.some(role => role.name === 'Guild Member')) var prefix = '❂ '
        else prefix = ''
        if (!message.member.permissions.has('ADMINISTRATOR')) message.member.setNickname(prefix + obj[0].ign + ' [' + obj[0].level + '] ' + tag).catch(err => {console.log(err)})
    }
}

module.exports.addXP = addXP
