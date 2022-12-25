import mongoose from 'mongoose'

const reqString = {
    type: String,
    required: true
}

const reqNum = {
    type: Number,
    required: true
}

const boolean = {
    type: Boolean,
    required: false,
    default: false
}

const serverStats = {
    timeSpentInVcs: Number,
    messagesSent: Number,
    commandsExecuted: Number,
    dailiesClaimed: Number,
    xpFromDailies: Number,
    xpFromVcs: Number,
    xpFromText: Number,
    indexGained: Number
}

const verifiedUsers = new mongoose.Schema({
    ign: reqString,
    uuid: reqString,
    memberid: reqString,
    stats: {
        duelswins: reqString,
        duelsdeaths: reqString,
        duelskills: reqString,
        bridgewins: reqString,
        bedwarsfinals: reqString,
        bedwarsstars: reqString,
    },
    customstats: {
        index: reqString,
        gm: boolean,
        total: serverStats,
        year: serverStats,
        lastyear: serverStats,
        day: serverStats,
        yesterday: serverStats,
        week: serverStats,
        lastweek: serverStats,
        month: serverStats,
        lastmonth: serverStats
    },
    form: {
        cquestion: {
            type: Number,
            required: false,
            default: 0,
        },
    },
    xp: {
        type: Number,
        default: 0,
    },
    level: {
        type: Number,
        default: 1,
    },
    customtag: {type: String},
    lastclaimed: reqNum,
})

export default mongoose.model('verifiedUsers', verifiedUsers, 'verifiedUsers')