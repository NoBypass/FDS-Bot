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
        gm: boolean
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
    lastclaimed: reqNum,
})

export default mongoose.model('verifiedUsers', verifiedUsers, 'verifiedUsers')