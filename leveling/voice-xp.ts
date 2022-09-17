import { AnyArray } from "mongoose";
import fetch from "node-fetch";
import verifiedUsers from '../schemas/verified-users'
import { guild } from "../index"

module.exports = (client) => {
    const channelArr = [
        '964660650481700904',
        '964659549594333194',
        '964660515991335012',
        '1000891686483939430',
    ]

    async function getAllMembers() {
        var allMembers = await verifiedUsers.distinct('memberid')
        return allMembers
    }

    async function isInVoice() {
        const allMembers = await getAllMembers() as any
        for (var i = 0; i < allMembers.length; i++) {
            let guild = client.guilds.cache.get('897642715368534027');
            let member = guild.members.cache.get(allMembers[i])
            if (member == undefined) continue

            if (channelArr.includes(member.voice.channelId)) var xpPer2Mins = 16
            else var xpPer2Mins = 8

            if (member.voice.channelId) {
                addXP(allMembers[i], xpPer2Mins)
            }
        }
    }

    setInterval(isInVoice, 1000 * 60 * 2)
}

var options = { upsert: true, new: true, setDefaultsOnInsert: true };

const addXP = async (memberid, xpToAdd) => {
    await verifiedUsers.findOneAndUpdate(
        {
            memberid: memberid,
        },
        {
            $inc: {
                xp : xpToAdd,
            },
        },
        options,
    )
}

module.exports.addXP = addXP